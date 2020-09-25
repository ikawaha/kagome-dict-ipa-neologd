package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJkBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJk, nil
}

func dictMod0IpaNeologdDictJk() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.jk", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
