package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNdBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNd, nil
}

func dictMod1IpaNeologdDictNd() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.nd", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
