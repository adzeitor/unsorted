package coin_change

func MakeChange(amount int, coins []int) int {
	memo := make(map[memoKey]int)
	return makeChangeWithMemo(amount, coins, memo)
}

func makeChangeWithMemo(amount int, coins []int, memo map[memoKey]int) int {
	if len(coins) == 0 {
		return 0
	}
	if value, ok := memo[newMemoKey(amount, coins)]; ok {
		return value
	}

	coin := coins[0]
	remainingCoins := coins[1:]

	n := 0
	remainingAmount := amount
	for remainingAmount >= 0 {
		if remainingAmount == 0 {
			n += 1
			break
		}
		if len(remainingCoins) > 0 {
			n += makeChangeWithMemo(remainingAmount, remainingCoins, memo)
		}
		remainingAmount -= coin
	}
	memo[newMemoKey(amount, coins)] = n
	return n
}

type memoKey struct {
	amount   int
	lenCoins int
}

func newMemoKey(amount int, coins []int) memoKey {
	return memoKey{
		amount:   amount,
		lenCoins: len(coins),
	}
}
