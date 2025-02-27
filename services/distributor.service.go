package services

import "qc_assignment/models"

var distributors map[string]*models.Distributor

func init() { // Initialize map to store all the distributors
	distributors = make(map[string]*models.Distributor)
}

func CreateDistributor(name string, d *models.Distributor) {
	distributors[name] = d
}

func CheckDistributor(parent *models.Distributor, permissions []models.Permission) bool { // Check if a sub-distributor can be created with the parent distributor.
	for _, permission := range permissions {
		action := permission.Action
		area := permission.Area

		if action == models.INCLUDE {
			if !parent.IsIncluded(area) {
				return false
			}
		}
		if action == models.EXCLUDE {
			if parent.IsExcluded(area) {
				return false
			}
		}
	}
	return true
}

func IsExistsDistributor(name string) bool {
	_, exists := distributors[name]
	return exists
}

func GetDistributor(name string) *models.Distributor {
	d, exists := distributors[name]
	if exists {
		return d
	}
	return nil
}
