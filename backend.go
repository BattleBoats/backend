package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	_ "backend/controllers"
	"backend/utils"
)

// The main function starts the web server on the specified port.
func main() {
	fmt.Printf("GOMAXPROCS set to: %v", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	// No Favicon this is an API Server. But god forbid you actually use this in a browser.
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Start listining on configured port
	log.Fatal(http.ListenAndServe(utils.Conf.GetString("port"), nil))
}
