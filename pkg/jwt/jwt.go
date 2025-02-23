package jwt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nuomizi-fw/stargazer/api"
)

const (
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"

	// Web client has shorter token duration for security
	WebAccessTokenDuration  = 2 * time.Hour
	WebRefreshTokenDuration = 7 * 24 * time.Hour

	// Native client has longer token duration for better UX
	NativeAccessTokenDuration  = 24 * time.Hour
	NativeRefreshTokenDuration = 30 * 24 * time.Hour

	JWKSPath = "/.well-known/jwks.json"
)

type JWK struct {
	Kty string `json:"kty"`
	Crv string `json:"crv"`
	X   string `json:"x"`
	Y   string `json:"y"`
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

type Claims struct {
	Username   string `json:"username"`
	ID         int    `json:"id"`
	TokenType  string `json:"token_type"`
	ClientType string `json:"client_type"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, privateKey *ecdsa.PrivateKey, id int, tokenType string, clientType string, t time.Duration) (string, error) {
	claims := Claims{
		Username:   username,
		ID:         id,
		TokenType:  tokenType,
		ClientType: clientType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "stargazer",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(privateKey)
	return signedToken, err
}

// get AccessToken for web client
func GetWebAccessToken(username string, privateKey *ecdsa.PrivateKey, id int) (string, error) {
	return GenerateToken(username, privateKey, id, TokenTypeAccess, string(api.Web), WebAccessTokenDuration)
}

// get RefreshToken for web client
func GetWebRefreshToken(username string, private *ecdsa.PrivateKey, id int) (string, error) {
	return GenerateToken(username, private, id, TokenTypeRefresh, string(api.Web), WebRefreshTokenDuration)
}

// get AccessToken for native client
func GetNativeAccessToken(username string, privateKey *ecdsa.PrivateKey, id int) (string, error) {
	return GenerateToken(username, privateKey, id, TokenTypeAccess, string(api.Native), NativeAccessTokenDuration)
}

// get RefreshToken for native client
func GetNativeRefreshToken(username string, private *ecdsa.PrivateKey, id int) (string, error) {
	return GenerateToken(username, private, id, TokenTypeRefresh, string(api.Native), NativeRefreshTokenDuration)
}

func ParseToken(signedToken string, publicKeyFunc func() (*ecdsa.PublicKey, error)) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return publicKeyFunc()
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func Validate(token string, publicKeyFunc func() (*ecdsa.PublicKey, error)) (bool, *Claims, error) {
	claims, err := ParseToken(token, publicKeyFunc)
	if err != nil {
		return false, nil, err
	}

	if claims == nil {
		return false, nil, errors.New("invalid token")
	}

	// Validate token type
	if claims.TokenType != TokenTypeAccess && claims.TokenType != TokenTypeRefresh {
		return false, nil, errors.New("invalid token type")
	}

	// Validate client type
	if claims.ClientType != string(api.Web) && claims.ClientType != string(api.Native) {
		return false, nil, errors.New("invalid client type")
	}

	// Additional validation for refresh tokens
	if claims.TokenType == TokenTypeRefresh {
		maxDuration := WebRefreshTokenDuration
		if claims.ClientType == string(api.Native) {
			maxDuration = NativeRefreshTokenDuration
		}
		if time.Since(claims.IssuedAt.Time) > maxDuration {
			return false, nil, errors.New("refresh token expired")
		}
	}

	return true, claims, nil
}

func GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating key pair: %w", err)
	}

	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}

func GenerateJwksJSON(publicKey *ecdsa.PublicKey) ([]byte, error) {
	jwk := JWK{
		Kty: "EC",
		Crv: "P-256",
		X:   base64.RawURLEncoding.EncodeToString(publicKey.X.Bytes()),
		Y:   base64.RawURLEncoding.EncodeToString(publicKey.Y.Bytes()),
	}

	jwks := JWKS{
		Keys: []JWK{jwk},
	}

	return json.Marshal(jwks)
}

func PublicKeyFromJwksJSON(jwksJSON []byte) (*ecdsa.PublicKey, error) {
	var jwks JWKS
	err := json.Unmarshal(jwksJSON, &jwks)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JWKS JSON: %w", err)
	}

	if len(jwks.Keys) == 0 {
		return nil, fmt.Errorf("no keys in JWKS")
	}

	jwk := jwks.Keys[0]

	x, err := base64.RawURLEncoding.DecodeString(jwk.X)
	if err != nil {
		return nil, fmt.Errorf("error decoding X: %w", err)
	}

	y, err := base64.RawURLEncoding.DecodeString(jwk.Y)
	if err != nil {
		return nil, fmt.Errorf("error decoding Y: %w", err)
	}

	publicKey := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     new(big.Int).SetBytes(x),
		Y:     new(big.Int).SetBytes(y),
	}

	return publicKey, nil
}

func JWKSHandler(jwksJSON []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(jwksJSON)
		if err != nil {
			http.Error(w, "Error writing JWKS JSON", http.StatusInternalServerError)
			return
		}
	})
}
