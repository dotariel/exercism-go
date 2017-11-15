package protein

var codons = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(codon string) string {
	return codons[codon]
}

func FromRNA(rna string) []string {
	protein := make([]string, 0)

	for i := 0; i < len(rna); i += 3 {
		c := rna[i : i+3]

		if codon, ok := codons[c]; ok {
			if codon == "STOP" {
				return protein
			}

			protein = append(protein, codon)
		}
	}

	return protein
}
