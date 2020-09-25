package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictFsBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictFs, nil
}

func dictMod0IpaNeologdDictFs() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictFsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.fs", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
