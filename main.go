package main

import (
	"flag"
)

func main() {

	optPtr := flag.Bool("p", false, "Lista cada 60 segundos los procesos corriendo en el sistema")
	flag.Parse()

	if *optPtr {
		processesStat()
	}
	generalStat('M')

}
