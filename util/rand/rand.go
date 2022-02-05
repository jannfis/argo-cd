package rand

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandString generates, from a given charset, a random string of a given length.
func RandString(n int) string {
	return RandStringCharset(n, letterBytes)
}

func RandStringCharset(n int, charset string) string {
	b := make([]byte, n)
	// RandInt63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, RandInt63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = RandInt63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(charset) {
			b[i] = charset[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// RandInt63 returns a random number with 63 bits of entropy, suitable for
// cryptographic operations. This method panics on errors returned from
// the underlying random implementation (usually, when there's not enough
// random data available), because it's better to be safe than sorry.
func RandInt63() int64 {
	const wantBytes = 8
	const tryRead = 10

	// b will store wantBytes of random data
	b := make([]byte, wantBytes)

	var n int
	var err error
	for i := 0; i < tryRead; i++ {
		n, err = rand.Read(b)
		if err == nil {
			break
		}
	}
	// Tried 10 times, it's time to panic if we couldn't read random data
	if err != nil {
		panic(fmt.Sprintf("could not read random data: %v", err))
	}

	// Ensure we read all of the required bytes
	if n != wantBytes {
		panic(fmt.Sprintf("did not read enough sufficient random data for seeding: wanted=%d, read=%d", wantBytes, n))
	}

	// Convert the data we have read to a signed 64 bit integer, which provides
	// us with 63 bits of entropy.
	return int64(binary.BigEndian.Uint64(b))
}
