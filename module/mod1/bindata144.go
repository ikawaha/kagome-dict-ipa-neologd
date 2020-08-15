package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPdBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPd, nil
}

func dictMod1IpaNeologdDictPd() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pd", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
