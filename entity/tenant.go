package entity

import "exl-server/constants"

type accessStorage[T interface{}] interface {
	getAccessStorage() T
}

type accessToken[T interface{}] interface {
	getAccessToken() T
}

type GetCloudType interface {
	getCloudType() constants.CloudType
}

type AbstractTenant[AS interface{}, AT interface{}] struct {
	id        string
	cloudType constants.CloudType
	accessStorage[AS]
	accessToken[AT]
}

func (t *AbstractTenant[AS, AT]) GetId() string {
	return t.id
}

type AzureTenant struct {
	TenantId 	string
	Storage     AzureStorage
	Credentials AzureCredentials
}

// func (azureTenant *AzureTenant) GetId() string {
// 	return azureTenant.Id
// }

type AzureStorage struct {
	StorageAccount string
	ContainerName string
}

type AzureCredentials struct {
	ClientID string
	ClientSecret string
}

type AwsTenant struct {
	TenantId string
	Storage AwsStorage
	Credentials AwsCredentials
}

type AwsStorage struct {
	AwsRegion string
	BucketName string
}

type AwsCredentials struct {
	AccessKeyId string
	SecretAccessKey string
}
