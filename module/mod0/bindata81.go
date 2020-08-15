package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDdBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDd, nil
}

func dictMod0IpaNeologdDictDd() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dd", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
