package ghostlib

import (
	"io/ioutil"
	"path/filepath"
	"os/exec"
	"strings"
	"math/rand"
	"regexp"
)

var (
	ghostFilenameReg *regexp.Regexp
)

func init() {
	//rand.Seed(time.Now().Unix())
	rand.Seed(1)
	ghostFilenameReg = regexp.MustCompile(`"/?[^-]+-[^-]+-[^-]+-[^-]+\.png"`)
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
func CreateGhosts(backs []Image, bodies []Image, faces []Image, hats []Image, count int, exclude []Image) []Ghost {
	exc := make(map[string]bool)
	for _, e := range exclude {
		exc[e.name] = true
	}

	//fmt.Println("len(exc):", len(exc))
	//fmt.Println("exc:", exc)

	var ghosts []Ghost
	for ; len(ghosts) < count; {
		ghost := Ghost{
			backs[rand.Intn(len(backs))],
			bodies[rand.Intn(len(bodies))],
			faces[rand.Intn(len(faces))],
			hats[rand.Intn(len(hats))],
		}
		if _, exists := exc[ghost.name()]; !exists {
			ghosts = append(ghosts, ghost)
		}
	}
	return ghosts

}
