package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictJqBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictJq, nil
}

func dictMod1IpaNeologdDictJq() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictJqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.jq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
