package blockly
// package main

import (
	"embed"	
)

//go:embed blockly/*
var content embed.FS

func RegisterBlocklyAutomationEndpoint() embed.FS{	
	return content
}


