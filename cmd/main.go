package main

import (
	"bufio"
	"fmt"
	"github.com/jgbz/rover/pkg/domain/entities"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFileEnvVariable = "INPUT_PATH"

func main() {
	var filePath string

	// Tries to read the input from the execution args, if it's empty then tries to read from the env var.
	argsFilePath := os.Args[1]
	if len(argsFilePath) == 0 {
		envFilePath := os.Getenv(inputFileEnvVariable)
		if len(envFilePath) == 0 {
			log.Fatal("unable to start application: input path not defined")
		}
		filePath = envFilePath
	} else {
		filePath = argsFilePath
	}

	fileLines, err := readFile(filePath)
	if err != nil {
		log.Fatalf("error while reading file: %v", err)
	}

	// The first line contains information about the grid size
	gridSize := fileLines[0]

	x, y, err := getGridSize(gridSize)
	if err != nil {
		log.Fatal(err)
	}

	grid := entities.NewGrid(x, y)
	fmt.Println(grid)

	// remove the fist line that contains the grid size
	fileLines = fileLines[1:]

	// validates if each rover will have it's initial position and a set of instructions
	if (len(fileLines) % 2) != 0 {
		log.Fatal("invalid domain data, missing initial position or instruction")
	}

	// loops every two lines
	for i := 0; i <= len(fileLines)-1; i += 2 {
		initialPosLine := fileLines[i]
		instructionsLine := fileLines[i+1]
		x, y, direction, err := getInitialCoordinates(initialPosLine)
		if err != nil {
			log.Printf("fail to getting initial coordinates: %v", err)
			continue
		}
		rover, err := entities.NewRover(x, y, direction, []rune(instructionsLine))
		if err != nil {
			log.Printf("unable to create new Rover: %v", err)
			continue
		}

		coordinates, err := rover.Explore(grid)
		if err != nil {
			log.Printf("domain failed to explore: %v", err)
		}
		fmt.Printf("Rover coordinates: %s\n", coordinates)
	}

}

func readFile(filePath string) ([]string, error) {
	readFile, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	// If the input file has less than 3 lines
	// one for the grid size, one for the initial domain position and the third one for the first domain instructions
	// then the application is unable to start.
	if len(fileLines) < 3 {
		return nil, fmt.Errorf("not enough data was provided to start the application")
	}

	return fileLines, nil
}

func getGridSize(grid string) (x, y int, err error) {
	gridSize := strings.Split(grid, " ")
	if len(gridSize) < 2 {
		log.Fatal("invalid grid size values")
	}

	x, err = strconv.Atoi(gridSize[0])
	if err != nil {
		return 0, 0, err
	}
	y, err = strconv.Atoi(gridSize[1])
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}

func getInitialCoordinates(coordinates string) (x, y int, direction string, err error) {
	coords := strings.Split(coordinates, " ")
	if len(coords) < 3 {
		return 0, 0, "", fmt.Errorf("not enough parameters for domain initial coordinates")
	}

	x, err = strconv.Atoi(coords[0])
	if err != nil {
		return 0, 0, "", err
	}
	y, err = strconv.Atoi(coords[1])
	if err != nil {
		return 0, 0, "", err
	}

	direction = coords[2]

	return x, y, direction, nil
}
