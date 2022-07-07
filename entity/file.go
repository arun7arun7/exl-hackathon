package entity

type File struct {
	objectId string
	tenantId string
	cloudType string
	fileExtension string
}

func NewFile(objectId, tenantId, cloudType, fileExtension string) *File {
	return &File{
		objectId: objectId,
		tenantId: tenantId,
		cloudType: cloudType,
		fileExtension: fileExtension,
	}
}

func (file *File) GetObjectId() string {
	return file.objectId
}

func (file *File) GetTenantId() string {
	return file.tenantId
}

func (file *File) GetFileExtension() string {
	return file.fileExtension
}

func (file *File) GetCloudType() string {
	return file.cloudType
}