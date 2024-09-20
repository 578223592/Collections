package Collections

import (
	"reflect"
	"testing"
)

func TestTransfer(t *testing.T) {
	type fields struct {
		InnerSlice          any
		ReflectValueOfInner reflect.Value
		Error               error
	}
	type args struct {
		transFunc func(in any) (out any)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Collection
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Collection{
				InnerSlice:          tt.fields.InnerSlice,
				ReflectValueOfInner: tt.fields.ReflectValueOfInner,
				Error:               tt.fields.Error,
			}
			if got := c.transfer(tt.args.transFunc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}
