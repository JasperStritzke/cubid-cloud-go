package server_group

import (
	"errors"
	"fmt"
	"github.com/AlexanderGrom/go-event"
	"github.com/JasperStritzke/cubid-cloud/controller/event_names"
	"github.com/JasperStritzke/cubid-cloud/util/config"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/datamodel"
	"github.com/JasperStritzke/cubid-cloud/util/fileutil"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type IGroupManager interface {
	Create(name, groupType string, memory, minOnline, maxOnline int, executors ...string) error
	Delete(name string) error
	ListLoaded() []ServerGroup
}

const (
	groupsFolder    = "./groups"
	groupFileSuffix = ".group.json"
)

func init() {
	_ = os.MkdirAll(groupsFolder, os.ModePerm)
}

type GroupManager struct {
	groups map[string]ServerGroup
}

func NewGroupManager() IGroupManager {
	groupManager := GroupManager{
		groups: make(map[string]ServerGroup),
	}

	logger.Info("Loading groups...")
	groupManager.groups = groupManager.loadAndGet()
	logger.Infof("Loaded %s group(s)", fmt.Sprint(len(groupManager.groups)))

	return &groupManager
}

func (g *GroupManager) loadAndGet() map[string]ServerGroup {
	files, err := ioutil.ReadDir(groupsFolder)

	if err != nil {
		return nil
	}

	var groups = make(map[string]ServerGroup)

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), groupFileSuffix) {
			continue
		}

		var group datamodel.ServerGroupInfo
		if err = config.LoadConfig(groupsFolder+"/"+file.Name(), fileutil.CodingJSON, &group); err != nil {
			logger.Errorf("Unable to load group from file %s", file.Name())
			continue
		}

		groupModel := ServerGroup{
			Servers: nil,
			Info:    group,
		}

		groups[strings.ToLower(group.Name)] = groupModel
	}

	return groups
}

func (g *GroupManager) reload() {
	newLoadedGroups := g.loadAndGet()

	if len(g.groups) > 0 {
		for name, groupModel := range g.groups {
			_, exists := newLoadedGroups[name]

			if !exists {
				logger.Warnf("Group %s was deleted. Stopping all servers", groupModel.Info.Name)
				_ = event.Go(event_names.ServerGroupDelete, groupModel.Info.Name)
			}
		}
	}
}

func (g *GroupManager) Create(name string, groupType string, memory, minOnline, maxOnline int, executors ...string) error {
	if len(executors) == 0 {
		return errors.New("A group can't start without any associated executor")
	}

	groupFilePath := path.Join(groupsFolder, strings.ToLower(name)+groupFileSuffix)

	if fileutil.ExistsFile(groupFilePath) {
		return errors.New("group with that name already exists")
	}

	groupInfo := datamodel.ServerGroupInfo{
		Name:              name,
		Type:              groupType,
		Memory:            memory,
		MinOnlineServices: minOnline,
		MaxOnlineServices: maxOnline,
		ExecutorNames:     executors,
	}

	if err := config.InitConfigIfNotExists(groupFilePath, fileutil.CodingJSON, groupInfo); err != nil {
		return err
	}

	logger.Infof(
		"Created group %s/%s running on %s.", name, groupType,
		strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(executors), "[", ""), "]", ""),
	)

	g.reload()
	return nil
}

func (g *GroupManager) Delete(name string) error {
	panic("implement me")
}

func (g *GroupManager) ListLoaded() []ServerGroup {
	panic("implement me")
}
