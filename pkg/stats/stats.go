package stats

import "github.com/fm2901/bank/v2/pkg/types"

// Calculates avg sum of payments
func Avg(payments []types.Payment) types.Money {
	allSum := types.Money(0)
	allCount := types.Money(0)

	if len(payments) < 1 {
		return allSum
	}

	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			allSum += payment.Amount
			allCount += 1
		}
	}

	if allCount < 1 {
		return allSum
	}

	avgSum := allSum / allCount

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

// Filter payments by category
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}

	return filtered
}

// Calculate sum in categories
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}

// Calculate avg sum of categories
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	sum := map[types.Category]types.Money{}
	count := map[types.Category]types.Money{}
	
	for _, payment := range payments {
		sum[payment.Category] += payment.Amount
		count[payment.Category] += 1
	}

	for category := range count {
		sum[category] = sum[category] / count[category]
	}

	return sum
}