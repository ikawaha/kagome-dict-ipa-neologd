package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictQbBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictQb, nil
}

func dictMod1IpaNeologdDictQb() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictQbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.qb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
