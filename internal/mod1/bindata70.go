package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMhBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMh, nil
}

func dictMod1IpaNeologdDictMh() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
