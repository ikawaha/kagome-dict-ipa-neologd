package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJnBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJn, nil
}

func dictMod0IpaNeologdDictJn() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJnBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.jn", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
