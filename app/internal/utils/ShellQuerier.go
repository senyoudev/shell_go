package utils

type ShellQuerier interface {
	IsBuiltin(name string) bool
}