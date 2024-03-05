package shared

type EndPoint string

const (
	// Case Service
	ValidateCaseSubject   EndPoint = "ValidateCase"
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
