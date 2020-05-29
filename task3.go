package main

import "fmt"

type cond struct{
    i int;
    p* cond;
    depth int;
}

var d [][]int;
var f [][]string;
var used, ts []int;
var id map[int]string;
var parents, q []cond;
var z int;

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

func find(v *cond) *cond {
    if *v == *v.p {
        return v;
    }
    v.p = find(v.p);
    return v.p;
}

func union(x, y *cond)  {
    rX := find(x);
    rY := find(y);
    if rX.depth < rY.depth {
        rX.p = rY;
    } else {
        rY.p = rX;
        if rX.depth == rY.depth && rX != rY {
            rX.depth += 1;
        }
    }
}

func split1() int {
    m := len(d);
    q = make([]cond, len(d));
    for i := 0; i < len(d); i++ {
        q[i].p = &q[i];
        q[i].depth = 0;
        q[i].i = i;
    }
    for i := 0; i < len(q); i++ {
        for j := 0; j < len(q); j++ {
            if *find(&q[i]) != *find(&q[j]) {
                eq := true;
                for t := 0; t < z; t++ {
                    if f[q[i].i][t] != f[q[j].i][t] {
                        eq = false;
                        break;
                    }
                }
                if eq == true {
                    union(&q[i], &q[j]);
                    m -= 1;
                }
            }
        }
    }
    for _, el := range q {
        parents[el.i] = *find(&el);
    }
    return m;
}

func split() int {
    m := len(d);
    q = make([]cond, len(d));
    for i := 0; i < len(d); i++ {
        q[i].p = &q[i];
        q[i].depth = 0;
        q[i].i = i;
    }
    for i := 0; i < len(q); i++ {
        for j := 0; j < len(q); j++ {
            if parents[q[i].i] == parents[q[j].i] && *find(&q[i]) != *find(&q[j]) {
                eq := true;
                for t := 0; t < z; t++ {
                    w1 := d[q[i].i][t];
                    w2 := d[q[j].i][t];
                    if parents[w1] != parents[w2] {
                        eq = false;
                        break;
                    }
                }
                if eq == true {
                    union(&q[i], &q[j]);
                    m -= 1;
                }
            }
        }
    }
    for _, el := range q {
        parents[el.i] = *find(&el);
    }
    return m;
}

func AufenkampHohn()  (map[cond]int, [][]cond, [][]string) {
    m := split1();
    var m1 int;
    for {
        m1 = split();
        if m == m1 {
            break;
        }
        m = m1;
    }
    Q1 := make(map[cond]int);
    var D1 [][]cond;
    var F1 [][]string;
    for _, el := range q {
        q1 := parents[el.i];
        _, check := Q1[q1];
        if check == false {
            D1 = append(D1, make([]cond, z));
            F1 = append(F1, make([]string, z));
            Q1[q1] = len(D1) - 1;
            for i := 0; i < z; i++ {
                D1[len(D1) - 1][i] = parents[d[el.i][i]];
                F1[len(F1) - 1][i] = f[el.i][i];
            }
        }
    }
    return Q1, D1, F1;
}

func main() {
    var start, n int;
    fmt.Scan(&n, &z, &start);
    d = make([][]int, n);
    f = make([][]string, n);
    id = make(map[int]string);
    parents = make([]cond, n);
    q = make([]cond, n);
    var a int;
    for i := 0; i < n; i++ {
        used = append(used, 0);
        for j := 0; j < z; j++ {
            fmt.Scan(&a);
            d[i] = append(d[i], a);
        }
    }
    var b string;
    for i := 0; i < n; i++ {
        for j := 0; j < z; j++ {
            fmt.Scan(&b);
            f[i] = append(f[i], b);
        }
    }
    for i := 0; i < z; i++ {
        id[i] = string('a' + i);
    }
    A, B, C := AufenkampHohn();
    fmt.Print("digraph {\n\trankdir = LR\n\tdummy [label = \"\", shape = none]\n");
    d = make([][]int, len(B));
    f = make([][]string, len(B));
    for ind, xs := range B {
        d[ind] = make([]int, z);
        f[ind] = make([]string, z);
        for i, e := range xs {
            d[ind][i] = A[e];
            f[ind][i] = C[ind][i];
        }
    }
    start = A[parents[start]];
    dfs(start);
    ts = append([]int{start}, ts...);
    ID := make(map[int]int);
    for ind, el := range ts {
        ID[el] = ind;
    }
    d1 := make([][]int, len(ts));
    f1 := make([][]string, len(ts));
    for i, el := range ts {
        d1[i] = make([]int, z);
        for j, e := range d[el] {
            d1[i][j] = ID[e];
        }
    }
    for i, el := range ts {
        f1[i] = make([]string, z);
        for j, e := range f[el] {
            f1[i][j] = e;
        }
    }
    d = d1;
    f = f1;
    for i := 0; i < len(d); i++ {
        fmt.Print("\t", i, " [shape = circle]\n");
    }
    fmt.Print("\tdummy -> 0\n");
    for ind, xs := range d {
        for i, e := range xs {
            fmt.Print("\t", ind, " -> ", e, " [label = \"", id[i], "(", f[ind][i], ")\"]\n");
        }
    }
    fmt.Print("\n}\n");
}

