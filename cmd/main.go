package main

import (
	"fmt"
	"github.com/adaptive-scale/dockyard/internal/configuration"
	"github.com/adaptive-scale/dockyard/internal/documentmanager"
	"github.com/adaptive-scale/dockyard/internal/server"
	"sync"
)

func main() {

	c := configuration.GetConfiguration()

	dManager := documentmanager.New(c)
	dManager.Generate()

	fmt.Println("completed")

	var wait sync.WaitGroup

	if c.Watch {
		wait.Add(1)

		go dManager.Watch()
	}

	if c.Serve {
		wait.Add(1)
		go func() {
			if err := server.New(documentmanager.OutputDir, c.Port).Start(); err != nil {
				panic(err)
			}
		}()
	}

	wait.Wait()
}
