package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")
}

var operationsPrompt = &survey.Select{
	Message: "Qué operación querés realizar?:",
	Options: []string{"Cambiar password", "Listar usuarios", "Salir"},
	Default: "Cambiar password",
}

func main() {
	info("Todo tranqui? Qué rompimos?\n")

	var op string
	safe(survey.AskOne(operationsPrompt, &op))

	switch op {
	case "Salir":
		info("Hasta luego")
		os.Exit(0)
	case "Listar usuarios":
		handleListUsers()
		os.Exit(0)
	case "Cambiar password":
		handlePasswordChange()
		os.Exit(0)
	default:
		out(fmt.Sprintf("Operación %q no soportada. Nos vimos.", op))
	}
}
