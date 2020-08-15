package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictLnBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictLn, nil
}

func dictMod1IpaNeologdDictLn() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictLnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ln", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
