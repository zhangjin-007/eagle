package test

import (
	"github.com/prometheus/common/log"
	"testing"
)

func TestMake(t *testing.T) {
	ch := make(chan int)
	log.Info(ch)
}