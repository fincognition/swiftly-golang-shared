package shared

type EndPoint string

const (
	// Case Service
	ValidateCaseSubject   EndPoint = "ValidateCase"
	WriteCaseFileSubject  EndPoint = "WriteCaseFile"
	WriteDocumentsSubject EndPoint = "WriteDocuments"
	// File Service
	ProcessFilesSubject EndPoint = "ProcessFiles"
	CleanupFilesSubject EndPoint = "CleanupFiles"
	// Content Classifier
	ClassifyPageSubject EndPoint = "ClassifyPage"
	// Model Service
	ClassifyImageSubject EndPoint = "ClassifyImage"
	ExtractPageSubject   EndPoint = "ExtractPage"
)
