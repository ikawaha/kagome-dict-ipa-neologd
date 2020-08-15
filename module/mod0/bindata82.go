package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDeBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDe, nil
}

func dictMod0IpaNeologdDictDe() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.de", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
