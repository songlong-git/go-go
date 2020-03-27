package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	ret := Split("abcdfe", "b")
	want := []string{"a", "", "c", "d", "f", "e"}
	if !reflect.DeepEqual(ret, want) {
		t.Errorf("want:%v but get:%v\n", want, ret)
	}
}
