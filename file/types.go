package file

// UploadedFile is a struct used to save information about an uploaded file.
type UploadedFile struct {
	FileSize         int64
	NewFileName      string
	OriginalFileName string
}
