package main

import (
	_ "go.uber.org/automaxprocs"

	"kubernetes-manager/internal/km"
	"os"
)

func main() {
	command := km.NewKuberManagerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
