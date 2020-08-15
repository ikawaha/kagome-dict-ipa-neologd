package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDcBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDc, nil
}

func dictMod0IpaNeologdDictDc() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
