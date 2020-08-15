package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOwBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOw, nil
}

func dictMod1IpaNeologdDictOw() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ow", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
