package textwrap
import "strings"

func Wrap(text string, width int, firstIndent, indent string) string {
	parts := strings.Split(text, " ")
	var lines []string

	var buffer string = firstIndent
	var bufferLen int = len(firstIndent)
	for _, p := range parts {
		var strlen = len(p) + 1

		if bufferLen + strlen > width {
			lines = append(lines, buffer)
			buffer = indent
			bufferLen = len(indent)
		}

		buffer += p + " "
		bufferLen += strlen
	}
	lines = append(lines, buffer)
	return strings.Join(lines, "\n")
}
