package main

import (
	"flag"
)

func main() {

	optPtr := flag.String("m", "g", "Modo: g (estadísticas generales), p (lista de procesos)")
	flag.Parse()

	switch *optPtr {
	case "g":
		generalStat('M')
	case "p":
		listProcesses()
	}

}
