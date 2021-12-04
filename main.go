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

const (
	cmdChangePassword = "Cambiar password"
	cmdListUsers      = "Listar usuarios"
	cmdCreateAdmin    = "Crear admin"
	cmdExit           = "Salir"
)

var operationsPrompt = &survey.Select{
	Message: "Qué operación querés realizar?:",
	Options: []string{cmdChangePassword, cmdListUsers, cmdCreateAdmin, cmdExit},
	Default: "Cambiar password",
}

func main() {
	info("Todo tranqui? Qué rompimos?\n")

	var op string
	safe(survey.AskOne(operationsPrompt, &op))

	switch op {
	case cmdExit:
		info("Hasta luego")
		os.Exit(0)
	case cmdListUsers:
		handleListUsers()
		os.Exit(0)
	case cmdChangePassword:
		handlePasswordChange()
		os.Exit(0)
	case cmdCreateAdmin:
		handleCreateAdmin()
		os.Exit(0)
	default:
		out(fmt.Sprintf("Operación %q no soportada. Nos vimos.", op))
	}
}
