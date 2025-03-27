package config

import "testing"

func TestValidWellknown(t *testing.T) {
	wk := WellKnown{
		Issuer:                "http://localhost:3000/",
		AuthorizationEndpoint: "http://localhost:3000/",
		TokenEndpoint:         "http://localhost:3000/",
		UserinfoEndpoint:      "http://localhost:3000/",
		RevocationEndpoint:    "http://localhost:3000/",
	}
	if !wk.isValid() {
		t.Errorf("invalid wellknown configuration should not be valid")
	}
}

func TestInvalidPartialWellknown(t *testing.T) {
	wk := WellKnown{
		Issuer:           "http://localhost:3000/",
		TokenEndpoint:    "http://localhost:3000/",
		UserinfoEndpoint: "http://localhost:3000/",
	}
	if wk.isValid() {
		t.Errorf("invalid wellknown configuration should not be valid")
	}
}

func TestInvalidMinimalWellknown(t *testing.T) {
	wk := WellKnown{}
	if wk.isValid() {
		t.Errorf("invalid wellknown configuration should not be valid")
	}
}
