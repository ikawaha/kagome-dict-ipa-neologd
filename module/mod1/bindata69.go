package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMgBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMg, nil
}

func dictMod1IpaNeologdDictMg() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
