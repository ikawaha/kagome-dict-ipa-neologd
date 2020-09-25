package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKjBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKj, nil
}

func dictMod1IpaNeologdDictKj() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
