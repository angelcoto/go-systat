package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

// generalStat muestra estadísticas generales de recursos cada 5 segundos
func generalStat(scale byte) {
	fmt.Printf("%s\t%s\t%s\t%s\n", "Mem Total", "Mem Disponible", "%Mem Utilizada", "%CPU tilizado")
	for {
		c, _ := cpu.Percent(time.Second*5, false)
		v, _ := mem.VirtualMemory()
		fmt.Printf("%.2f\t%.2f\t%.2f\t%.2f\n",
			convScale(v.Total, scale),
			convScale(v.Available, scale),
			v.UsedPercent,
			c[0],
		)

	}
}