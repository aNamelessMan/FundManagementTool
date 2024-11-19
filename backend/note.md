# 创建项目
1. 初始化
> cd backend  
> $ GO111MODULE=on  
> $ go mod init github.com/aNamelessMan/FundManagementTool
2. 验证
在backend/main.go里写个helloworld
```
package main

import "fmt"

func main() {
  fmt.Println("Fund Tool App v0.01")
}
```
> $ go run main.go
