package s

import (
	"fmt"
	"os"

	utils "github.com/harryvince/s/cli/pkg"
)

func Setup(environment ...string) {
	env := "dev"
	if len(environment) > 0 {
		env = environment[0]
	}

	ssm_client := utils.NewSSMClient()

	params, err := ssm_client.GetAllSecrets(env)
	if err != nil {
		fmt.Println("Error occured while trying to load environment variables")
	}

	for _, param := range params {
		os.Setenv(param.Name, param.Value)
	}

	fmt.Println("s: Environment variables loaded")
}
