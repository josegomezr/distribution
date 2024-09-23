package s3

import (
	storagedriver "github.com/distribution/distribution/v3/registry/storage/driver"
	"time"
)

type s3FileInfoFieldsContainer struct {
	storagedriver.FileInfoFields
	prefix *string
}

// Path provides the full path of the target of this file info.
func (wi s3FileInfoFieldsContainer) Path() string {
	return wi.FileInfoFields.Path
}

// Size returns current length in bytes of the file. The return value can
// be used to write to the end of the file at path. The value is
// meaningless if IsDir returns true.
func (wi s3FileInfoFieldsContainer) Size() int64 {
	return wi.FileInfoFields.Size
}

// ModTime returns the modification time for the file. For backends that
// don't have a modification time, the creation time should be returned.
func (wi s3FileInfoFieldsContainer) ModTime() time.Time {
	return wi.FileInfoFields.ModTime
}

// IsDir returns true if the path is a directory.
func (wi s3FileInfoFieldsContainer) IsDir() bool {
	return wi.FileInfoFields.IsDir
}
