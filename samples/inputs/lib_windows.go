package main

import (
	"github.com/sqweek/dialog"
)


func PickImage() string {
	filename, err := dialog.File().Filter("Passport file", "jpg").Load()
	if err != nil {
		return
	}
	return filename
}