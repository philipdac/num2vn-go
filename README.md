# num2vn-go
A go package that converts number to string in Vietnamese language

## Usages

Code example:

```
package main

import (
	"fmt"
	"num2vn"
)

func main() {
	fmt.Println("%v >>> %s", 1234567890, num2vn.Int2Vn(1234567890))
	// 1234567890 >>> Một tỷ hai trăm ba tư triệu năm trăm sáu bảy nghìn tám trăm chín mươi

	fmt.Println("%v >>> %s", 12345.6789, num2vn.Float2Vn(12345.6789))
	// 12345.67890 >>> Mười hai nghìn ba trăm bốn lăm phẩy sáu nghìn bảy trăm tám chín
}
```

## Conversion highlights

#### Outcome string as short as posible
- "lẻ" instead of "linh"
- "hai lăm" instead of "hai mươi lăm"
- "ba tư" instead of "ba mươi tư"

#### Numeric range:
- Int64
- Float64
