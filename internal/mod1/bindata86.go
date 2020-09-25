package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMxBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMx, nil
}

func dictMod1IpaNeologdDictMx() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mx", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
