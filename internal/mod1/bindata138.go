package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOxBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOx, nil
}

func dictMod1IpaNeologdDictOx() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ox", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
