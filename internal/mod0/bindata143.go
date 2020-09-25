package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFnBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFn, nil
}

func dictMod0IpaNeologdDictFn() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fn", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
