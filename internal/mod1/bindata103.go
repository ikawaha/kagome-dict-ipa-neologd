package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNoBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNo, nil
}

func dictMod1IpaNeologdDictNo() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.no", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
