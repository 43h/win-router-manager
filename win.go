package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

// 主窗口大小
const MainWinMinWidth = 600
const MainWinMinHeight = 400

var searchTE *walk.LineEdit
var tableView *walk.TableView
var sbi *walk.StatusBarItem

type Item struct {
	index int
	sip   string
	sport uint16
	dip   string
	dport uint16
	pkt   uint32
	seq   string
}

type DataModel struct {
	walk.TableModelBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Item
}

func (m *DataModel) RowCount() int {
	return len(m.items)
}

func NewDataModel() *DataModel {
	m := new(DataModel)
	m.ResetRows()
	return m
}

func (m *DataModel) ResetRows() {
	// Create some random data.
	m.items = make([]*Item, 1)

	for i := range m.items {
		m.items[i] = &Item{
			index: i,
		}
	}

	m.PublishRowsReset()
}

//func (m *DataModel) FlushRows() {
//	// Create some random data.
//
//	m.items = make([]*Item, len(PktStat))
//	i := 0
//	for pktinfo, head := range PktStat {
//		m.items[i] = &Item{
//			index: i,
//		}
//		m.items[i].sip = pktinfo.SrcIP
//		m.items[i].sport = pktinfo.SrcPort
//		m.items[i].dip = pktinfo.DstIP
//		m.items[i].dport = pktinfo.DstPort
//		m.items[i].pkt = head.num
//
//		for node := head.list; node != nil; node = node.next {
//			m.items[i].seq += fmt.Sprintf("%d--->%d;", node.seqS, node.SeqE)
//		}
//		i += 1
//	}
//	m.PublishRowsReset()
//}

//func (m *DataModel) FilterRows(str string) (bool, error) {
//
//	if len(str) == 0 {
//		if len(m.items) != len(PktStat) {
//			m.FlushRows()
//		}
//		return true, nil
//	}
//
//	var port uint16
//	number, err := strconv.Atoi(str)
//	if err == nil && number > 0 && number < 65536 {
//		port = uint16(number)
//	} else {
//		return false, errors.New("无效过滤条件,目前仅支持端口过滤")
//	}
//
//	var num int32
//	for pktinfo, _ := range PktStat {
//		if port != 0 && port != pktinfo.SrcPort && port != pktinfo.DstPort {
//			continue
//		} else {
//			num += 1
//		}
//	}
//
//	if num == 0 { //无满足条件
//		return false, errors.New("无满足条件流")
//	}
//
//	m.items = make([]*Item, num)
//	i := 0
//	for pktinfo, head := range PktStat {
//		if port != 0 && port != pktinfo.SrcPort && port != pktinfo.DstPort {
//			continue
//		}
//
//		m.items[i] = &Item{
//			index: i,
//			sip:   pktinfo.SrcIP,
//			sport: pktinfo.SrcPort,
//			dip:   pktinfo.DstIP,
//			dport: pktinfo.DstPort,
//			pkt:   head.num,
//		}
//
//		for node := head.list; node != nil; node = node.next {
//			m.items[i].seq += fmt.Sprintf("%d--->%d;", node.seqS, node.SeqE)
//		}
//		i += 1
//	}
//	m.PublishRowsReset()
//	return true, nil
//}

func (m *DataModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.index
	case 1:
		return item.sip
	case 2:
		return item.sport
	case 3:
		return item.dip
	case 4:
		return item.dport
	case 5:
		return item.pkt
	case 6:
		return item.seq
	}

	panic("unexpected col")
}

func main() {
	newWin := new(walk.MainWindow)
	dataModel := NewDataModel()
	MainWindow{
		Title:    "Windows路由管理",
		AssignTo: &newWin,
		MinSize:  Size{MainWinMinWidth, MainWinMinHeight},
		Layout:   VBox{},

		MenuItems: []MenuItem{
			Menu{
				Text: "网口",
				Items: []MenuItem{
					Action{
						Text: "退出",
						OnTriggered: func() {
							newWin.Close()
						},
					},
				},
			},
			Menu{
				Text: "路由",
				Items: []MenuItem{
					Action{
						Text: "退出",
						OnTriggered: func() {
							newWin.Close()
						},
					},
				},
			},
			Menu{
				Text: "端口转发",
				Items: []MenuItem{
					Action{
						Text: "退出",
						OnTriggered: func() {
							newWin.Close()
						},
					},
				},
			},
			Menu{
				Text: "帮助",
				Items: []MenuItem{
					Action{
						Text: "关于",
						OnTriggered: func() {
							walk.MsgBox(newWin, "关于", BuildVersion, walk.MsgBoxIconInformation)
						},
					},
					Action{
						Text: "退出",
						OnTriggered: func() {
							newWin.Close()
						},
					},
				},
			},
		},

		Children: []Widget{
			TableView{
				AssignTo:      &tableView,
				Model:         dataModel,
				StretchFactor: 2,
				Columns: []TableViewColumn{
					TableViewColumn{
						DataMember: "No.",
						Alignment:  AlignCenter,
						Width:      128,
					},
					TableViewColumn{
						DataMember: "SrcIP",
						Alignment:  AlignCenter,
						Width:      128,
					},
					TableViewColumn{
						DataMember: "SrcPort",
						Alignment:  AlignCenter,
						Width:      128,
					},
					TableViewColumn{
						DataMember: "DstIP",
						Alignment:  AlignCenter,
						Width:      128,
					},
					TableViewColumn{
						DataMember: "DstPort",
						Alignment:  AlignCenter,
						Width:      128,
					},
					TableViewColumn{
						DataMember: "PKT",
						Alignment:  AlignCenter,
						Width:      128,
					},
					TableViewColumn{
						DataMember: "Seq",
						Alignment:  AlignCenter,
						Width:      500,
					},
				},
			},
		},
		StatusBarItems: []StatusBarItem{
			StatusBarItem{
				AssignTo: &sbi,
				Text:     "状态栏",
				Width:    MainWinMinWidth,
				OnClicked: func() {
					if sbi.Text() == "click" {
						sbi.SetText("again")
					} else {
						sbi.SetText("click")
					}
				},
			},
		},
	}.Create()

	newWin.Run()
}
