package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPbBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPb, nil
}

func dictMod1IpaNeologdDictPb() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
