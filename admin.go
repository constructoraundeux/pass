package main

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
)

func handleCreateAdmin() {
	prompt := []*survey.Question{
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Nombre del nuevo admin:",
			},
		},
		{
			Name: "email",
			Prompt: &survey.Input{
				Message: "Email del nuevo admin:",
			},
		},
		{
			Name:   "password",
			Prompt: &survey.Password{Message: "Password"},
			Validate: func(val interface{}) error {
				if str, ok := val.(string); !ok || len(str) > 32 || len(str) < 12 {
					return errors.New("La password debe contener entre 12 y 32 caracteres.")
				}
				return nil
			},
		},
	}

	answers := struct {
		Name     string
		Email    string
		Password string
	}{}

	safe(survey.Ask(prompt, &answers))
	safe(createAdmin(answers.Name, answers.Email, answers.Password))
	info("Usuario con rol `admin` creado exitosamente!")
}
