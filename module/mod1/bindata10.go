package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictJzBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictJz, nil
}

func dictMod1IpaNeologdDictJz() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictJzBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.jz", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
