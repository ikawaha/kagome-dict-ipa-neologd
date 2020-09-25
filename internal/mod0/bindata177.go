package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGvBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGv, nil
}

func dictMod0IpaNeologdDictGv() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGvBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gv", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
