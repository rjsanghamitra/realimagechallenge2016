package models

import (
	"errors"
	"strings"
)

type Distributor struct {
	Name       string
	ParentName *Distributor
	Included   map[string]struct{}
	Excluded   map[string]struct{}
}

func CreateNewEmptyDistributor() *Distributor {
	return &Distributor{
		Name:       "",
		ParentName: &Distributor{},
		Included:   map[string]struct{}{},
		Excluded:   map[string]struct{}{},
	}
}

type Permission struct {
	Action string
	Area   string
}

// The area string will be of the form country, state-country, or city-state-country

func (d *Distributor) AddPermissions(permissions []Permission) error {
	for _, permission := range permissions {
		action := permission.Action
		area := permission.Area
		areas := strings.Split(area, "-")
		if action == INCLUDE {
			for _, i := range areas {
				d.Included[i] = struct{}{}
			}
		} else if action == EXCLUDE {
			for _, i := range areas {
				d.Excluded[i] = struct{}{}
				delete(d.Included, i)
			}
		} else {
			return errors.New("invalid action")
		}
	}
	return nil
}

func (d *Distributor) IsIncluded(area string) bool {
	areas := strings.Split(area, "-")
	for _, i := range areas {
		_, exists := d.Included[i]
		if exists {
			return true
		}
	}
	return false
}

func (d *Distributor) IsExcluded(area string) bool {
	areas := strings.Split(area, "-")
	for _, i := range areas {
		_, exists := d.Excluded[i]
		if !exists {
			return false
		}
	}
	return true
}
