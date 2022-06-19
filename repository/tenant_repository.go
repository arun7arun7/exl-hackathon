package repository

import "exl-server/entity"

type TenantRepository interface {
	FindByAzureTenantId(tenantId string) (*entity.AzureTenant, error)
}

type TenantRepositoryImpl struct {
}

func (tenantRepository *TenantRepositoryImpl) FindByAzureTenantId(tenantid string) (*entity.AzureTenant, error) {
	return &entity.AzureTenant{
		TenantId: tenantid,
		Storage: entity.AzureStorage{
			StorageAccount: "prac",
			ContainerName: "exl",
		},
		Credentials: entity.AzureCredentials{
			ClientID: "1ff1fc50-373f-4a9d-9553-631b58f7e97a",
			ClientSecret: "_7u8Q~Pzz8-tt2DNeX_zvNYGZN~.efEoVJ~tgbfK",
		},
	}, nil
}
