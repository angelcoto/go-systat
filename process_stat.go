package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Estructura para almacenar información sobre un proceso
type processInfo struct {
	id    string
	Name  string
	State string
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
	p.Name = strings.Trim(fields[1], "()")
	p.State = fields[2]

	// Cálculo de memoria virtual del proceso en MB
	vsize, _ := strconv.Atoi(fields[22])
	vsize = vsize / (1024 * 1024) // El valor leído está en bytes
	p.vsize = fmt.Sprintf("%d", vsize)

	// Cálculo de memoria residente en MB
	RSS, _ := strconv.Atoi(fields[23])
	RSS = RSS * 4 / 1024 // Se asume un tamaño de página de 4KB
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
			}
			fmt.Println(process)
		} else {
			fmt.Println(err)

		}
	}
	return nil
}

// listProcesses recorre los procesos en directorio "/proc"
// Cada proceso en /proc tiene subprocesos en sudirectorio "task"
// Para cada proceso se llama la función list
func listProcesses() {
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
