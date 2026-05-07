package buttrbase

// ValidateCouponOptions holds optional parameters for ValidateCoupon.
type ValidateCouponOptions struct {
	UserID         *int    `json:"user_id,omitempty"`
	OrderTotalCents *int64 `json:"order_total_cents,omitempty"`
}

// CouponValidation is the response from ValidateCoupon.
type CouponValidation struct {
	Valid          bool    `json:"valid"`
	Code           string  `json:"code,omitempty"`
	DiscountCents  int64   `json:"discount_cents,omitempty"`
	DiscountType   string  `json:"discount_type,omitempty"`
	Reason         string  `json:"reason,omitempty"`
	Raw            map[string]any `json:"-"`
}

// GiftCardValidation is the response from ValidateGiftCard.
type GiftCardValidation struct {
	Valid        bool   `json:"valid"`
	Code         string `json:"code,omitempty"`
	BalanceCents int64  `json:"balance_cents,omitempty"`
	Reason       string `json:"reason,omitempty"`
	Raw          map[string]any `json:"-"`
}

// GiftCardRedemption is the response from RedeemGiftCard.
type GiftCardRedemption struct {
	Success            bool   `json:"success"`
	Code               string `json:"code,omitempty"`
	RedeemedCents      int64  `json:"redeemed_cents,omitempty"`
	RemainingCents     int64  `json:"remaining_cents,omitempty"`
	Raw                map[string]any `json:"-"`
}

// SendMagicLinkOptions holds optional parameters for SendMagicLink.
type SendMagicLinkOptions struct {
	RedirectURL string `json:"redirect_url,omitempty"`
	TTLSeconds  *int64 `json:"ttl_seconds,omitempty"`
}

// MagicLinkSend is the response from SendMagicLink.
type MagicLinkSend struct {
	Sent  bool   `json:"sent"`
	Email string `json:"email,omitempty"`
	Raw   map[string]any `json:"-"`
}

// MagicLinkVerify is the response from VerifyMagicLink.
type MagicLinkVerify struct {
	Valid  bool   `json:"valid"`
	Email  string `json:"email,omitempty"`
	UserID *int   `json:"user_id,omitempty"`
	Raw    map[string]any `json:"-"`
}

// MfaStatus is the response from MfaStatus.
type MfaStatus struct {
	Enrolled bool   `json:"enrolled"`
	Active   bool   `json:"active"`
	Label    string `json:"label,omitempty"`
	Raw      map[string]any `json:"-"`
}

// MfaEnrollment is the response from MfaEnroll.
type MfaEnrollment struct {
	Secret    string `json:"secret,omitempty"`
	OtpauthURL string `json:"otpauth_url,omitempty"`
	Label     string `json:"label,omitempty"`
	Raw       map[string]any `json:"-"`
}

// MfaStatusResponse is the response from MfaActivate.
type MfaStatusResponse struct {
	Active bool   `json:"active"`
	Label  string `json:"label,omitempty"`
	Raw    map[string]any `json:"-"`
}

// OrgSignResponse is the response from OrgSign.
type OrgSignResponse struct {
	Token string `json:"token"`
	Kid   string `json:"kid,omitempty"`
	Raw   map[string]any `json:"-"`
}

// JWKSResponse is the response from OrgJWKS.
type JWKSResponse struct {
	Keys []map[string]any `json:"keys"`
	Raw  map[string]any   `json:"-"`
}

// SecretGet is the response from GetSecret.
type SecretGet struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	Raw         map[string]any `json:"-"`
}

// SecretSummary is the response from PutSecret.
type SecretSummary struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Raw         map[string]any `json:"-"`
}
