package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOuBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOu, nil
}

func dictMod1IpaNeologdDictOu() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ou", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
