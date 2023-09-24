package db

import (
	"context"
	"fmt"

	"github.com/google/go-tika/tika"
	"github.com/spf13/viper"
)

var TikaClient *tika.Client

func ClientToTika(ctx context.Context, v *viper.Viper) error {
	TikaClient = tika.NewClient(nil,
		fmt.Sprintf("http://%s:%s/", (*v).GetString("tika.host"), (*v).GetString("tika.port")))
	_, err := TikaClient.Version(ctx)
	if err != nil {
		return err
	}
	return nil
}
