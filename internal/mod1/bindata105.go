package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNqBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNq, nil
}

func dictMod1IpaNeologdDictNq() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.nq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
