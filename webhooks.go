package buttrbase

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"strconv"
	"time"
)

// VerifyWebhookSignature verifies a webhook signature.
//
// The signed payload is "<timestamp>.<body>" and the expected signature is the
// hex-encoded HMAC-SHA256 of that payload using the shared secret. The
// signatureHeader may optionally be prefixed with "sha256=". If toleranceSeconds
// is positive, the timestamp must be within that many seconds of now.
func VerifyWebhookSignature(body []byte, signatureHeader, timestampHeader, secret string, toleranceSeconds int) bool {
	if signatureHeader == "" || timestampHeader == "" || secret == "" {
		return false
	}
	ts, err := strconv.ParseInt(timestampHeader, 10, 64)
	if err != nil {
		return false
	}
	if toleranceSeconds > 0 {
		now := time.Now().Unix()
		diff := now - ts
		if diff < 0 {
			diff = -diff
		}
		if diff > int64(toleranceSeconds) {
			return false
		}
	}
	sig := signatureHeader
	if len(sig) > 7 && sig[:7] == "sha256=" {
		sig = sig[7:]
	}
	provided, err := hex.DecodeString(sig)
	if err != nil {
		return false
	}
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestampHeader))
	mac.Write([]byte("."))
	mac.Write(body)
	expected := mac.Sum(nil)
	return subtle.ConstantTimeCompare(provided, expected) == 1
}
