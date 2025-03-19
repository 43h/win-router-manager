package main

import (
	"main/logger"
	"main/windows"
)

func main() {
	if rst := logger.InitLog(true); rst == false {
		return
	}
	defer logger.CloseLog()

	if rst := windows.CreateWin(); rst == false {
		return
	}
	windows.RunWin()
	defer windows.DestroyWin()
}
