package fig

import "strings"
import "fmt"

func findFiles(fs FileSystem, baseDir string, pattern string) []string {
    var fsArr []string
    modifiers := []string{"?", "*"}

    // We're using a wildcard
    if strings.Contains(pattern, modifiers[0]) || strings.Contains(pattern, modifiers[1]) {
        // Wildcard matching method here
        pattArr := strings.Split(pattern, "/", -1)
        fsArr, _ = fs.List(baseDir)

        if len(fsArr) <= 0 {
            fsArr, _ = fs.List(pattArr[0])
        }

        return recursiveMatch(fsArr, pattArr, 0)
    } else {
        // Not using a wildcard
        if fs.Exists(pattern) {
            return []string{strings.Trim(pattern, "")}
        }
    }

	return []string{}
}

// Matches a given string against a wildcard pattern
func wildcardMatch(text string, pattern string) bool {
    cards := strings.Split(pattern, "*", -1);

    for _, str := range cards {
        idx := strings.Index(text, str)

        if idx == -1 {
            return false
        }

        text = strings.TrimLeft(text, str + "*")
    }

    return true
}

func recursiveMatch(fsArr []string, pArr []string, pPos int) []string {
    fmt.Println("fsArr: \t", fsArr)
    fmt.Println("pArr: \t", pArr)
    fmt.Println("pPos: \t", pPos)

    var retArr []string
    var pClean string

    // Track the differences in fsArr length
    fsArrLen := len(fsArr)

    if pPos == len(pArr) {
        return fsArr
    }

    for i := 0; i < len(fsArr); i++ {
        fsSlashArr := strings.Split(fsArr[i], "/", -1)

        for j := pPos; j < len(fsSlashArr); j++ {
            pClean = strings.Replace(strings.Replace(pArr[pPos], "?", "", -1), "*", "", -1)
            fmt.Println("pClean: \t", pClean)

            if strings.TrimSpace(pClean) == "" {
                continue
            }

            if strings.Index(fsSlashArr[j], pClean) >= 0 {
                retArr = append(retArr, fsArr[i])
            }
        }
    }

    if fsArrLen == len(retArr) {
        return recursiveMatch(retArr, pArr, pPos + 1)
    }

    return retArr
}
