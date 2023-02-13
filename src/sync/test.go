// 创建分两种
package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	eg1 := new(errgroup.Group)

	// 执行
	for _, v := range urls {
		eg2.Go(func() error {
			resp, err := http.Get(v)
			if err != nil {
				resp.Body.Close()
			}
			return err
		})
	}

	if err := eg2.Wait(); err != nil {
		fmt.Println(err)
	}
}
