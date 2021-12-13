package main

 

import (

    "encoding/json"

    "fmt"

    "io/ioutil"

    "log"

    "net/http"

    "sort"

    "strconv"

 

    "github.com/gorilla/mux"

)

 

type Fquester struct {

    Name string `json:"name"`

    Qty  int    `json:"qty"`

}

 

func displayDetails(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)

    minQty, _ := strconv.Atoi(vars["quantity"])

    flag := false

    var allItems []Fquester

    fts, _ := http.Get(https://f8776af4-e760-4c93-97b8-70015f0e00b3.mock.pstmn.io/fruits)

    vgs, _ := http.Get(https://f8776af4-e760-4c93-97b8-70015f0e00b3.mock.pstmn.io/vegetables)

    grs, _ := http.Get(https://f8776af4-e760-4c93-97b8-70015f0e00b3.mock.pstmn.io/grains)

    response1, _ := ioutil.ReadAll(fts.Body)

    var fruit []Fquester

    json.Unmarshal(response1, &fruit)

    for i := 0; i < len(fruit); i++ {

        allItems = append(allItems, fruit[i])

    }

    response2, _ := ioutil.ReadAll(vgs.Body)

    var vegetable []Fquester

    json.Unmarshal(response2, &vegetable)

    for i := 0; i < len(vegetable); i++ {

        allItems = append(allItems, vegetable[i])

    }

    response3, _ := ioutil.ReadAll(grs.Body)

    var grain []Fquester

    json.Unmarshal(response3, &grain)

    for i := 0; i < len(grain); i++ {

        allItems = append(allItems, grain[i])

    }

    sort.Slice(allItems, func(i, j int) bool {

        return allItems[i].Name < allItems[j].Name

    })

    for _, item := range allItems {

        if item.Qty <= minQty {

            flag = true

            fmt.Println(w, item)

        }

    }

    if flag == false {

        fmt.Fprintln(w, "Not_Found")

    }

}

func handleRequests() {

    myRouter := mux.NewRouter()

    myRouter.HandleFunc("/quest/{quantity}", displayDetails).Methods(http.MethodGet)

    log.Fatal(http.ListenAndServe(":8086", myRouter))

}

func main() {

    handleRequests()

 

}