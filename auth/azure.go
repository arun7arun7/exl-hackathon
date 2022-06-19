package auth

import (
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func Authenticate(clientID string) (azcore.TokenCredential, error) {
	options := &azidentity.ManagedIdentityCredentialOptions{
		ID: azidentity.ClientID(clientID),
	}
	cred, err := azidentity.NewManagedIdentityCredential(options)
	if err != nil {
		log.Println("Authentication failure")
		return nil, err
	}
	return cred, nil
}