package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "v1.0.1"

func main() {

	optPtr := flag.Bool("p", false, "Lista periodicamente los procesos corriendo en el sistema")
	verPtr := flag.Bool("v", false, "Imprime la versión del programa")
	flag.Parse()

	if *verPtr {
		fmt.Printf("go-systat %s. Copyright (c) 2024-2025 Ángel Coto.  GPU 3 License.\n", version)
		os.Exit(0)
	}

	if *optPtr {
		processesStat()
	}
	generalStat('M')

}
