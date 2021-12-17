package ghostlib

import (
	"testing"
)

func eq(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
    t.Errorf("expected %v, but got %v", expected, actual)
  }
}

func TestGhostName(t *testing.T) {
	g := Ghost{	Image{ "back", "" }, Image{ "body", "" }, Image{ "face", "" }, Image{ "hat", "" } }
	eq(t, "back-body-face-hat", g.name())
}

func TestCreateGhosts(t *testing.T) {
	backs  := []Image {{ "back", "back_path" }}
	bodies := []Image {{ "body", "body_path" }}
	faces  := []Image {{ "face", "face_path" }}
	hats   := []Image {{ "hat", "hat_path" }}
	ghosts := CreateGhosts(backs, bodies, faces, hats, 1)
	eq(t, "back-body-face-hat", ghosts[0].name())
}
