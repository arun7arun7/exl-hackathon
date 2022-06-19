package cloud

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

type AzureService struct {
	storageAccount string
	containerName string
	tenantId string
	clientID string
	clientSecret string
}

const (
	url = "https://%s.blob.core.windows.net/%s/%s"
)

func NewAzureService(storageAccount, containerName, tenantId, clientID, clientSecret string) *AzureService {
	return &AzureService {
		storageAccount: storageAccount,
		containerName: containerName,
		tenantId: tenantId,
		clientID: clientID,
		clientSecret: clientSecret,
	}
}

// func (azureService *AzureService) Authenticate() (azcore.TokenCredential, error) {
// 	options := &azidentity.ManagedIdentityCredentialOptions{
// 		ID: azidentity.ClientID(azureService.userManagedClientId),
// 	}
// 	cred, err := azidentity.NewManagedIdentityCredential(options)
// 	if err != nil {
// 		log.Println("Authentication failure")
// 		return nil, err
// 	}
// 	return cred, nil
// }

func (azureService *AzureService) Authenticate() (azcore.TokenCredential, error) {
	cred, err := azidentity.NewClientSecretCredential(azureService.tenantId, azureService.clientID, azureService.clientSecret, nil)
	if err != nil {
		log.Printf("Error Authenticating %s \n", err)
		return nil, err
	}
	return cred, nil
}

func (azureService *AzureService) FileUploadSync(ctx context.Context, blobName string, data []byte) error {
	cred, err := azureService.Authenticate()
	if err != nil {
		log.Println("Error Authenticating")
		return err
	}
	blobUrl := fmt.Sprintf(url, azureService.storageAccount, azureService.containerName, blobName)
	log.Printf("BlobURL: %s", blobUrl)
	blobClient, err := azblob.NewBlockBlobClient(blobUrl, cred, nil)
	if err != nil {
		log.Printf("Failed to create blob client: %s \n", err)
		return err
	}
	_, err = blobClient.UploadBuffer(ctx, data, azblob.UploadOption{})
	if err != nil {
		log.Printf("Failure to upload blob: %s \n", err)
		return err
	}
	return nil
}