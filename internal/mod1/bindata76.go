package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMnBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMn, nil
}

func dictMod1IpaNeologdDictMn() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mn", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
