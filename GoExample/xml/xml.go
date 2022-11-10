package main

import "encoding/xml"

// Go 通过 encoding.xml 包为 XML 和 类-XML 格式提供了内建支持

// Plant 结构将被映射到 XML
// 与 JSON 示例类似，字段标签包含用于编码器和解码器的指令
// 这里使用了 XML 包的一些特性： XMLName 字段名称规定了该结构的 XML 元素名称； id,attrr 表示 Id 字段是一个 XML 属性，而不是嵌套元素
type Plant struct {
	XMLName xml.Name `xml:"plant"`
}

func main() {

}
