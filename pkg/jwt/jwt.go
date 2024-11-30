package jwt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
	"github.com/golang-jwt/jwt/v5"
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

const JWKSPath = "/.well-known/jwks.json"

type Claims struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, privateKey *ecdsa.PrivateKey, id int, issuer string, t time.Duration) (string, error) {
	claims := Claims{
		username,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(privateKey)
	return signedToken, err
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

// get AccessToken
func GetAccessToken(username string, privateKey *ecdsa.PrivateKey, id int) (string, error) {
	return GenerateToken(username, privateKey, id, "casaos", 3*time.Hour)
}

func GetRefreshToken(username string, private *ecdsa.PrivateKey, id int) (string, error) {
	return GenerateToken(username, private, id, "refresh", 7*24*time.Hour)
}

func Validate(token string, publicKeyFunc func() (*ecdsa.PublicKey, error)) (bool, *Claims, error) {
	claims, err := ParseToken(token, publicKeyFunc)
	if err != nil {
		return false, nil, err
	}

	if claims != nil {
		return true, claims, nil
	}

	return false, nil, errors.New("invalid token")
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

	return sonic.Marshal(jwks)
}

func PublicKeyFromJwksJSON(jwksJSON []byte) (*ecdsa.PublicKey, error) {
	var jwks JWKS
	err := sonic.Unmarshal(jwksJSON, &jwks)
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
