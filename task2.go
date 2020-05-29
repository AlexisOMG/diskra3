package main

import "fmt"

var d [][]int;
var f [][]string;
var used, ts []int;
var id map[int]string;

func main() {
    var n, m, start, a int;
    fmt.Scan(&n, &m, &start);
    d = make([][]int, n);
    f = make([][]string, n);
    id = make(map[int]string);
    for i := 0; i < n; i++ {
        used = append(used, 0);
        for j := 0; j < m; j++ {
            fmt.Scan(&a);
            d[i] = append(d[i], a);
        }
    }
    var b string;
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            fmt.Scan(&b);
            f[i] = append(f[i], b);
        }
    }
    for i := 0; i < m; i++ {
        id[i] = string('a' + i);
    }
    fmt.Print("digraph {\n\trankdir = LR\n\tdummy [label = \"\", shape = none]\n");
    for i := 0; i < n; i++ {
        fmt.Print("\t", i, " [shape = circle]\n");
    }
    fmt.Print("\tdummy -> ", start, "\n");
    for ind, el := range d[start] {
        fmt.Print("\t", start, " -> ", el, " [label = \"", id[ind], "(", f[start][ind], ")\"]\n");
    }
    for ind, xs := range d {
        if ind == start {
            continue;
        }
        for i, e := range xs {
            fmt.Print("\t", ind, " -> ", e, " [label = \"", id[i], "(", f[ind][i], ")\"]\n");
        }
    }
    fmt.Print("\n}\n");
}

