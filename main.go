/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/mir-asif01/simple_cli_app/cmd"
)

type Quote struct {
	Id     string `json:"ID"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func main() {
	cmd.Execute()
}
