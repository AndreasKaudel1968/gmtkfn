package gmtkfn

import "testing"

func TestGematikGoKit_RandomString(t *testing.T) {
	var gematikGoKit Gmtkfn

	s := gematikGoKit.RandomString(15)

	if len(s) != 15 {
		t.Error("wrong length random string returned")
	}
}
