package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

type agent struct {
	agentOptions

	ID string
}

type agentOptions struct {
	CommandPrefix  string
	InvocationMode string
}

const (
	agentCommandPrefixDefault  = "x-"
	agentInvocationModeDefault = "none"

	agentInvocationModeAuto    = "auto"
	agentInvocationModeCommand = "command"
)

var (
	agentIDRegexp = regexp.MustCompile(`^[a-z-]+$`)

	agentInvocationModes = []string{
		agentInvocationModeAuto,
		agentInvocationModeCommand,
	}
)

func (a *agent) init(id string, options *agentOptions) error {
	if a.ID == "" {
		a.ID = id
	}

	if !agentIDRegexp.MatchString(a.ID) {
		return fmt.Errorf(".agents.%s.id=%q must match `%s`", id, a.ID, agentIDRegexp.String())
	}

	if a.CommandPrefix == "" {
		if options.CommandPrefix == "" {
			a.CommandPrefix = agentCommandPrefixDefault
		} else {
			a.CommandPrefix = options.CommandPrefix
		}
	}

	if a.InvocationMode == "" {
		if options.InvocationMode == "" {
			a.InvocationMode = agentInvocationModeDefault
		} else {
			a.InvocationMode = options.InvocationMode
		}
	}

	if !slices.Contains(agentInvocationModes, a.InvocationMode) {
		return fmt.Errorf(
			".agents.%s.invocationMode=%q must be one of %q",
			id,
			a.InvocationMode,
			agentInvocationModes,
		)
	}

	return nil
}

func (a *agent) path() string {
	return filepath.Join(claudeAgentsDir, a.ID+claudeAgentFileExt)
}

func (a *agent) markdown() []byte {
	description := ""
	if a.InvocationMode == agentInvocationModeAuto {
		description = "MUST BE USED "
	}

	return []byte(strings.TrimPrefix(fmt.Sprintf(`
---
name: %s
description: %s
---

# %s
`, a.ID, description, a.ID), "\n"))
}

func (a *agent) commandPath() string {
	return filepath.Join(claudeCommandsDir, a.CommandPrefix+a.ID+claudeCommandFileExt)
}

func (a *agent) commandMarkdown() []byte {
	return []byte(strings.TrimPrefix(fmt.Sprintf(`
---
description: Invocate the %s subagent.
argument-hint: <What you want to delegate to the %s subagent.>
---

# %s

Use the %s subagent to accomplish the following.

$ARGUMENTS
`, a.ID, a.ID, a.ID, a.ID), "\n"))
}

func writeAgent(agent *agent) error {
	err := _fs.write(agent.path(), agent.markdown())
	if err != nil {
		return err
	}

	if agent.InvocationMode == agentInvocationModeCommand {
		return _fs.write(agent.commandPath(), agent.commandMarkdown())
	}

	return nil
}
