package fig

import "strings"
import "fmt"

func findFiles(fs FileSystem, baseDir string, pattern string) []string {
	var fsArr []string
	var pattArr []string

	fmt.Println("--- NEW FIND...")

	// Deal with patterns that don't have wildcards quickly
	if !hasWildcards(pattern) {
		// We've got a direct file, not a directory structure
		if !strings.Contains(pattern, "/") && fs.Exists(pattern) {
			return []string{pattern}
		}

		// We've got a directory structure
		withoutWildcards, _ := fs.List(pattern)
		return withoutWildcards
	}

	// Wildcard matching method here
	if strings.Contains(pattern, "/") {
		pattArr = strings.Split(pattern, "/", -1)
	} else {
		pattArr = []string{pattern}
	}

	fsArr, _ = fs.List(baseDir)
	fmt.Println("#1 fsArr:", fsArr)

	if len(fsArr) <= 0 && baseDir == "" {
		fsArr, _ = fs.List(trimWildcardRight(pattArr[0]))
		fmt.Println("#2 fsArr:", fsArr)
	}

	fmt.Println("fsArr val:", trimWildcardRight(pattArr[0]))
	fmt.Println("pattArr[0]", pattArr[0])
	fmt.Println("fsArr:", fsArr)
	fmt.Println("pattArr:", pattArr)

	matches := recursiveMatch(fsArr, pattArr, 0)

	// Strip out baseDir from the return array since we know it already
	if baseDir != "" {
		var retArr []string

		for _, m := range matches {
			retArr = append(retArr, strings.TrimLeft(m, baseDir+"/"))
		}

		return retArr
	}

	return matches
}

// Are we dealing with a pattern using wildcards?
func hasWildcards(pattern string) bool {
	return strings.IndexAny(pattern, "*?") != -1
}

// Strips away the string to find the first complete string
func trimWildcardRight(str string) string {
	var splitStr []string

	splitStr = strings.Split(str, "*", -1)
	splitStr = strings.Split(splitStr[0], "?", -1)

	return splitStr[0]
}

func recursiveMatch(fsArr []string, pArr []string, pPos int) []string {
	fmt.Println("fsArr: \t", fsArr)
	fmt.Println("pArr: \t", pArr)
	fmt.Println("pPos: \t", pPos)

	var retArr []string

	// Track the differences in fsArr length
	fsArrLen := len(fsArr)

	if pPos >= len(pArr) {
		return fsArr
	}

	for i := 0; i < fsArrLen; i++ {
		fsSlashArr := strings.Split(fsArr[i], "/", -1)

		for j := pPos; j < len(fsSlashArr); j++ {
			if wildcardMatch(fsSlashArr[j], pArr[pPos]) {
				fmt.Println("Wildcard Match!")
				retArr = append(retArr, fsArr[i])
			}
		}
	}

	if fsArrLen >= len(retArr) {
		return recursiveMatch(retArr, pArr, pPos+1)
	}

	return retArr
}

// Matches a given string against a wildcard pattern
func wildcardMatch(text string, pattern string) bool {
	cards := strings.Split(pattern, "*", -1)

	/*
	   for _, sc := range starCards {
	       scQuestionMark := strings.Split(sc, "?", -1)

	       for _, splitSc := range scQuestionMark {
	           cards = append(cards, splitSc)
	       }
	   }
	*/

	fmt.Println("WC Cards:", cards)

	for _, str := range cards {
		idx := strings.Index(text, str)

		if idx == -1 {
			return false
		}

		text = strings.TrimLeft(text, str+"*")
	}

	return true
}
