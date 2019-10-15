# regioncn-golang-client

### 安装

```
go get github.com/yingzhuo/regioncn-golang-client
```

### 使用例

```golang
package main

import (
	"fmt"
	cli "github.com/yingzhuo/regioncn-golang-client"
)

func main() {
	client := cli.NewProtobufClient("localhost", 8080)

	for _, v := range client.GetAllProvince() {
		fmt.Println(v)
	}

	for _, v := range client.GetCityByProvinceCode("430000") {
		fmt.Println(v)
	}

	for _, v := range client.GetAreaByCityCode("430100") {
		fmt.Println(v)
	}

	for _, v := range client.GetStreetByAreaCode("430102") {
		fmt.Println(v)
	}
}

```
