package Frontend

import (
	"awesomeProject/srcs/Backend/JsonStruct"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

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

		if value == "" {
			_ = tmpl.Execute(w, "")
		} else if ok != nil {
			_ = tmpl.Execute(w, "Не корректное значение!")
		} else if id >= len(js.GetSlice()) || id < 0 {
			_ = tmpl.Execute(w, "Значение выходит за пределы допустимого!")
		} else {
			marshal, _ := json.Marshal(js.GetSlice()[id])
			b := new(bytes.Buffer)
			_ = json.Indent(b, marshal, "", "\t")
			_ = tmpl.Execute(w, b.String())
		}
	})

	ok := http.ListenAndServe(":8080", nil)
	if ok != nil {
		log.Panic(ok)
	}
}
