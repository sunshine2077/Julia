package db

import (
	"fmt"

	"github.com/google/go-tika/tika"
	"github.com/spf13/viper"
)

var TikaClient *tika.Client

func ClientToTika(v *viper.Viper) error {
	TikaClient = tika.NewClient(nil,
		fmt.Sprintf("http://%s:%s/", (*v).GetString("tika.host"), (*v).GetString("tika.port")))
	return nil
}
