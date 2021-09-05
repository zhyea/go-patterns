package main

import "fmt"

func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "maverick" {
		return newMaverick(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
