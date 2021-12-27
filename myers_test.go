package myers

import (
	"reflect"
	t "testing"
)

type TestCase struct {
	l1  []string
	l2  []string
	exp []Op
}

func TestDiff(t *t.T) {
	testCases := []TestCase{
		{[]string{}, []string{}, []Op{}},
		{[]string{}, []string{"foo"}, []Op{{OpInsert, 0, 0, "foo"}}},
		{[]string{"foo", "bar", "baz"}, []string{"foo", "bar", "baz"}, []Op{}},
		{[]string{"foo", "bar", "baz"}, []string{"foo", "baz"}, []Op{{OpDelete, 1, -1, "bar"}}},
		{[]string{"baz"}, []string{"foo", "baz"}, []Op{{OpInsert, 0, 0, "foo"}}},
		{[]string{"bar", "baz"}, []string{"foo", "baz"}, []Op{{OpDelete, 0, -1, "bar"}, {OpInsert, 1, 0, "foo"}}},
		{[]string{"foo", "bar", "baz"}, []string{"foo", "bar"}, []Op{{OpDelete, 2, -1, "baz"}}},
	}
	for _, c := range testCases {
		act := DiffStr(c.l1, c.l2)
		if !reflect.DeepEqual(c.exp, act) {
			t.Errorf("Failed diff, expected %v actual %v\n", c.exp, act)
		}
	}
}
