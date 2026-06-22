package sandbox

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"os"
	"path/filepath"
)

type FileMetadata struct {
	FileName  string
	FileSize  int64
	Extension string
	MIMEType  string
	MD5       string
}

func GetFileMetadata(
	filePath string,
) FileMetadata {

	info, _ :=
		os.Stat(
			filePath,
		)

	data, _ :=
		os.ReadFile(
			filePath,
		)

	md5Hash :=
		md5.Sum(
			data,
		)

	mimeType :=
		http.DetectContentType(
			data,
		)

	return FileMetadata{
		FileName: filepath.Base(
			filePath,
		),
		FileSize: info.Size(),
		Extension: filepath.Ext(
			filePath,
		),
		MIMEType: mimeType,
		MD5: hex.EncodeToString(
			md5Hash[:],
		),
	}
}