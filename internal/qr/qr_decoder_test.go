package qr

import "testing"

func TestDecodeQRImage_InvalidFile(t *testing.T) {

	result := DecodeQRImage("does_not_exist.png")

	if len(result) != 0 {
		t.Errorf("expected empty result got %v", result)
	}
}