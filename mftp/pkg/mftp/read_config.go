package mftp

import (
	"encoding/json"
	"mihaimiuta/mftp/pkg/mftp/types"
	"os"
	"path/filepath"
)

func readConfig(filePath string) (types.Config, error) {
	absolutePath, err := filepath.Abs(filePath)

	if err != nil {
		return types.Config{}, err
	}

	config := types.Config{
		LogFormat: "",
		Driver:    nil,
	}

	data, err := os.ReadFile(absolutePath)

	if err != nil {
		return types.Config{}, err
	}

	unmarshallError := json.Unmarshal(data, &config)

	if unmarshallError != nil {
		return types.Config{}, err
	}

	return config, nil
}
