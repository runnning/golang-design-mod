package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {

	for i := 0; i < 30; i++ {
		go getInstance()
	}
}
