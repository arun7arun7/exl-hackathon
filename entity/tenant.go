package entity

type AzureTenant struct {
	TenantId 	string
	Storage     AzureStorage
	Credentials AzureCredentials
}

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
