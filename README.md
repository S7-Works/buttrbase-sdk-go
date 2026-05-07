# buttrbase-go

Go SDK for the Buttrbase API. Stdlib only.

## Install

```
go get github.com/buttrbase/buttrbase-go
```

Requires Go 1.21+.

## Validate a coupon

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/buttrbase/buttrbase-go"
)

func main() {
    client := buttrbase.New("YOUR_API_KEY")

    res, err := client.ValidateCoupon(context.Background(), "WELCOME10", nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("valid=%v discount_cents=%d\n", res.Valid, res.DiscountCents)
}
```

## Verify a webhook signature

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/buttrbase/buttrbase-go"
)

func handler(w http.ResponseWriter, r *http.Request) {
    body := make([]byte, r.ContentLength)
    r.Body.Read(body)

    ok := buttrbase.VerifyWebhookSignature(
        body,
        r.Header.Get("X-Buttrbase-Signature"),
        r.Header.Get("X-Buttrbase-Timestamp"),
        "YOUR_WEBHOOK_SECRET",
        300, // tolerance in seconds
    )
    if !ok {
        http.Error(w, "invalid signature", http.StatusUnauthorized)
        return
    }
    fmt.Fprintln(w, "ok")
}
```
