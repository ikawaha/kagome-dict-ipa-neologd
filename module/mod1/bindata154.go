package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPnBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPn, nil
}

func dictMod1IpaNeologdDictPn() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pn", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
