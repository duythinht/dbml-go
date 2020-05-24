package annotations

func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }
func isAnnotation(ch rune) bool { return ch == '@' }
func isLetter(ch rune) bool     { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }
func isDigit(ch rune) bool      { return (ch >= '0' && ch <= '9') }
