package repository

import (
	"database/sql"
	"exl-server/entity"
	// "fmt"
	"log"
)

type TenantRepository interface {
	FindByAzureTenantId(tenantId string) (*entity.AzureTenant, error)
	FindByAwsTenantId(tenantId string) (*entity.AwsTenant, error)
}

type TenantRepositoryImpl struct {
	db *sql.DB
}

func NewTenantRepositoryImpl(db *sql.DB) *TenantRepositoryImpl {
	return &TenantRepositoryImpl{
		db: db,
	}
}

func (tenantRepository *TenantRepositoryImpl) FindByAzureTenantId(tenantid string) (*entity.AzureTenant, error) {
	// query := fmt.Sprintf("SELECT tenant_id, storage_account, container_name, client_id, client_secret from azure_tenant where tenant_id = '%s'", tenantid)
	row, err := tenantRepository.db.Query("SELECT tenant_id, storage_account, container_name, client_id, client_secret from azure_tenant where tenant_id = ?", tenantid)
	if err != nil {
		log.Printf("error retrieving tenant id from azure_tenant : %s", err)
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		log.Printf("Debug : %v\n", row)
		var result entity.AzureTenant
		row.Scan(&result.TenantId, &result.Storage.StorageAccount, &result.Storage.ContainerName, &result.Credentials.ClientID, &result.Credentials.ClientSecret)
		return &result, nil
	}
	return nil, nil

	// return &entity.AzureTenant{
	// 	TenantId: tenantid,
	// 	Storage: entity.AzureStorage{
	// 		StorageAccount: "prac",
	// 		ContainerName: "exl",
	// 	},
	// 	Credentials: entity.AzureCredentials{
	// 		ClientID: "1ff1fc50-373f-4a9d-9553-631b58f7e97a",
	// 		ClientSecret: "_7u8Q~Pzz8-tt2DNeX_zvNYGZN~.efEoVJ~tgbfK",
	// 	},
	// }, nil
}

func (tenantRepositoryImpl *TenantRepositoryImpl) FindByAwsTenantId(tenantId string) (*entity.AwsTenant, error) {
	return &entity.AwsTenant{
		TenantId: tenantId,
		Storage: entity.AwsStorage{
			AwsRegion: "us-east-2",
			BucketName: "exl-storage",
		},
		Credentials: entity.AwsCredentials{
			AccessKeyId: "AKIA52EZ53Y4GJP6CEX2",
			SecretAccessKey: "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",
		},
	}, nil
}