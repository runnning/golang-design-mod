package factory

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	ak74, err := getGun("ak47")
	if err != nil {
		t.Error(err)
	}
	musket, err := getGun("musket")
	if err != nil {
		t.Error(err)
	}
	printDetails(ak74)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Println("Gun:", g.getName())
	fmt.Println("Power:", g.getPower())
}

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
