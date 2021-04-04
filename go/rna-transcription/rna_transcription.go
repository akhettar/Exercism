package strand

// ToRNA converts DNA strand to its transcribed RNA strand according to
// conversion rule below
// G -> C
// C -> G
// T -> A
// A -> U
func ToRNA(dna string) string {

	ch := make(chan string, len(dna))

	// launch
	go convert(dna, ch)

	var rna string
	for c := range ch {
		rna += c
	}
	return rna
}

func convert(dna string, ch chan<- string) {
	defer close(ch)
	dnaRna := map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}
	for _, s := range dna {
		ch <- dnaRna[string(s)]
	}
}
