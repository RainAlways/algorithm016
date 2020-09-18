package other

import (
	"log"
	"testing"
)

func TestDecodeString(t *testing.T) {
	log.Println(decodeString("3[a2[c]]"))
}
