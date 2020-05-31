package main

import (
    "fmt"
    "os"
    "sort"
)

type cond struct {
    target int;
    sign string;
}

type stack struct {
    arr [][]int;
}

func (obj *stack) push(x []int) {
    obj.arr = append(obj.arr, x);
}

func (obj *stack) pop() []int {
    if len(obj.arr) == 0 {
        fmt.Println("OMG This is The ENDD!!");
        os.Exit(0);
    }
    res := obj.arr[len(obj.arr) - 1];
    obj.arr = obj.arr[:len(obj.arr) - 1];
    return res;
}

func (obj *stack) empty() bool {
    return len(obj.arr) == 0;
}

type mili struct{
    n, m, start int;
    d [][]cond;
    final, tmp []int;
    X map[string]int;
}

func (obj *mili) contain(x int) bool {
    for _, el := range obj.tmp {
        if x == el {
            return true;
        }
    }
    return false;
}

func (obj *mili) read() {
    fmt.Scan(&obj.n, &obj.m);
    obj.X = make(map[string]int);
    obj.d = make([][]cond, obj.n);
    obj.final = make([]int, obj.n);
    var a, b int;
    var s string;
    t := 0;
    for i := 0; i < obj.m; i++ {
        fmt.Scan(&a, &b, &s);
        if s != "lambda" {
            _, ok := obj.X[s];
            if ok == false {
                obj.X[s] = t;
                t += 1;
            }
        }
        obj.d[a] = append(obj.d[a], cond{target: b, sign: s});
    }
    for i := 0; i < obj.n; i++ {
        fmt.Scan(&obj.final[i]);
    }
    fmt.Scan(&obj.start);
}

func (obj *mili) closure(z []int) []int {
    obj.tmp = []int{};
    for _, el := range z {
        obj.dfs(el);
    }
    return obj.tmp;
}

func (obj *mili) extract(q int, s string) (res[]int) {
    for _, el := range obj.d[q] {
        if el.sign == s {
            res = append(res, el.target);
        }
    }
    return;
}

func (obj *mili) dfs(q int)  {
    if obj.contain(q) == false {
        obj.tmp = append(obj.tmp, q);
        for _, el := range obj.extract(q, "lambda") {
            obj.dfs(el);
        }
    }
}

func check(xs [][]int, x []int) (int, bool) {
    res := true;
    for ind, elements := range xs {
        if len(x) != len(elements) {
            continue;
        }
        for i, el := range elements {
            if el != x[i] {
                res = false;
                break;
            }
        }
        if res == true {
            return ind, true;
        }
        res = true;
    }
    return -1, false;
}

func (obj *mili) det() ([][]int, []map[int]map[string]bool, []int) {
    q0 := obj.closure([]int{obj.start});
    sort.Ints(q0);
    Q := [][]int{q0};
    var D []map[int]map[string]bool;
    var F []int;
    var s stack;
    s.push(q0);
    for s.empty() == false {
        z := s.pop();
        indCond, _ := check(Q, z);
        for _, el := range z {
            if obj.final[el] == 1 {
                F = append(F, indCond);
                break;
            }
        }
        for len(D) < indCond + 1 {
            D = append(D, make(map[int]map[string]bool));
        }
        for key,_ := range obj.X {
            var z1 []int;
            for _, el := range z {
                z1 = append(z1, obj.extract(el, key)...);
            }
            z1 = obj.closure(z1);
            sort.Ints(z1);
            ind, ok := check(Q, z1);
            if ok == false {
                Q = append(Q, z1);
                ind = len(Q) - 1;
                s.push(z1);
            }
            _, ok = D[indCond][ind];
            if ok == false {
                D[indCond][ind] = make(map[string]bool);
            }
            D[indCond][ind][key] = true;
        }
    }
    return Q, D, F;
}

func (obj *mili) print()  {
    Q, D, F := obj.det();
    fmt.Print("digraph {\n\trankdir = LR\n\tdummy [label = \"\", shape = none]\n");
    for i := 0; i < len(Q); i++ {
        fmt.Print("\t", i, " [label = \"[");
        for j := 0; j < len(Q[i]) - 1; j++ {
            fmt.Print(Q[i][j], " ");
        }
        if len(Q[i]) > 0 {
            fmt.Print(Q[i][len(Q[i]) - 1]);
        }
        fmt.Print("]\", shape = ");
        ch := false;
        for _, el := range F {
            if el == i {
                ch = true;
            }
        }
        if ch == true {
            fmt.Print("doublecircle]\n");
        } else {
            fmt.Print("circle]\n");
        }
    }
    fmt.Print("\tdummy -> 0\n");
    for i := 0; i < len(D); i++ {
        for key, value := range D[i] {
            fmt.Print("\t", i, " -> ", key, " [label = \"");
            var labels []string;
            for k, _ := range value {
                labels = append(labels, k);
            }
            sort.Slice(labels, func(i, j int) bool {return obj.X[labels[i]] < obj.X[labels[j]]});
            for j := 0; j < len(labels) - 1; j++ {
                fmt.Print(labels[j], ", ");
            }
            fmt.Print(labels[len(labels) - 1], "\"]\n");
        }
    }
    fmt.Print("}\n");
}

func main() {
    var a mili;
    a.read();
    a.print();
}

