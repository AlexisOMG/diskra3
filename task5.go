package main

import "fmt"

type cond struct {
    first int;
    second string;
}

type mili struct{
    alphIn, alphOut []string;
    d, mooreD [][]int;
    f [][]string;
    cntCond, cntSign int;
    mooreCond map[cond]int;
    reversed map[int]cond;

}

func (obj *mili) read() {
    obj.mooreCond = make(map[cond]int);
    obj.reversed = make(map[int]cond);
    fmt.Scan(&obj.cntSign);
    obj.alphIn = make([]string, obj.cntSign);
    for i := 0; i < obj.cntSign; i++ {
        fmt.Scan(&obj.alphIn[i]);
    }
    var n int;
    fmt.Scan(&n);
    obj.alphOut = make([]string, n);
    for i := 0; i < n; i++ {
        fmt.Scan(&obj.alphOut[i]);
    }
    fmt.Scan(&obj.cntCond);
    obj.d = make([][]int, obj.cntCond);
    obj.f = make([][]string, obj.cntCond);
    //obj.id = make(map[int]string);
    var a int;
    for i := 0; i < obj.cntCond; i++ {
        for j := 0; j < obj.cntSign; j++ {
            fmt.Scan(&a);
            obj.d[i] = append(obj.d[i], a);
        }
    }
    var b string;
    for i := 0; i < obj.cntCond; i++ {
        for j := 0; j < obj.cntSign; j++ {
            fmt.Scan(&b);
            obj.f[i] = append(obj.f[i], b);
        }
    }
}

func (obj *mili) moore() {
    t := 0;
    for i := 0; i < len(obj.d); i++ {
        for j := 0; j < obj.cntSign; j++ {
            vertex := cond{first: obj.d[i][j], second: obj.f[i][j]};
            _, ok := obj.mooreCond[vertex];
            if ok == false {
                obj.mooreCond[vertex] = t;
                obj.reversed[t] = vertex;
                t += 1;
            }
        }
    }
    obj.mooreD = make([][]int, len(obj.mooreCond));
    for key, val := range obj.mooreCond {
        for i := 0; i < obj.cntSign; i++ {
            obj.mooreD[val] = append(obj.mooreD[val], obj.mooreCond[cond{first: obj.d[key.first][i], second: obj.f[key.first][i]}]);
        }
    }
}

func (obj *mili) print()  {
    obj.moore();
    // fmt.Println(obj.mooreCond);
    // fmt.Println(obj.mooreD);
    fmt.Print("digraph {\n\trankdir = LR\n");
    for i := 0; i < len(obj.mooreCond); i++ {
        fmt.Print("\t", i, " [label = \"(", obj.reversed[i].first, ",", obj.reversed[i].second, ")\"]\n");
        for ind, e := range obj.mooreD[i] {
            fmt.Print("\t", i, " -> ", e, "[label = \"", obj.alphIn[ind], "\"]\n")
        }
    }
    fmt.Println("}");
}

func main() {
    var a mili;
    a.read();
    a.print();
}

