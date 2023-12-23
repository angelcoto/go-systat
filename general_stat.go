package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// generalStat muestra estadísticas generales de recursos cada 5 segundos
func generalStat(scale byte) {
	counter := 0
	fmt.Printf("Hora\tMem Total\tMem Disponible\tMem Utilizada\tCPU tilizado\n")
	for {
		c, _ := cpu.Percent(time.Second*5, false)
		v, _ := mem.VirtualMemory()
		fmt.Printf("%s\t%.2f\t%.2f\t%.2f\t%.2f\n",
			time.Now().Format(time.RFC3339),
			convScale(v.Total, scale),
			convScale(v.Available, scale),
			v.UsedPercent,
			c[0],
		)
		counter++
		if counter == 48 {
			runtime.GC() // Garbage collected
			counter = 0
		}
	}
}
