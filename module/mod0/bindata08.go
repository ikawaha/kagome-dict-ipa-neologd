package data

import(
	"os"
	"time"
)


func dictMod0IpaNeologdDictAiBytes() ([]byte, error) {
	return _dictMod0IpaNeologdDictAi, nil
}

func dictMod0IpaNeologdDictAi() (*asset, error) {
	bytes, err := dictMod0IpaNeologdDictAiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dict/mod0/ipa-neologd-dict.ai", size: 524288, mode: os.FileMode(420), modTime: time.Unix(1597496709, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
