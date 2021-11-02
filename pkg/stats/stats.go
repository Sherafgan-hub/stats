package stats

import (
	"github.com/Sherafgan-hub/bank1/v2/pkg/types"
)

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) (avgSum types.Money) {
	var total types.Money
	var sum []types.Money

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
	} else {
		total += payment.Amount
	}

	sum = append(sum, total)
}

	avgSum = total / types.Money(len(sum))
	return
}

//TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) (specCatSum types.Money) {
	for _, payment := range payments {
		if category == payment.Category {
			if payment.Status == types.StatusFail {
				continue
			} else {
				specCatSum += payment.Amount
			}
		}
	}

	return
}

//FilterByCategory возврашает платежи в указаной категории.
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered=append(filtered, payment)
		}
	}
	return filtered
}

//CatigoriesTotal возвращает сумму платежей по каждой катгори.
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
} 

// CategoriesAvg рассчитывает среднюю сумму платежа по каждой категории.
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	total := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount

		total[payment.Category]++
	}
	
	for key, category :=range categories {
		categories[key] = category/total[key]
	}

	return categories
}