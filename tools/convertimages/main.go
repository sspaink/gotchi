package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
)

const binDir = "../../bin"

func filenameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
}

func buildUtility(s *spinner.Spinner) {
	s.Start()
	defer s.Stop()
	cmd := exec.Command("go", "build", "-C", "gopherbadgeimg")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func runUtility(s *spinner.Spinner, filename string, ratio string) {
	s.Start()
	defer s.Stop()

	cmd := exec.Command("./gopherbadgeimg/gopherbadgeimg.exe", "-outmode", "bin", "-ratio", ratio, filename)
	if output, err := cmd.CombinedOutput(); err != nil {
		panic("error: " + err.Error() + "output: " + string(output))
	}

	// gopherbadgeimg names the file using the provided ratio, TODO propose adding a filename flag
	if err := os.MkdirAll(binDir, os.ModePerm); err != nil {
		panic(err)
	}

	if err := os.Rename(ratio+".bin", filepath.Join(binDir, filenameWithoutExtension(filename)+".bin")); err != nil {
		panic(err)
	}
}

func main() {
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)

	if err := os.RemoveAll(binDir); err != nil {
		panic(err)
	}

	buildUtility(s)

	err := filepath.WalkDir("../../assets", func(path string, d os.DirEntry, err error) error {
		if filepath.Ext(path) != ".png" {
			return nil
		}

		fmt.Println("processing ", path, "....")
		runUtility(s, path, "64x64")

		return nil
	})
	if err != nil {
		panic(err)
	}
}
