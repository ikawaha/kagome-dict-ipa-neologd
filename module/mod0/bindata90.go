package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDmBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDm, nil
}

func dictMod0IpaNeologdDictDm() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDmBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dm", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
