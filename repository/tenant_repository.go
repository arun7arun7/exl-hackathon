package repository

import (
	"database/sql"
	"exl-server/entity"
	"log"
)

type TenantRepository interface {
	FindByAzureOrgId(orgId string) (*entity.AzureTenant, error)
	FindByAwsOrgId(orgId string) (*entity.AwsTenant, error)
}

type TenantRepositoryImpl struct {
	db *sql.DB
}

func NewTenantRepositoryImpl(db *sql.DB) *TenantRepositoryImpl {
	return &TenantRepositoryImpl{
		db: db,
	}
}

func (tenantRepository *TenantRepositoryImpl) FindByAzureOrgId(orgId string) (*entity.AzureTenant, error) {
	row, err := tenantRepository.db.Query("SELECT tenant_id, storage_account, container_name, client_id, client_secret from azure_tenant where organization_id = ?", orgId)
	if err != nil {
		log.Printf("error retrieving organization id from azure_tenant : %s", err)
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
}

func (tenantRepository *TenantRepositoryImpl) FindByAwsOrgId(orgId string) (*entity.AwsTenant, error) {
	row, err := tenantRepository.db.Query("SELECT aws_region, bucket_name, access_key_id, secret_access_key from aws_tenant where organization_id = ?", orgId)
	if err != nil {
		log.Printf("error retrieving organization id from aws_tenant : %s", err)
		return nil, err
	}
	defer row.Close()
	if row.Next() {
		log.Printf("Debug : %v\n", row)
		var result entity.AwsTenant
		row.Scan(&result.Storage.AwsRegion, &result.Storage.BucketName, &result.Credentials.AccessKeyId, &result.Credentials.SecretAccessKey)
		return &result, nil
	}
	return nil, nil
}