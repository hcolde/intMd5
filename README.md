# intMd5

## How to get

```shell
go get github.com/hcolde/intMd5
```

## Demo

1. Getting MD5 as number instead of string.

```go
package main

import (
    "fmt"
    "github.com/hcolde/intMd5"
)

func main() {
    md5 := "AE439DEDC6A1F7FA7C4C37A01D8A4297"
    md5Stu, err := intMd5.GetMd5(md5)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%#v", md5Stu)
    //intMd5.MD5{Md51:1543945436116761, Md52:7343371014814297, Temp:3349893712}
}

```

2. intMd5.MD5 struct to 32-char string.

```go
package main

import (
    "fmt"
    "github.com/hcolde/intMd5"
)

func main() {
    var md5Stu intMd5.MD5
    md5Stu.Md51 = 1543945436116761
    md5Stu.Md52 = 7343371014814297
    md5Stu.Temp = 3349893712
    md5 = intMd5.Reset(md5Stu)
    fmt.Println(md5) // "AE439DEDC6A1F7FA7C4C37A01D8A4297"
}

```

## Note

**In my broken English**

1. Please make sure to pass in the correct md5!At least it is a 32-char string!
2. In fact,you can pass in a string likes "A153GI...",even if it does not look like a md5 string,but make sure that the character is between 'a' and 'i',and it is Case-insensitive.

**注意**

1. 请确保传入正确的md5，至少保证是一个32位字符组成的字符串！
2. 实际上你传入的字符串中的字符可以为a至i之间，并且不区分大小写。
