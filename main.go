package main

import (
	"io/ioutil"
	"log"

	"github.com/skeidderf/codoner/codon"
)

func main() {

	dnaSeq, err := ioutil.ReadFile("NC_045512.txt")
	if err != nil {
		log.Fatal(err)
	}

	rnaSeq := []*codon.Nucleotide{}
	for _, r := range string(dnaSeq) {
		n, err := codon.Parse(r)
		if err != nil || n == nil {
			log.Fatal(err)
		}
		rnaSeq = append(rnaSeq, n)
	}

	var (
		n1 codon.Nucleotide
		n2 codon.Nucleotide
		n3 codon.Nucleotide
	)
	protSeq := []*codon.AminoAcid{}
	for i, r := range rnaSeq {
		n, err := codon.TranslateDNA(*r)
		if err != nil || n == nil {
			log.Fatal(err)
		}
		p := i % 3
		switch p {
		case 0:
			if i != 0 {
				prot := codon.Codons[n1][n2][n3]
				if prot != nil {
					protSeq = append(protSeq, prot)
				}
			}
			n1 = *n
		case 1:
			n2 = *n
		case 2:
			n3 = *n
		}
	}

	for _, p := range protSeq {
		print(string(*p))
	}
	println()
}
