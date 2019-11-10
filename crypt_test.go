package crypt

import (
	"testing"
)

func TestCrypt(t *testing.T) {
	stored := "K.kG5h6z2d2Ms"
	crypted := Crypt("datchandjjjjjjj", stored[:2])

	if stored != crypted {
		t.Fatal("crypted value not match", stored, crypted)
	}
}
