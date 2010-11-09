package fig

import "strings"

func findFiles(fs FileSystem, baseDir string, pattern string) []string {
	// When there's no wildcard specified and only a single string provided
    if !strings.Contains(pattern, "*") {
		if fs.Exists(pattern) {
			return []string{strings.Trim(pattern, "")}
		}
	}
	return []string{}
}
