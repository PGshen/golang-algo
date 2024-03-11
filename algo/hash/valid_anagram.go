package algohash

func validAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := map[rune]int{}
	tMap := map[rune]int{}
	for _, ch := range s {
		sMap[ch]++
	}
	for _, ch := range t {
		tMap[ch]++
	}
	for k, cnt := range sMap {
		if tMap[k] != cnt {
			return false
		}
	}
	return true
}
