package main

import (
  "github.com/oneself/ghost/internal/ghostlib"

	"flag"
	"os"
	"path/filepath"
	"fmt"
)

var (
	basedir *string
	count *int
)


func init() {
	basedir = flag.String("dir", "dir", "base directory for image files containing Bodies, Faces, and Hats")
	count   = flag.Int("count", 0, "number of ghosts to generate")
}

func main() {
	flag.Parse()
	fmt.Println("basedir:", *basedir)

	backs, backsErr := ghostlib.GetImageFilenames(filepath.Join(*basedir, "backs"))
	bodies, bodiesErr := ghostlib.GetImageFilenames(filepath.Join(*basedir, "bodies"))
	faces, facesErr := ghostlib.GetImageFilenames(filepath.Join(*basedir, "faces"))
	hats, hatsErr := ghostlib.GetImageFilenames(filepath.Join(*basedir, "hats"))
	handleErrors(backsErr, bodiesErr, facesErr, hatsErr)

	//fmt.Println("bodies:", bodies)
  //fmt.Println("faces:", faces)
	//fmt.Println("hats:", hats)
	//fmt.Println("backs:", backs)

	outdir := filepath.Join(*basedir, "out")
	os.Mkdir(outdir, 0755)
	ghosts := ghostlib.CreateGhosts(backs, bodies, faces, hats, *count)
	cmds := ghostlib.CreateGhostImageCommands(outdir, ghosts)

	for _, cmd := range cmds {
		fmt.Println(cmd.String())
		err := cmd.Run()
		handleErrors(err)
	}
}

// If a none-nil error is found, panic
func handleErrors(errs... error) {
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}
