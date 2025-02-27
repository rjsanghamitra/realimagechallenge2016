package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"qc_assignment/models"
	"qc_assignment/services"
	"strconv"
	"strings"
)

func main() {
	err := services.LoadCSVIntoStruct("cities.csv")
	if err != nil {
		log.Println("error loading csv file")
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the operation you want to perform:")
		fmt.Println("1 - Create new Distributor")
		fmt.Println("2 - Create Sub-Distributor")
		var optionStr string
		optionStr, _ = reader.ReadString('\n')
		optionStr = strings.TrimSpace(optionStr)
		option, _ := strconv.Atoi(optionStr)

		if option == 1 {
			fmt.Println("Enter the name of the distributor: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			if services.IsExistsDistributor(name) {
				fmt.Println("Error: Distributor already exists. Please try again.")
				continue
			}
			fmt.Println("Enter the permissions in the form <ACTION> <CITY>-<STATE>-<COUNTRY> \nOR <ACTION> <STATE>-<COUNTRY> \nOR <ACTION> <COUNTRY>.\nEnter 'DONE' to stop:")
			var permissions []models.Permission
			for {
				line, _ := reader.ReadString('\n')
				line = strings.TrimSpace(line)
				if line == "DONE" {
					break
				}
				lineSplit := strings.Split(line, " ")
				permissions = append(permissions, models.Permission{
					Action: lineSplit[0],
					Area:   lineSplit[1],
				})
			}
			distributor := models.CreateNewEmptyDistributor()
			services.CreateDistributor(name, distributor)
			distributor.AddPermissions(permissions)
			fmt.Printf("Distributor %s created.", name)
		} else if option == 2 {
			fmt.Println("Enter the name of the primary distributor: ")
			primaryDistributorName, _ := reader.ReadString('\n')
			primaryDistributorName = strings.TrimSpace(primaryDistributorName)
			if !services.IsExistsDistributor(primaryDistributorName) {
				fmt.Println("Error: Primary Distributor does not exist. Please try again.")
				continue
			}
			fmt.Println("Enter the name of the Secondary Distributor: ")
			secondaryDistributorName, _ := reader.ReadString('\n')
			secondaryDistributorName = strings.TrimSpace(secondaryDistributorName)
			if services.IsExistsDistributor(secondaryDistributorName) {
				fmt.Println("Error: Distributor already exists. Please try again.")
				continue
			}
			fmt.Println("Enter the permissions in the form <ACTION> <CITY>-<STATE>-<COUNTRY> \nOR <ACTION> <STATE>-<COUNTRY> \nOR <ACTION> <COUNTRY>.\nEnter 'DONE' to stop:")
			var permissions []models.Permission
			for {
				line, _ := reader.ReadString('\n')
				line = strings.TrimSpace(line)
				if line == "DONE" {
					break
				}
				lineSplit := strings.Split(line, " ")
				permissions = append(permissions, models.Permission{
					Action: lineSplit[0],
					Area:   lineSplit[1],
				})
			}
			parentDistributor := services.GetDistributor(primaryDistributorName)
			check := services.CheckDistributor(parentDistributor, permissions)
			if !check {
				fmt.Println("Error: Parent distributor can't create distributor with these permissions. Please try again.")
				continue
			}
			secondaryDistributor := models.CreateNewEmptyDistributor()
			secondaryDistributor.AddPermissions(permissions)
			fmt.Printf("Distributor %s created.", secondaryDistributorName)
		} else {
			fmt.Println("Invalid option entered. Please try again.")
		}
	}
}
