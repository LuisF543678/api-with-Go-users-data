package structures

type User struct{
	id int `json:"id"`
	Username string `json: "username"`
	First_name string `json: "fisrt_name"`
	Last_name string `json: "last_name"`
}

type structures struct {
	Status string `json:"status"`
	Data User `json:"data"`
	Message string `json:"message"`
	
}