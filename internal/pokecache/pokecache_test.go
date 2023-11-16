package pokecache

import (
	"testing"
	"time"
)

func TestCreateCashe(t *testing.T) {
	cashe := NewCashe(time.Millisecond)

	if cashe.cashe == nil {
		t.Error("cashe is nil")
	}
}

func TestAddGetCashe(t *testing.T) {
	cashe := NewCashe(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "key3",
			inputVal: []byte("val3"),
		},
	}

	for _, cas := range cases {
		cashe.Add(cas.inputKey, cas.inputVal)
		actual, ok := cashe.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s does not match %s", string(actual), string(cas.inputVal))
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cashe := NewCashe(interval)

	keyOne := "key1"
	cashe.Add(keyOne, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cashe.Get("key1")
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}
