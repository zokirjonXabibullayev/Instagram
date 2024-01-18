package Hendler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"Instagram-lite/helper"
	"Instagram-lite/models"
	"os"
)

func CommentHendler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllComments(w,r)

	case "POST":
		PostComment(w,r)
	case "PUT":
		UpdateComment(w,r)
	case "DELETE":
		DeleteComment(w,r)
	}
	
}

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	// var CommentData [] models.CommentModel
	// CommentByte,_:=os.ReadFile("db/comments.json")
	// json.Unmarshal(CommentByte, &CommentData)
	var CommentData [] models.CommentModel
	CommentByte,_:=os.ReadFile("db/Comments.json")
	json.Unmarshal(CommentByte, &CommentData)


	for i := 0; i < len(CommentData); i++ {
		fmt.Fprintln(w, "Comment's ID:      ", CommentData[i].Id)
		fmt.Fprintln(w, "Comment's UserID:  ", CommentData[i].UserID)
		fmt.Fprintln(w, "Comment's PostID:  ", CommentData[i].PostID)
		fmt.Fprintln(w, "Comment's Content: ", CommentData[i].Content)
		
		
	}
	// finished time  13:53 16.01.2024




}

func UpdateComment (w http.ResponseWriter, r *http.Request) {

	var deletecomment models.CommentModel
	json.NewDecoder(r.Body).Decode(&deletecomment)

	var CommentData [] models.CommentModel
	CommentByte,_:=os.ReadFile("db/Comments.json")
	json.Unmarshal(CommentByte, &CommentData)

	commentFound:=false
	for i := 0; i <len(CommentData); i++ {
		if CommentData[i].Id== deletecomment.Id {
			CommentData[i].Id= deletecomment.Id
			// CommentData[i].PostID = CommentData[i].PostID
			// CommentData[i].UserID = CommentData[i].UserID
			CommentData[i].Content= deletecomment.Content
			commentFound=true 
			break



		
		}
		
	}
	if !commentFound {
		fmt.Fprintln(w, "Comment's ID not found")
		fmt.Println("Comment not found")
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintln(w, "Comment Updated")
	fmt.Println("'Comment Updated")
	w.WriteHeader(http.StatusOK)

	res,_:=json.Marshal(CommentData)
	os.WriteFile("db/comments.json", res, 0)

	// Update comment finished 14:20 16.01.2024
	
}
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	
	var deletecomment models.CommentModel
	json.NewDecoder(r.Body).Decode(&deletecomment)

	var CommentData [] models.CommentModel
	CommentByte,_:=os.ReadFile("db/Comments.json")
	json.Unmarshal(CommentByte, &CommentData)

	commentFound:=false
	for i := 0; i <len(CommentData); i++ {
		if CommentData[i].Id== deletecomment.Id {
			CommentData=append(CommentData[:i], CommentData[i+1:]...)
			break



		
		}
		
	}
	if !commentFound {
		fmt.Fprintln(w, "Comment's ID not found")
		fmt.Println("Comment not found")
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintln(w, "Comment deleted")
	fmt.Println("'Comment deleted")
	w.WriteHeader(http.StatusOK)

	res,_:=json.Marshal(CommentData)
	os.WriteFile("db/comments.json", res, 0)

	
}
func PostComment(w http.ResponseWriter, r *http.Request) {

	var NewComment models.CommentModel
	json.NewDecoder(r.Body).Decode(&NewComment)

	var CommentData [] models.CommentModel
	CommentByte,_:=os.ReadFile("db/Comments.json")
	json.Unmarshal(CommentByte, &CommentData)

	var NewData []models.UserModel 
	UserByte, _:=os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &NewData)

	var PostData []models.PostModel
	PostByte,_:=os.ReadFile("db/posts.json")
	json.Unmarshal(PostByte, &PostData)
	 
	var UserFound bool
	NewComment.Id=helper.MaxIdComment(CommentData)
	for i := 0; i < len(NewData); i++ {
		if NewData[i].ID==NewComment.UserID {
		UserFound = true 
		break 
		}
		
	} 

	var PostFound bool
	for j := 0; j < len(PostData); j++ {
		if PostData[j].ID==NewComment.PostID {
		PostFound=true
		break
		}

		
	} 
	if ! UserFound { 
		fmt.Fprintln(w, "User not found  with ID: ", )
		fmt.Println("User not found with ID: ")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// PostID topilmagan holatda 
	if !PostFound {
		fmt.Fprintln(w, "Post not found with ID: ")
		fmt.Println("Post not found with ID: ")
		w.WriteHeader(http.StatusNotFound)
		return
		
	}

	// endi ekranga chiqarish kerak amiyot bajarilgani haqida

	fmt.Fprintln(w, "New comment Created!")
	fmt.Println("New Comment Created!")
	w.WriteHeader(http.StatusOK)

	// // Userdatani qayta json ga o'rash
	// res, _:=json.Marshal(NewData)
	// os.WriteFile("db/user.json", res, 0)
	

	// // Postdatani qayta jsonga o'rash
	// res, _:=json.Marshal(PostData)
	// os.WriteFile("db/posts.json",res , 0)

	kal, _:=json.Marshal(CommentData)
	os.WriteFile("db/Comments.json", kal, 0)
}
