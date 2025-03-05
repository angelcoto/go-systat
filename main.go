package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "v1.0.2"

func main() {

	optPtr := flag.Bool("p", false, "Lista periodicamente los procesos corriendo en el sistema")
	verPtr := flag.Bool("v", false, "Imprime la versión del programa")
	delayPtr := flag.Int("t", 3, "Tiempo entre cada lectura de estadísticas de procesos (3 default)")
	flag.Parse()

	if *verPtr {
		fmt.Printf("go-systat %s. Copyright (c) 2024-2025 Ángel Coto.  GPU 3 License.\n", version)
		os.Exit(0)
	}

	if *optPtr {
		processesStat(*delayPtr)
	}
	generalStat('M')

}
