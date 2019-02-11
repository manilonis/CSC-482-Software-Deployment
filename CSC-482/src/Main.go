package main

import "fmt"
import "os"
import "os/exec"
import "encoding/json"
import "strings"
//import "reflect"

type arrival struct{
		time int
}
type stop_time_update struct{
		stop_time []arrival
}
type trip_update struct{
		trip []stop_time_update
}
type entity struct{
		up []trip_update
}
type head struct{
	ent entity  `json:"entity"`
}

func main(){

	cmd := exec.Command("./try.sh")
	out, err := cmd.Output()

	if err != nil {
    println(err.Error())
    return
}
    o := string(out)

    //fmt.Println(o)

    f, err := os.Create("stuff.json")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString(o)
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }

    store := head{}
    err = json.Unmarshal([]byte(o), &store)

    fmt.Println(store)

    for i:=0; i < len(store.ent.up); i++{
 		fmt.Println(i)
    }

    var result map[string]interface{}
	json.Unmarshal([]byte(o), &result)

	keys := make([]string, 0)
  	for key := range result{
    keys = append(keys, key)
  	}

  	fmt.Println(strings.Join(keys, ","))

  	ent := result["entity"]

  	m, ok := ent.([]interface{})
  	fmt.Println(ok)

  	m1, ok1 := m[0].(map[string]interface{})
  	fmt.Println(ok1)


  	//fmt.Println(reflect.TypeOf(ent))
  	//fmt.Println(reflect.TypeOf(fmt.Sprintf("%v", ent)))
  	//entstr := "{" + fmt.Sprintf("%v", ent) + "}"

  	//var ent1 map[string]interface{}
  	//json.Unmarshal([]byte(entstr),&ent1)

  	more_keys := make([]string, 0)
  	for more_key := range m1{
  		more_keys = append(more_keys, more_key)
  	}

  	fmt.Println(strings.Join(more_keys, ","))


}