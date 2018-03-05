package test

import (
	"crypto/rand"
	"fmt"
)

// uuid generates a UUID that isn't particularly compliant
// to any spec but is more than good enough for our use case
func uuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	str := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return str
}
