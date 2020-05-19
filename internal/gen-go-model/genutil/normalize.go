package genutil

import "strings"

// NormalizeTypeName return name of model type, which normalize
// eg:
//	users => User
// 	categories => category
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

// Normalize return name in underscore
func Normalize(s string) string {
	return strings.ReplaceAll(s, " ", "_")
}

// NormalLizeGoName return normalize GoName
func NormalLizeGoName(s string) string {
	return GoInitialismCamelCase(Normalize(s))
}

//NormalizeGoTypeName return Normalize for Go Type Name
func NormalizeGoTypeName(s string) string {
	return GoInitialismCamelCase(NormalizeTypeName(s))
}
