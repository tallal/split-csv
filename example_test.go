package split_csv_test

import (
	"fmt"
	splitCsv "github.com/tolik505/split-csv"
)

func ExampleSplitCsv() {
	splitter := splitCsv.New()
	splitter.FileChunkSize = 800 //in bytes
	result, _ := splitter.Split("testdata/test.csv", "testdata/")
	fmt.Println(result)
	// Output: [testdata/test_1.csv testdata/test_2.csv testdata/test_3.csv]
}

func ExampleSplitCsvFullWindowsPath() {
	splitter := splitCsv.New()
	splitter.FileChunkSize = 800 //in bytes
	result, _ := splitter.Split("C:\\dev\\go\\src\\split-csv\\testdata\\test.csv", "C:\\dev\\go\\src\\split-csv\\testdata\\")
	fmt.Println(result)
	// Output: [C:\dev\go\src\split-csv\testdata\test_1.csv C:\dev\go\src\split-csv\testdata\test_2.csv C:\dev\go\src\split-csv\testdata\test_3.csv]
}

func ExampleSplitCsvWindowsUNCPath1() {
	splitter := splitCsv.New()
	splitter.FileChunkSize = 6 * 1024 * 1024 //in bytes
	result, _ := splitter.Split("\\\\oberon\\steamBackups\\100000test.csv", "\\\\oberon\\steamBackups\\")
	fmt.Println(result)
	// Output: [\\oberon\steamBackups\100000test_1.csv \\oberon\steamBackups\100000test_2.csv \\oberon\steamBackups\100000test_3.csv]
}

func ExampleSplitCsvWindowsUNCPath2() {
	splitter := splitCsv.New()
	splitter.FileChunkSize = 50 * 1024 * 1024 //in bytes
	result, _ := splitter.Split("\\\\oberon\\steamBackups\\1500000 Sales Records.csv", "\\\\oberon\\steamBackups\\")
	fmt.Println(result)
	// Output: [\\oberon\steamBackups\1500000 Sales Records_1.csv \\oberon\steamBackups\1500000 Sales Records_2.csv \\oberon\steamBackups\1500000 Sales Records_3.csv \\oberon\steamBackups\1500000 Sales Records_4.csv]
}
