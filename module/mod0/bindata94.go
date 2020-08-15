package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDqBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDq, nil
}

func dictMod0IpaNeologdDictDq() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDqBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.dq", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
