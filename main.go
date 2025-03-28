package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func processMarkdownFiles(rootDir string) error {
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
      newPath := filepath.Join(imageDir,info.Name())
      os.Rename(path, newPath)
    }

    // update markdown files 
    if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
      // update markdown images
      err := updateMarkdownImages(path)
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

// updateMarkdown images
func updateMarkdownImages(mdFile string) error {
  content, err := os.ReadFile(mdFile)
  if err != nil {
    return err
  }

  // find all images in the file
  re := regexp.MustCompile(`!\[.*?\]\((.*?)\)`)
  updatedContent := re.ReplaceAllFunc(content, func(match []byte) []byte {
    // Extract the image URL from the match
    url := string(match[2 : len(match)-1])
    // Check if the URL is a relative path
    if !strings.HasPrefix(url, "http") {
      // Replace relative path with absolute path
      baseDir := filepath.Dir(mdFile)
      newURL := filepath.Join(baseDir, url)
      return []byte(fmt.Sprintf("![%s](%s)", match[1], newURL))
    }
    return match
  })

  // write the updated content back to the file
  err = os.WriteFile(mdFile, updatedContent, 0644)
  if err != nil {
    return err
  }

  return nil
}

func main() {
  println("Hello there, I am Kishor Rathva");
  rootDir := "."
  err := processMarkdownFiles(rootDir)

  if err != nil {
    fmt.Println("Error processing files:", err)
    os.Exit(1);
  }
  fmt.Println("Markdown files formatted and images organized successfully.")
}
