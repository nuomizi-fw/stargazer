package keystore

import (
	"crypto/ecdsa"
	"sync"

	"github.com/nuomizi-fw/stargazer/pkg/jwt"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"keystore",
	fx.Provide(
		NewKeyStore,
	),
)

type KeyStore struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	mu         sync.RWMutex
}

// NewKeyStore creates a new in-memory KeyStore instance
func NewKeyStore() (*KeyStore, error) {
	ks := &KeyStore{}
	if err := ks.initialize(); err != nil {
		return nil, err
	}
	return ks, nil
}

func (ks *KeyStore) initialize() error {
	ks.mu.Lock()
	defer ks.mu.Unlock()

	// Generate new key pair
	privateKey, publicKey, err := jwt.GenerateKeyPair()
	if err != nil {
		return err
	}

	ks.privateKey = privateKey
	ks.publicKey = publicKey
	return nil
}

func (ks *KeyStore) GetKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	ks.mu.RLock()
	defer ks.mu.RUnlock()
	return ks.privateKey, ks.publicKey
}
