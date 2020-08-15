package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBsBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBs, nil
}

func dictMod0IpaNeologdDictBs() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.bs", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
