package subcmd

import (
	"slices"

	"github.com/2754github/ccfw/cmd/ccfw/config"
	"github.com/2754github/ccfw/cmd/ccfw/model"
	"github.com/2754github/ccfw/cmd/ccfw/util/file"
	"github.com/2754github/ccfw/cmd/ccfw/util/heredoc"
)

func Help() {
	panic("subcommand must be one of [init, sync]")
}

func Init() {
	err := file.Write(config.CcfwSettingsFile, heredoc.Format(`
{
  "version": 0,
  "agents": {
    "designer": {},
    "implementer": {
      "commandPrefix": "y-"
    },
    "reviewer": {
      "invocationMode": "auto"
    }
  },
  "options": {
    "agents": {
      "commandPrefix": "x-",
      "invocationMode": "command"
    }
  }
}
`))
	if err != nil {
		panic(err)
	}
}

//nolint:cyclop
func Sync() {
	settings, err := model.ReadSettings()
	if err != nil {
		panic(err)
	}

	agentsToRemove, err := file.Paths(config.ClaudeAgentsDir)
	if err != nil {
		panic(err)
	}

	commandsToRemove, err := file.Paths(config.ClaudeCommandsDir)
	if err != nil {
		panic(err)
	}

	for _, agent := range settings.Agents {
		agentsToRemove = deleteElem(agentsToRemove, agent.Path())

		if agent.HasCommand() {
			commandsToRemove = deleteElem(commandsToRemove, agent.CommandPath())
		}

		err := model.WriteAgent(&agent)
		if err != nil {
			panic(err)
		}
	}

	for _, path := range agentsToRemove {
		err := file.Remove(path)
		if err != nil {
			panic(err)
		}
	}

	for _, path := range commandsToRemove {
		err := file.Remove(path)
		if err != nil {
			panic(err)
		}
	}
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
