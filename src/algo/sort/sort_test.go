package sort

import (
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	array := []int{1, 4, 3, 5, 6, 7, 3, 1}
	res := BubbleSort(array)
	fmt.Println(res)
}
func TestInsertSort(t *testing.T) {
	array := []int{1, 4, 3, 5, 6, 7, 3, 1}
	res := InsertSort(array)
	fmt.Println(res)
}
func TestSelectSort(t *testing.T) {
	array := []int{1, 4, 3, 5, 6, 7, 3, 1}
	res := SelectSort(array)
	fmt.Println(res)
}
func TestShellSort(t *testing.T) {
	array := []int{1, 4, 3, 5, 6, 7, 3, 1}
	res := ShellSort(array)
	fmt.Println(res)
}
func TestMergeSort(t *testing.T) {
	array := []int{1, 4, 15, 10, 3, 5, 6, 7, 3, 1}
	res := MergeSort(array)
	fmt.Println(res)
}
func TestQuickSort(t *testing.T) {
	array := []int{1, 4, 15, 10, 3, 5, 6, 7, 3, 1}
	res := QuickSort(array)
	fmt.Println(res)
}

func TestMergeArray(t *testing.T) {
	ary1 := []int{1, 3, 5, 36}
	ary2 := []int{2, 4, 6, 9, 24}
	res := MergeArray(ary1, ary2)
	fmt.Println(res)
}

func BenchmarkInsertSort(b *testing.B) {
	b.StartTimer()
	fmt.Printf("[GOMAXPROCS=%d, NUM_CPU=%d, NUM_GOROUTINE=%d]",
		runtime.GOMAXPROCS(-1), runtime.NumCPU(), runtime.NumGoroutine())
	for i := 0; i < b.N; i++ {
		array := []int{1, 4, 3, 5, 6, 7, 3, 1}
		_ = BubbleSort(array)
		go func() {
			fmt.Println(1)
		}()
	}

	b.StopTimer()
}
func BenchmarkTest(b *testing.B) {
	curstomTimerTag := true
	if curstomTimerTag {
		b.StopTimer()
	}
	time.Sleep(time.Nanosecond)
	if curstomTimerTag {
		b.StartTimer()
		b.SetBytes(1)
	}
}

func ExampleTest() {
	fmt.Printf("hello")
	// Output: hello
}
