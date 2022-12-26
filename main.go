// It creates a new Echo instance, adds some middleware, creates a new WhyPFS node, creates a new GatewayHandler, and then
// adds a route to the Echo instance
package main

import (
	logging "github.com/ipfs/go-log/v2"
	_ "net/http"
	"whypfs-gateway/cmd"
)

var (
	log = logging.Logger("api")
)

func main() {

	cmd.DaemonCmd() // daemon
}
