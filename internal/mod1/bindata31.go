package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictKuBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictKu, nil
}

func dictMod1IpaNeologdDictKu() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictKuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ku", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
