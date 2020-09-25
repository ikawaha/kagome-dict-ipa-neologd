package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKyBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKy, nil
}

func dictMod1IpaNeologdDictKy() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ky", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
