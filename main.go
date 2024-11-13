package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	s "winners.com/recipes/Server"
)

const DEBUG_MODE = false

var wait_group sync.WaitGroup = sync.WaitGroup{}

func main() {

	initialize_server(&wait_group)
	initialize_client(&wait_group)

	wait_group.Wait()
	os.Exit(3)
}

func initialize_server(wait_group *sync.WaitGroup) {
	wait_group.Add(1)
	go s.StartServer(DEBUG_MODE) // go routine so tests can be done concurrently

    //Check if server is ready before preceeding
	for {
		if s.ServerReady() {
			break
		}
	}
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 1)
		if DEBUG_MODE {
			fmt.Print(".")
		}
	}
	fmt.Print("\033[H\033[2J") // clear screen with ASCII
	wait_group.Done()
}

func initialize_client(wait_group *sync.WaitGroup) {
	wait_group.Add(1)
	go view_prompt(wait_group)
}
