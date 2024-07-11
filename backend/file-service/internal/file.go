package file_upload

import (
	"encoding/base64"

	"github.com/google/uuid"
)

func DecodeUUID(base64UUID string) (uuid.UUID, error) {
	// Decode the Base64 string into bytes
	uuidBytes, err := base64.StdEncoding.DecodeString(base64UUID)
	if err != nil {
		return uuid.Nil, err
	}

	// Create a UUID from the decoded bytes
	u, err := uuid.FromBytes(uuidBytes)
	if err != nil {
		return uuid.Nil, err
	}

	return u, nil
}
