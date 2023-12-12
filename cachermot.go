package hangman

import "math/rand"

func CacherMot(mot string) string {
	n := len(mot)/2 - 1
	if n < 1 {
		n = 1
	}
	motCache := make([]byte, len(mot))
	for i := range mot {
		motCache[i] = '_'
	}
	for i := 0; i < n; i++ {
		// Révéler une lettre aléatoire
		for {
			index := rand.Intn(len(mot))
			if motCache[index] == '_' {
				motCache[index] = mot[index]
				break
			}
		}
	}
	return string(motCache)
}
