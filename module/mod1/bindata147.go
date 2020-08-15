package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPgBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPg, nil
}

func dictMod1IpaNeologdDictPg() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
