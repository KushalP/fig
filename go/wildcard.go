package fig

import "strings"

func findFiles(fs FileSystem, baseDir string, pattern string) []string {
	modifiers := []string{"?", "*"}

    // We're using a wildcard
    if strings.Contains(pattern, modifiers[0]) || strings.Contains(pattern, modifiers[1]) {
        // Wildcard matching method here
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
    cards := strings.Split(pattern, "*", 2000);

    for _, str := range cards {
        idx := strings.Index(text, str)

        if idx == -1 {
            return false
        }

        text = strings.TrimLeft(text, str + "*")
    }

    return true
}
