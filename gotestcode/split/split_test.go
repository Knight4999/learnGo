package split_test

import (
	"learnGo/gotestcode/split"
	"reflect"
	"testing"
)

/*
	func TestSplit(t *testing.T) {
		got := split.Split("我爱你", "爱")
		want := []string{"我", "你"}
		//使用reflect包中函数来比较slice
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want：%v got:%v", want, got)
		}
	}
*/
//测试函数
func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := map[string]test{
		"中文":  test{"我爱你", "爱", []string{"我", "你"}},
		"abc": test{"a:b:c", ":", []string{"a", "b", "c"}},
		"a-g": test{"abcdefg", "cde", []string{"ab", "fg"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { //通过t.run 查看测试用例细节
			got := split.Split(tc.input, tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("name:%v want:%v got:%v", name, tc.want, got)
			}
		})
	}
}

// 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		split.Split("北京人北京魂北京骗子北", "北")
	}
}
