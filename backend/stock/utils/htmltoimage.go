package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/google/uuid"
)

// UUID UUID
func HtmlToImage(html string) (string, error) {
	fileName := uuid.New().String()
	filePath := fmt.Sprintf("/tmp/%s.png", fileName)
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("echo '%s' | wkhtmltoimage - - %s", html, filePath))
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return filePath, nil
}

func HtmlToImageDelete(path string) error {
	err := os.Remove(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
