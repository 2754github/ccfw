package main

import (
	"fmt"
)

type settings struct {
	Version int
	Agents  map[string]agent
	Options struct {
		Agents agentOptions
	}
}

func readSettings() (*settings, error) {
	data, err := _fs.read(ccfwSettingsFile)
	if err != nil {
		return nil, err
	}

	var v settings
	err = _jsonc.unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	err = v.init()
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func (s *settings) init() error {
	if s.Version != 0 {
		return fmt.Errorf(".version=%d must be 0", s.Version)
	}

	for id, agent := range s.Agents {
		err := agent.init(id, &s.Options.Agents)
		if err != nil {
			return err
		}

		s.Agents[id] = agent
	}

	return nil
}
