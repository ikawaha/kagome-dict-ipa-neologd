package data

import(
	"os"
	"time"
)


func dictMod1IpaNeologdDictMuBytes() ([]byte, error) {
	return _dictMod1IpaNeologdDictMu, nil
}

func dictMod1IpaNeologdDictMu() (*asset, error) {
	bytes, err := dictMod1IpaNeologdDictMuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod1/ipa-neologd-dict.mu", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
