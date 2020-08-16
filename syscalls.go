package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func commit(date time.Time) {
	filePath := fmt.Sprintf("%d/%d-%d.txt", date.Year(), date.Month(), date.Day())
	text := fmt.Sprintf("commit for " + dateFormat(date))

	//append some dummy changes to the file
	appendToFile(filePath, text)

	env := os.Environ()
	env = append(env, fmt.Sprintf("GIT_COMMITTER_DATE=\"%s\"", dateFormat(date)))

	cmd := exec.Command("git", "add", ".")
	cmd.Env = env
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

	cmd = exec.Command("git", "commit", "--date", fmt.Sprintf("\"%s\"", dateFormat(date)), "-m", fmt.Sprintf("\"%s\"", text))
	cmd.Env = env
	err = cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}

func appendToFile(filePath, text string) {
	folderPath := filepath.Dir(filePath)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, 0755)
		if err != nil {
			fmt.Println(err)
			panic("can't create commit-folder, make sure path exists ")
		}
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(text + "\n"); err != nil {
		log.Println(err)
	}
}
