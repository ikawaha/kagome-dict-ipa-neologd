package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictPtBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictPt, nil
}

func dictMod1IpaNeologdDictPt() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictPtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.pt", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
