package prices

import (
	"fmt"
	"go-project/price-calculator/conversion"
	"go-project/price-calculator/iomanager"
)

type TaxedPricesJob struct {
	IOManager   iomanager.IOManager `json:"-"`
	TaxRate     float64             `json:"tax_rate"`
	Prices      []float64           `json:"prices"`
	TaxedPrices map[string]string   `json:"taxed_prices"`
}

func (job *TaxedPricesJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}
	job.Prices = prices
	return nil
}

func (job *TaxedPricesJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]string)
	for _, price := range job.Prices {
		taxedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}
	job.TaxedPrices = result
	return job.IOManager.WriteResult(job)
}

func NewTaxedPricesJob(iom iomanager.IOManager, taxRate float64) *TaxedPricesJob {
	return &TaxedPricesJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}
