package main

import (
	"errors"
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
	"strings"
)




func main() {

	// result:=ReadString("metadata.name","test.yaml")
	result:=ReadString("spec.template.spec.containers[0].name","test.yaml")
	fmt.Println(result)

	result2:=ReadInt("spec.replica","test.yaml")
	fmt.Println(result2)

}

func ReadString(clause string,path string) string {
	m:=marshal(path)
	seq:=syntaxParser(clause)
	result,_:=readString(seq,m)

	return result
}

func ReadInt(clause string,path string) int {
	m:=marshal(path)
	seq:=syntaxParser(clause)
	result,_:=readInt(seq,m)

	return result
}

func marshal(path string) map[interface{}]interface{}{
	content:=loadFileContent(path)
	m := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(content), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil
	}
	return m
}

func readInt(seq []string,data map[interface{}]interface{}) (int,error){
	seqResult:=readValue(seq,data)
	if seqResult!=nil {
		if s,ok:=seqResult.(int);ok{
			return s,nil
		}
	}
	return 0,errors.New("wrong type")
}

func readString(seq []string,data map[interface{}]interface{}) (string,error){
	seqResult:=readValue(seq,data)
	if seqResult!=nil {
		if s,ok:=seqResult.(string);ok{

			return s,nil
		}
	}

	return "",errors.New("wrong type")
}

//metadata.name
func readValue(seq []string,data map[interface{}]interface{}) interface{}{
    var seqResult interface{}
    seqResult=data
	for _,k :=range seq{

		// array conv
		if strings.Contains(k,"[") && strings.Contains(k,"]"){
			sub:=strings.Trim(k,"[]")
			value,err:=strconv.ParseInt(sub,10,64)
			if err != nil {
				fmt.Println(err.Error())
				return nil
			}
			if arr,ok:=seqResult.([]interface{});ok {
				seqResult=arr[value]
				continue
			}else{
				fmt.Println("wrong index path")
				return nil
			}
		}
		if mii,ok:=seqResult.(map[interface{}]interface{});ok{
			if _,exist:=mii[k];exist {
				seqResult = mii[k]
				continue
			}else{
				log.Fatalf("bad path")
			}

		}
		if msi,ok:=seqResult.(map[string]interface{});ok{
			if _,exist:=msi[k];exist{
				seqResult=msi[k]
				continue
			}else{
				log.Fatalf("bad path")
			}

		}


	}
	return seqResult
}

// parse metadata.name => ["metadata","name"]
func syntaxParser(clause string) []string{
	seq:=strings.Split(clause,".")
	resultSeq:=make([]string,0)
	for _,k:=range seq {

		// containers
		if strings.Contains(k,"[") && strings.Contains(k,"]"){
			idx:=strings.LastIndex(k,"[")
			object:=k[:idx]
			index:=k[idx:]
			fmt.Println("object:",object,"index:",index)
			resultSeq=append(resultSeq,object,index)
		}else{
			resultSeq=append(resultSeq,k)
		}
	}
	return resultSeq
}

func loadFileContent(path string) string{

	if _,err:=os.Stat(path);err!=nil{
		fmt.Println(err.Error())
	}

	bytes,err:=os.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(bytes)
}