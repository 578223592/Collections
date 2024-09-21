package Collections

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
	"testing"
)

type Struct1 struct {
	Age string
}
type Struct2 struct {
	Age int
}

func TestTransfer(t *testing.T) {

	type args struct {
		originSlice any
		transFuncs  []func(any) any
	}
	tests := []struct {
		name           string
		args           args
		wantInputSlice []Struct2
		wantSlice      []Struct2
	}{
		{
			name: "",
			args: args{
				originSlice: []Struct1{
					{Age: "1"},
					{Age: "2"},
					{Age: "3"},
				},
				transFuncs: []func(any) any{
					func(in any) any {
						return Struct2{Age: cast.ToInt(in.(Struct1).Age)}
					},
				},
			},
			wantInputSlice: make([]Struct2, 0),
			wantSlice: []Struct2{
				{Age: 1},
				{Age: 2},
				{Age: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCollection(tt.args.originSlice)
			for _, transFunc := range tt.args.transFuncs {
				c.transfer(transFunc)
			}
			c.Get(&tt.wantInputSlice)
			if !reflect.DeepEqual(tt.wantInputSlice, tt.wantSlice) {
				fmt.Println(c.Error)
				t.Errorf("got = %v, want %v", tt.wantInputSlice, tt.wantSlice)
			}
		})
	}
}

func TestWhere(t *testing.T) {

	type args struct {
		originSlice any
		whereFuncs  []func(any) bool
	}
	tests := []struct {
		name           string
		args           args
		wantInputSlice []Struct2
		wantSlice      []Struct2
	}{
		{
			name: "",
			args: args{
				originSlice: []Struct2{
					{Age: 1},
					{Age: 2},
					{Age: 3},
				},
				whereFuncs: []func(any) bool{
					func(in any) bool {
						return in.(Struct2).Age >= 2
					},
				},
			},
			wantInputSlice: make([]Struct2, 0),
			wantSlice: []Struct2{
				{Age: 2},
				{Age: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCollection(tt.args.originSlice)
			for _, whereFunc := range tt.args.whereFuncs {
				c.where(whereFunc)
			}
			c.Get(&tt.wantInputSlice)
			if !reflect.DeepEqual(tt.wantInputSlice, tt.wantSlice) {
				fmt.Printf("c.Error:%s\n", c.Error)
				t.Errorf("got = %v, want %v", tt.wantInputSlice, tt.wantSlice)
			}
		})
	}
}

// todo 添加一些性能测试，用于对比不同版本迭代之间的性能差距
