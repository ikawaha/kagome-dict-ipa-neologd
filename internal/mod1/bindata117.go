package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictOcBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictOc, nil
}

func dictMod1IpaNeologdDictOc() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictOcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.oc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
