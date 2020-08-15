package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFpBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFp, nil
}

func dictMod0IpaNeologdDictFp() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fp", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
