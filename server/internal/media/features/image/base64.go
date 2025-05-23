package image

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
)

func (s *service) Base64(img image.Image) (string, error) {

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return "", fmt.Errorf("encode: %v", err)
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
