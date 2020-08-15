package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHnBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHn, nil
}

func dictMod0IpaNeologdDictHn() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hn", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
