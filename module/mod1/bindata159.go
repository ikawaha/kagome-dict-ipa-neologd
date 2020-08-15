package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPsBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPs, nil
}

func dictMod1IpaNeologdDictPs() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ps", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
