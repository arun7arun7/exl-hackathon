package constants

type FileStatus string

const (
	CREATED FileStatus = "CREATED"
	ENQUEUED FileStatus = "ENQUEUED"
	UPLOADED FileStatus = "UPLOADED"
)