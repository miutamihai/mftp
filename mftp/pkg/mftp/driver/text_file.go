package driver

import (
	"errors"
	"io/fs"
	"mihaimiuta/mftp/pkg/mftp/types"
	"os"
	"path/filepath"
)

type TextFileDriver struct {
	FilePath string
}

func (driver *TextFileDriver) Write(logs []types.Log, encodeLog LogEncoder) error {
	absolutePath, err := filepath.Abs(driver.FilePath)

	if err != nil {
		return err
	}

	file, err := os.OpenFile(absolutePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			file, err = os.Create(absolutePath)

			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	for _, log := range logs {
		_, err := file.WriteString(encodeLog(log))

		if err != nil {
			return err
		}

	}

	fileCloseError := file.Close()

	if fileCloseError != nil {
		return fileCloseError
	}

	return nil
}

func (driver *TextFileDriver) GetBufferSize() int {
	return 10
}

func (driver *TextFileDriver) SupportsANSIColors() bool {
	return false
}
