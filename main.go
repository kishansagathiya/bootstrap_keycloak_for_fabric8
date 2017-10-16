package main

import (
	"fmt"

	"github.com/bootstrap_keycloak_for_fabric8/bootstrap"
)

func main() {
	var config bootstrap.Config
	config, err := bootstrap.ParseConfigJSON()

	fmt.Println(config)
	fmt.Println(err)

	token, err := bootstrap.GetKeyCloakAdminToken()
	if err == nil {
		fmt.Println(token)
		result, err := bootstrap.CreateRealm(token)

		if result == true {
			fmt.Println("A new realm has been created")
		}
		fmt.Println(err)
	}

}
