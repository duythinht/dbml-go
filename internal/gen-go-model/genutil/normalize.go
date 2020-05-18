package genutil

import "strings"

func NormalizeTypeName(s string) string {
	s1 := strings.ReplaceAll(s, " ", "_")
	if strings.HasSuffix(s1, "us") {
		return s1
	}

	if strings.HasSuffix(s1, "ies") {
		return s1[:len(s1)-3] + "y"
	}

	for _, suffix := range []string{"oes", "ses", "zes", "xes", "shes", "ches"} {
		if strings.HasSuffix(s1, suffix) {
			return s1[:len(s1)-2]
		}
	}
	if strings.HasSuffix(s1, "s") {
		return s1[:len(s1)-1]
	}

	return s1
}

func Normalize(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}

func NormalLizeGoName(s string) string {
	return GoInitialismCamelCase(Normalize(s))
}

func NormalizeGoTypeName(s string) string {
	return GoInitialismCamelCase(NormalizeTypeName(s))
}
