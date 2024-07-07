package slices

import "fmt"

func Includes[T comparable](list []T, value T) bool {
	for _, v := range list {
		if v == value {
			log(fmt.Sprintf("%s: found %v into the list", pkgName, value))
			return true
		}
	}

	log(fmt.Sprintf("%s: not ound %v into the list", pkgName, value))
	return false
}
