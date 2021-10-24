package strand

var transcription = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	RNA := make([]rune, len(dna))
	for i, nuc := range dna {
		RNA[i] = transcription[nuc]
	}
	return string(RNA)
}
