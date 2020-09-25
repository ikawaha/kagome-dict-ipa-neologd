package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMjBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMj, nil
}

func dictMod1IpaNeologdDictMj() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
