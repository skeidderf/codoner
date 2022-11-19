package codon

import "fmt"

// Nucleotide is an enum of possible nucleotides
type Nucleotide rune

var (
	// Guanine is G nucleotide
	Guanine Nucleotide = 'G'

	// Adenine is A nucleotide
	Adenine Nucleotide = 'A'

	// Thymine is U nucleotide
	Thymine Nucleotide = 'T'

	// Cytosine is C nucleotide
	Cytosine Nucleotide = 'C'

	// Uracil is U nucleotide
	Uracil Nucleotide = 'U'
)

// TranslateDNA translates a DNA nucleotide into the RNA counterpart
func TranslateDNA(n Nucleotide) (*Nucleotide, error) {
	switch n {
	case Thymine:
		return &Adenine, nil
	case Adenine:
		return &Uracil, nil
	case Cytosine:
		return &Guanine, nil
	case Guanine:
		return &Cytosine, nil
	}
	return nil, fmt.Errorf("Non translatable nucleotide")
}

// Parse converts rune into Nucleotides
func Parse(r rune) (*Nucleotide, error) {
	switch r {
	case rune(Guanine):
		return &Guanine, nil
	case rune(Adenine):
		return &Adenine, nil
	case rune(Thymine):
		return &Thymine, nil
	case rune(Cytosine):
		return &Cytosine, nil
	case rune(Uracil):
		return &Uracil, nil
	}
	return nil, fmt.Errorf("failed to parse rune into nucleotide")
}
