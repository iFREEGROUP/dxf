package main

import (
	"encoding/json"
	"github.com/iFREEGROUP/dxf"
	"log"
	"os"
)

func main() {
	dxfPath := "/Users/kingzcheung/go/src/autocad/BGS8.dxf"
	drawing, err := dxf.FromFile(dxfPath)
	if err != nil {
		//fmt.Println(err)
		return
	}

	marshal, err := json.Marshal(drawing)
	if err != nil {
		log.Fatalln(err)
		return
	}
	file, err := os.OpenFile("test.json", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln(err)
		return
	}
	_, _ = file.Write(marshal)

}
