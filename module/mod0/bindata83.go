package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictDfBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictDf, nil
}

func dictMod0IpaNeologdDictDf() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictDfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.df", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
