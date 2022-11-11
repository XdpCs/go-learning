package main

// Go 提供了对 base64 编解码的内建支持

// 这个语法引入了 encoding/base64 包， 并使用别名 b64 代替默认的 base64。这样可以节省点空间
import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	// 这是要编解码的字符串
	data := "abc123!?$*&()'-=@~"
	// Go 同时支持标准 base64 以及 URL 兼容 base64
	// 这是使用标准编码器进行编码的方法
	// 编码器需要一个 []byte，因此将 string 转换为该类型
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码可能会返回错误，如果不确定输入信息格式是否正确， 那么，就需要进行错误检查了
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 使用 URL base64 格式进行编解码
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
	// 标准 base64 编码和 URL base64 编码的 编码字符串存在稍许不同（后缀为 + 和 -）， 但是两者都可以正确解码为原始字符串
}
