package jsonplaceholder

//const BaseURLJsonPlaceholder = "https://jsonplaceholder.typicode.com"

type Task struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

/*func GetTask() Task {

	key := "1"

	URL := fmt.Sprintf("%s/todos/%s", BaseURLJsonPlaceholder, key)
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	bytes, err := io.ReadAll(resp.Body)
	log.Println(string(bytes))
	var response Task
	errUnmarshal := json.Unmarshal(bytes, &response)
	if errUnmarshal != nil {
		log.Fatal(errUnmarshal)
	}

	return response
}
*/
