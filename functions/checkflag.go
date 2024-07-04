package handler

import (
	"regexp"
)
// CheckFlag checks if a given string represents a valid flag (output file or color option).
// Returns true if the flag is valid, otherwise false.
func CheckFlag(Flag string) bool {
	Checkflag := regexp.MustCompile(`^(--output=((.*).txt)|--color=(?i)(Red|Blue|Yellow|Green|Magenta|Cyan|Gray|orange))$`)
	flag := Checkflag.FindStringSubmatch(Flag)
	return flag != nil
}
