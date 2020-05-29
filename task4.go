package main

import (
    "fmt"
    "os"
)

type cond struct{
    i int;
    p* cond;
    depth int;
}

type mili struct{
    d [][]int;
    f [][]string;
    used, ts []int;
    id map[int]string;
    parents, q []cond;
    start, n, z int;
}

func (obj *mili)dfs(v int)  {
    if obj.used[v] == 1 {
        return;
    }
    obj.used[v] = 1;
    for _, el := range obj.d[v] {
        if obj.used[el] == 0 {
            obj.ts = append(obj.ts, el);
            obj.dfs(el);
        }
    }
}

func (obj *mili) find(v *cond) *cond {
    if *v == *v.p {
        return v;
    }
    v.p = obj.find(v.p);
    return v.p;
}

func (obj *mili) union(x, y *cond)  {
    rX := obj.find(x);
    rY := obj.find(y);
    if rX.depth < rY.depth {
        rX.p = rY;
    } else {
        rY.p = rX;
        if rX.depth == rY.depth && rX != rY {
            rX.depth += 1;
        }
    }
}

func (obj *mili) split1() int {
    m := len(obj.d);
    obj.q = make([]cond, len(obj.d));
    for i := 0; i < len(obj.d); i++ {
        obj.q[i].p = &obj.q[i];
        obj.q[i].depth = 0;
        obj.q[i].i = i;
    }
    for i := 0; i < len(obj.q); i++ {
        for j := 0; j < len(obj.q); j++ {
            if *obj.find(&obj.q[i]) != *obj.find(&obj.q[j]) {
                eq := true;
                for t := 0; t < obj.z; t++ {
                    if obj.f[obj.q[i].i][t] != obj.f[obj.q[j].i][t] {
                        eq = false;
                        break;
                    }
                }
                if eq == true {
                    obj.union(&obj.q[i], &obj.q[j]);
                    m -= 1;
                }
            }
        }
    }
    for _, el := range obj.q {
        obj.parents[el.i] = *obj.find(&el);
    }
    return m;
}

func (obj *mili) split() int {
    m := len(obj.d);
    obj.q = make([]cond, len(obj.d));
    for i := 0; i < len(obj.d); i++ {
        obj.q[i].p = &obj.q[i];
        obj.q[i].depth = 0;
        obj.q[i].i = i;
    }
    for i := 0; i < len(obj.q); i++ {
        for j := 0; j < len(obj.q); j++ {
            if obj.parents[obj.q[i].i] == obj.parents[obj.q[j].i] && *obj.find(&obj.q[i]) != *obj.find(&obj.q[j]) {
                eq := true;
                for t := 0; t < obj.z; t++ {
                    w1 := obj.d[obj.q[i].i][t];
                    w2 := obj.d[obj.q[j].i][t];
                    if obj.parents[w1] != obj.parents[w2] {
                        eq = false;
                        break;
                    }
                }
                if eq == true {
                    obj.union(&obj.q[i], &obj.q[j]);
                    m -= 1;
                }
            }
        }
    }
    for _, el := range obj.q {
        obj.parents[el.i] = *obj.find(&el);
    }
    return m;
}

func (obj *mili) AufenkampHohn()  (map[cond]int, [][]cond, [][]string) {
    m := obj.split1();
    var m1 int;
    for {
        m1 = obj.split();
        if m == m1 {
            break;
        }
        m = m1;
    }
    Q1 := make(map[cond]int);
    var D1 [][]cond;
    var F1 [][]string;
    for _, el := range obj.q {
        q1 := obj.parents[el.i];
        _, check := Q1[q1];
        if check == false {
            D1 = append(D1, make([]cond, obj.z));
            F1 = append(F1, make([]string, obj.z));
            Q1[q1] = len(D1) - 1;
            for i := 0; i < obj.z; i++ {
                D1[len(D1) - 1][i] = obj.parents[obj.d[el.i][i]];
                F1[len(F1) - 1][i] = obj.f[el.i][i];
            }
        }
    }
    return Q1, D1, F1;
}

func (obj *mili) minimization() {
    fmt.Scan(&obj.n, &obj.z, &obj.start);
    obj.d = make([][]int, obj.n);
    obj.f = make([][]string, obj.n);
    obj.id = make(map[int]string);
    obj.parents = make([]cond, obj.n);
    obj.q = make([]cond, obj.n);
    var a int;
    for i := 0; i < obj.n; i++ {
        obj.used = append(obj.used, 0);
        for j := 0; j < obj.z; j++ {
            fmt.Scan(&a);
            obj.d[i] = append(obj.d[i], a);
        }
    }
    var b string;
    for i := 0; i < obj.n; i++ {
        for j := 0; j < obj.z; j++ {
            fmt.Scan(&b);
            obj.f[i] = append(obj.f[i], b);
        }
    }
    for i := 0; i < obj.z; i++ {
        obj.id[i] = string('a' + i);
    }
    A, B, C := obj.AufenkampHohn();
    obj.d = make([][]int, len(B));
    obj.f = make([][]string, len(B));
    for ind, xs := range B {
        obj.d[ind] = make([]int, obj.z);
        obj.f[ind] = make([]string, obj.z);
        for i, e := range xs {
            obj.d[ind][i] = A[e];
            obj.f[ind][i] = C[ind][i];
        }
    }
    obj.start = A[obj.parents[obj.start]];
    obj.dfs(obj.start);
    obj.ts = append([]int{obj.start}, obj.ts...);
    ID := make(map[int]int);
    for ind, el := range obj.ts {
        ID[el] = ind;
    }
    d1 := make([][]int, len(obj.ts));
    f1 := make([][]string, len(obj.ts));
    for i, el := range obj.ts {
        d1[i] = make([]int, obj.z);
        for j, e := range obj.d[el] {
            d1[i][j] = ID[e];
        }
    }
    for i, el := range obj.ts {
        f1[i] = make([]string, obj.z);
        for j, e := range obj.f[el] {
            f1[i][j] = e;
        }
    }
    obj.d = d1;
    obj.f = f1;
}

func main() {
    var a, b mili;
    a.minimization();
    b.minimization();
    if len(a.d) != len(b.d) || len(a.f) != len(b.f) || a.z != b.z {
        fmt.Println("NOT EQUAL");
    } else {
        for i := 0; i < len(a.d); i++ {
            for j := 0; j < a.z; j++ {
                if a.d[i][j] != b.d[i][j] || a.f[i][j] != b.f[i][j] {
                    fmt.Println("NOT EQUAL");
                    os.Exit(0);
                }
            }
        }
        fmt.Println("EQUAL");
    }

}

