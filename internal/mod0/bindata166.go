package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGkBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGk, nil
}

func dictMod0IpaNeologdDictGk() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gk", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
