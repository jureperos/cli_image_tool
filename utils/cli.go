package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func YesNoPrompt(question string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		log.Printf("%s [y/n]: ", question)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v\n", err)
			return false
		}
		response = strings.TrimSpace(response)

		switch strings.ToLower(response) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			log.Println("Please enter 'y' or 'n'")
		}
	}
}

func ArePathsSame(path1, path2 string) (bool, error) {
	// Resolve symlinks and clean paths
	abs1, err := filepath.EvalSymlinks(path1)
	if err != nil {
		return false, fmt.Errorf("error resolving symlink for %q: %v", path1, err)
	}

	abs2, err := filepath.EvalSymlinks(path2)
	if err != nil {
		return false, fmt.Errorf("error resolving symlink for %q: %v", path2, err)
	}

	// Clean paths to normalize them
	clean1 := filepath.Clean(abs1)
	clean2 := filepath.Clean(abs2)

	return clean1 == clean2, nil
}
