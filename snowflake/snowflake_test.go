package snowflake

import (
	"testing"
)

func TestNewSnowflake(t *testing.T) {
	var i, j int64
	for i = 0; i < 32; i++ {
		for j = 0; j < 32; j++ {
			_, err := NewSnowflake(i, j)
			if err != nil {
				t.Error(err)
			}
		}
	}
}

func TestNextVal(t *testing.T) {
	s, err := NewSnowflake(0, 0)
	if err != nil {
		t.Error(err)
	}
	var i int64
	for i = 0; i < sequenceMask*10; i++ {
		val := s.NextVal()
		if val == 0 {
			t.Fail()
		}
	}
}
