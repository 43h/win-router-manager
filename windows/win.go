package windows

import (
	"log"
	"main/logger"
	"syscall"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// 主窗口大小
const MainWinMinWidth = 600
const MainWinMinHeight = 400

var tableView *walk.TableView
var sbi *walk.StatusBarItem

type Windows struct {
	mWin              *walk.MainWindow
	mNI               *walk.NotifyIcon
	mContentContainer *walk.Composite
}

var windows = Windows{}

func CreateWin() bool {

	// 创建主窗口
	mainWindow := MainWindow{
		AssignTo: &windows.mWin,
		Title:    "Golang Walk 示例",
		MinSize:  Size{Width: 600, Height: 400},
		Layout:   HBox{Margins: Margins{10, 10, 10, 10}},
		Children: []Widget{
			Composite{
				Layout:  VBox{Margins: Margins{0, 0, 10, 0}},
				MaxSize: Size{Width: 150},
				Children: []Widget{
					PushButton{
						Text: "页面 1",
						OnClicked: func() {
							updateContent(windows.mContentContainer, "这是页面 1 的内容")
						},
					},
					PushButton{
						Text: "页面 2",
						OnClicked: func() {
							updateContent(windows.mContentContainer, "这是页面 2 的内容")
						},
					},
				},
			},
			Composite{
				AssignTo: &windows.mContentContainer,
				Layout:   VBox{},
				Children: []Widget{
					Label{Text: "请选择左侧菜单项"},
				},
			},
		},
	}
	// 修改窗口关闭事件
	windows.mWin.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		*canceled = true    // 阻止默认关闭行为
		windows.mWin.Hide() // 隐藏窗口
		log.Println("窗口已隐藏到托盘")
	})

	//// 窗口最小化时隐藏到托盘
	//windows.mWin.Minimized().Attach(func() {
	//	windows.mWin.Hide()
	//})
	// 创建窗口并初始化托盘
	if err := mainWindow.Create(); err != nil {
		log.Fatal("创建窗口失败: ", err)
		return false
	}
	return true
}

func initWin() bool {
	if err := windows.mWin.SetTitle("路由管理"); err != nil {
		logger.LOGE(err)
		return false
	}
	if err := windows.mWin.SetLayout(walk.NewVBoxLayout()); err != nil {
		logger.LOGE(err)
		return false
	}

	//set windows location
	screenWidth, screenHeight := getScreenSize()
	windowWidth := MainWinMinWidth
	windowHeight := MainWinMinHeight
	x := (screenWidth - windowWidth) / 2
	y := (screenHeight - windowHeight) / 2
	if err := windows.mWin.SetBounds(walk.Rectangle{
		X:      x,
		Y:      y,
		Width:  windowWidth,
		Height: windowHeight,
	}); err != nil {
		logger.LOGE(err)
		return false
	}

	if err := windows.mWin.SetSize(walk.Size{Width: MainWinMinWidth, Height: MainWinMinHeight}); err != nil {
		logger.LOGE(err)
		return false
	}
	windows.mWin.SetVisible(true)
	if err := windows.mWin.SetFocus(); err != nil {
		logger.LOGE(err)
		return false
	}
	return true
}

func getScreenSize() (int, int) {
	user32 := syscall.NewLazyDLL("user32.dll")
	getSystemMetrics := user32.NewProc("GetSystemMetrics")

	// 0 = SM_CXSCREEN (屏幕宽度)
	// 1 = SM_CYSCREEN (屏幕高度)
	screenWidth, _, _ := getSystemMetrics.Call(0)
	screenHeight, _, _ := getSystemMetrics.Call(1)

	return int(screenWidth), int(screenHeight)
}

//
//// Copyright 2011 The Walk Authors. All rights reserved.
//// Use of this source code is governed by a BSD-style
//// license that can be found in the LICENSE file.
//
//type Foo struct {
//	Index   int
//	Bar     string
//	Baz     float64
//	Quux    time.Time
//	checked bool
//}
//
//type FooModel struct {
//	walk.TableModelBase
//	walk.SorterBase
//	sortColumn int
//	sortOrder  walk.SortOrder
//	items      []*Foo
//}
//
//func NewFooModel() *FooModel {
//	m := new(FooModel)
//	m.ResetRows()
//	return m
//}
//
//// Called by the TableView from SetModel and every time the model publishes a
//// RowsReset event.
//func (m *FooModel) RowCount() int {
//	return len(m.items)
//}
//
//// Called by the TableView when it needs the text to display for a given cell.
//func (m *FooModel) Value(row, col int) interface{} {
//	item := m.items[row]
//
//	switch col {
//	case 0:
//		return item.Index
//
//	case 1:
//		return item.Bar
//
//	case 2:
//		return item.Baz
//
//	case 3:
//		return item.Quux
//	}
//
//	panic("unexpected col")
//}
//
//// Called by the TableView to retrieve if a given row is checked.
//func (m *FooModel) Checked(row int) bool {
//	return m.items[row].checked
//}
//
//// Called by the TableView when the user toggled the check box of a given row.
//func (m *FooModel) SetChecked(row int, checked bool) error {
//	m.items[row].checked = checked
//
//	return nil
//}
//
//// Called by the TableView to sort the model.
//func (m *FooModel) Sort(col int, order walk.SortOrder) error {
//	m.sortColumn, m.sortOrder = col, order
//
//	sort.SliceStable(m.items, func(i, j int) bool {
//		a, b := m.items[i], m.items[j]
//
//		c := func(ls bool) bool {
//			if m.sortOrder == walk.SortAscending {
//				return ls
//			}
//
//			return !ls
//		}
//
//		switch m.sortColumn {
//		case 0:
//			return c(a.Index < b.Index)
//
//		case 1:
//			return c(a.Bar < b.Bar)
//
//		case 2:
//			return c(a.Baz < b.Baz)
//
//		case 3:
//			return c(a.Quux.Before(b.Quux))
//		}
//
//		panic("unreachable")
//	})
//
//	return m.SorterBase.Sort(col, order)
//}
//
//func (m *FooModel) ResetRows() {
//	// Create some random data.
//	m.items = make([]*Foo, rand.Intn(50000))
//
//	now := time.Now()
//
//	for i := range m.items {
//		m.items[i] = &Foo{
//			Index: i,
//			Bar:   strings.Repeat("*", rand.Intn(5)+1),
//			Baz:   rand.Float64() * 1000,
//			Quux:  time.Unix(rand.Int63n(now.Unix()), 0),
//		}
//	}
//
//	// Notify TableView and other interested parties about the reset.
//	m.PublishRowsReset()
//
//	m.Sort(m.sortColumn, m.sortOrder)
//}
//
//func main() {
//	rand.Seed(time.Now().UnixNano())
//
//	boldFont, _ := walk.NewFont("Segoe UI", 9, walk.FontBold)
//	goodIcon, _ := walk.Resources.Icon("../img/check.ico")
//	badIcon, _ := walk.Resources.Icon("../img/stop.ico")
//
//	barBitmap, err := walk.NewBitmap(walk.Size{100, 1})
//	if err != nil {
//		panic(err)
//	}
//	defer barBitmap.Dispose()
//
//	canvas, err := walk.NewCanvasFromImage(barBitmap)
//	if err != nil {
//		panic(err)
//	}
//	defer barBitmap.Dispose()
//
//	canvas.GradientFillRectangle(walk.RGB(255, 0, 0), walk.RGB(0, 255, 0), walk.Horizontal, walk.Rectangle{0, 0, 100, 1})
//
//	canvas.Dispose()
//
//	model := NewFooModel()
//
//	var tv *walk.TableView
//
//	MainWindow{
//		Title:  "Walk TableView Example",
//		Size:   Size{800, 600},
//		Layout: VBox{MarginsZero: true},
//		Children: []Widget{
//			PushButton{
//				Text:      "Reset Rows",
//				OnClicked: model.ResetRows,
//			},
//			PushButton{
//				Text: "Select first 5 even Rows",
//				OnClicked: func() {
//					tv.SetSelectedIndexes([]int{0, 2, 4, 6, 8})
//				},
//			},
//			TableView{
//				AssignTo:         &tv,
//				AlternatingRowBG: true,
//				CheckBoxes:       true,
//				ColumnsOrderable: true,
//				MultiSelection:   true,
//				Columns: []TableViewColumn{
//					{Title: "#"},
//					{Title: "Bar"},
//					{Title: "Baz", Alignment: AlignFar},
//					{Title: "Quux", Format: "2006-01-02 15:04:05", Width: 150},
//				},
//				StyleCell: func(style *walk.CellStyle) {
//					item := model.items[style.Row()]
//
//					if item.checked {
//						if style.Row()%2 == 0 {
//							style.BackgroundColor = walk.RGB(159, 215, 255)
//						} else {
//							style.BackgroundColor = walk.RGB(143, 199, 239)
//						}
//					}
//
//					switch style.Col() {
//					case 1:
//						if canvas := style.Canvas(); canvas != nil {
//							bounds := style.Bounds()
//							bounds.X += 2
//							bounds.Y += 2
//							bounds.Width = int((float64(bounds.Width) - 4) / 5 * float64(len(item.Bar)))
//							bounds.Height -= 4
//							canvas.DrawBitmapPartWithOpacity(barBitmap, bounds, walk.Rectangle{0, 0, 100 / 5 * len(item.Bar), 1}, 127)
//
//							bounds.X += 4
//							bounds.Y += 2
//							canvas.DrawText(item.Bar, tv.Font(), 0, bounds, walk.TextLeft)
//						}
//
//					case 2:
//						if item.Baz >= 900.0 {
//							style.TextColor = walk.RGB(0, 191, 0)
//							style.Image = goodIcon
//						} else if item.Baz < 100.0 {
//							style.TextColor = walk.RGB(255, 0, 0)
//							style.Image = badIcon
//						}
//
//					case 3:
//						if item.Quux.After(time.Now().Add(-365 * 24 * time.Hour)) {
//							style.Font = boldFont
//						}
//					}
//				},
//				Model: model,
//				OnSelectedIndexesChanged: func() {
//					fmt.Printf("SelectedIndexes: %v\n", tv.SelectedIndexes())
//				},
//			},
//		},
//	}
//}

func RunWin() {
	windows.mWin.Run()
}

func DestroyWin() {
	if windows.mNI != nil {
		err := windows.mNI.Dispose()
		if err != nil {
			logger.LOGE(err.Error())
		}
	}
}

// 更新右侧内容
func updateContent(container *walk.Composite, text string) {
	children := container.Children()
	// 遍历并清理现有子控件
	for i := 0; i < children.Len(); i++ {
		child := children.At(i)
		child.Dispose()
	}
	// 添加新内容
	label, err := walk.NewLabel(container)
	if err != nil {
		logger.LOGE("创建 Label 失败: %v" + err.Error())
		return
	}
	label.SetText(text)
}