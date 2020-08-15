package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJfBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJf, nil
}

func dictMod0IpaNeologdDictJf() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.jf", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
