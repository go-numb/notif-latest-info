package main

import (
	"errors"
	"fmt"
	mod "notif_latest_info"
	"time"

	"log"
)

const (
	MINITE = 60          // 60s
	HOUR   = MINITE * 60 // 3600s = 1Hour
	DAY    = HOUR * 24   // 3600s * 24h = 1Day
)

const (
	// 設定する項目群
	TARGETTERM = HOUR

	// 通知先
	WHERETO = ""
)

func main() {
	t := time.NewTicker(time.Duration(TARGETTERM) * time.Second)
	defer t.Stop()

	// 対象の型
	target := mod.NewSubsidies()

	for {
		// 取得タイミング
		<-t.C

		// 情報取得
		if err := target.Get(); err != nil {
			log.Println(errors.Join(err, fmt.Errorf("get target info")))
			continue
		}

		// 文字列生成
		msg, err := target.MakeMsg()
		if err != nil {
			log.Println(errors.Join(err, errors.New("make notif string")))
			continue
		}

		// 通知実行
		if err := mod.ToLine(WHERETO, msg); err != nil {
			log.Println(errors.Join(err, fmt.Errorf("notification info to customers")))
			continue
		}

	}

}
