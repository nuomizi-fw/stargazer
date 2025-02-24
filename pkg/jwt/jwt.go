package jwt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenTypeAccess  = "stargazer-access"
	TokenTypeRefresh = "stargazer-refresh"

	AccessTokenDuration  = 3 * time.Hour
	RefreshTokenDuration = 7 * 24 * time.Hour

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
	Username string `json:"username"`
	ID       int    `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(username string, privateKey *ecdsa.PrivateKey, id int, issuer string, t time.Duration) (string, error) {
	claims := Claims{
		Username: username,
		ID:       id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}
	return signedToken, nil
}

func GetAccessToken(username string, privateKey *ecdsa.PrivateKey, id int, t time.Duration) (string, error) {
	return GenerateToken(username, privateKey, id, TokenTypeAccess, t)
}

func GetRefreshToken(username string, privateKey *ecdsa.PrivateKey, id int, t time.Duration) (string, error) {
	return GenerateToken(username, privateKey, id, TokenTypeRefresh, t)
}

func ParseToken(signedToken string, publicKeyFunc func() (*ecdsa.PublicKey, error)) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
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

	// Validate expiration
	if claims.ExpiresAt == nil || claims.ExpiresAt.Before(time.Now()) {
		return false, nil, errors.New("token has expired")
	}

	return true, claims, nil
}

func GenerateKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating key pair: %w", err)
	}
	return privateKey, &privateKey.PublicKey, nil
}

func GenerateJwksJSON(publicKey *ecdsa.PublicKey) JWKS {
	jwk := JWK{
		Kty: "EC",
		Crv: "P-256",
		X:   base64.RawURLEncoding.EncodeToString(publicKey.X.Bytes()),
		Y:   base64.RawURLEncoding.EncodeToString(publicKey.Y.Bytes()),
	}

	return JWKS{
		Keys: []JWK{jwk},
	}
}
