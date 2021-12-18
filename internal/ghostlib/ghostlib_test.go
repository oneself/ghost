package ghostlib

import (
	"testing"
)

func eq(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
    t.Errorf("expected %v, but got %v", expected, actual)
  }
}

func ghostNames(ghosts []Ghost) (names []string) {
	for _, g := range ghosts {
		names = append(names, g.name())
	}
	return
}

func TestGhostName(t *testing.T) {
	g := Ghost{	Image{ "back", "" }, Image{ "body", "" }, Image{ "face", "" }, Image{ "hat", "" } }
	eq(t, "back-body-face-hat", g.name())
}

func TestCreateGhosts(t *testing.T) {
	backs  := []Image {
		{ "back1", "back_path1" },
  	{ "back2", "back_path2" }}
	bodies := []Image {{ "body", "body_path" }}
	faces  := []Image {{ "face", "face_path" }}
	hats   := []Image {{ "hat", "hat_path" }}
	ghosts := CreateGhosts(backs, bodies, faces, hats, 1, []Image{})
	eq(t, 1, len(ghosts))
}


func TestCreateGhostsExclude(t *testing.T) {
	backs  := []Image {
		{ "back1", "back_path1" },
		{ "back2", "back_path2" }}
	bodies  := []Image {
		{ "body1", "body_path1" },
  	{ "body2", "body_path2" }}
	faces  := []Image {{ "face", "face_path" }}
	hats   := []Image {{ "hat", "hat_path" }}
	outs   := []Image {
		{ "back1-body1-face-hat", "out_path1" },
  	{ "back1-body2-face-hat", "out_path2" },
		{ "back2-body1-face-hat", "out_path3" }}
	ghosts := CreateGhosts(backs, bodies, faces, hats, 1, outs)
	eq(t, 1, len(ghosts))
	eq(t, "back2-body2-face-hat", ghosts[0].name())
}
