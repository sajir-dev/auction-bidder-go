package app

import (
	"github.com/sajir-dev/auction-bidder-go/handler"
)

func mapUrls() {
	router.GET("/ping", handler.Ping)
	router.GET("/bid", handler.Bid)
	router.GET("/list", handler.List)
}
