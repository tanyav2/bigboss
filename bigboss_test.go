package bigboss

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestBigBoss(t *testing.T) {
	a := AmIBigBoss()
	b := AmIBigBoss()
	if !cmp.Equal(a, b, compOpts) {
		t.Fatal("failure to roundtrip")
	}
}
