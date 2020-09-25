package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNtBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNt, nil
}

func dictMod1IpaNeologdDictNt() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.nt", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
