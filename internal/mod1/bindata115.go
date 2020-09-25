package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOaBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOa, nil
}

func dictMod1IpaNeologdDictOa() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.oa", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
