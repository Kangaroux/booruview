package api

import (
	"os"

	"github.com/valkey-io/valkey-go"
)

var (
	hasInit      bool
	valkeyClient valkey.Client
)

const KeyTtl = 10 * 60 // seconds

func InitValkey() error {
	if hasInit {
		return nil
	}

	vc, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{os.Getenv("VALKEY_ADDR")}})
	if err != nil {
		return err
	}

	valkeyClient = vc
	hasInit = true

	return nil
}

func Valkey() valkey.Client {
	return valkeyClient
}
