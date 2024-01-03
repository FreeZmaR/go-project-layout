package load

import "encoding/base64"

func Production(path string, dst any) error {
	encContent, err := loadFile(path)
	if err != nil {
		return err
	}

	decContent, err := decodeContent(encContent)
	if err != nil {
		return err
	}

	return unmarshal(decContent, dst)
}

// decodeContent - simple example for decode security content
func decodeContent(encContent []byte) ([]byte, error) {
	decContent, err := base64.StdEncoding.DecodeString(string(encContent))
	if err != nil {
		return nil, newDecodeError(err.Error())
	}

	return decContent, nil
}
