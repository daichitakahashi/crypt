package crypt

import (
	"testing"
)

func TestCrypt(t *testing.T) {
	stored := "us8yT8zzfb8Jg"
	crypted := Crypt("passsample", stored[:2])

	if stored != crypted {
		t.Fatal("crypted value not match", stored, crypted)
	}
}
