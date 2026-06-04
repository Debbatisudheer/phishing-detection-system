package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func CalculateSHA256(
	filePath string,
) string {

	data, err :=
		os.ReadFile(
			filePath,
		)

	if err != nil {
		return ""
	}

	hash :=
		sha256.Sum256(
			data,
		)

	return hex.EncodeToString(
		hash[:],
	)
}