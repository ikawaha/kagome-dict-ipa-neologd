package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNxBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNx, nil
}

func dictMod1IpaNeologdDictNx() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.nx", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
