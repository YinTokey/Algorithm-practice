import "math"

func commonChars(A []string) []string {
	array := make([][]int, len(A))

	for i := 0; i < len(A); i++ {
		array[i] = make([]int, 26)
	}

	for section, word := range A {

		for row := 0; row < len(word); row++ {
			c := word[row]
			array[section][c-'a']++
		}
	}

	result := make([]string, 0)
	for j := 0; j < 26; j++ {
		var shouldOutput bool = true
		minCount := math.MaxInt64
		for i := 0; i < len(A); i++ {
			if array[i][j] == 0 {
				shouldOutput = false
			} else if array[i][j] < minCount {
				minCount = array[i][j]
			}
		}

		if shouldOutput {
			for x := 0; x < minCount; x++ {
				result = append(result, string('a'+j))
			}
		}

	}

	return result

}