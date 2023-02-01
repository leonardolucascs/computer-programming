/*

    Topic: STRING PADDING IN GO
    https://gosamples.dev/string-padding/



    The Go standard library also includes a useful text/tabwriter package that creates 
    properly aligned tables from strings separated by \t characters. This type of formatting 
    is particularly useful for creating command-line applications that print tabular data. 
    In the example below, we initialize tabwriter.Writer by setting output to the standard 
    output os.Stdout, \t as the padding character and 4 as the width of the tab character. 
    For printing, the line cells should be concatenated into a single string separated by tabs. 
    We do this using the strings.Join() function. The tabwriter.Writer requires every column 
    cell to be tab-terminated, so we also add the \t for the last cell in the row.

*/

package main

import (
    "fmt"
    "os"
    "strings"
    "text/tabwriter"
)

func printTable(table [][]string) {
    writer := tabwriter.NewWriter(os.Stdout, 0, 4, 0, '\t', 0)
    for _, line := range table {
        fmt.Fprintln(writer, strings.Join(line, "\t")+"\t")
    }
    writer.Flush()
}

func main() {
    var table = [][]string{
        {"vegetables", "fruits", "rank"},
        {"potato", "strawberry", "1"},
        {"lettuce", "raspberry", "2"},
        {"carrot", "apple", "3"},
        {"broccoli", "pomegranate", "4"},
    }
    printTable(table)
}
