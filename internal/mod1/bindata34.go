package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKxBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKx, nil
}

func dictMod1IpaNeologdDictKx() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kx", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
