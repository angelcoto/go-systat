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
	fmt.Printf("Anio\tMes\tDia\tHora\tMinuto\tSec\tMem_tot\tMem_disp\tMem_util_per\tCPU_util_per\n")
	for {

		fecha := time.Now()
		c, _ := cpu.Percent(time.Second*10, false)
		v, _ := mem.VirtualMemory()
		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%.2f\t%.2f\t%.2f\t%.2f\n",
			// time.Now().Format(time.RFC3339),
			fecha.Format("2006"),
			fecha.Format("01"),
			fecha.Format("02"),
			fecha.Format("15"),
			fecha.Format("04"),
			fecha.Format("05"),
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
