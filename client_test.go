package api


import (
	"fmt"
	"testing"
)

func TestGetByKeyword(t *testing.T) {
	fmt.Println("TestGetByKeyword")
	GetByKeyword("shirt")
}