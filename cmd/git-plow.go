package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	args "github.com/integrii/flaggy"
)

const (
	TempDirName = "gmm_tmp_*"
)

var Version string = "0.0.0"

// Runs the command with the given args in the given dir or current dir if not specified
func run(dir string, name string, args ...string) (data string, ok bool, err error) {
	cmd := exec.Command(name, args...)

	if len(dir) > 0 {
		cmd.Dir = dir
	}

	out, err := cmd.Output()
	outStr := strings.Trim(string(out), " \n\t")

	if err != nil {
		return outStr, false, err
	}

	return outStr, true, err
}

func main() {
	var subdir string
	var url string
	var checkoutRef string
	var keepStructure bool

	fmt.Print("git-plow v", Version, "\n")

	args.AddPositionalValue(&url, "url", 1, true, "The repo to clone")
	args.AddPositionalValue(&subdir, "subdir", 2, true, "The subdirectory to checkout")

	// args.String(&checkoutRef, "r", "ref", "Specify a branch/tag or any tree ref to checkout")
	args.Bool(&keepStructure, "k", "keep", "Keep originial structure and do not elevate the subdir")
	args.Parse()

	tmpDir, err := os.MkdirTemp("", TempDirName)

	if err != nil {
		fmt.Println("Cannot create temp folder")
		os.Exit(1)
	}

	subdir = strings.TrimPrefix(subdir, "/")
	subdir = strings.TrimSuffix(subdir, "/")

	if _, ok, _ := run(tmpDir, "git", "clone", "-n", "--depth=1", "--filter=tree:0", url, "."); !ok {
		fmt.Println("Cannot clone")
		os.Exit(1)
	}

	if _, ok, _ := run(tmpDir, "git", "sparse-checkout", "set", "--no-cone", subdir); !ok {
		fmt.Println("Cannot sparse checkout")
		os.Exit(1)
	}

	if len(checkoutRef) == 0 {
		if _, ok, _ := run(tmpDir, "git", "checkout"); !ok {
			fmt.Println("Cannot checkout")
			os.Exit(1)
		}
	} else {
		if _, ok, _ := run(tmpDir, "git", "fetch", "--tags"); !ok {
			fmt.Println("Cannot fetch all tags")
			os.Exit(1)
		}

		if _, ok, _ := run(tmpDir, "git", "remote", "set-branches", "origin", "'*'"); !ok {
			fmt.Println("Cannot set-branches")
			os.Exit(1)
		}

		if _, ok, _ := run(tmpDir, "git", "fetch", "-v", "--depth=1"); !ok {
			fmt.Println("Cannot fetch remote")
			os.Exit(1)
		}

		if _, ok, _ := run(tmpDir, "git", "checkout", checkoutRef); !ok {
			fmt.Println("Cannot checkout", checkoutRef)
			os.Exit(1)
		}
	}

	path := tmpDir

	if !keepStructure {
		path = filepath.Join(tmpDir, subdir)
	}

	if err := os.CopyFS(".", os.DirFS(path)); err != nil {
		fmt.Println("Cannot copy content of", path)
		os.Exit(1)
	}
}
