package buttrbase

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strconv"
	"testing"
	"time"
)

func smokeClient(t *testing.T) *Client {
	t.Helper()
	key := os.Getenv("BUTTRBASE_SMOKE_API")
	if key == "" {
		t.Skip("BUTTRBASE_SMOKE_API not set; skipping smoke test")
	}
	opts := []Option{}
	if base := os.Getenv("BUTTRBASE_BASE_URL"); base != "" {
		opts = append(opts, WithBaseURL(base))
	}
	return New(key, opts...)
}

func TestValidateCoupon_Nonexistent(t *testing.T) {
	c := smokeClient(t)
	res, err := c.ValidateCoupon(context.Background(), "NONEXISTENT", nil)
	if err != nil {
		t.Fatalf("ValidateCoupon error: %v", err)
	}
	if res.Valid {
		t.Fatalf("expected valid=false for NONEXISTENT, got true")
	}
}

func TestValidateGiftCard_Nonexistent(t *testing.T) {
	c := smokeClient(t)
	res, err := c.ValidateGiftCard(context.Background(), "NONEXISTENT")
	if err != nil {
		t.Fatalf("ValidateGiftCard error: %v", err)
	}
	if res.Valid {
		t.Fatalf("expected valid=false for NONEXISTENT, got true")
	}
}

func TestVerifyWebhookSignature_RoundTrip(t *testing.T) {
	secret := "test-secret"
	body := []byte(`{"event":"test","data":{"id":1}}`)
	ts := strconv.FormatInt(time.Now().Unix(), 10)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(ts))
	mac.Write([]byte("."))
	mac.Write(body)
	sig := hex.EncodeToString(mac.Sum(nil))

	if !VerifyWebhookSignature(body, sig, ts, secret, 300) {
		t.Fatalf("expected signature to verify")
	}
	if !VerifyWebhookSignature(body, "sha256="+sig, ts, secret, 300) {
		t.Fatalf("expected prefixed signature to verify")
	}
	if VerifyWebhookSignature(body, sig, ts, "wrong-secret", 300) {
		t.Fatalf("expected wrong secret to fail")
	}
	if VerifyWebhookSignature([]byte("tampered"), sig, ts, secret, 300) {
		t.Fatalf("expected tampered body to fail")
	}
	oldTs := strconv.FormatInt(time.Now().Unix()-10000, 10)
	mac2 := hmac.New(sha256.New, []byte(secret))
	mac2.Write([]byte(oldTs))
	mac2.Write([]byte("."))
	mac2.Write(body)
	oldSig := hex.EncodeToString(mac2.Sum(nil))
	if VerifyWebhookSignature(body, oldSig, oldTs, secret, 300) {
		t.Fatalf("expected stale timestamp to fail")
	}
}
