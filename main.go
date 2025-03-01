package main

import (
	"fmt"
	"os"

	"github.com/52funny/pikpakcli/conf"
	"github.com/52funny/pikpakcli/internal/pikpak"
	"github.com/sirupsen/logrus"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: mycli <username> <password> <directory>")
		os.Exit(1)
	}

	username := os.Args[1]
	password := os.Args[2]
	dir := os.Args[3]

	// Initialize the PikPak client
	p := pikpak.NewPikPak(username, password)

	// Perform the login
	err := p.Login()
	if err != nil {
		logrus.Errorln("Login Failed:", err)
		os.Exit(1)
	}

	// List files and directories
	err = listFiles(dir, p)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func listFiles(dir string, p *pikpak.PikPak) error {
	// Get the folder ID for the specified path
	parentId, err := p.GetPathFolderId(dir)
	if err != nil {
		return err
	}

	// Get the file list for the specified folder
	files, err := p.GetFolderFileStatList(parentId)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.Kind == "drive#folder" {
			fmt.Printf("[DIR] %s\n", file.Name)
		} else {
			fmt.Printf("[FILE] %s\n", file.Name)
		}
	}
	return nil
}
