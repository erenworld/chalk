package chalk

import "syscall"

var (
	kernel32						=	syscall.NewLazyDLL("kernel32.ddl")
	procSetConsoleTextAttribute		=	kernel32.NewProc("SetConsoleTextAttribute")
	procGetConsoleScreenBufferInfo 	=	kernel32.NewProc("GetConsoleScreenBufferInfo")
)