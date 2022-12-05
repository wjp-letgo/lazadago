# lazada go language API

API interface for lazada go language version

my email :474790700@qq.com

## Contents

- [lazada go language API](#lazada-go)
  - [Installation](#installation)
  - [Quick start](#quick-start)

## Installation

To install lazadago package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed (**version 1.12+ is required**), then you can use the below Go command to install lazadago.

```sh
$ go get -u github.com/wjpxxx/lazadago
```

2. Import it in your code:

```go
import (
	"github.com/wjpxxx/lazadago"
)
```
## Quick start

## API call

```go
package main

import (
	"github.com/wjpxxx/lazadago"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	"fmt"
)

func main() {
	api:=lazadago.NewApi(&lazadaConfig.Config{
		AppKey:"your app key",
		AccessToken:"your Access Token", //刚开始可以为空字符串
		AppSecret:"your app Secret",
		Country:"ph",
	})
	//first Seller authorization introduction
	fmt.Println(api.AuthorizationURL("redirect_uri"))
	//generate access_token for call api
	rs:=api.GenerateAccessToken("code","")
	fmt.Println(rs)
	//set access token
	api.SetAccessToken("your Access Token")

	//Use this API to get the list of items for a single order.
	order:=api.GetOrder("order_id")
	fmt.Println(order)
	//Use this API to get the list of items for a range of orders1
	orders:=api.GetOrders("2018-02-10T16:00:00+08:00","DESC",0,10,"2017-02-10T09:00:00+08:00","updated_at","","","shipped")
	fmt.Println(orders)
}

```
