package internal

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
)

func HandleCoin() error {
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
		fmt.Printf("uid: %d, coin: %d, exp: %d, handle_exp: %d, handle_coin: %d\n",
			v.Uid, v.Coin, v.Exp, v.HandleExp, v.HandleCoin)
		err := HandleOneCoin(v, i+2)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	return sf.Save()
}

func HandleOneCoin(info *UserGameInfo, row int) error {
	// 已经处理过的就不再处理了
	if info.HandleCoin == 1 {
		return nil
	}
	err := AddCoin(context.Background(), info.Uid, info.Coin)
	if err != nil {
		return err
	}
	dir, _ := os.Getwd()
	path := filepath.Join(dir, "user_data.xlsx")
	sf, err := GetExcelizeFile(path)
	if err != nil {
		return err
	}
	SetCellValue(sf, sheet, fmt.Sprintf("D%d", row), 1)

	return nil
}
