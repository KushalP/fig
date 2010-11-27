package fig

import "strings"
import "fmt"

func findFiles(fs FileSystem, baseDir string, pattern string) []string {
	var fsArr []string
	var pattArr []string

	// Wildcard matching method here
	if strings.Contains(pattern, "/") {
		pattArr = strings.Split(pattern, "/", -1)
	} else {
		pattArr = []string{pattern}
	}

	fsArr, _ = fs.List(baseDir)

	if len(fsArr) <= 0 {
		fsArr, _ = fs.List(pattArr[0])
	}

	fmt.Println("--- NEW FIND...")
	fmt.Println("pattArr[0]", pattArr[0])
	fmt.Println("fsArr:", fsArr)
	fmt.Println("pattArr:", pattArr)

	matches := recursiveMatch(fsArr, pattArr, 0)

	if baseDir != "" {
		var retArr []string

		for _, m := range matches {
			retArr = append(retArr, strings.TrimLeft(m, baseDir+"/"))
		}

		return retArr
	}

	return matches
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
