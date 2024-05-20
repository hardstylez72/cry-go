package defi

import (
	"math/big"
	"reflect"
	"testing"
)

func TestSlippage(t *testing.T) {
	type args struct {
		v               *big.Int
		slippagePercent SlippagePercent
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "",
			args: args{
				v:               big.NewInt(100),
				slippagePercent: "1",
			},
			want:    big.NewInt(99),
			wantErr: false,
		},
		{
			name: "",
			args: args{
				v:               big.NewInt(100_000),
				slippagePercent: "1",
			},
			want:    big.NewInt(99_000),
			wantErr: false,
		},
		{
			name: "",
			args: args{
				v:               big.NewInt(100_000),
				slippagePercent: "10",
			},
			want:    big.NewInt(90_000),
			wantErr: false,
		},
		{
			name: "",
			args: args{
				v:               big.NewInt(100_000),
				slippagePercent: "0.1",
			},
			want:    big.NewInt(99_900),
			wantErr: false,
		},
		{
			name: "",
			args: args{
				v:               big.NewInt(100_000),
				slippagePercent: "0.01",
			},
			want:    big.NewInt(99_990),
			wantErr: false,
		},
		{
			name: "",
			args: args{
				v:               big.NewInt(100_000),
				slippagePercent: "0.001",
			},
			want:    big.NewInt(99_999),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Slippage(tt.args.v, tt.args.slippagePercent)
			if (err != nil) != tt.wantErr {
				t.Errorf("Slippage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slippage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slippageSpecial(t *testing.T) {
	type args struct {
		v *big.Int
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := slippageSpecial(tt.args.v, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("slippageSpecial() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("slippageSpecial() got = %v, want %v", got, tt.want)
			}
		})
	}
}
