package digits

// DigitsAPI required protocol and domain
const DigitsAPI = "https://api.digits.com"

// Account is a Digits user account
type Account struct {
	AccessToken AccessToken `json:"access_token"`
	ID          int64       `json:"id"`
	IDStr       string      `json:"id_str"`
	CreatedAt   string      `json:"created_at"`
	PhoneNumber string      `json:"phone_number"`
}

// AccessToken is a Digits OAuth1 access token and secret
type AccessToken struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
}
