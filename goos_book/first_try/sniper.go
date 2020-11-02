package first_try

type PriceEvent struct {
	MinimalIncrement int
	NewPrice         int
	IsOurBid         bool
}

type State string

const (
	Joining State = "JOINING"
	Bidding State = "BIDDING"
	Losing  State = "LOSING"
	Lost    State = "LOST"
	Winning State = "WINNING"
	Won     State = "WON"
)

type Sniper struct {
	ItemID    string
	StopPrice int
	State     State
	LastPrice int
	LastBid   int
}

func NewSniper(itemID string, stopPrice int) Sniper {
	return Sniper{
		ItemID:    itemID,
		StopPrice: stopPrice,
		State:     Joining,
	}
}

func (sniper *Sniper) PriceReceived(event PriceEvent) {
	if event.IsOurBid {
		sniper.State = Winning
		return
	}

	bid := event.NewPrice + event.MinimalIncrement
	if bid > sniper.StopPrice {
		sniper.State = Losing
		return
	}

	sniper.State = Bidding
	sniper.LastPrice = event.NewPrice
	sniper.LastBid = bid
}

func (sniper *Sniper) AuctionClosed() {
	switch sniper.State {
	case Joining:
		sniper.State = Lost
	case Bidding:
		sniper.State = Lost
	case Losing:
		sniper.State = Lost
	case Winning:
		sniper.State = Won
	}
}
