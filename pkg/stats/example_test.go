package stats

import (
	"github.com/Sherafgan-hub/bank1/v2/pkg/types"
	"fmt"
)

func ExampleAvg() {
	myPay := []types.Payment {
		{
			ID: 1,
			Amount: 100,
			Status: types.StatusFail,
		},
		{
			ID: 2,
			Amount: 200,
			Status: types.StatusOk,
		},

		{
			ID: 3,
			Amount: 300,
			Status: types.StatusOk,
		},
		{
			ID: 4,
			Amount: 100,
			Status: types.StatusOk,
		},
	}
	fmt.Println(Avg(myPay))

	// Output:
	// 200

}

func ExampleTotalInCategory() {
	myPay := []types.Payment {
		{
			ID: 1,
			Amount: 100,
			Category: "test",
			Status: types.StatusFail,
		},
		{
			ID: 2,
			Amount: 200,
			Category: "notest",
			Status: types.StatusOk,
		},

		{
			ID: 3,
			Amount: 300,
			Category: "test",
			Status: types.StatusOk,
		},
		{
			ID: 4,
			Amount: 50,
			Category: "test",
			Status: types.StatusOk,
		},
	}
	fmt.Println(TotalInCategory(myPay, "test"))

	// Output:
	// 350

}
