package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictBeBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictBe, nil
}

func dictMod0IpaNeologdDictBe() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictBeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.be", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1601000133, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
