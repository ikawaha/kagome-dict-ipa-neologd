package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKpBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKp, nil
}

func dictMod1IpaNeologdDictKp() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.kp", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
