package main

import "strconv"

func IntToHex(i int) string {

	return strconv.FormatInt(int64(i), 16)
}
