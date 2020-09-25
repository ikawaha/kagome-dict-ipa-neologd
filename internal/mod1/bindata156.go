package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPpBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPp, nil
}

func dictMod1IpaNeologdDictPp() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pp", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
