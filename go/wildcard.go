package fig

import "strings"

func findFiles(fs FileSystem, baseDir string, pattern string) []string {
	if !strings.Contains(pattern, "*") {
		if fs.Exists(pattern) {
			return []string{strings.Trim(pattern, "")}
		}
	}
	return []string{}
}
