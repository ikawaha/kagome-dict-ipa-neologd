module github.com/ikawaha/kagome-dict-ipa-neologd

go 1.15

replace (
	github.com/ikawaha/kagome-dict-ipa-neologd/internal/mod0 => ./internal/mod0
	github.com/ikawaha/kagome-dict-ipa-neologd/internal/mod1 => ./internal/mod1
)

require (
	github.com/ikawaha/kagome-dict-ipa-neologd/internal/mod0 v0.0.0-00010101000000-000000000000
	github.com/ikawaha/kagome-dict-ipa-neologd/internal/mod1 v0.0.0-00010101000000-000000000000
	github.com/ikawaha/kagome/v2 v2.0.4
)
