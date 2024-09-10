package bot

import "fmt"

func Teste() error {
	token, err := getToken()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Token: ", token.AccessJwt)

	return nil
}
