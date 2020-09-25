package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFqBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFq, nil
}

func dictMod0IpaNeologdDictFq() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
