package provider

import "github.com/franela/goreq"

const (
	dictionary = `https://www.moedict.tw/`
)

func Query(word string) (str string, err error){
	req := goreq.Request{
		Uri: dictionary + string(word),
	}

	res, err := req.Do()
	if err != nil {
		return
	}

	return res.Body.ToString()
}
