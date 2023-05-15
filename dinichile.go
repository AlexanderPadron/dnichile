package dnichile

import (
	"strings"
    "strconv"
)

func validarRut(rut string) bool {
	// Eliminar puntos y guiones y convertir a mayúsculas
	rut = strings.ToUpper(strings.ReplaceAll(rut, ".", ""))
	rut = strings.ReplaceAll(rut, "-", "")

	// Separar número y dígito verificador
	numero := rut[:len(rut)-1]
	dv := rut[len(rut)-1]

	// Validar que el número sea numérico
	_, err := strconv.Atoi(numero)
	if err != nil {
		return false
	}

	// Calcular dígito verificador esperado
	var suma int
	var multiplicador int = 2
	for i := len(numero) - 1; i >= 0; i-- {
		suma += int(numero[i]-'0') * multiplicador
		multiplicador++
		if multiplicador > 7 {
			multiplicador = 2
		}
	}
	var dvEsperado int = 11 - (suma % 11)
	if dvEsperado == 10 {
		dvEsperado = 'K'
	} else if dvEsperado == 11 {
		dvEsperado = '0'
	}
	dvInt, err := strconv.Atoi(string(dv))
	if err != nil {
		return false
	}
	// Comparar dígito verificador esperado con el recibido
	return dvInt == dvEsperado
}