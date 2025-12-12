package model

import (
	"fmt"
	"path/filepath"
	"regexp"
	"slices"

	"github.com/2754github/ccfw/cmd/ccfw/config"
	"github.com/2754github/ccfw/cmd/ccfw/util/file"
	"github.com/2754github/ccfw/cmd/ccfw/util/heredoc"
)

type agent struct {
	agentOptions

	Name string
}

type agentOptions struct {
	CommandPrefix  string
	InvocationMode string
}

const (
	agentInvocationModeAuto    = "auto"
	agentInvocationModeCommand = "command"
)

var (
	agentNameRegexp = regexp.MustCompile(`^[a-z-]+$`)

	//nolint:gochecknoglobals
	agentInvocationModes = []string{
		agentInvocationModeAuto,
		agentInvocationModeCommand,
	}
)

func (a *agent) init(name string, options *agentOptions) error {
	if a.Name == "" {
		a.Name = name
	}

	if !agentNameRegexp.MatchString(a.Name) {
		return fmt.Errorf(
			".agents.%s.name=%q must match `%s`",
			name,
			a.Name,
			agentNameRegexp.String(),
		)
	}

	if a.CommandPrefix == "" {
		a.CommandPrefix = options.CommandPrefix
	}

	if a.InvocationMode == "" {
		a.InvocationMode = options.InvocationMode
	}

	if !slices.Contains(agentInvocationModes, a.InvocationMode) {
		return fmt.Errorf(
			".agents.%s.invocationMode=%q must be one of %q",
			name,
			a.InvocationMode,
			agentInvocationModes,
		)
	}

	return nil
}

func (a *agent) Path() string {
	return filepath.Join(config.ClaudeAgentsDir, a.Name+config.ClaudeAgentFileExt)
}

func (a *agent) HasCommand() bool {
	return a.InvocationMode == agentInvocationModeCommand
}

func (a *agent) markdown() []byte {
	description := ""
	if a.InvocationMode == agentInvocationModeAuto {
		description = "MUST BE USED "
	}

	return heredoc.Format(`
---
name: %s
description: %s
---
`, a.Name, description)
}

func (a *agent) CommandPath() string {
	return filepath.Join(
		config.ClaudeCommandsDir,
		a.CommandPrefix+a.Name+config.ClaudeCommandFileExt,
	)
}

func (a *agent) commandMarkdown() []byte {
	return heredoc.Format(`
---
description: Invocate the %s subagent.
argument-hint: <What you want to delegate to the %s subagent.>
---

Use the %s subagent to accomplish the following.

$ARGUMENTS
`, a.Name, a.Name, a.Name)
}

func WriteAgent(agent *agent) error {
	err := file.Write(agent.Path(), agent.markdown())
	if err != nil {
		return err
	}

	if agent.HasCommand() {
		return file.Write(agent.CommandPath(), agent.commandMarkdown())
	}

	return nil
}
