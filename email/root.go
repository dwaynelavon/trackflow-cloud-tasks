package email

import (
	"os"

	sendinblue "github.com/sendinblue/APIv3-go-library/lib"
)

type email struct {
	Client *sendinblue.APIClient
}

func Email() *email {
	cfg := sendinblue.NewConfiguration()

	apiKey := os.Getenv("SENDINBLUE_API_KEY")
	cfg.AddDefaultHeader("api-key", apiKey)

	sib := sendinblue.NewAPIClient(cfg)
	return &email{
		Client: sib,
	}
}
