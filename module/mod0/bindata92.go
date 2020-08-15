package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDoBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDo, nil
}

func dictMod0IpaNeologdDictDo() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.do", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
