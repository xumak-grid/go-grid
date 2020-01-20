package pgen_test

import (
	"fmt"
	"testing"

	"github.com/xumak-grid/go-grid/pkg/pgen"
)

func TestGeneratePassword(t *testing.T) {
	psec := pgen.NewGenerator()

	pss, err := psec.GeneratePassword(pgen.MinimumLength+5, true)
	if err != nil {
		t.Error("the err GeneratePassword should be nil")
	}

	if len(pss) != pgen.MinimumLength+5 {
		t.Error("Len of password should be MinimumLength+5", len(pss))
	}

	// Test minimum password
	pss, err = psec.GeneratePassword(pgen.MinimumLength-1, true)
	if err == nil {
		t.Error("the GeneratePassword should return an error")
		return
	}
	if err != pgen.ErrMinLengthRequired {
		t.Error("the error type should be ErrMinLengthRequired type")
	}
}

// Example generates 3 passwords
func Example() {
	g := pgen.NewGenerator()
	p1, _ := g.GeneratePassword(10, true)
	p2, _ := g.GeneratePassword(15, true)
	fmt.Println(p1, p2) 

	// Changing specials characters (numerics, letters, specials)
	g.CharacterTables["specials"] = "*/-$"
	p3, _ := g.GeneratePassword(20, true)
	fmt.Println(p3) 
}
