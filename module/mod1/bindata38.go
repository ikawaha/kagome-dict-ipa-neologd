package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLbBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLb, nil
}

func dictMod1IpaNeologdDictLb() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
