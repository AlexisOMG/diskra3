package main

import "fmt"

var d [][]int;
var f [][]string;
var used, ts []int;
var id map[int]int;

func dfs(v int)  {
    if used[v] == 1 {
        return;
    }
    used[v] = 1;
    for _, el := range d[v] {
        if used[el] == 0 {
            ts = append(ts, el);
            dfs(el);
        }
    }
}

func main() {
    var n, m, start, a int;
    fmt.Scan(&n, &m, &start);
    d = make([][]int, n);
    f = make([][]string, n);
    id = make(map[int]int);
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
    dfs(start);
    ts = append([]int{start}, ts...);
    for ind, el := range ts {
        id[el] = ind;
    }
    fmt.Print(len(ts), "\n", m, "\n0\n");
    for _, el := range ts {
        for _, e := range d[el] {
            fmt.Print(id[e], " ");
        }
        fmt.Print("\n");
    }
    for _, el := range ts {
        for _, x := range f[el] {
            fmt.Print(x, " ");
        }
        fmt.Print("\n");
    }
}

