package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sqweek/dialog"
)

func DirectoryExists(dir string) (bool, error) {
	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return info.IsDir(), nil
}

func WritePromptFile(proopt *Proompt, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(proopt.text)
	if err != nil {
		return err
	}

	return nil
}

func DoWindowExe() {
	startDir := "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Counter-Strike Global Offensive\\csgo\\replays"
	if exists, err := DirectoryExists(startDir); err == nil && !exists {
		startDir = "C:\\Program Files (x86)\\Steam\\steamapps\\common\\Counter-Strike Global Offensive\\csgo"
	}

	filePath, err := dialog.File().
		Title("Select a demo file").
		Filter("Demo files (*.dem)", "dem").
		SetStartDir(startDir).
		Load()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	f, err := os.Open(filePath)

	if err != nil {
		log.Panic("failed to open demo file: ", err)
	}
	defer f.Close()

	proompt := GenerateProompt(f)

	fileName := filepath.Base(filePath) + "-proompt.txt"
	WritePromptFile(&proompt, fileName)
}

func main() {
	//filePath := "demos/rolled-16-0.dem"
	//filepath := "demos/og-vs-natus-vincere-m2-mirage.dem"
	DoWindowExe()
}
