package stats

import "github.com/fm2901/bank/v2/pkg/types"

// Calculates avg sum of payments
func Avg(payments []types.Payment) types.Money {
	allSum := types.Money(0)

	if len(payments) < 1 {
		return allSum
	}

	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			allSum += payment.Amount
		}
	}

	avgSum := allSum / types.Money(len(payments))

	return avgSum
}

// Calculate sum of payments in category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	allSum := types.Money(0)

	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		if payment.Status == types.StatusFail {
			continue
		}
		allSum += payment.Amount
	}

	return allSum
}
