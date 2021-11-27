package config

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var Prop *Properties

//初始化配置文件
func InitProp() {
	p := NewProperties()
	p.LoadFromFile("conf/zihao.properties")
	Prop = p
}

// 创建一个空属性列表。

// Create an empty property list.
func NewProperties() *Properties {
	return &Properties{
		object: make(map[string]string),
	}
}

type Properties struct {
	m      sync.Mutex
	object map[string]string
}

// 用指定的键在此属性列表中搜索属性。
//
// Searches for the property with the specified key in this property list.
func (p *Properties) Property(key string) (value string, isExist bool) {
	if value, ok := p.object[key]; !ok {
		return "", false
	} else {
		return value, true
	}
}

// 用指定的键在此属性列表中搜索属性，把","连接的多个属性转换为切片返回。
//
// Search for attributes in this attribute list using the specified key to return multiple attributes connected by "," converted to slices.
func (p *Properties) PropertySlice(key string) (values []string, isExist bool) {
	strs, isExist := p.Property(key)
	if !isExist {
		return nil, false
	}
	return strings.Split(strs, ","), true
}

// 为指定的键设置多个属性，把多个属性值转换成“，”连接的属性字符串。
//
// Set  multiple attributes for the specified key, converting multiple attribute values into a "," concatenated attribute string.
func (p *Properties) SetPropertySlice(key string, values ...string) {
	value := strings.Join(values, ",")
	p.SetProperty(key, value)
}

// 更新指定的键和属性,如果键不存在就新建。
//
// Update the specified key and properties. If the key does not exist, create a new one.
func (p *Properties) SetProperty(key, value string) {
	p.m.Lock()
	p.object[key] = value
	p.m.Unlock()
}

// 从输入流读取属性列表。
//
// Reads a property list (key and element pairs) from the input character stream in a simple line-oriented format.
func (p *Properties) Load(r io.Reader) error {
	br := bufio.NewReader(r)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		p.parseLine(line)
	}
	return nil
}

// 从文件中读取属性列表。
//
// Reads a property list from a file
func (p *Properties) LoadFromFile(filePath string) error {
	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		return fmt.Errorf("os.Open failed: %v", err)
	}
	return p.Load(f)
}

// 将属性列表写入输出流。
//
// Writes this property list (key and element pairs) in this Properties table to the output stream in a format suitable for loading into a Properties table using the Load() method.
func (p *Properties) Store(w io.Writer) error {
	var buf bytes.Buffer
	for k, v := range p.object {
		p.line(k, v, &buf)
	}
	_, err := w.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("w.Write failed: %v", err)
	}
	return nil
}

// 将属性列表写入文件。
//
// Writes a list of property to a file.
func (p *Properties) StoreToFile(filePath string) error {
	f, err := os.Create(filePath)
	defer f.Close()

	if err != nil {
		return fmt.Errorf("os.Open failed: %v", err)
	}
	return p.Store(f)

}

func (p *Properties) line(key, value string, buf *bytes.Buffer) {
	buf.WriteString(key)
	buf.WriteString(" = ")
	buf.WriteString(value)
	buf.WriteByte('\n')
}

func (p *Properties) parseLine(line []byte) {
	lineStr := strings.TrimSpace(string(line))
	if isCommentline(lineStr) {
		return
	}
	splitStrs := strings.Split(lineStr, "=")
	key := strings.TrimSpace(splitStrs[0])
	value := strings.TrimSpace(splitStrs[1])
	p.SetProperty(key, value)
}

// 返回属性列表中所有键的枚举。
//
// Returns an enumeration of all keys in the property list.
func (p *Properties) PropertyNames() []string {
	var names []string
	for k := range p.object {
		names = append(names, k)
	}
	return names
}

// 过滤注释的行
func isCommentline(line string) bool {
	if len(line) == 0 {
		return true
	}
	if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "/*") {
		return true
	}
	return false
}
