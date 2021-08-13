package bigboss

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestBigBoss(t *testing.T) {
	a := amIBigBoss()
	b := amIBigBoss()
	if !cmp.Equal(a, b) {
		t.Fatal("failure to roundtrip")
	}
}
