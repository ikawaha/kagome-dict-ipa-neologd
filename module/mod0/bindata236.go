package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJcBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJc, nil
}

func dictMod0IpaNeologdDictJc() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.jc", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
