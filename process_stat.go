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
	id      string
	name    string
	cmdline string
	state   string
	vsize   string
	rss     string
}

// getProcessInfo obtiene información sobre un proceso dado su PID
func (p *processInfo) getProcessInfo(pid int, dir string) error {
	// Lee el archivo "/proc/{pid}/stat" para obtener información sobre el proceso
	statPath := fmt.Sprintf(dir+"/%d/stat", pid)

	statContent, err := os.ReadFile(statPath)
	if err != nil {
		return err
	}

	/* 	// Divide el contenido del archivo en campos
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
	   	RSS = RSS * 4 / 1024 // Se usa tamaño de página de 4KB
	   	p.rss = fmt.Sprintf("%d", RSS)
	*/
	fields := strings.Fields(string(statContent))

	p.id = fields[0]

	statStr := string(statContent) // Convertir []byte a string

	openParen := strings.Index(statStr, "(")
	closeParen := strings.LastIndex(statStr, ")")

	// Extraer el nombre del proceso correctamente
	p.name = statStr[openParen+1 : closeParen]

	// Extraer los demás campos después del ")"
	fields = strings.Fields(statStr[closeParen+2:])

	// Asignaciones
	//	p.id = fields[0]
	p.state = fields[1]

	// Memoria virtual (VSZ) y residente (RSS)
	vsize, _ := strconv.Atoi(fields[20])
	p.vsize = fmt.Sprintf("%d", vsize/(1024*1024)) // Convertir a MB

	RSS, _ := strconv.Atoi(fields[21])
	p.rss = fmt.Sprintf("%d", RSS*4/1024) // Convertir a MB (asumiendo páginas de 4 KB)

	// Leer el archivo "/proc/{pid}/cmdline" para obtener la línea de comando completa
	cmdlinePath := fmt.Sprintf(dir+"/%d/cmdline", pid)
	cmdlineContent, err := os.ReadFile(cmdlinePath)
	if err == nil {
		// En /proc/<PID>/cmdline, los argumentos están separados por bytes nulos ('\0')
		p.cmdline = strings.ReplaceAll(string(cmdlineContent), "\x00", " ")
	} else {
		p.cmdline = "N/A"
	}

	return nil
}

// listSubProcesses recorre los subprocesos de cada proceso en el directorio "/proc/<PID>/task"
func listSubProcesses(pid int) error {
	dir := fmt.Sprintf("/proc/%d/task", pid)
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error al leer el directorio %s: %v\n", dir, err)
		return err
	}

	process := processInfo{}
	for _, file := range files {
		// Verifica si el nombre del archivo es un número (identificador de proceso)
		pid, err := strconv.Atoi(file.Name())
		if err == nil {
			// Obtener información del proceso
			err := process.getProcessInfo(pid, "/proc")
			if err != nil {
				fmt.Println(err)
				continue
			}

			fecha := time.Now()
			fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
				// time.Now().Format(time.RFC3339),
				// fecha.Format("2006-01-02 15:04:05"),
				fecha.Format("2006"),
				fecha.Format("01"),
				fecha.Format("02"),
				fecha.Format("15"),
				fecha.Format("04"),
				process.id,
				process.name,
				process.cmdline,
				process.state,
				process.vsize,
				process.rss)
		}
	}
	return nil
}

// getProcesses recorre los procesos en "/proc"
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

func processesStat(delay int) {

	delaySec := delay * 60
	fmt.Printf("Anio\tMes\tDia\tHora\tMinuto\tPID\tProceso\tComando\tEstado\tVSize_MB\tRSS_MB\n")
	for {
		getProcesses()
		time.Sleep(time.Second * time.Duration(delaySec))
	}
}
