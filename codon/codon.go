package codon

// Codons are the codon triplet translation into amino acid
var Codons = map[Nucleotide]map[Nucleotide]map[Nucleotide]*AminoAcid{
	Guanine: {
		Uracil: map[Nucleotide]*AminoAcid{
			Uracil:   &Valine,
			Cytosine: &Valine,
			Adenine:  &Valine,
			Guanine:  &Valine,
		},
		Cytosine: map[Nucleotide]*AminoAcid{
			Uracil:   &Alanine,
			Cytosine: &Alanine,
			Adenine:  &Alanine,
			Guanine:  &Alanine,
		},
		Adenine: map[Nucleotide]*AminoAcid{
			Uracil:   &Alanine,
			Cytosine: &Alanine,
			Adenine:  &GlutamicAcid,
			Guanine:  &GlutamicAcid,
		},
		Guanine: map[Nucleotide]*AminoAcid{
			Uracil:   &Glycine,
			Cytosine: &Glycine,
			Adenine:  &Glycine,
			Guanine:  &Glycine,
		},
	},
	Uracil: {
		Uracil: map[Nucleotide]*AminoAcid{
			Uracil:   &Phenylalanine,
			Cytosine: &Phenylalanine,
			Adenine:  &Leucine,
			Guanine:  &Leucine,
		},
		Cytosine: map[Nucleotide]*AminoAcid{
			Uracil:   &Serine,
			Cytosine: &Serine,
			Adenine:  &Serine,
			Guanine:  &Serine,
		},
		Adenine: map[Nucleotide]*AminoAcid{
			Uracil:   &Tyrosine,
			Cytosine: &Tyrosine,
			Adenine:  nil,
			Guanine:  nil,
		},
		Guanine: map[Nucleotide]*AminoAcid{
			Uracil:   &Cysteine,
			Cytosine: &Cysteine,
			Adenine:  nil,
			Guanine:  &Tryptophan,
		},
	},
	Cytosine: {
		Uracil: map[Nucleotide]*AminoAcid{
			Uracil:   &Leucine,
			Cytosine: &Leucine,
			Adenine:  &Leucine,
			Guanine:  &Leucine,
		},
		Cytosine: map[Nucleotide]*AminoAcid{
			Uracil:   &Proline,
			Cytosine: &Proline,
			Adenine:  &Proline,
			Guanine:  &Proline,
		},
		Adenine: map[Nucleotide]*AminoAcid{
			Uracil:   &Histidine,
			Cytosine: &Histidine,
			Adenine:  &Glutamine,
			Guanine:  &Glutamine,
		},
		Guanine: map[Nucleotide]*AminoAcid{
			Uracil:   &Arginine,
			Cytosine: &Arginine,
			Adenine:  &Arginine,
			Guanine:  &Arginine,
		},
	},
	Adenine: {
		Uracil: map[Nucleotide]*AminoAcid{
			Uracil:   &Isoleucine,
			Cytosine: &Isoleucine,
			Adenine:  &Isoleucine,
			Guanine:  &Methionine,
		},
		Cytosine: map[Nucleotide]*AminoAcid{
			Uracil:   &Threonine,
			Cytosine: &Threonine,
			Adenine:  &Threonine,
			Guanine:  &Threonine,
		},
		Adenine: map[Nucleotide]*AminoAcid{
			Uracil:   &Asparagine,
			Cytosine: &Asparagine,
			Adenine:  &Lysine,
			Guanine:  &Lysine,
		},
		Guanine: map[Nucleotide]*AminoAcid{
			Uracil:   &Serine,
			Cytosine: &Serine,
			Adenine:  &Arginine,
			Guanine:  &Arginine,
		},
	},
}
