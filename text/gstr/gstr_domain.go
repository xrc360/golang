package gstr

import "strings"

// IsSubDomain checks whether `subDomain` is sub-domain of mainDomain.
// It supports '*' in `mainDomain`.
func IsSubDomain(subDomain string, mainDomain string) bool {
	if p := strings.IndexByte(subDomain, ':'); p != -1 {
		subDomain = subDomain[0:p]
	}
	if p := strings.IndexByte(mainDomain, ':'); p != -1 {
		mainDomain = mainDomain[0:p]
	}
	var (
		subArray   = strings.Split(subDomain, ".")
		mainArray  = strings.Split(mainDomain, ".")
		subLength  = len(subArray)
		mainLength = len(mainArray)
	)
	// Eg:
	// "goxrc.org" is not sub-domain of "s.goxrc.org".
	if mainLength > subLength {
		for i := range mainArray[0 : mainLength-subLength] {
			if mainArray[i] != "*" {
				return false
			}
		}
	}

	// Eg:
	// "s.s.goxrc.org" is not sub-domain of "*.goxrc.org"
	// but
	// "s.s.goxrc.org" is sub-domain of "goxrc.org"
	if mainLength > 2 && subLength > mainLength {
		return false
	}
	minLength := subLength
	if mainLength < minLength {
		minLength = mainLength
	}
	for i := minLength; i > 0; i-- {
		if mainArray[mainLength-i] == "*" {
			continue
		}
		if mainArray[mainLength-i] != subArray[subLength-i] {
			return false
		}
	}
	return true
}
