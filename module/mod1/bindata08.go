package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictJxBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictJx, nil
}

func dictMod1IpaNeologdDictJx() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictJxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.jx", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
