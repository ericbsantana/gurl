package formatter

import (
	"fmt"
	"strings"
)

func SplitResponse(response string) (header, body string) {
	parts := strings.Split(response, "\r\n\r\n")
	return parts[0], parts[1]
}

func PrintHeaderLines(header string) {
	lines := strings.Split(header, "\n")
	for _, line := range lines {
		fmt.Println("<" + line)
	}
}
func PrintRequestLines(request string) {
	lines := strings.Split(request, "\n")
	for _, line := range lines[:len(lines)-1] {
		fmt.Println(">" + line)
	}
}
