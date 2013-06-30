package main

import (
    // Standard imports
    "fmt"
    "io"
    "os"
    "strings"
    "text/template"

    // Imports the library defined in the containing package
    "github.com/munckymagik/gosort"
)

type sortAlgorithm struct {
    // The name that the user can specify on the command line
    // to select this algorithm
    Name string

    // The algorithm to apply to the input
    Run gosort.IntegerSorter

    // The usage description
    Description string
}

var algorithmsList = []sortAlgorithm{
    {"quicksort", gosort.QuickSort, "Sort using the Quick Sort algorithm"},
    {"mergesort", gosort.MergeSort, "Sort using the Merge Sort algorithm"},
    {"insertionsort", gosort.InsertionSort, "Sort using the Insetion Sort algorithm"},
    {"bubblesort", gosort.BubbleSort, "Sort using the Bubble Sort algorithm"},
    {"heapsort", gosort.HeapSort, "Sort using the Heap Sort algorithm"},
}

var algorithmsMap = make(map[string]sortAlgorithm)

var usageTemplate = `gosort is a test front-end for common sorting algorithms.

USAGE gosort <algorithm>

gosort reads integers from STDIN, one per line, then sorts and outputs to
STDOUT. It expects one positional parameter, <algorithm>, used to specify
which algorithm to apply to the input.

Available algorithms:
{{range .}}
    {{.Name | printf "%-14s"}} {{.Description}}{{end}}
    help           Don't sort, show this help message

`

func init() {
    for _, choice := range algorithmsList {
        algorithmsMap[choice.Name] = choice
    }
}

func usage(){
    t := template.Must(template.New("usage").Parse(usageTemplate))
    t.Execute(os.Stderr, algorithmsList)
    os.Exit(2)
}

func main() {
    if len(os.Args) != 2 {
        usage()
    }

    requestedAlgorithm := os.Args[1]
    requestedAlgorithm = strings.ToLower(requestedAlgorithm)

    if strings.HasSuffix(requestedAlgorithm, "help") {
        usage()
    }

    if theAlgorithm, ok := algorithmsMap[requestedAlgorithm]; ok {
        fmt.Fprintln(os.Stderr, "Chosen algorithm", theAlgorithm.Name)

        buffer := make([]int, 0)
        var itemBuffer int

        itemsRead, err := fmt.Scanln(&itemBuffer)
        for itemsRead == 1 && err == nil {
            buffer = append(buffer, itemBuffer)
            itemsRead, err = fmt.Scanln(&itemBuffer)
        }

        if err != io.EOF {
            fmt.Println("ERROR reading input:", err)
            os.Exit(1)
        }

        theAlgorithm.Run(buffer)

        for _, element := range buffer {
            fmt.Println(element)
        }

    } else {
        fmt.Println("ERROR Unrecognised algorithm/command:", requestedAlgorithm)
        fmt.Println("Run gosort help to see the list of supported choices")
        os.Exit(1)
    }

}
