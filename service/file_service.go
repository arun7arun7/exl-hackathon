package service

import (
	"context"
	"errors"
	"exl-server/cloud"
	"exl-server/constants"
	"exl-server/repository"
	"io"
	"io/ioutil"
	"log"

	"github.com/google/uuid"
)

var (
	ErrTenantIdNotFound = errors.New("tenant Id is not found")
	ErrReadingBody      = errors.New("error reading body")
)

type FileService interface {
	UploadSync(parentCtx context.Context, tenantId string, cloudType constants.CloudType, body io.ReadCloser, fileExtension string) (string, error)
}

type FileServiceImpl struct {
	fileRepository repository.FileRepository
	tenantService  TenantService
}

func NewFileServiceImpl(fp repository.FileRepository, ts TenantService) *FileServiceImpl {
	return &FileServiceImpl{
		fileRepository: fp,
		tenantService:  ts,
	}
}

func (fs *FileServiceImpl) UploadSync(parentCtx context.Context, tenantId string, cloudType constants.CloudType, body io.ReadCloser, fileExtension string) (string, error) {
	switch cloudType {
	case constants.AZURE:
		azureTenant, err := fs.tenantService.GetByAzureTenantID(tenantId)
		if err != nil {
			log.Println("Error fetching AzureTenant ID")
			return "", err
		}
		if azureTenant == nil {
			log.Println("Tenant ID is not found")
			return "", ErrTenantIdNotFound
		}
		fileUniqueIdentifier := uuid.New().String()
		fileId := fileUniqueIdentifier + fileExtension

		log.Printf("FileId: %s", fileId)
		data, err := ioutil.ReadAll(body)
		if err != nil {
			log.Println("Error reading the body")
			return "", ErrReadingBody
		}
		azureService := cloud.NewAzureService(azureTenant.Storage.StorageAccount, azureTenant.Storage.ContainerName, azureTenant.TenantId, azureTenant.Credentials.ClientID, azureTenant.Credentials.ClientSecret)
		err = azureService.FileUploadSync(parentCtx, fileId, data)
		if err != nil {
			log.Println("Error Uploading data")
			return "", err
		}
		err = fs.fileRepository.Create(fileUniqueIdentifier, fileExtension, tenantId)
		if err != nil {
			log.Printf("Error creating file entry %s", err)
		}
		log.Printf("Successfully uploaded %s", fileId)
		return fileUniqueIdentifier, nil
	}
	return "", nil
}
