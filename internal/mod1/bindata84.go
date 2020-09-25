package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMvBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMv, nil
}

func dictMod1IpaNeologdDictMv() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
