package anagram

import ("unicode/utf8"
        "reflect"
        "strings"
        )
// func Detect is used to find anagrams for given subject word
func Detect(subject string, candidates []string) []string {
	if subject == "" || len(candidates) == 0 {
        return []string{}
    }
    
	anagrams := make([]string, 0)
    lettersInSubject := make(map[rune]int)
	subjectLength, candidateWordLength := 0, 0
    subject = strings.ToLower(subject) // to ignore case
    
    for _, letter := range subject { // Fill out the lettersInSubject map
        _, exists := lettersInSubject[letter]
        if !exists {
            lettersInSubject[letter] = 1
        } else {
            lettersInSubject[letter]++
        }
    }

    subjectLength = utf8.RuneCountInString(subject)
    for _, word := range candidates {
        candidateWordLength = utf8.RuneCountInString(word)
        // if words are equal or has different length - skip
        if candidateWordLength != subjectLength || subject == strings.ToLower(word){
            continue
        }
        
        lettersInCandidateWord := make(map[rune]int) // Init for every word to compare
        for _, letter := range strings.ToLower(word) {
            _, exists := lettersInCandidateWord[letter]
            if !exists {
                lettersInCandidateWord[letter] = 1
            } else {
                lettersInCandidateWord[letter]++
            }
        }
        // check equality between subject word and candidate word
        if reflect.DeepEqual(lettersInSubject, lettersInCandidateWord) {
            anagrams = append(anagrams, word)
        }
    }

    return anagrams
}
