package executor

import (
	"github.com/JasperStritzke/cubid-cloud/util/config"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/fileutil"
	"os"
	"path"
	"strings"
)

const (
	executorFolder = "executors/"
	executorSuffix = ".data.json"
)

type Manager struct {
	executors []*Executor
}

func NewManager() *Manager {
	manager := &Manager{}

	_ = os.Mkdir(executorFolder, os.ModePerm)

	files, err := os.ReadDir(executorFolder)

	if err != nil {
		return nil
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), executorSuffix) {
			continue
		}

		var executorInfo Info
		err = config.LoadConfig(path.Join(executorFolder, file.Name()), fileutil.CodingJSON, &executorInfo)

		if err != nil {
			logger.Errorf("Unable to load executor configuration from %s.", file.Name())
			continue
		}

		manager.executors = append(
			manager.executors,
			&Executor{
				Info:     executorInfo,
				endpoint: nil,
			},
		)
	}

	if len(manager.executors) == 0 {
		logger.Warn("No executor configured.")
	}

	return manager
}
