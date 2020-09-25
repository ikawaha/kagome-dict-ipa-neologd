package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLoBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLo, nil
}

func dictMod1IpaNeologdDictLo() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lo", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
