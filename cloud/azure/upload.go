package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

// type FileUploadSync func(data []byte) error

const (
	url = "https://storageblobsgo.blob.core.windows.net/"
)

func FileUploadSync(ctx context.Context, credential azcore.TokenCredential, containerName string, blobName string, data []byte) error {
	blobClient, err := azblob.NewBlockBlobClient(url+containerName+"/"+blobName, credential, nil)
	if err != nil {
		log.Println("Failed to create blob client")
		return err
	}
	_, err = blobClient.UploadBuffer(ctx, data, azblob.UploadOption{})
	if err != nil {
		log.Println("Failure to upload blob")
		return err
	}
	return nil
}