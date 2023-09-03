package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_randFloatRange(t *testing.T) {
	type args struct {
		min float64
		max float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{
				min: 0.01,
				max: 0.02,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, randFloatRange(tt.args.min, tt.args.max), "randFloatRange(%v, %v)", tt.args.min, tt.args.max)
		})
	}
}
