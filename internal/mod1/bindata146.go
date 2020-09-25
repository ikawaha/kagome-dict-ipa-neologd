package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPfBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPf, nil
}

func dictMod1IpaNeologdDictPf() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pf", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
