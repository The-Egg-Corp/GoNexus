package util

import "github.com/sanity-io/litter"

func PrettyPrint(i interface{}) {
	litter.Config.StripPackageNames = true
	litter.Dump(i)
}
