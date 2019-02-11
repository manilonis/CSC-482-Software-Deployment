package main

import "fmt"
import "os"
import "os/exec"
import "encoding/json"
import "strings"

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
  	fmt.Println(len(m))

  	m1, ok1 := m[400].(map[string]interface{})
  	fmt.Println(ok1)


  	more_keys := make([]string, 0)
  	for more_key := range m1{
  		more_keys = append(more_keys, more_key)
  	}

  	fmt.Println(strings.Join(more_keys, ","))


}