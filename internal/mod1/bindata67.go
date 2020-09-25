package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMeBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMe, nil
}

func dictMod1IpaNeologdDictMe() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.me", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
