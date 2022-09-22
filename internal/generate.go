package internal

import (
	"os"

	"gopkg.in/yaml.v3"
)

func GenerateEmptyYML(path string) error {
	content := []byte("---\n")
	return os.WriteFile(path, content, 0644)
}

func GenerateSettings(settings JobSettings, path string) error {

	if err := GenerateEmptyYML(path); err != nil {
		return err
	}

	ymlData, err := yaml.Marshal(settings)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, ymlData, 0755)
	if err != nil {
		return err
	}

	return nil
}
