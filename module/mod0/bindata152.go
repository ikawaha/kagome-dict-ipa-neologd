package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFwBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFw, nil
}

func dictMod0IpaNeologdDictFw() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFwBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fw", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
