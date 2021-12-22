package config

import (
	"github.com/JasperStritzke/cubid-cloud/util/fileutil"
	"os"
)

func WrapExistingConfig(cfg interface{}) func() interface{} {
	return func() interface{} {
		return cfg
	}
}

func InitConfigIfNotExists(path string, codeType string, config interface{}) error {
	file, err, created := fileutil.OpenFileOrCreate(path)

	if err != nil {
		return err
	}

	if created {
		encoder := fileutil.NewPrettyEncoder(file, codeType)

		encodeErr := encoder.Encode(config)

		if encodeErr != nil {
			return encodeErr
		}

		closeErr := file.Close()
		if closeErr != nil {
			return closeErr
		}

		return nil
	}

	return nil
}

func WriteConfig(path string, codeType string, cfg interface{}) error {
	bytes, err := fileutil.MarshalIndent(cfg, codeType)

	if err != nil {
		return err
	}

	return os.WriteFile(path, bytes, os.ModePerm)
}

func LoadConfig(path string, codeType string, config interface{}) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	decoder := fileutil.NewDecoder(file, codeType)

	decodeErr := decoder.Decode(config)

	if decodeErr != nil {
		return nil
	}

	closeErr := file.Close()
	if closeErr != nil {
		return closeErr
	}

	return nil
}
