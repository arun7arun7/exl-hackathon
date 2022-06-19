package constants

type CloudType string

const (
	AZURE CloudType = "AZURE"
	AWS CloudType = "AWS"
	GCP CloudType = "GCP"
)

var cloudTypeMap = map[string]CloudType {
	"AZURE" : AZURE,
	"AWS" : AWS,
	"GCP" : GCP,
}

func GetCloudType(cloudType string) CloudType {
	if res, present := cloudTypeMap[cloudType]; present {
		return res
	}
	return CloudType("")
}