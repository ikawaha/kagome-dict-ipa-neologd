package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDrBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDr, nil
}

func dictMod0IpaNeologdDictDr() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dr", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
