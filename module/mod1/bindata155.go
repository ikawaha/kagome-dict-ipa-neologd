package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPoBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPo, nil
}

func dictMod1IpaNeologdDictPo() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.po", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
