package repository

import (
	"database/sql"
	"exl-server/entity"
)

type TenantRepository interface {
	FindByAzureTenantId(tenantId string) (*entity.AzureTenant, error)
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
	query := "SELECT tenant_id, storage_account, container_name, client_id, client_secret from azure_tenant where tenant_id = ?"
	row := tenantRepository.db.QueryRow(query, tenantid)
	if row.Scan() == sql.ErrNoRows {
		return nil, nil
	}
	var result *entity.AzureTenant
	row.Scan(&result.TenantId, &result.Storage.StorageAccount, &result.Storage.ContainerName, &result.Credentials.ClientID, &result.Credentials.ClientSecret)
	return result, nil

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
