package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPkBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPk, nil
}

func dictMod1IpaNeologdDictPk() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pk", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
