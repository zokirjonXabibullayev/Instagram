package main 
import (
	"fmt"
	"Instagram-lite/hendler"
	"net/http"
)

func main (){
	fmt.Println("Server is working... :8080")


	http.HandleFunc("/user", Hendler.UserHendler)
	http.HandleFunc("/post", Hendler.PostHendler)
	http.HandleFunc("/comment", Hendler.CommentHendler)

	http.ListenAndServe(":8080", nil)

}