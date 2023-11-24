package task

import (
	"reflect"
	"testing"

	v1 "github.com/hardstylez72/cry/internal/pb/gen/proto/go/v1"
	"github.com/stretchr/testify/assert"
)

func TestExtractItems(t *testing.T) {

	arg := PermuteArg{
		In: []RandomTask{
			{
				Type: v1.TaskType_SyncSwap,
				Kind: v1.TaskKind_TKSwap,
				P: &v1.RPswapItem{
					Network: v1.Network_ZKSYNCERA, From: v1.Token_ETH, To: v1.Token_USDC,
				},
			},
			{
				Type: v1.TaskType_SyncSwap,
				Kind: v1.TaskKind_TKSwap,
				P: &v1.RPswapItem{
					Network: v1.Network_ZKSYNCERA, From: v1.Token_USDC, To: v1.Token_ETH,
				},
			},
			{
				Type: v1.TaskType_SyncSwap,
				Kind: v1.TaskKind_TKSwap,
				P: &v1.RPswapItem{
					Network: v1.Network_ZKSYNCERA, From: v1.Token_USDC, To: v1.Token_USDT,
				},
			},
			{
				Type: v1.TaskType_SyncSwap,
				Kind: v1.TaskKind_TKSwap,
				P: &v1.RPswapItem{
					Network: v1.Network_ZKSYNCERA, From: v1.Token_VC, To: v1.Token_USDT,
				},
			},
			{
				Type: v1.TaskType_SyncSwap,
				Kind: v1.TaskKind_TKSwap,
				P: &v1.RPswapItem{
					Network: v1.Network_ZKSYNCERA, From: v1.Token_USDT, To: v1.Token_VC,
				},
			},
		},
		Limit: map[v1.TaskType]int{
			v1.TaskType_SyncSwap: 2,
		},
		StartFromNetwork: v1.Network_ZKSYNCERA,
		MaxCount:         1000,
		MaxItemLenght:    2,
	}

	for i := 0; i < 10; i++ {
		out, err := Permute(arg)
		assert.NoError(t, err)
		for _, tasks := range out {
			for _, task := range tasks {
				print(" " + task.InKey() + "_" + task.OutKey() + " ")
			}
			println("n")
		}
	}

}

func TestPermute(t *testing.T) {
	type args struct {
		arg PermuteArg
	}
	tests := []struct {
		name    string
		args    args
		want    [][]RandomTask
		wantErr bool
	}{
		{
			args: args{
				arg: PermuteArg{
					In: []RandomTask{
						{
							Type: v1.TaskType_SyncSwap,
							Kind: v1.TaskKind_TKSwap,
							P: &v1.RPswapItem{
								Network: v1.Network_ZKSYNCERA, From: v1.Token_ETH, To: v1.Token_USDC,
							},
						},
						{
							Type: v1.TaskType_SyncSwap,
							Kind: v1.TaskKind_TKSwap,
							P: &v1.RPswapItem{
								Network: v1.Network_ZKSYNCERA, From: v1.Token_USDC, To: v1.Token_ETH,
							},
						},
						{
							Type: v1.TaskType_SyncSwap,
							Kind: v1.TaskKind_TKSwap,
							P: &v1.RPswapItem{
								Network: v1.Network_ZKSYNCERA, From: v1.Token_USDC, To: v1.Token_USDT,
							},
						},
						{
							Type: v1.TaskType_SyncSwap,
							Kind: v1.TaskKind_TKSwap,
							P: &v1.RPswapItem{
								Network: v1.Network_ZKSYNCERA, From: v1.Token_VC, To: v1.Token_USDT,
							},
						},
						{
							Type: v1.TaskType_SyncSwap,
							Kind: v1.TaskKind_TKSwap,
							P: &v1.RPswapItem{
								Network: v1.Network_ZKSYNCERA, From: v1.Token_USDT, To: v1.Token_VC,
							},
						},
					},
					Limit: map[v1.TaskType]int{
						v1.TaskType_SyncSwap: 2,
					},
					StartFromNetwork: v1.Network_ZKSYNCERA,
					MaxCount:         1000,
					MaxItemLenght:    2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Permute(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Permute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permute() got = %v, want %v", got, tt.want)
			}
		})
	}
}
