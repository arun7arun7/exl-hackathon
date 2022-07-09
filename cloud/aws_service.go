package cloud

import (
	"context"
	"io"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AwsService struct {
	awsRegion string
	bucketName string
	accessKeyId string
	secretAccessKey string
}

func NewAwsService(awsRegion, bucketName, accessKeyId, secretAccessKey string) (*AwsService) {
	return &AwsService{
		awsRegion: awsRegion,
		bucketName: bucketName,
		accessKeyId: accessKeyId,
		secretAccessKey: secretAccessKey,
	}
}

func (awsService *AwsService) Authenticate() (*session.Session, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(awsService.awsRegion),
			Credentials: credentials.NewStaticCredentials(
				awsService.accessKeyId,
				awsService.secretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})
	if err != nil {
		log.Printf("error Authenticate AWS: %s\n", err)
		return nil, err
	}
	return sess, nil
}

func (awsService *AwsService) FileUploadSync(ctx context.Context, fileName string, data io.Reader) error {
	sess, err := awsService.Authenticate()
	if err != nil {
		log.Printf("error Authenticating")
		return err
	}
	uploader := s3manager.NewUploader(sess)
	log.Printf("uploading...\n")
	// Upload the file to S3.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(awsService.bucketName),
		Key:    aws.String(fileName),
		Body:   data,
	})
	if err != nil {
		log.Printf("error uploading data : %s", err)
		return err
	}
	return nil
}

func (awsService *AwsService) FileDownloadSync(ctx context.Context, fileName string) (io.ReadCloser, error) {
	sess, err := awsService.Authenticate()
	if err != nil {
		log.Printf("error Authenticating")
		return nil, err
	}
	downloader := s3manager.NewDownloader(sess)
	file, err := ioutil.TempFile("", "*")
	if err != nil {
		log.Printf("error creating temporary file for download: %s \n", err)
		return nil, err
	}
	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(awsService.bucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		log.Printf("failed to download file, %s", err)
		return nil, err
	}
	return file, nil
}
