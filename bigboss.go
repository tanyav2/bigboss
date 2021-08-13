package bigboss

import "github.com/google/go-cmp/cmp"

var compOpts = cmp.AllowUnexported()

func AmIBigBoss() bool {
	return false
}
