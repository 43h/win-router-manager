// Copyright 2011 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

import (
	"log"
)

import (
	"github.com/lxn/walk"
)

func InitTray(window *walk.MainWindow) {
	// We need either a walk.MainWindow or a walk.Dialog for their message loop.
	// We will not make it visible in this example, though.

	// We load our icon from a file.
	icon, err := walk.Resources.Icon("resource/ok.ico")
	if err != nil {
		log.Fatal(err)
	}

	// Create the notify icon and make sure we clean it up on exit.
	ni, err := walk.NewNotifyIcon(window)
	if err != nil {
		log.Fatal(err)
	}

	// Set the icon and a tool tip text.
	if err := ni.SetIcon(icon); err != nil {
		log.Fatal(err)
	}
	//if err := ni.SetToolTip("Click for info or use the context menu to exit."); err != nil {
	//	log.Fatal(err)
	//}

	// When the left mouse button is pressed, bring up our balloon.
	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}

		if err := ni.ShowCustom(
			"Walk NotifyIcon Example",
			"There are multiple ShowX methods sporting different icons.",
			icon); err != nil {

			log.Fatal(err)
		}
	})

	// We put an exit action into the context menu.
	exitAction := walk.NewAction()
	if err := exitAction.SetText("退出"); err != nil {
		log.Fatal(err)
	}
	exitAction.Triggered().Attach(func() { walk.App().Exit(0) })
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		log.Fatal(err)
	}

	// The notify icon is hidden initially, so we have to make it visible.
	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	// Now that the icon is visible, we can bring up an info balloon.
	if err := ni.ShowInfo("Walk NotifyIcon Example", "Click the icon to show again."); err != nil {
		log.Fatal(err)
	}
}

// 初始化系统托盘
func initTray(w *walk.MainWindow) {
	//ni.MouseDown().Attach(func(x, y int, btn walk.MouseButton) {
	//	if btn == walk.RightButton {
	//		menu := walk.NewMenu()
	//		menu.Items().Add(walk.NewMenuItem(&walk.Action{
	//			Text: "显示窗口",
	//			OnTriggered: func() {
	//				w.Show()
	//				w.SetWindowState(walk.WindowStateNormal)
	//			},
	//		}))
	//		menu.Items().Add(walk.NewMenuItem(&walk.Action{
	//			Text: "退出",
	//			OnTriggered: func() {
	//				walk.App().Exit(0)
	//			},
	//		}))
	//		menu.ShowPopupMenu(w, x, y)
	//	}
	//})

	log.Println("托盘初始化完成")
}