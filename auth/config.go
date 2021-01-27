package auth

import "github.com/manicminer/hamilton/environments"

type TokenVersion int

const (
	TokenVersion2 TokenVersion = iota
	TokenVersion1
)

type Config struct {
	// Specifies the national cloud environment to use
	Environment environments.Environment

	// Version specifies the token version  to acquire from Microsoft Identity Platform.
	// Ignored when using Azure CLI authentication.
	Version TokenVersion

	// Azure Active Directory tenant to connect to, should be a valid UUID
	TenantID string

	// Client ID for the application used to authenticate the connection
	ClientID string

	// Enables authentication using Azure CLI
	EnableAzureCliToken bool

	// Enables authentication using managed service identity.
	EnableMsiAuth bool

	// Specifies a custom MSI endpoint to connect to
	MsiEndpoint string

	// Enables client certificate authentication using client assertions
	EnableClientCertAuth bool

	// Specifies the path to a client certificate bundle in PFX format
	ClientCertPath string

	// Specifies the encryption password to unlock a client certificate
	ClientCertPassword string

	// Enables client secret authentication using client credentials
	EnableClientSecretAuth bool

	// Specifies the password to authenticate with using client secret authentication
	ClientSecret string
}
