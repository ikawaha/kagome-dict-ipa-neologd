package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGfBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGf, nil
}

func dictMod0IpaNeologdDictGf() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.gf", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
