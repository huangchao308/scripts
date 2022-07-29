package internal

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

var expConfig = []uint32{0, 30, 170, 350, 580, 920, 1310, 1760, 2270, 2840, 3600, 4440, 5350, 6340, 7400, 8540, 9750, 11040, 12400, 13840, 15660, 17580, 19590, 21690, 23890, 26180, 28570, 31060, 33640, 36320, 39630, 43060, 46600, 50260, 54040, 57940, 61960, 66100, 70360, 74740, 80070, 85540, 91160, 96930, 102840, 108900, 115110, 121470, 127980, 134630, 142640, 150830, 159200, 167750, 176480, 185390, 194490, 203770, 213230, 222870, 234390, 246130, 258090, 270270, 282680, 295310, 308160, 321240, 334540, 348070, 365010, 382240, 399750, 417550, 435640, 454010, 472670, 491620, 510860, 530390}

func HandleExp() error {
	dir, _ := os.Getwd()
	path := filepath.Join(dir, "user_data.xlsx")
	sf, err := GetExcelizeFile(path)
	if err != nil {
		return err
	}
	data, err := OpenExcel(sf, sheet)
	if err != nil {
		return err
	}
	for i, v := range data {
		fmt.Printf("uid: %d, coin: %d, exp: %d, handle_exp: %d, handle_coin: %d, cur_lv: %d, cur_exp: %d\n",
			v.Uid, v.Coin, v.Exp, v.HandleExp, v.HandleCoin, v.CurLv, v.CurExp)
		err := handleOneExp(v, i+2)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return sf.Save()
}

func handleOneExp(info *UserGameInfo, row int) error {
	// 已经处理过的就不再处理了
	if info.HandleExp == 1 {
		return nil
	}
	err := AddExp(context.Background(), info.Uid, info.Exp)
	if err != nil {
		return err
	}
	dir, _ := os.Getwd()
	path := filepath.Join(dir, "user_data.xlsx")
	sf, err := GetExcelizeFile(path)
	if err != nil {
		return err
	}
	SetCellValue(sf, sheet, fmt.Sprintf("E%d", row), 1)

	return nil
}

func CalExp() error {
	dir, _ := os.Getwd()
	path := filepath.Join(dir, "user_data.xlsx")
	sf, err := GetExcelizeFile(path)
	if err != nil {
		return err
	}
	data, err := OpenExcel(sf, sheet)
	if err != nil {
		return err
	}
	for i, v := range data {
		fmt.Printf("uid: %d, coin: %d, exp: %d, handle_exp: %d, handle_coin: %d, cur_lv: %d, cur_exp: %d\n",
			v.Uid, v.Coin, v.Exp, v.HandleExp, v.HandleCoin, v.CurLv, v.CurExp)
		err := calOneExp(v, i+2)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return sf.Save()
}

func calOneExp(info *UserGameInfo, row int) error {
	if info.CurLv == 0 {
		return nil
	}
	newExp := getExpByLevel(info.CurLv, info.CurExp)
	dir, _ := os.Getwd()
	path := filepath.Join(dir, "user_data.xlsx")
	sf, err := GetExcelizeFile(path)
	if err != nil {
		return err
	}
	SetCellValue(sf, sheet, fmt.Sprintf("C%d", row), int(newExp))
	return nil
}

func getExpByLevel(level int, cur uint32) uint32 {
	if level == 0 {
		return 0
	}
	return expConfig[level-1] + cur
}
