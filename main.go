package randomizer

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func RandomString(length int) string {
	// Generate random string use Unix Time
	var time = fmt.Sprintf("%d", time.Now().UnixNano())

	sha := sha256.New()
	sha.Write([]byte(time))

	shaString := fmt.Sprintf("%x", sha.Sum(nil))

	return shaString[:length]
}
