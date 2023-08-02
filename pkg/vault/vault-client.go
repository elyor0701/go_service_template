package vaultclient

import (
	"context"

	vault "github.com/hashicorp/vault/api"
)

type VaultClient interface {
	RenewToken(ctx context.Context) error
	Get(ctx context.Context, secretPath string) (map[string]interface{}, error)
	Put(ctx context.Context, secretPath string, data map[string]interface{}) error
	GetSecretPath(ctx context.Context) (string, error)
}

type vaultClient struct {
	client     *vault.Client
	roleID     string
	secretID   string
	mountPath  string
	secretPath string
}

func NewClient(ctx context.Context, args NewClientArgs) (VaultClient, error) {
	config := vault.DefaultConfig()
	config.Address = args.Address

	client, err := vault.NewClient(config)
	if err != nil {
		return vaultClient{}, err
	}

	return vaultClient{
		client:     client,
		roleID:     args.RoleID,
		secretID:   args.SecretID,
		mountPath:  args.MountPath,
		secretPath: args.SecretPath,
	}, nil
}

type NewClientArgs struct {
	Address    string
	RoleID     string
	SecretID   string
	MountPath  string
	SecretPath string
}
