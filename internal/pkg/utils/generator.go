package utils

import (
	
)

func GenerateRandomString(length int) string {
	//temporary
	rstring := ""
	for i := range length {
		rstring += string(i)
	}
	return rstring
 }