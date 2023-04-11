package centrifugo

import (
	"github.com/TotrazOrstO/Imapro/pkg/config"
	"github.com/centrifugal/gocent"
	"net/http"
)

func InitClient(cfg config.CentrifugoConfig) *gocent.Client {
	client := gocent.New(gocent.Config{
		Addr:       cfg.URL,
		Key:        "YOUR_API_KEY",
		HTTPClient: http.DefaultClient,
	})

	return client
}
