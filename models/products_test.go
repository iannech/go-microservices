package models

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "ian",
		Price: 2.00,
		SKU:   "abc-abc-abc",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
