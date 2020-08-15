package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHtBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHt, nil
}

func dictMod0IpaNeologdDictHt() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ht", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
