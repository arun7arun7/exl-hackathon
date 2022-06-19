package repository

import "log"

type FileRepository interface {
	Create(fileId, fileExtension, tenantId string) error
}

type FileRepositoryImpl struct {
}

func (fileRepository *FileRepositoryImpl) Create(fileId, fileExtension, tenantId string) error {
	log.Printf("Created fileId %s fileExtension %s for tenantid %s", fileId, fileExtension, tenantId)
	return nil
}