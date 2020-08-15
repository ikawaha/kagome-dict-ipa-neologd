package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFgBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFg, nil
}

func dictMod0IpaNeologdDictFg() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFgBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fg", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
