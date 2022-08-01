package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// CSV : csv
type CSV struct { //estrutura que receberá os dados do CSV
	Id         string `json:"id"`
	Tipo       string `json:"tipo"`
	Modelo     string `json:"modelo"`
	Observacao string `json:"observacao"`
}

func checkErr(err error) { //checa erros
	if err != nil {
		log.Panic("ERROR: " + err.Error())
	}
}

func WriteCSV(arquivocsv string, dados string) (string, error) {

	csvFile, err := os.OpenFile(arquivocsv, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm) //abre arquivo

	checkErr(err)

	//csvFile.Write([]byte(dados))
	bytesWrited, err := csvFile.WriteString(dados)
	checkErr(err)

	csvFile.Close()

	fmt.Printf("%d bytes gravados", bytesWrited)
	return fmt.Sprintf("%d bytes gravados", bytesWrited), err

}

func ReadCSV(arquivocsv string) ([]byte, error) {

	csvFile, err := os.Open(arquivocsv) //abre arquivo
	checkErr(err)

	reader := csv.NewReader(bufio.NewReader(csvFile)) //lê arquivo
	reader.Comma = ';'                                //define delimitador

	var patrs []CSV

	for {
		line, err := reader.Read() //para cada linha
		if err == io.EOF {
			break
		} else if err != nil {
			checkErr(err)
		}
		patrs = append(patrs, CSV{ //adiciona uma pessoa
			Id:         line[0],
			Tipo:       line[1],
			Modelo:     line[2],
			Observacao: line[3],
		})

	}

	var res []byte

	res, err = json.Marshal(patrs)
	checkErr(err)

	// fmt.Printf("%v\n", res)
	return res, err

}
