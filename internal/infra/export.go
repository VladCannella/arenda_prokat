package infra

import (
	"encoding/json"
	"os"
)

func ExportJson(path string, data any) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := file.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}
