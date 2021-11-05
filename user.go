package main

import "fmt"

func handleListUsers() {
	users, err := listUsers()
	if err != nil {
		out(err.Error())
	}

	for _, user := range users {
		info(fmt.Sprintf("%d\t%s\t%s\trol: %s", user.ID, user.Name, user.Email, user.Role))
	}
}
