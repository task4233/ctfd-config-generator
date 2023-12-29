package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	_ "embed"

	"github.com/manifoldco/promptui"
)

var (
	//go:embed templates/challenge.yml.tmpl
	challengeTemplate string

	//go:embed templates/writeup.md.tmpl
	writeupTemplate string

	//go:embed templates/README.md.tmpl
	readmeTemplate string

	genres []string = []string{"crypto", "hardware", "misc", "network", "osint", "pwn", "rev", "web"}

	challengeFormat = "^[a-z0-9_!?]+$"
	challengeRegExp = regexp.MustCompile(challengeFormat)

	flagFormat = "^HogeCTF23{[^{}]+}$"
	flagRegExp = regexp.MustCompile(flagFormat)
)

type challengeInfo struct {
	ChallengeName string
	Author        string
	Genre         string
	Flag          string
}

func main() {
	// get genre
	promptForSelect := promptui.Select{
		Label: "genre",
		Items: genres,
	}
	_, genre, err := promptForSelect.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed promptForSelect.Run(): %s", err.Error())
		os.Exit(1)
	}

	// get challenge name
	prompt := promptui.Prompt{
		Label: "challenge name",
		Validate: func(input string) error {
			if !challengeRegExp.MatchString(input) {
				return fmt.Errorf("challenge name should meet %s", challengeFormat)
			}
			return nil
		},
	}
	challengeName, err := prompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed prompt.Run() for challengeName: %s", err.Error())
		os.Exit(1)
	}

	// get author name
	prompt.Label = "author name"
	author, err := prompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed prompt.Run() for author: %s", err.Error())
		os.Exit(1)
	}

	// get flag
	prompt = promptui.Prompt{
		Label: "flag",
		Validate: func(input string) error {
			if !flagRegExp.MatchString(input) {
				return fmt.Errorf("flag should meet %s", flagFormat)
			}
			return nil
		},
	}
	flag, err := prompt.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed prompt.Run(): %s", err.Error())
		os.Exit(1)
	}

	info := challengeInfo{
		ChallengeName: challengeName,
		Author:        author,
		Genre:         genre,
		Flag:          flag,
	}

	// ready a directory structure
	// - make directory(./genre/challengeName)
	//   - directory: build, files, solver
	//   - file: README.md, flag.txt, challenge.yml, writeup/README.md
	err = os.MkdirAll(filepath.Join(genre, challengeName), os.ModePerm)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed os.MkdirAll(genre/challengeName): %s", err.Error())
		os.Exit(1)
	}

	dirs := []string{"build", "public", "solver", "writeup"}
	for _, dirName := range dirs {
		err = os.MkdirAll(filepath.Join(genre, challengeName, dirName), os.ModePerm)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed os.MkdirAll(genre/challengeName/%s): %s", dirName, err.Error())
			os.Exit(1)
		}
	}

	// write default description for each file
	files := []string{"README.md", "flag.txt", "challenge.yml", "writeup/README.md"}
	for _, fileName := range files {
		if err := readyFile(fileName, info); err != nil {
			fmt.Fprintf(os.Stderr, "failed readyFile: %s", err.Error())
			os.Exit(1)
		}
	}
}

func generateMarkdown(templateName string, templateStr string, info challengeInfo) (string, error) {
	tpl, err := template.New(templateName).Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("failed template.New: %w", err)
	}
	writer := &bytes.Buffer{}
	err = tpl.Execute(writer, info)
	return writer.String(), err
}

func readyFile(fileName string, info challengeInfo) error {
	fp, err := os.Create(filepath.Join(info.Genre, info.ChallengeName, fileName))
	if err != nil {
		return fmt.Errorf("failed os.Create(genre/challengeName/%s): %w", fileName, err)
	}
	defer fp.Close()

	// write template message
	switch fileName {
	case "README.md":
		readmeMd, err := generateMarkdown("README", readmeTemplate, info)
		if err != nil {
			return fmt.Errorf("failed generateMarkdown for challenge.yml: %w", err)
		}
		fmt.Fprint(fp, readmeMd)
		break
	case "flag.txt":
		fmt.Fprintln(fp, info.Flag)
		break
	case "challenge.yml":
		challengeYaml, err := generateMarkdown("challenge", challengeTemplate, info)
		if err != nil {
			return fmt.Errorf("failed generateMarkdown for challenge.yml: %w", err)
		}
		fmt.Fprint(fp, challengeYaml)
		break
	case "writeup/README.md":
		writeupMd, err := generateMarkdown("writeup", writeupTemplate, info)
		if err != nil {
			return fmt.Errorf("failed generateMarkdown for challenge.yml: %w", err)
		}
		fmt.Fprint(fp, writeupMd)
		break
	}

	return nil
}
