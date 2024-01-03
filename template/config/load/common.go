package load

import (
	"gopkg.in/yaml.v3"
	"os"
)

func loadFile(path string) ([]byte, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, newLoadFile(err.Error())
	}

	return content, nil
}

func unmarshal(content []byte, dst any) error {
	err := yaml.Unmarshal(content, dst)
	if err != nil {
		return newUnmarshalError(err.Error())
	}

	return nil
}
