package main

// convScale devuelve el valor en la escala requerida:
// KB, MB o GB
func convScale(val uint64, scale byte) (newVal float32) {
	fval := float32(val)
	switch scale {
	case 'K':
		newVal = fval / 1024
	case 'M':
		newVal = fval / (1024 * 1024)
	case 'G':
		newVal = fval / (1024 * 1024 * 1024)
	}
	return newVal
}
