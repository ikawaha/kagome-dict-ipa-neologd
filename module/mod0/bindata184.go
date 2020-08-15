package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictHcBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictHc, nil
}

func dictMod0IpaNeologdDictHc() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictHcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.hc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
