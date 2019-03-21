/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 4.6.0
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Account currency detail
type MarginAccountCurrency struct {
	// Currency name
	Currency string `json:"currency,omitempty"`
	// Amount suitable for margin trading.
	Available string `json:"available,omitempty"`
	// Locked amount, used in margin trading
	Locked string `json:"locked,omitempty"`
	// Borrowed amount
	Borrowed string `json:"borrowed,omitempty"`
}
