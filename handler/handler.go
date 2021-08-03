package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajir-dev/auction-bidder-go/service"
	"github.com/sajir-dev/auction-bidder-go/service/contract"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}

func Bid(c *gin.Context) {
	ctx := context.Background()
	bidreq := &contract.BidReq{}
	bids, _ := service.Bid(ctx, bidreq)
	c.JSON(http.StatusOK, bids)
}

func List(c *gin.Context) {
	ctx := context.Background()
	list, _ := service.List(ctx, &contract.ListReq{})
	c.JSON(
		http.StatusOK,
		list,
	)
}

func Register(c *gin.Context) {
	// Register when bidder comes online
}
