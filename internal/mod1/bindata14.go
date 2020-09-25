package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKdBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKd, nil
}

func dictMod1IpaNeologdDictKd() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kd", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
