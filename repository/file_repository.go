package repository

import (
	"exl-server/entity"
	"log"
)

type FileRepository interface {
	Create(fileId, fileExtension, tenantId string) error
	Get(objectId string) (*entity.File, error)
}

type FileRepositoryImpl struct {
}

func (fileRepository *FileRepositoryImpl) Create(fileId, fileExtension, tenantId string) error {
	log.Printf("Created fileId %s fileExtension %s for tenantid %s", fileId, fileExtension, tenantId)
	return nil
}

func (fileRepository *FileRepositoryImpl) Get(objectId string) (*entity.File, error) {
	file := entity.NewFile(objectId, "de18e1ee-5536-4959-961c-bcfb59c93e26", "AZURE", ".pdf")
	return file, nil
}