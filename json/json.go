package json

import (
	"bytes"
	"encoding/json"
)

var newLineBytes = []byte("\n")

// Marshal 标准库默认的 json.Marshal 会对一些字符进行转义（<、>、&等），使用此函数可避免转义。
//
// 参考：How to stop json.Marshal from escaping < and >?(https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and)
func Marshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	var retBytes []byte
	if err == nil {
		retBytes = bytes.TrimSuffix(buffer.Bytes(), newLineBytes) // encoder.Encode() 会在末尾添加换行符
	}
	return retBytes, err
}
