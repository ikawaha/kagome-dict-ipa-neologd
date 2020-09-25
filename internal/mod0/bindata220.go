package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictImBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIm, nil
}

func dictMod0IpaNeologdDictIm() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictImBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.im", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
