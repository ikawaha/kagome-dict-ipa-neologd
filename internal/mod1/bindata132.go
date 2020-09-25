package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOrBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOr, nil
}

func dictMod1IpaNeologdDictOr() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.or", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
