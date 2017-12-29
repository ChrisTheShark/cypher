package cypher_test

import (
	"testing"
	. "github.com/ChrisTheShark/cypher"
)

// Illustrating simplified Caesar shift to encode a message by shifting
// character value by a desired amount. Receiver would need to know the shift
// to decode properly.
func TestShift(t *testing.T) {
	message := "thisistheencodedmessage"

	message, err := Shift(message, 1)
	if err != nil {
		t.Errorf("Shift invocation failed due to: %v", err.Error())
	} else if message != "uijtjtuiffodpefenfttbhf" {
		t.Errorf("Shift invocation resulted in unexpected output.")
	}
	t.Logf("Encoded message: %v", message)
}

func TestShiftOverlap(t *testing.T) {
	message := "tuvwxyzabcdefghijkl"

	message, err := Shift(message, 1)
	if err != nil {
		t.Errorf("Shift invocation failed due to: %v", err.Error())
	} else if message != "uvwxyzabcdefghijklm" {
		t.Errorf("Shift invocation resulted in unexpected output.")
	}
	t.Logf("Encoded message: %v", message)
}