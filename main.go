package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	var (
		path, dir string
	)

	flag.StringVar(&path, "path", "", "Usage: go run main.go <path/to/file.vtt>")
	flag.StringVar(&dir, "dir", "", "Usage: go run main.go <dir/to/vttDir>")

	flag.Parse()

	if path == "" && dir == "" {
		log.Panic("atleast specify one vtt file or directory containing vtt files")
	}
	if path != "" {
		if err := vtt2srt(path); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("yay! successfully converted %s into %s\n", path, strings.TrimSuffix(path, ".vtt")+".srt")
	}
	if dir != "" {
		root := os.DirFS(dir)

		vttFiles, err := fs.Glob(root, "*.vtt")
		if err != nil {
			log.Fatal(err)
		}

		for _, vtt := range vttFiles {
			if err := vtt2srt(vtt); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("yay! successfully converted %s into %s\n", path, strings.TrimSuffix(vtt, ".vtt")+".srt")
		}
	}
}

func vtt2srt(vttPath string) error {
	vtt, err := os.Open(vttPath)
	if err != nil {
		return err
	}
	defer vtt.Close()

	srtPath := strings.TrimSuffix(vttPath, ".vtt") + ".srt"

	srt, err := os.Create(srtPath)
	if err != nil {
		return err
	}
	defer srt.Close()

	scanner := bufio.NewScanner(vtt)
	writer := bufio.NewWriter(srt)

	timestampRegex := regexp.MustCompile(`(\d{2}:\d{2}:\d{2})\.(\d{3}) --> (\d{2}:\d{2}:\d{2})\.(\d{3})`)

	var lineCount int

	for scanner.Scan() {
		line := scanner.Text()

		if line == "WEBVTT" {
			continue
		}
		lineCount++

		//ignore first line if its empty
		if lineCount == 1 {
			if strings.TrimSpace(line) == "" {
				continue
			}
		}

		if timestampRegex.MatchString(line) {
			line = strings.ReplaceAll(line, ".", ",")
		}

		if _, err := writer.WriteString(line + "\n"); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	if err := writer.Flush(); err != nil {
		return err
	}
	return nil
}
