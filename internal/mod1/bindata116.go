package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictObBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOb, nil
}

func dictMod1IpaNeologdDictOb() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictObBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.ob", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
