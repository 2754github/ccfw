package main

import (
	"github.com/2754github/ccfw/cmd/ccfw/model"
)

func main() {
	settings, err := model.ReadSettings()
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
