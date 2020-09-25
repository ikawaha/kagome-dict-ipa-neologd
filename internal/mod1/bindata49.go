package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLmBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLm, nil
}

func dictMod1IpaNeologdDictLm() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lm", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
