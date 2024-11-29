package davinci

import (
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

type DaVinciEnvironment struct {
	EnvironmentID string
	Client        *davinci.APIClient
}
