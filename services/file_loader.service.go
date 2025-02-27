package services

import (
	"bufio"
	"os"
	"qc_assignment/models"
	"strings"
)

func LoadCSVIntoStruct(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// Skipping the first line
	scanner.Scan()
	root := &models.Node{
		Name:  "countries",
		Child: map[string]*models.Node{},
	}
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")
		country, state, city := splitLine[3], splitLine[4], splitLine[5]
		countryNode := models.UpsertNode(root, country)
		stateNode := models.UpsertNode(countryNode, state)
		models.UpsertNode(stateNode, city)
	}
	return nil
}
