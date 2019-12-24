# intMd5

## How to get

```shell
go get github.com/hcolde/intMd5
```

## Demo

```go
import (
    "fmt"
    "github.com/hcolde/intMd5"
)

func main() {
    md5 := "AE439DEDC6A1F7FA7C4C37A01D8A4297"
    md5Stu := intMd5.getMd5(md5)
    fmt.Println(md5Stu.md51) //1543945436116761
    fmt.Println(md5Stu.md52) //7343371014814297
    fmt.Println(md5Stu.temp) //3349893712
    
    md5 = intMd5.reset(md5Stu)
    fmt.Println(md5) // "AE439DEDC6A1F7FA7C4C37A01D8A4297"
}
```