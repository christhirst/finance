package helper

import (
	"reflect"
	"testing"
)

func TestRandom(t *testing.T) {
	max := 5
	for i := 1; i <= max; i++ {
		rand := random(i)
		if rand < 0 && rand < max {
			t.Errorf("%d", rand)
		}
	}
}

func TestRandomInRange(t *testing.T) {
	for i := 1; i <= 50; i += 10 {
		if RandomInRange(1, i+1) < i && RandomInRange(1, i+1) > i {
			t.Errorf("%d", RandomInRange(1, i+1))
		}
	}
}

func TestRandomString(t *testing.T) {
	for i := 1; i <= 5; i++ {
		s := RandomString(random(i))
		if reflect.TypeOf(s).String() != reflect.String.String() {
			t.Error(reflect.TypeOf(s).String())
		}
	}
}
