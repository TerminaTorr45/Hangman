package hangmanpackage

func RevelerLettres(mot, motCache, lettre string) string {
	revele := []byte(motCache)
	for i := range mot {
		if mot[i:i+1] == lettre {
			revele[i] = mot[i]
		}
	}
	return string(revele)
}
