package util

import "fmt"

func FormatCurrency(amount float64) string {
    return fmt.Sprintf("Rp %.3f", amount)
}