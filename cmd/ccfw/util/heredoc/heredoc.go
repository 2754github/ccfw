package heredoc

import (
	"fmt"
	"strings"
)

func Format(format string, a ...any) []byte {
	return []byte(strings.TrimPrefix(fmt.Sprintf(format, a...), "\n"))
}
