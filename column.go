package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println(`Usage:
  echo "First Second Third\nMyFirstField 2ndf ThisIsTheThirdField" | column`)
	} else {
		tabLenPtr := flag.Int("t", 8, "Tabulation length")
		minWidthPtr := flag.Int("w", 0, "Minimum cell width")
		paddingPtr := flag.Int("p", 0, "Cell padding")
		dividerChar := flag.String("d", " ", "The divider char")
		paddingByte := flag.String("c", "\t", "The padding byte")

		flag.Parse()

		reader := bufio.NewReader(os.Stdin)

		w := new(tabwriter.Writer)

		// Format in tab-separated columns with a tab stop of 8.
		//w.Init(os.Stdout, 0, 8, 0, '\t', 0)
		w.Init(os.Stdout, *minWidthPtr, *tabLenPtr, *paddingPtr, byte((*paddingByte)[0]), 0)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			lineStr := strings.Replace(string(line), *dividerChar, "\t", -1)
			fmt.Fprintln(w, lineStr)
		}
		fmt.Fprintln(w)
		w.Flush()

	}

}
