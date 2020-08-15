package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBfBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBf, nil
}

func dictMod0IpaNeologdDictBf() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bf", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
