package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDjBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDj, nil
}

func dictMod0IpaNeologdDictDj() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
