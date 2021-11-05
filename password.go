package main

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
)

func handlePasswordChange() {
	users, err := listUsers()
	if err != nil {
		out(err.Error())
	}

	var emails []string
	for _, user := range users {
		emails = append(emails, user.Email)
	}

	prompt := []*survey.Question{
		{
			Name: "email",
			Prompt: &survey.Select{
				Message: "Qué usuario querés modificar?:",
				Options: emails,
			},
		},
		{
			Name:   "password",
			Prompt: &survey.Password{Message: "Nueva password"},
			Validate: func(val interface{}) error {
				if str, ok := val.(string); !ok || len(str) > 32 || len(str) < 12 {
					return errors.New("La password debe contener entre 12 y 32 caracteres.")
				}
				return nil
			},
		},
	}

	answers := struct {
		Email    string
		Password string
	}{}

	safe(survey.Ask(prompt, &answers))
	safe(updatePassword(answers.Email, answers.Password))
	info("Password actualizada exitosamente!")
}

func handlePasswordChangeLegacy() {
	prompt := []*survey.Question{
		{
			Name:      "email",
			Prompt:    &survey.Input{Message: "Email del usuario a modificar:"},
			Validate:  survey.Required,
			Transform: survey.ToLower,
		},
		{
			Name:   "password",
			Prompt: &survey.Password{Message: "Nueva password"},
			Validate: func(val interface{}) error {
				if str, ok := val.(string); !ok || len(str) > 32 || len(str) < 12 {
					return errors.New("La password debe contener entre 12 y 32 caracteres.")
				}
				return nil
			},
		},
	}

	answers := struct {
		Email    string
		Password string
	}{}

	safe(survey.Ask(prompt, &answers))
	safe(updatePassword(answers.Email, answers.Password))
	info("Password actualizada exitosamente!")
}
