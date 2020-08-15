package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOqBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOq, nil
}

func dictMod1IpaNeologdDictOq() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.oq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
