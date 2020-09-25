package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLjBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLj, nil
}

func dictMod1IpaNeologdDictLj() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
