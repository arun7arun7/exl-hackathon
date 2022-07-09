package entity

type File struct {
	ObjectId string
	TenantId string
	CloudType string
	FileExtension string
}

// func NewFile(objectId, fileExtension, tenantId, cloudType string) *File {
// 	return &File{
// 		objectId: objectId,
// 		tenantId: tenantId,
// 		cloudType: cloudType,
// 		fileExtension: fileExtension,
// 	}
// }

func (file *File) GetObjectId() string {
	return file.ObjectId
}

func (file *File) GetTenantId() string {
	return file.TenantId
}

func (file *File) GetFileExtension() string {
	return file.FileExtension
}

func (file *File) GetCloudType() string {
	return file.CloudType
}