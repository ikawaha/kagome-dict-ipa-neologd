package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCkBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCk, nil
}

func dictMod0IpaNeologdDictCk() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ck", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
