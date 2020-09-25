package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictCuBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictCu, nil
}

func dictMod0IpaNeologdDictCu() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictCuBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.cu", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
