package cmd

// Simple strings to print out
var (
	greeting = "Welcome to ProtocolZero. What would you like to wipe clean?"
	options  = `
1) Reddit
2) ???

OPTION: `
)

// Website type for enum services
type Website int

// Have constant reddit
const (
	Reddit Website = 1 + iota
)
