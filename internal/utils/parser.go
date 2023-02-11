package utils

import (
	"math"
	"strconv"
	"strings"
)

// ParseAmount  returns amount of net package as bytes
func ParseAmount(raw string) (int, error) {
	// گیگابایت 15
	// first trim raw string cause it has many spaces in the end of string.
	raw = strings.TrimSpace(raw)
	// split raw text with spaces between amount and label
	splitRawText := strings.Split(raw, " ")
	amountStr := splitRawText[0]
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		if strings.Contains(err.Error(), "invalid") {
			if amountStr == "۱" {
				amount = 1
			}
		} else {
			return -1, err
		}
	}
	label := splitRawText[1]
	if strings.Contains(label, "گیگ") {
		return int(math.Pow(10, 9)) * amount, nil
	} else {
		return int(math.Pow(10, 6)) * amount, nil
	}
}

func ParseDuration(raw string) (int, error) {
	rawSeparatedBySpaces := trimAndSplit(raw)
	days, err := strconv.Atoi(rawSeparatedBySpaces[0])
	if err != nil {
		return 1, nil
	}
	return days, nil

}

func ParsePrice(raw string) (int, error) {
	//raw := "                                                    62,000 تومان"
	rawSeparatedBySpaces := trimAndSplit(raw)
	price, err := strconv.Atoi(strings.ReplaceAll(rawSeparatedBySpaces[0], ",", ""))
	if err != nil {
		return -1, err
	}
	return price, nil
}

// Helper Functions
func trimAndSplit(raw string) []string {
	raw = strings.TrimSpace(raw)
	return strings.Split(raw, " ")
}

func convertToInteger(numberAsStr string) (int, error) {
	return strconv.Atoi(strings.ReplaceAll(numberAsStr, ",", ""))
}
