package service

import (
	"context"
	"errors"
	"exl-server/cloud"
	"exl-server/constants"
	"exl-server/dto"
	"exl-server/repository"
	"io"
	"io/ioutil"
	"log"

	"github.com/google/uuid"
)

var (
	ErrTenantIdNotFound = errors.New("tenant Id is not found")
	ErrReadingBody      = errors.New("error reading body")
	ErrObjectIdNotFound = errors.New("object id is not found")
	ErrUnknown = errors.New("unknown error occurred")
)

type FileService interface {
	UploadSync(parentCtx context.Context, orgId string, cloudType constants.CloudType, body io.ReadCloser, fileExtension string) (string, error)
	DownloadSync(parentCtx context.Context, objectId string) (io.ReadCloser, *dto.FileMetadata, error)
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

func (fs *FileServiceImpl) UploadSync(parentCtx context.Context, orgId string, cloudType constants.CloudType, body io.ReadCloser, fileExtension string) (string, error) {
	switch cloudType {
	case constants.AZURE:
		azureTenant, err := fs.tenantService.GetByAzureOrgID(orgId)
		if err != nil {
			log.Println("Error fetching AzureTenant ID")
			return "", err
		}
		if azureTenant == nil {
			log.Println("Tenant ID is not found")
			return "", ErrTenantIdNotFound
		}
		log.Printf("AzureTenant : %v", azureTenant)

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
		err = fs.fileRepository.Create(fileUniqueIdentifier, fileExtension, orgId, string(cloudType))
		if err != nil {
			log.Printf("Error creating file entry %s", err)
		}
		log.Printf("Successfully uploaded %s", fileId)
		return fileUniqueIdentifier, nil

	case constants.AWS:
		log.Printf("Inside AWS")
		awsTenant, err := fs.tenantService.GetByAwsOrgID(orgId)
		if err != nil {
			log.Println("Error fetching AwsTenant ID")
			return "", err
		}
		if awsTenant == nil {
			log.Println("Tenant ID is not found")
			return "", ErrTenantIdNotFound
		}
		fileUniqueIdentifier := uuid.New().String()
		fileId := fileUniqueIdentifier + fileExtension
		log.Printf("FileId: %s", fileId)
		
		awsService := cloud.NewAwsService(awsTenant.Storage.AwsRegion, awsTenant.Storage.BucketName, awsTenant.Credentials.AccessKeyId, awsTenant.Credentials.SecretAccessKey)
		err = awsService.FileUploadSync(parentCtx, fileId, body)
		if err != nil {
			log.Println("Error Uploading data")
			return "", err
		}
		err = fs.fileRepository.Create(fileUniqueIdentifier, fileExtension, orgId, string(cloudType))
		if err != nil {
			log.Printf("Error creating file entry %s", err)
		}
		log.Printf("Successfully uploaded %s", fileId)
		return fileUniqueIdentifier, nil
	default:
		return "", ErrUnknown
	}
}

func (fs *FileServiceImpl) DownloadSync(parentCtx context.Context, objectId string) (io.ReadCloser, *dto.FileMetadata, error) {
	file, err := fs.fileRepository.Get(objectId)
	if err != nil {
		log.Printf("Error fetching file %s\n", err)
		return nil, nil, err
	}
	if file == nil {
		log.Printf("Object Id is not found\n")
		return nil, nil, ErrObjectIdNotFound
	}

	log.Printf("File : %v\n", file)
	log.Printf("ObjectId: %s\n FileExtension: %s\n Cloud: %s\n", file.ObjectId, file.FileExtension, file.CloudType)
	cloudType := constants.GetCloudType(file.CloudType)
	if cloudType == constants.CloudType("") {
		log.Printf("CloudType Mismatch\n")
		return nil, nil, ErrUnknown
	}
	switch cloudType {
		case constants.AZURE:
			azureTenant, err := fs.tenantService.GetByAzureOrgID(file.OrganizationId)
			if err != nil {
				log.Printf("Error fetching AzureTenant ID\n")
				return nil, nil, err
			}
			if azureTenant == nil {
				log.Printf("Tenant ID is not found\n")
				return nil, nil, ErrUnknown
			}
			blobName := file.ObjectId + file.FileExtension
			azureService := cloud.NewAzureService(azureTenant.Storage.StorageAccount, azureTenant.Storage.ContainerName, azureTenant.TenantId, azureTenant.Credentials.ClientID, azureTenant.Credentials.ClientSecret)
			body, err := azureService.FileDownloadSync(parentCtx, blobName)
			if err != nil {
				log.Printf("Error in downloading blob\n")
				return nil, nil, err
			}
			log.Printf("Successfully downloaded %s\n", blobName)
			metadata := &dto.FileMetadata{
				Name: blobName,
				FileExtension: file.FileExtension,
			}
			return body, metadata, nil

		case constants.AWS:
			awsTenant, err := fs.tenantService.GetByAwsOrgID(file.OrganizationId)
			if err != nil {
				log.Printf("Error fetching AwsTenant ID\n")
				return nil, nil, err
			}
			if awsTenant == nil {
				log.Printf("Tenant ID is not found\n")
				return nil, nil, ErrUnknown
			}
			fileName := file.ObjectId + file.FileExtension
			awsService := cloud.NewAwsService(awsTenant.Storage.AwsRegion, awsTenant.Storage.BucketName, awsTenant.Credentials.AccessKeyId, awsTenant.Credentials.SecretAccessKey)
			body, err := awsService.FileDownloadSync(parentCtx, fileName)
			if err != nil {
				log.Printf("Error in downloading file\n")
				return nil, nil, err
			}
			log.Printf("Successfully downloaded %s\n", fileName)
			metadata := &dto.FileMetadata{
				Name: fileName,
				FileExtension: file.FileExtension,
			}
			return body, metadata, nil
		default:
			return nil, nil, ErrUnknown
	}
}
