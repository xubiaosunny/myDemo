package main
import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "io"
    "strconv"
    // "reflect"
    "time"
    "github.com/xubiaosunny/myDemo/sorter/algorithms/bubblesort"
    "github.com/xubiaosunny/myDemo/sorter/algorithms/qsort"
)

var infile  = flag.String("i","infile","File contains values for sorting")
var outfile = flag.String("o","outfile","File to receive sorted values")
var algorithm = flag.String("a","qsort","sort algorithm")

func readValues(infile string) (values [] int, err error) {
    f, err := os.Open(infile)
    if err != nil {
        fmt.Println("Failed open input file",infile)
        return
    }
    defer f.Close()

    br := bufio.NewReader(f)
    values = make([] int,0)
    for {
        line, isPrefix, err1 := br.ReadLine()
        if err1 != nil {
            if err1 != io.EOF {
                err = err1
            }
            break
        }
        // fmt.Println(string(line),isPrefix,err1)
        if isPrefix {
            fmt.Println("A too lang line, seems unexpexted")
        }
        value, err1 := strconv.Atoi(string(line))
        if err1 != nil {
            err = err1
            return
        }
        values = append(values,value)
    }
    return
}

func writeValues(values [] int, outfile string) error {
    f, err := os.Create(outfile)
    if err != nil {
        fmt.Println("Failed create outfile:", outfile)
        return err
    }
    defer f.Close()
    for _, value := range values {
        // println(reflect.TypeOf(value))
        str := strconv.Itoa(value)
        f.WriteString(str + "\n")
    }
    return nil
}
func main() {
    fmt.Println(time.Now().Format("2006.01.02 15:04:05"))
    flag.Parse()
    if infile != nil {
        fmt.Println("infile=",*infile,"outfile=",*outfile,"algorithm=",*algorithm)
    }
    values, err := readValues(*infile)
    if err == nil {
        // fmt.Println("Read values:",values)
        // fmt.Println(reflect.TypeOf(values))
        // writeValues(values,*outfile)
        t1 := time.Now()

        switch *algorithm {
        case "qsort":
            qsort.QuickSort(values)
        case "bubblesort":
            bubblesort.BubbleSort(values)
        default:
            fmt.Println("sorting algorithms ",*algorithm," is unknown or unsupported") 
        }

        t2 := time.Now()
        fmt.Println("the sorting process costs ",t2.Sub(t1))
        writeValues(values,*outfile)
    } else{
        fmt.Println(err)
    }
}
