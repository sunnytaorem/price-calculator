package main

import (
	"fmt"
	"go-project/price-calculator/filemanager"
	"go-project/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.15, 0.30}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxedPricesJob(*fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process job.", err)
		}
	}
}
