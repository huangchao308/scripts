package internal

import (
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xuri/excelize/v2"
)

const sheet = "Sheet1"

type UserGameInfo struct {
	Uid        uint32
	Exp        uint32
	Coin       uint32
	HandleExp  int
	HandleCoin int
	CurLv      int    // 上一次登录时日志打印的用户等级
	CurExp     uint32 // 上一次登录时日志打印的用户经验
}

var ef *excelize.File

func GetExcelizeFile(path string) (*excelize.File, error) {
	lock := sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()
	if ef == nil {
		f, err := excelize.OpenFile(path)
		if err != nil {
			return nil, err
		}
		ef = f
	}

	return ef, nil
}

func OpenExcel(f *excelize.File, sheet string) ([]*UserGameInfo, error) {
	results := make([]*UserGameInfo, 0)
	for i := 2; i < 123; i++ {
		f.GetCellValue(sheet, fmt.Sprintf("A%d", i))
	}
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}
	for j, row := range rows {
		// 第一行是表头, 跳过
		if j == 0 {
			continue
		}
		item := &UserGameInfo{}
		for i, r := range row {
			switch i {
			case 0:
				item.Uid = gconv.Uint32(r)
			case 1:
				item.Coin = gconv.Uint32(r)
			case 2:
				item.Exp = gconv.Uint32(r)
			case 3:
				item.HandleCoin = gconv.Int(r)
			case 4:
				item.HandleExp = gconv.Int(r)
			case 5:
				item.CurLv = gconv.Int(r)
			case 6:
				item.CurExp = gconv.Uint32(r)
			}
		}
		results = append(results, item)
	}
	return results, nil
}

func SetCellValue(f *excelize.File, sheet, axis string, value int) error {
	return f.SetCellInt(sheet, axis, value)
}
