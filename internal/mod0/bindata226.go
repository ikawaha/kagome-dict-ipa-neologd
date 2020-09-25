package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictIsBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictIs, nil
}

func dictMod0IpaNeologdDictIs() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictIsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.is", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
