package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKcBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKc, nil
}

func dictMod1IpaNeologdDictKc() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
