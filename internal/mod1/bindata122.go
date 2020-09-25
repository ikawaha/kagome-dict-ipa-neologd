package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOhBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOh, nil
}

func dictMod1IpaNeologdDictOh() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOhBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.oh", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
