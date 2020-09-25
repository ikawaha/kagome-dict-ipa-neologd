package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPxBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPx, nil
}

func dictMod1IpaNeologdDictPx() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPxBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.px", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
