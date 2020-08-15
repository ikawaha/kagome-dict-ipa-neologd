package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictJtBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictJt, nil
}

func dictMod1IpaNeologdDictJt() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictJtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.jt", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
