// Package validators provides ...
package validators

import (
	"fmt"
	"regexp"
)

// Validates if a file is indented correctly if indentStyle is set to "space"
func Tab(line string, indentStyle string) bool {
	if indentStyle == "tab" && len(line) > 0 {
		regexpPattern := fmt.Sprintf("^\t*[^ \t]")
		matched, err := regexp.MatchString(regexpPattern, line)

		if err != nil {
			panic(err)
		}

		if matched {
			return true
		}

		return false
	}

	return true
}
