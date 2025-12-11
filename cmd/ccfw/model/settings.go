package model

import (
	"fmt"

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
