package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBrBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBr, nil
}

func dictMod0IpaNeologdDictBr() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBrBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.br", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
