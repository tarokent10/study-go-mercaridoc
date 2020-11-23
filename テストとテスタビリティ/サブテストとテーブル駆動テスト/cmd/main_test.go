package main

import (
	"testing"
)

var hexTests = []struct {
	name string
	in   int32
	out  string
}{
	// テストデータにラベリングができる.データ量が多いときやパターン示したいとき役立つかも
	{"test:10", 10, "a"}, {"test:16", 16, "10"}, {"test:80", 80, "50"},
}

func TestHex(t *testing.T) {
	for _, tt := range hexTests {
		t.Run(tt.name, func(t *testing.T) {
			if result := Hex(tt.in).string(); result != tt.out {
				testhelper(t, result, tt.out)
			}
		})
	}
}

func testhelper(t *testing.T, expect, actual string) {
	t.Helper() //アサーションで行数が呼び元で表示される.共通で使うヘルパーで使うとトレースしやすい.
	t.Errorf("failed: expect=%s but actual=%s", expect, actual)
}
