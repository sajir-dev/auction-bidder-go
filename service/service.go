package service

import (
	"context"
	"fmt"

	"github.com/sajir-dev/auction-bidder-go/bidder"
	"github.com/sajir-dev/auction-bidder-go/service/contract"
)

var biddersData = map[int]bidder.Bidder{
	1: {Id: "a1"},
	2: {Id: "a2"},
	3: {Id: "b1"},
	4: {Id: "c1"},
	5: {Id: "c2"},
}

type Bids struct {
	AuctionId string
	BidderId  string
	Bidamount float64
}

var bidData = make(map[int]Bids)

func Register(ctx context.Context, req *contract.RegisterReq) (*contract.RegisterRes, error) {

	return nil, nil
}

func Bid(ctx context.Context, req *contract.BidReq) (*contract.BidRes, error) {
	i := 0
	for _, v := range biddersData {
		bidRes, _ := v.NewBid()
		bidData[i] = Bids{
			AuctionId: "123",
			BidderId:  bidRes.BidderId,
			Bidamount: bidRes.Bidamount,
		}
		i++
	}

	fmt.Println(bidData)

	var max float64 = 0
	var maxBidderId string

	for _, v := range bidData {
		if v.Bidamount > max && v.AuctionId == "123" {
			max = v.Bidamount
			maxBidderId = v.BidderId
		}
	}

	return &contract.BidRes{
		BidderId: maxBidderId,
		MaxValue: max,
	}, nil
}

func List(ctx context.Context, req *contract.ListReq) (*contract.ListRes, error) {
	// domain.InsertOne()
	bidders := []string{}
	for _, v := range biddersData {
		bidders = append(bidders, v.Id)
	}
	return &contract.ListRes{
		Bidders: bidders,
	}, nil
}
