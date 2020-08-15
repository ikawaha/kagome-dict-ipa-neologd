package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictNbBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictNb, nil
}

func dictMod1IpaNeologdDictNb() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictNbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.nb", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
