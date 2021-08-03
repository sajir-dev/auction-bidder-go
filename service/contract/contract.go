package contract

type RegisterReq struct {
}

type RegisterRes struct {
}

type ListReq struct {
}

type ListRes struct {
	Bidders []string
}

type BidReq struct {
}

type BidRes struct {
	BidderId string
	MaxValue float64
}
