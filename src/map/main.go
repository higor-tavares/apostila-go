package main
import "fmt"

func main() {
	//empty := make(map[int]string)
	//declaring capacity from map to improve performance
	httpStatus := make(map[string]int, 32)
	httpStatus["OK"] = 200
	httpStatus["CREATED"] = 201
	httpStatus["ACCEPTED"] = 202
	httpStatus["NO_CONTENT"] = 204
	httpStatus["BAD_REQUEST"] = 400
	httpStatus["UNAUTHRIZED"] = 401
	httpStatus["FORBIDEN"] = 403
	httpStatus["NOT_FOUND"] = 404
	httpStatus["INTERNAL_SERVER_ERROR"] = 500
	httpStatus["NOT_EXISTENT"] = 700
	delete(httpStatus, "NOT_EXISTENT")
	fmt.Println(httpStatus)
	fmt.Printf("Status: %s Code: %d\n", "NOT_FOUND", httpStatus["NOT_FOUND"])
	fmt.Println("                 o/      ")
	fmt.Println("-------------------------")
	fmt.Println("                 o\\     ")
	for status, code := range httpStatus {
		fmt.Printf("Status: %s -> code: %d\n", status, code)
	}
}