package stats

import (
	"github.com/Sherafgan-hub/bank1/pkg/types"
)


// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) (avgSum types.Money) {
	var total types.Money

	for _, payment := range payments {
		total += payment.Amount
	}

	avgSum = total / types.Money(len(payments))
	return
}

//TotalInCategory находит сумму покупок в определённой категории
func TotalInCategory(payments []types.Payment, cateory types.Category) (specCatSum types.Money) {
	// var total types.Money

	for _, payment := range payments {
		if cateory == payment.Category {
			specCatSum +=payment.Amount
		}
	}
	return
}