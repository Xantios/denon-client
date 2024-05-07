package client

import (
	"fmt"
	"testing"
)

func TestMute(t *testing.T) {

	// @TODO: Not hardcoded this
	client := New("10.0.61")

	client.SetMute(true)
	mute := client.GetMute()
	fmt.Println("Mute :: ", mute)

	if mute != true {
		t.Errorf("Mute is false, expected true")
	}
}
