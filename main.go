package main

import (
	"flag"
	"log"
	"web/app/root"
)

var listen = flag.String("listen", ":9000", "The api server listen on.")

func main() {
	flag.Parse()
	if err := root.Router.Run(*listen); err != nil {
		log.Fatal(err)
	}
}
