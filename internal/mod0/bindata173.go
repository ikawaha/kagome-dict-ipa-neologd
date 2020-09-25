package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGrBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGr, nil
}

func dictMod0IpaNeologdDictGr() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gr", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
