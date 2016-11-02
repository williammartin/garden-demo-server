package commands

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Serve struct {
	Port int `short:"p" long:"port" description:"port to serve requests on" default:"8080"`
}

func (command *Serve) Execute(args []string) error {
	fmt.Printf("Serving requests on port %d\n", command.Port)

	http.HandleFunc("/", handleWebSocketCommands)
	if err := http.ListenAndServe(fmt.Sprintf("localhost:%d", command.Port), nil); err != nil {
		return err
	}

	return nil
}

var upgrader = websocket.Upgrader{} // use default options

func handleWebSocketCommands(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
}
