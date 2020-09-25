package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictGaBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictGa, nil
}

func dictMod0IpaNeologdDictGa() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictGaBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ga", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
