package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHsBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHs, nil
}

func dictMod0IpaNeologdDictHs() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hs", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
