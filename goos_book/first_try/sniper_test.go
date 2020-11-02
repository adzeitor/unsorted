package first_try

import "testing"

func TestSniper(t *testing.T) {
	itemID := "1"

	t.Run("Lost if auction is closed just after joining", func(t *testing.T) {
		sniper := NewSniper(itemID, 100)
		sniper.AuctionClosed()
		assert(t, Lost, sniper.State)
	})

	t.Run("Loosing if item is overpriced", func(t *testing.T) {
		sniper := NewSniper(itemID, 100)

		sniper.PriceReceived(PriceEvent{
			MinimalIncrement: 5,
			NewPrice:         96,
		})
		assert(t, Losing, sniper.State)

		sniper.AuctionClosed()
		assert(t, Lost, sniper.State)
	})

	t.Run("Bidding on stop price and won", func(t *testing.T) {
		sniper := NewSniper(itemID, 100)
		sniper.PriceReceived(PriceEvent{
			MinimalIncrement: 5,
			NewPrice:         95,
		})
		assert(t, Bidding, sniper.State)

		sniper.PriceReceived(PriceEvent{
			MinimalIncrement: 5,
			NewPrice:         100,
			IsOurBid:         true,
		})
		assert(t, Winning, sniper.State)

		sniper.AuctionClosed()
		assert(t, Won, sniper.State)
	})

	t.Run("Lost if auction is closed on bidding", func(t *testing.T) {
		sniper := NewSniper(itemID, 100)
		sniper.PriceReceived(PriceEvent{
			MinimalIncrement: 5,
			NewPrice:         95,
		})
		assert(t, 100, sniper.LastBid)
		assert(t, Bidding, sniper.State)

		sniper.AuctionClosed()
		assert(t, Lost, sniper.State)
	})

	t.Run("Try beat price if our price was beaten", func(t *testing.T) {
		sniper := NewSniper(itemID, 100)
		sniper.PriceReceived(PriceEvent{
			MinimalIncrement: 5,
			NewPrice:         50,
		})
		assert(t, Bidding, sniper.State)

		sniper.PriceReceived(PriceEvent{
			MinimalIncrement: 5,
			NewPrice:         55,
			IsOurBid:         true,
		})
		assert(t, Winning, sniper.State)

		sniper.PriceReceived(PriceEvent{
			MinimalIncrement: 5,
			NewPrice:         80,
		})
		assert(t, 85, sniper.LastBid)
		assert(t, Bidding, sniper.State)

		sniper.PriceReceived(PriceEvent{
			MinimalIncrement: 5,
			NewPrice:         85,
			IsOurBid:         true,
		})
		assert(t, Winning, sniper.State)

		sniper.AuctionClosed()
		assert(t, Won, sniper.State)
	})
}
