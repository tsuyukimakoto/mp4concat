package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"gopkg.in/djherbis/times.v1"
)

const extMP4 = ".MP4"
const commandFfmpeg = "ffmpeg"

func basePath() string {
	// mp4concat uses ~/Desktop/mp4concat , HARD CODED for now.
	home, err := os.UserHomeDir()
	if err != nil {
		log.Print("Home directory not found")
		log.Fatal(err)
	}
	dir := fmt.Sprintf("%s/Desktop/mp4concat_work", home)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}
	return dir
}

func extractMP4Path(files []string) []string {
	result := make([]string, 0, len(files))
	for _, v := range files {
		if extMP4 == strings.ToUpper(filepath.Ext(v)) {
			result = append(result, v)
		}
	}
	return result
}

func creationTime(t time.Time) string {
	return fmt.Sprintf(
		"%4d-%02d-%02dT%02d:%02d:%02d.000000Z",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
	)
}

func buildFFMPEGCommandArguments(filesMP4 []string, inputFileName string, outputPath string) []string {
	arguments := make([]string, 0, 12)
	// ffmpeg -f concat -i input.txt -c copy -metadata creation_time="2021-01-02T15:56:00.000000Z" -y output.mp4
	var minTime = time.Now()
	for _, filePath := range filesMP4 {
		t, err := times.Stat(filePath)
		if err != nil {
			log.Fatal(err)
		}
		if t.HasChangeTime() {
			fileTime := t.ChangeTime()
			if fileTime.Before(minTime) {
				minTime = fileTime
			}
		}
	}
	arguments = append(arguments, "-f")
	arguments = append(arguments, "concat")
	arguments = append(arguments, "-safe") // Absolute file path needs this...
	arguments = append(arguments, "0")
	arguments = append(arguments, "-i")
	arguments = append(arguments, inputFileName)
	arguments = append(arguments, "-c")
	arguments = append(arguments, "copy")      // codec is copy
	arguments = append(arguments, "-metadata") // Set the oldest update date in the selected video file as the shooting date.
	arguments = append(arguments, fmt.Sprintf("creation_time=%s", creationTime(minTime)))
	arguments = append(arguments, "-y")
	arguments = append(arguments, outputPath)
	return arguments
}

func checkFFMPEGCommandExist() {
	cmd := exec.Command(commandFfmpeg, "--help")
	// fmt.Println("checking ffmpeg command installed.")
	if err := cmd.Run(); err != nil {
		log.Printf("%s command not found", commandFfmpeg)
		log.Fatal(err)
	}
}

func createInputFile(filesMP4 []string, inputFileNameAbsolute string) {
	f, err := os.Create(inputFileNameAbsolute)
	if err != nil {
		log.Printf("can't create file %s", inputFileNameAbsolute)
		log.Fatal(err)
	}
	for _, v := range filesMP4 {
		_, err := f.WriteString(fmt.Sprintf("file %s\n", v))
		if err != nil {
			log.Fatal(err)
		}
	}
	f.Sync()
}

func main() {
	// mp4concat needs ffmpeg command callable
	checkFFMPEGCommandExist()

	// Prompt users for input
	fmt.Print("Drag and Drop files here to concat: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	files := strings.Fields(scanner.Text())
	filesMP4 := extractMP4Path(files)
	if len(filesMP4) == 0 {
		log.Print("Target MP4 file not found.")
		os.Exit(0)
	}

	// Create input file
	mp4concatBasePath := basePath()
	inputFileName := fmt.Sprintf("%s/%s.txt", mp4concatBasePath, uuid.New().String())
	createInputFile(filesMP4, inputFileName)

	// Build arguments for the ffmpeg command
	now := time.Now()
	outputFileName := fmt.Sprintf(
		"%s/output_%04d%02d%02d_%02d%02d%02d.mp4",
		mp4concatBasePath,
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second(),
	)
	arguments := buildFFMPEGCommandArguments(filesMP4, inputFileName, outputFileName)

	// Exec ffmpeg command
	cmd := exec.Command(commandFfmpeg, arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
