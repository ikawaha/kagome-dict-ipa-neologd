package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictJjBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictJj, nil
}

func dictMod0IpaNeologdDictJj() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictJjBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.jj", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
