package data

import "testing"

func TestChecksValidation(t *testing.T) {
  p := &Product{
    Name: "jon",
    Price: 1.00,
    SKU: "abs-qwer-zxc",
  }

  err := p.Validate()
  if err != nil {
    t.Fatal(err)
  }
}
