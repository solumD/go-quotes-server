package model

// Quote represents a quote.
type Quote struct {
	ID     int64  `json:"id"`
	Text   string `json:"quote"`
	Quthor string `json:"author"`
}
