package constants

var ContentTypeToFileExtension  = map[string]string {
		"application/pdf" : ".pdf",
		"text/csv": ".csv",
}

var FileExtensionToContentType = map[string]string {
	".pdf" : "application/pdf",
	".csv" : "text/csv",
}

