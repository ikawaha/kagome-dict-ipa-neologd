package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOzBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOz, nil
}

func dictMod1IpaNeologdDictOz() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.oz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
