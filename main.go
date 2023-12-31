package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "v1.0.0"

func main() {

	optPtr := flag.Bool("p", false, "Lista cada 60 segundos los procesos corriendo en el sistema")
	verPtr := flag.Bool("v", false, "Imprime la versión")
	flag.Parse()

	if *verPtr {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	if *optPtr {
		processesStat()
	}
	generalStat('M')

}
