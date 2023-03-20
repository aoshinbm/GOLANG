package main
import "fmt"
message User {
	int32 id = 1;
	string name = 2;
	string email = 3;
	}
	func main() {
	user := &User{ID: 1, Name: "John Smith", Email:
	"john.smith@example.com"}
	data, _ := proto.Marshal(user) // Encode the struct
	newUser := &User{}
	proto.Unmarshal(data, newUser) // Decode the struct
	fmt.Println(newUser)
	}