package main

func main() {
	settings, err := readSettings()
	if err != nil {
		panic(err)
	}

	for _, agent := range settings.Agents {
		err := writeAgent(&agent)
		if err != nil {
			panic(err)
		}
	}
}
