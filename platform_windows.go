//go:build windows
// +build windows

package lumberjack

func Umask(newmask int) (oldmask int) {
	return 0
}
