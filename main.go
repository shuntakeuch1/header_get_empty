package main

import (
	"net/http"
	"fmt"
)

// name ...
func main()  {
	// 取得したいURL
	targeturl := "https://b.hatena.ne.jp/"
	// 取得したいHTTPヘッダー
	headkey := "X-Amz-Cf-Id"

	resp, err := http.Get(targeturl)
	if err != nil {

	}
	// 必ずCloseする
	defer resp.Body.Close()

	cfid := resp.Header.Get(headkey)
	if cfid != "" {
		fmt.Printf("%sが存在しました: %s\n", headkey, cfid)
	}
}
