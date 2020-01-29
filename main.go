package main

import (
	"net/http"
	"fmt"
)

// name ...
func main()  {
	targeturl := "https://b.hatena.ne.jp/"
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
