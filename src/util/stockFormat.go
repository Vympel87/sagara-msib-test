package util

import "strconv"

func FormatStock(stock *int) string {
    if stock == nil {
        return "N/A"
    }
    return strconv.Itoa(*stock)
}
