package Frontend

import (
	"awesomeProject/srcs/Backend/JsonStruct"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type InfoModels struct {
	Model  string
	Length string
}

func ModelToStriong(js *JsonStruct.JsonSlice, id int) string {
	buf := bytes.Buffer{}
	marshal, _ := json.Marshal(js.GetSlice()[id])
	_ = json.Indent(&buf, marshal, "", "\t")
	return buf.String()
}

func Handler(js *JsonStruct.JsonSlice) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("srcs/Frontend/html/index.html")
		if err != nil {
			log.Panic(err)
			return
		}
		//_ = tmpl.Execute(w, len(js.GetSlice()))

		value := r.FormValue("input_id")
		id, ok := strconv.Atoi(value)

		info := InfoModels{"", fmt.Sprint(len(js.GetSlice()))}
		if value == "" {
			info.Model = ""
		} else if ok != nil {
			info.Model = "Не корректное значение!"
		} else if id >= len(js.GetSlice()) || id < 0 {
			info.Model = "Значение выходит за пределы допустимого!"
		} else {
			info.Model = ModelToStriong(js, id)
		}
		_ = tmpl.Execute(w, info)
	})

	ok := http.ListenAndServe(":8080", nil)
	if ok != nil {
		log.Panic(ok)
	}
}
