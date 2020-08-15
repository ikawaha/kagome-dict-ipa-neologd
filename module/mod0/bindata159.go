package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGdBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGd, nil
}

func dictMod0IpaNeologdDictGd() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gd", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
