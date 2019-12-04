package hello

import "testing"

func TestGreetsGitHub(t *testing.T) {
	result := Greet("home")
	if result != "Hello GitHub Actions, home!" {
		t.Errorf("Greet() = %s; want Hello GitHub Actions, home!", result)
	}
}
