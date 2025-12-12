package model

import (
	"fmt"
	"slices"

	"github.com/2754github/ccfw/cmd/ccfw/config"
	"github.com/2754github/ccfw/cmd/ccfw/util/file"
	"github.com/2754github/ccfw/cmd/ccfw/util/jsonc"
)

type settings struct {
	Version int
	Agents  map[string]agent
	Options struct {
		Agents agentOptions
	}
}

func (s *settings) init() error {
	if s.Version != 0 {
		return fmt.Errorf(".version=%d must be 0", s.Version)
	}

	for name, agent := range s.Agents {
		err := agent.init(name, &s.Options.Agents)
		if err != nil {
			return err
		}

		s.Agents[name] = agent
	}

	return nil
}

func ReadSettings() (*settings, error) {
	data, err := file.Read(config.CcfwSettingsFile)
	if err != nil {
		return nil, err
	}

	var v settings
	err = jsonc.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	err = v.init()
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func RemoveUntrackedFiles(settings *settings) error {
	agents, err := file.Paths(config.ClaudeAgentsDir)
	if err != nil {
		return err
	}

	commands, err := file.Paths(config.ClaudeCommandsDir)
	if err != nil {
		return err
	}

	for _, agent := range settings.Agents {
		agents = deleteElem(agents, agent.path())

		if agent.hasCommand() {
			commands = deleteElem(commands, agent.commandPath())
		}
	}

	for _, path := range agents {
		err := file.Remove(path)
		if err != nil {
			return err
		}
	}

	for _, path := range commands {
		err := file.Remove(path)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteElem[T comparable](slice []T, elem T) []T {
	c := make([]T, len(slice))
	copy(c, slice)

	i := slices.Index(c, elem)
	if i != -1 {
		c = slices.Delete(c, i, i+1)
	}

	return c
}
