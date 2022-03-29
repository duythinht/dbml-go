package annotations

import "strings"

// Parse ...
func Parse(s *Scanner) map[string]string {
	ret := make(map[string]string)
	for {
		// Get annotation
		tok, annotation := s.Scan()
		if tok != ANNOTATION {
			break
		}

		tok, _ = s.Scan()
		if tok != LPAREN {
			break
		}

		tok, value := s.Scan()
		if tok != VALUE {
			break
		}

		tok, _ = s.Scan()
		if tok != RPAREN {
			break
		}

		ret[strings.ToLower(annotation)] = value

		tok, _ = s.Scan()
		if tok != COMMA {
			break
		}
	}
	return ret
}
