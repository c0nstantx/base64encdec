package encoder

import (
	"encoding/base64"
)

func Encode(fileContents []byte) string {
	return base64.StdEncoding.EncodeToString(fileContents)
}
