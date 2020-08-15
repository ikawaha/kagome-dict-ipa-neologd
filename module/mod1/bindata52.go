package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLpBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLp, nil
}

func dictMod1IpaNeologdDictLp() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.lp", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
