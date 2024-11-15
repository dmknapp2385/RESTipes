package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	s "winners.com/recipes/Server"
)

//The main package includes the startup for the server, the view and controller
//wait_group is used so that cocurrent processes of starting the server and
//prompting the user for input are not overlapping. This allows the server to
//start up and run before the user is allowed to input data.

const DEBUG_MODE = false

var wait_group sync.WaitGroup = sync.WaitGroup{}
var file = flag.String("file", "", "a string");

func main() {
	//parse any flags from command line
    flag.Parse();

	// use wait_group to delay prompts until server has started
	initialize_server(&wait_group)
	initialize_client(&wait_group, file)

	wait_group.Wait()
	os.Exit(3)
}

//start server as goroutine, wait until server running before stopping
//wait group
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

//start view prompts in wait_group to delay system exit
func initialize_client(wait_group *sync.WaitGroup, file *string) {
	wait_group.Add(1)
	go view_prompt(wait_group, file)
}
