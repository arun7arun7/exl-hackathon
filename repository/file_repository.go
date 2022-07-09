package repository

import (
	"database/sql"
	"exl-server/entity"
	"log"
)

type FileRepository interface {
	Create(objectId, fileExtension, tenantId, cloudType string) error
	Get(objectId string) (*entity.File, error)
}

type FileRepositoryImpl struct {
	db *sql.DB
}

func NewFileRepositoryImpl(db *sql.DB) *FileRepositoryImpl {
	return &FileRepositoryImpl{
		db: db,
	}
}

func (fileRepository *FileRepositoryImpl) Create(objectId, fileExtension, tenantId, cloudType string) error {
	query := "INSERT INTO files(object_id, file_extension, tenant_id, cloud_type) VALUES(?, ?, ?, ?)" 						
	log.Printf("Inserting into files table\n")
	_, err := fileRepository.db.Exec(query, objectId, fileExtension, tenantId, cloudType)
	if err != nil {
		return err	
	}
	log.Printf("Created objectId %s fileExtension %s for tenantid %s", objectId, fileExtension, tenantId)
	return nil
}

func (fileRepository *FileRepositoryImpl) Get(objectId string) (*entity.File, error) {
	// query := "SELECT object_id, file_extension, tenant_id, cloud_type FROM files where object_id=?"
	row, err := fileRepository.db.Query("SELECT object_id, file_extension, tenant_id, cloud_type FROM files where object_id= ?", objectId)
	if err != nil {
		log.Printf("error retrieving object id: %s", err)
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		log.Printf("Debug : %v\n", row)
		var file entity.File
		row.Scan(&file.ObjectId, &file.FileExtension, &file.TenantId, &file.CloudType)
		return &file, nil
	}
	return nil, nil
	// file := entity.NewFile(objectId, "de18e1ee-5536-4959-961c-bcfb59c93e26", "AZURE", ".pdf")
	// return file, nil
}