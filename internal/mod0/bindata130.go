package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFaBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFa, nil
}

func dictMod0IpaNeologdDictFa() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fa", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
