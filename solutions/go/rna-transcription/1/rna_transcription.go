package strand

func ToRNA(dna string) string {
    rna := ""
    
    dnaToRna := map[string]string {
        "G" : "C",
        "C" : "G",
        "T" : "A",
        "A" : "U",
    }

    for _, val := range dna {
        sVal := string(val)
        rnaVal, ok := dnaToRna[sVal]

        if !ok { // incorrect nucleotide
            rna = ""
            break
        }

        rna += rnaVal 
    }
    return rna
}
