package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictJyBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictJy, nil
}

func dictMod1IpaNeologdDictJy() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictJyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.jy", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
