package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func processMarkdownFiles(rootDir string, githubRoot string) error {
  imageDir := filepath.Join(rootDir, "images")
  err := os.MkdirAll(imageDir, os.ModePerm);

  if err != nil {
    return err
  }

  err = filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
    if err != nil {
      return err
    }

    // Mode images
    if !info.IsDir() && isImageFile(path){
      updatedName := strings.ReplaceAll(info.Name(), " ", "_")
      newPath := filepath.Join(imageDir,updatedName)
      os.Rename(path, newPath)
    }

    // update markdown files 
    if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
      // update markdown images
      fmt.Println(path);
      err := convertObsidianToGitHub(path, githubRoot)
      if err != nil {
        return err
      }
    }

    return nil
  });

  return err
}

// Check if the file is an image 
func isImageFile(file string) bool {
  ext := strings.ToLower(filepath.Ext(file));
  return ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".gif"
}

// Converts Obsidian-style image references to GitHub Markdown format
func convertObsidianToGitHub(mdFile string, githubRoot string) error {
	content, err := ioutil.ReadFile(mdFile)
	if err != nil {
		return err
	}

	// Regex for Obsidian image format: ![[image.png]]
	re := regexp.MustCompile(`!\[\[(.*?)\]\]`)
	updatedContent := re.ReplaceAllStringFunc(string(content), func(match string) string {
		matches := re.FindStringSubmatch(match)
		if len(matches) > 1 {
			imagePath := matches[1]
      updatedImageName := strings.ReplaceAll(imagePath, " ", "_") 
			imageName := filepath.Base(updatedImageName) // Extract just the filename

      // GitHub absolute URL
      githubImageURL := fmt.Sprintf("%s%s/%s", githubRoot ,"images", imageName);
			return fmt.Sprintf("![%s](%s)", imageName, githubImageURL) // Convert to GitHub format
		}
		return match
	})

	return ioutil.WriteFile(mdFile, []byte(updatedContent), 0644)
}

func main() {
  var githubRoot string
	if len(os.Args) > 1 {
		githubRoot = os.Args[1]
	} else {
		fmt.Println("Error: GitHub root URL is required (e.g., https://github.com/user/repo/blob/main/)")
		os.Exit(1)
	}
  rootDir := "."
  err := processMarkdownFiles(rootDir, githubRoot)

  if err != nil {
    fmt.Println("Error processing files:", err)
    os.Exit(1);
  }
  fmt.Println("Markdown files formatted and images organized successfully.")
}
