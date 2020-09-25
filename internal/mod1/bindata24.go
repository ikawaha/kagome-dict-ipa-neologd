package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKnBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKn, nil
}

func dictMod1IpaNeologdDictKn() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kn", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
