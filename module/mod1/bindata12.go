package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKbBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKb, nil
}

func dictMod1IpaNeologdDictKb() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
