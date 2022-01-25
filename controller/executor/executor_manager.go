package executor

import (
	"github.com/AlexanderGrom/go-event"
	"github.com/JasperStritzke/cubid-cloud/controller/event_names"
	"github.com/JasperStritzke/cubid-cloud/util/config"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/fileutil"
	"github.com/JasperStritzke/cubid-cloud/util/random"
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

func (m *Manager) Executors() []*Executor {
	return m.executors
}

type simpleExecutorInfo struct {
	Name        string `json:"name"`
	MaxMemory   int    `json:"maxMemory"`
	MaxCpuUsage int    `json:"maxCpuUsage"`
}

func (m *Manager) ExecutorInfos() []simpleExecutorInfo {
	var names []simpleExecutorInfo

	for _, executor := range m.executors {
		names = append(names, simpleExecutorInfo{
			Name:        executor.Info.Name,
			MaxMemory:   executor.Info.MaxMemory,
			MaxCpuUsage: executor.Info.MaxCPUUsage,
		})
	}

	return names
}

func (m *Manager) GetExecutor(name string) (int, *Executor) {
	for i, executor := range m.Executors() {
		if executor.Info.Name == name {
			return i, executor
		}
	}

	return -1, nil
}

func (m *Manager) CreateExecutor(name string, maxCpuUsage, maxMemory int) bool {
	index, _ := m.GetExecutor(name)

	if index != -1 {
		return false
	}

	for _, executor := range m.executors {
		if executor.Info.Name == name {
			return false
		}
	}

	executorFile := path.Join(executorFolder, name+executorSuffix)

	if fileutil.ExistsFile(executorFile) {
		return false
	}

	key, _ := random.GenerateRandomString(64)

	info := Info{
		Name:        name,
		MaxMemory:   maxMemory,
		MaxCPUUsage: maxCpuUsage,
		SecretKey:   key,
	}

	err := config.InitConfigIfNotExists(executorFile, fileutil.CodingJSON, &info)

	if err != nil {
		return false
	}

	m.executors = append(m.executors, &Executor{
		Info:     info,
		endpoint: nil,
	})

	logger.WarnfAsImportant("Executor %s was created", name)
	return true
}

func (m *Manager) DeleteExecutor(name string) bool {
	var index = -1
	for i, executor := range m.executors {
		if executor.Info.Name == name {
			index = i
		}
	}

	if index == -1 {
		return false
	}

	executors := m.executors
	executors[index] = executors[len(executors)-1]
	executors[len(executors)-1] = nil
	m.executors = executors[:len(executors)-1]

	executorFile := path.Join(executorFolder, name+executorSuffix)
	_ = os.Remove(executorFile)
	_ = event.Go(event_names.ExecutorDelete)

	logger.WarnfAsImportant("Executors %s was deleted.", name)
	return true
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

		if strings.Split(file.Name(), executorSuffix)[0] != executorInfo.Name {
			logger.Errorf(
				"Unable to load executor configuration from %s: File name doesn't match executor name",
				executorInfo.Name,
			)
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
