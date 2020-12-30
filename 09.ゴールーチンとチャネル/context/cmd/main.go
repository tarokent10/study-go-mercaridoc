package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	url string = "" // バックエンドのURL
)

// サーバがリクエストを受けると、バックエンドのサーバへアクセスする処理をイメージ
// リクエストタイムアウト３秒とし、その際にバックエンドへの接続資源を解放する処理を実装してみる
func main() {
	handler()
}
func handler() {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Microsecond)
	defer cancel() // not forget!

	errChan := make(chan error, 1)
	go func() {
		errChan <- requestWithContext(ctx)
	}()
	select {
	case err := <-errChan:
		if err != nil {
			fmt.Printf("faild: %+v\n", err)
			return
		}
	}
	fmt.Println("success!")

}

func requestWithContext(ctx context.Context) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}

	// バックエンドにGoルーチンでリクエスト
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	errCh := make(chan error, 1)
	go func() {
		res, err := client.Do(req)
		if err == nil {
			bytes, _ := ioutil.ReadAll(res.Body)
			fmt.Printf("%s\n", string(bytes))
		}
		errCh <- errors.Wrap(err, "wraped")
	}()

	// http requestとcontext doneの２つのチャンネルをWatch（いずれかだけ処理）
	// キャンセル時はコネクションをキャンセル
	var ret error
	select {
	case err := <-errCh:
		ret = err
	case <-ctx.Done():
		tr.CancelRequest(req) // キャンセル処理
		fmt.Printf("context done: %+v\n", <-errCh)
		ret = ctx.Err()
	}
	return ret
}
