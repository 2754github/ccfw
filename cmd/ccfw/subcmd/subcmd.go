package subcmd

import (
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

func Sync() {
	settings, err := model.ReadSettings()
	if err != nil {
		panic(err)
	}

	err = model.RemoveUntrackedFiles(settings)
	if err != nil {
		panic(err)
	}

	for _, agent := range settings.Agents {
		err := model.WriteAgent(&agent)
		if err != nil {
			panic(err)
		}
	}
}
