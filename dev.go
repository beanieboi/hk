//go:build !release
// +build !release

package main

const (
	Version = "dev"
)

var updater *Updater
