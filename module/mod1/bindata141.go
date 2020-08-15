package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPaBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPa, nil
}

func dictMod1IpaNeologdDictPa() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pa", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
