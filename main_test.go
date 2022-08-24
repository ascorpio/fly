package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)
	<-timer.C
	fmt.Println("3 second")
}
