package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Estructura para almacenar información sobre un proceso
type processInfo struct {
	id    string
	name  string
	state string
	vsize string
	rss   string
}

// getProcessInfo obtiene información sobre un proceso dado su PID,
// imprimiendo la información en pantalla
func (p *processInfo) getProcessInfo(pid int, dir string) error {

	// Lee el archivo "/proc/{pid}/stat" para obtener información sobre el proceso
	statPath := fmt.Sprintf(dir+"/%d/stat", pid)

	statContent, err := os.ReadFile(statPath)
	if err != nil {
		return err
	}

	// Divide el contenido del archivo en campos
	fields := strings.Fields(string(statContent))

	p.id = fields[0]
	p.name = strings.Trim(fields[1], "()")
	p.state = fields[2]

	// Cálculo de memoria virtual del proceso en MB
	vsize, _ := strconv.Atoi(fields[22])
	vsize = vsize / (1024 * 1024) // El valor leído está en bytes
	p.vsize = fmt.Sprintf("%d", vsize)

	// Cálculo de memoria residente en MB
	RSS, _ := strconv.Atoi(fields[23])
	RSS = RSS * 4 / 1024 // Determinar con getconf PAGESIZE (4K asumido)
	p.rss = fmt.Sprintf("%d", RSS)

	return nil
}

// listSubProcesses recorre el listado de subprocesos de cada proceso
// en el directorio "/proc", obteniendo la información de cada uno de ellos
func listSubProcesses(pid int) error {
	dir := fmt.Sprintf("/proc/%d/task", pid)
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error al leer el directorio %s%v: ", dir, err)
		return err
	}

	process := processInfo{}
	for _, file := range files {
		// Verifica si el nombre del archivo es un número (identificador de proceso)
		pid, err := strconv.Atoi(file.Name())
		if err == nil {
			// Imprime información sobre el proceso
			err := process.getProcessInfo(pid, dir)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\n",
				time.Now().Format(time.RFC3339),
				process.id,
				process.name,
				process.state,
				process.vsize,
				process.rss)
		}
	}
	return nil
}

// getProcesses recorre los procesos en directorio "/proc".
// Cada proceso en /proc tiene subprocesos en sudirectorio "task".
// Para cada proceso se llama la función listSubProcesses
func getProcesses() {
	dir := "/proc"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error al leer el directorio /proc:", err)
		os.Exit(1)
	}

	// Itera sobre los archivos en el directorio /proc
	for _, file := range files {
		// Verifica si el nombre del archivo es un número (identificador de proceso)
		pid, err := strconv.Atoi(file.Name())
		if err == nil {
			_ = listSubProcesses(pid)
		}
	}

}

func processesStat() {
	fmt.Printf("Hora\tPID\tNombre\tEstado\tVSize(MB)\tRSS(MB)\n")
	for {
		getProcesses()
		time.Sleep(time.Second * 60)
	}
}
