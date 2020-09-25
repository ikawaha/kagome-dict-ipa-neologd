package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKoBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKo, nil
}

func dictMod1IpaNeologdDictKo() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ko", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
