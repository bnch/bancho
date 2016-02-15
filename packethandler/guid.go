package packethandler

import (
	"math/rand"
	"time"
)

const gs = "-"

// GenerateGUID makes a random GUID.
func GenerateGUID() string {
	return "" +
		genRandomString(8) + gs +
		genRandomString(4) + gs +
		genRandomString(4) + gs +
		genRandomString(4) + gs +
		genRandomString(12)
}

const letterBytes = "0123456789abcdef"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var randSrc = rand.NewSource(time.Now().UnixNano())

// Gen generates a random string.
func genRandomString(n int) string {
	b := make([]byte, n)
	// A randSrc.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
