package ghostlib

import (
	"io/ioutil"
	"path/filepath"
	"os/exec"
	"strings"
	"math/rand"
	"regexp"
)


// out/back04-yellow_body-smile-irish_cap_gray_irish_cap.png
// out/back06-white_body-surprised-party_hat_brown_party_hat.png
// out/back08-blue_body-cat-viking_helmet_pink_viking_helmet.png
// out/back08-pink_body-cat-sombrero_purple_sombrero.png
// out/back08-white_body-sad-viking_helmet_red_viking_helmet.png
// out/back09-green_body-surprised-baseball_cap_orange_baseball_cap.png
// out/back10-brown_body-strait-sombrero_red.png
// out/back11-beige_body-cat-wizard_hat_brown_wizard_hat.png
// out/back13-blue_body-sad-santa_hat_yellow_santa_hat.png
// out/back13-purple_body-very_happy-wizard_hat_red_wizard_hat.png

var (
	ghostFilename *regexp.Regexp
)

func init() {
	//rand.Seed(time.Now().Unix())
	rand.Seed(1)
	ghostFilename = regexp.MustCompile(`"/?[^-]+-[^-]+-[^-]+-[^-]+\.png"`)
}

type Image struct {
	name string
	path string
}

type Ghost struct {
	back Image
	body Image
	face Image
	hat Image
}

//func NewGhost(filename fn) Ghost {
//
//}

// Name the ghost based on its parts
func (g Ghost) name() string {
	return g.back.name + "-" + g.body.name + "-" + g.face.name + "-" + g.hat.name
}

// Scan dir and find all images.
func GetImageFilenames(dir string) ([]Image, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var images []Image
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".png") {
			name := file.Name()[:len(file.Name()) - 4]
			path := filepath.Join(dir, file.Name())
			images = append(images, Image{name, path})
		}
	}
	return images, nil
}

// Create a ghost images on disk under `outdir/out`.
func CreateGhostImageCommands(outdir string, ghosts []Ghost) []exec.Cmd {
	cmds := make([]exec.Cmd, len(ghosts))
	for i, ghost := range ghosts {
		cmd := exec.Command(
			"convert",
			ghost.back.path,
			ghost.body.path,
			ghost.hat.path,
			ghost.face.path,
			"-gravity",
			"Center",
			"-background",
			"None",
			"-layers",
			"Flatten",
			filepath.Join(outdir, ghost.name() + ".png"))
		cmds[i] = *cmd
	}
	return cmds
}

// Create a list of ghosts based on lists of ghost parts.
// The list will contain `count` ghosts picked at random.
func CreateGhosts(backs []Image, bodies []Image, faces []Image, hats []Image, count int) []Ghost {
	var ghosts = make([]Ghost, count)
	for i := 0; i < len(ghosts); i++ {
		ghosts[i] = Ghost{
			backs[rand.Intn(len(backs))],
			bodies[rand.Intn(len(bodies))],
			faces[rand.Intn(len(faces))],
			hats[rand.Intn(len(hats))],
		}
	}
	return ghosts
}
