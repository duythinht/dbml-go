package genutil

import "testing"

func TestNormalizeType(t *testing.T) {
	for s, s1 := range map[string]string{
		"heroes":     "hero",
		"countries":  "country",
		"watches":    "watch",
		"the_heroes": "the_hero",
	} {
		if out := NormalizeTypeName(s); out != s1 {
			t.Fatalf("in: %s, out: %s, expected: %s", s, out, s1)
		}
	}
}

func TestNormalizeGoType(t *testing.T) {
	for s, s1 := range map[string]string{
		"heroes":     "Hero",
		"countries":  "Country",
		"watches":    "Watch",
		"the_heroes": "TheHero",
	} {
		if out := NormalizeGoTypeName(s); out != s1 {
			t.Fatalf("in: %s, out: %s, expected: %s", s, out, s1)
		}
	}

}
