package fiat_shamir

import (
	"math/big"
	"testing"
)

func TestProove(t *testing.T) {

	G := Secp256k1Order()
	msg := []byte("hey, you look beautiful today.")
	x := big.NewInt(0).SetBytes(msg)

	type args struct {
		x *big.Int
		g *big.Int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"secpk1",
			args{
				x: x,
				g: G,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Proove(tt.args.x, tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !got.Verify() {
				t.Error("failed to verify proof")
			}
		})
	}
}
