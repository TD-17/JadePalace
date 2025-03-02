package services

import "github.com/TD17/jade-palace/pkg"

func GetTokenTemp() string {
	token, _, err := pkg.GenerateAllTokens("deipayan@gmail.com", "Deip", "dash", "admin", "356ghg")

	if err != nil {
		panic(err)
	}

	return token
}
