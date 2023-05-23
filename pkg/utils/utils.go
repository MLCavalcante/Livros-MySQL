package utils //o processo de "unmarshalling" é frequentemente usado para converter os dados recebidos em uma requisição em um formato serializado, como JSON, em uma estrutura de dados que possa ser manipulada pelo controller, essa file faz isso

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody (r http.Request, X interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), X); err != nil{
			return
		}

		
	}

}
