package group_manager

import (
	"fmt"
	"github.com/AlexanderGrom/go-event"
	"github.com/JasperStritzke/cubid-cloud/controller/event_names"
	"github.com/JasperStritzke/cubid-cloud/controller/model"
	"github.com/JasperStritzke/cubid-cloud/util/config"
	"github.com/JasperStritzke/cubid-cloud/util/console/logger"
	"github.com/JasperStritzke/cubid-cloud/util/datamodel"
	"io/ioutil"
	"os"
	"strings"
)

type IGroupManager interface {
	Create(name, groupType string, memory, minOnline, maxOnline int) error
	Delete(name string) error
	ListLoaded() []model.ServerGroup
}

const groupsFolder = "./groups"

func init() {
	_ = os.MkdirAll(groupsFolder, os.ModePerm)
}

type GroupManager struct {
	groups map[string]model.ServerGroup
}

func NewGroupManager() IGroupManager {
	groupManager := GroupManager{
		groups: make(map[string]model.ServerGroup),
	}

	logger.Info("Loading groups...")
	groupManager.groups = groupManager.loadAndGet()
	logger.Infof("Loaded %s group(s)", fmt.Sprint(len(groupManager.groups)))

	return &groupManager
}

func (g *GroupManager) loadAndGet() map[string]model.ServerGroup {
	files, err := ioutil.ReadDir(groupsFolder)

	if err != nil {
		return nil
	}

	var groups = make(map[string]model.ServerGroup)

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		var group datamodel.ServerGroup
		err = config.LoadConfig(groupsFolder+"/"+file.Name(), &group)

		if err != nil {
			logger.Errorf("Unable to load group from file %s", file.Name())
			continue
		}

		groupModel := model.ServerGroup{
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

func (g *GroupManager) Create(name string, groupType string, memory, minOnline, maxOnline int) error {
	panic("implement me")
}

func (g *GroupManager) Delete(name string) error {
	panic("implement me")
}

func (g *GroupManager) ListLoaded() []model.ServerGroup {
	panic("implement me")
}
