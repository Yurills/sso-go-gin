package utils

import "strconv"

func GenerateRandomString(length int) string {
	//temporary
	rstring := ""
	for i := range length {
		rstring += strconv.Itoa(i % 10) // Generate a string of digits
	}
	return rstring
}
