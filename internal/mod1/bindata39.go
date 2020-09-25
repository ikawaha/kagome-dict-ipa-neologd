package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLcBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLc, nil
}

func dictMod1IpaNeologdDictLc() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
