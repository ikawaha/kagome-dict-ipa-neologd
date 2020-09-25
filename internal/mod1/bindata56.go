package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLtBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLt, nil
}

func dictMod1IpaNeologdDictLt() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lt", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
