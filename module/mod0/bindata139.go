package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFjBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFj, nil
}

func dictMod0IpaNeologdDictFj() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
