package bidder

import (
	"math"
	"math/rand"
	"time"
)

type Bidder struct {
	Id string
}

type BidResponse struct {
	BidderId  string
	Bidamount float64
}

// = {{1: Bidder{id: "a1"}}{2: Bidder{id: "a2"}}}
// for _, v := range bidders {
// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))
// 	amount = r.float64()*1000
// }

func (b *Bidder) NewBid() (*BidResponse, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	amount := math.Floor((r.Float64() * 1000) * 100 / 100)
	return &BidResponse{
		BidderId:  b.Id,
		Bidamount: amount,
	}, nil
}
