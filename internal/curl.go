package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"scripts/global"
	"scripts/pb"
	"time"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

const issueDate = "2022-07-28"

func AddCoin(ctx context.Context, uid, amount uint32) error {
	url := fmt.Sprintf("%s%s", global.Conf.AddCoinClient.Host, global.Conf.AddCoinClient.Path)
	req := &pb.ReqCoin{
		OrderId:   fmt.Sprintf("add_%d_%d_%s", uid, amount, issueDate),
		UserId:    uid,
		MoneyType: 0,
		Op:        0,
		OpDetail:  0,
		Amount:    int32(amount),
		Ts:        time.Now().Unix(),
		Desc:      "恢复用户金币",
	}
	rsp, err := g.Client().Post(ctx, url, req)
	if err != nil {
		return err
	}
	data := rsp.ReadAll()
	fmt.Printf("rsp: %s\n", data)
	rspd := &pb.RspCoin{}
	json.Unmarshal(data, rspd)
	if !rspd.GetSuccess() {
		return gerror.New(rspd.GetMsg())
	}

	return nil
}

func AddExp(ctx context.Context, uid, exp uint32) error {
	url := fmt.Sprintf("%s%s", global.Conf.AddExpClient.Host, global.Conf.AddExpClient.Path)
	req := &pb.ReqAddExp{
		UserId: uid,
		Exp:    exp,
		Desc:   "恢复用户经验值",
	}
	rsp, err := g.Client().Post(ctx, url, req)
	if err != nil {
		return err
	}
	data := rsp.ReadAll()
	fmt.Printf("rsp: %s\n", data)
	rspd := &pb.RspAddExp{}
	json.Unmarshal(data, rspd)
	if !rspd.GetSuccess() {
		return gerror.New(rspd.GetMsg())
	}

	return nil
}
