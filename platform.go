//go:build !windows
// +build !windows

package lumberjack

import "syscall"

func Umask(newmask int) (oldmask int) {
	return syscall.Umask(newmask)
}
