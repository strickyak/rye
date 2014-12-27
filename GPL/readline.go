package GPL

import "io"
import "github.com/shavac/readline"

func ReadLine(prompt string) string {
	maybeLine := readline.ReadLine(&prompt)
	if maybeLine == nil {
		panic(io.EOF)
	}
	readline.AddHistory(*maybeLine)
	return *maybeLine
}

func ReadHistoryFile(filename string) {
	readline.ReadHistoryFile(filename)
}

func WriteHistoryFile(filename string) {
	readline.WriteHistoryFile(filename)
}
