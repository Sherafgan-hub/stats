package stats

import (
	"github.com/Sherafgan-hub/bank1/pkg/types"
	"fmt"
)

func ExampleAvg() {
	myPay := []types.Payment {
		{
			ID: 1,
			Amount: 100,
		},
		{
			ID: 2,
			Amount: 200,
		},

		{
			ID: 3,
			Amount: 300,
		},
	}
	fmt.Println(Avg(myPay))

	// Output:
	// 200

}

func ExampleTotalInCategory() {
	paySrc := []types.Payment {
		{
			ID: 1,
			Amount: 100,
			Category: "test",
		},
		{
			ID: 2,
			Amount: 200,
			Category: "notest",
		},

		{
			ID: 3,
			Amount: 300,
			Category: "test",
		},
	}
	fmt.Println(TotalInCategory(paySrc, "test"))

	// Output:
	// 400

}
