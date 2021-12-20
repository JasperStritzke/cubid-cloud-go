package template

import (
	"errors"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"os"
)

func InitEmptyTemplate(value VersionValue, serverGroup, templateName string) error {
	logger.Warnf("Creating template %s/%s ...", serverGroup, templateName)

	dirName := "templates/" + serverGroup + "/" + templateName + "/"
	if _, err := os.Stat(dirName); err != nil {
		logger.Error("Failed to create template: Template already exists.")
		return err
	}

	_ = os.MkdirAll(dirName, os.ModePerm)

	dir, err := os.Stat(dirName)

	if !dir.IsDir() {
		return errors.New(dir.Name() + " is not a directory")
	}

	err = value.DownloadToPath(dirName)

	if err != nil {
		logger.Errorf("Unable to download version to %s", dirName)
		return err
	}

	if value.IsProxy() {
		return completeProxy(value, dirName)
	}

	return completeServer(value, dirName)
}

func completeProxy(value VersionValue, dirName string) error {

	var err error
	return err
}

func completeServer(value VersionValue, dirName string) error {
	var err error
	return err
}
