package Hendler

import (
	"Instagram-lite/models"
	"Instagram-lite/helper"
	


	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func PostHendler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllPosts(w,r)

	case "POST":
		PostPost(w,r)
	case "PUT":
		UpdatePost(w,r)
	case "DELETE":
		DeletePost(w,r)
	}
	
}
// GetAllPost ni boshladim time 21:50 14.04.2024
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	var PostData []models.PostModel
	PostByte,_:=os.ReadFile("db/posts.json")
	json.Unmarshal(PostByte, &PostData)

	// commentni ekranga chiqarish uchun uni jsondan yechib olishimiz kerak
	var CommentData []models.CommentModel
	CommentByte,_:=os.ReadFile("db/comments.json")
	json.Unmarshal(CommentByte, &CommentData)





	// Postlarni for bn ekranga chiqaramiz
	for i := 0; i <len(PostData); i++ {
		fmt.Fprintln(w, "Post's ID:          ", PostData[i].ID)
		fmt.Fprintln(w, "Post's UserID:      ", PostData[i].UserID)
		fmt.Fprintln(w, "Post's Title:       ", PostData[i].Title)
		fmt.Fprintln(w, "Post's Content:     ", PostData[i].Content)
		fmt.Fprintln(w, "Post's Likes Count: ", PostData[i].Likes)
		fmt.Fprintln(w, "Post's Comments.... ")
		fmt.Fprintln(w,"-----------------------------------")
			for j := 0; j < len(CommentData); j++ {
				if CommentData[j].PostID==PostData[i].ID {
					fmt.Fprintln(w, "  Comment's ID:      ", CommentData[j].Id)
					fmt.Fprintln(w, "  Comment's PostID:  ", CommentData[j].PostID)
					fmt.Fprintln(w, "  Comment's Content: ", CommentData[j].Content)
					fmt.Fprintln(w, "------------------------------------------")

				}

				
			}
	}


	// ekranga chiqarish tugadi endi yana qaytib jsonga o'rash kerak
	// res,_:=json.Marshal(PostData)
	// os.Writefile("db/posts.json",res, 0)

	// res,_:=json.Marshal(CommentData)
	// os.Writefile("db/comments/json", res, 0)



}

func UpdatePost (w http.ResponseWriter, r *http.Request) {

	var updatePost models.PostModel
	json.NewDecoder(r.Body).Decode(&updatePost)
	

	var PostData []models.PostModel
	PostByte,_:=os.ReadFile("db/posts.json")
	json.Unmarshal(PostByte, &PostData)

	//----------------------------------------

	var UserData []models.UserModel
	UserByte,_:=os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	var  UserFound bool
	
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID==updatePost.UserID {
			UserFound = true
			break
			
		}
	}	

	// User topilmagan holatda ishlaydi
	if !UserFound {
		fmt.Fprintln(w, "User ID can not found ID: ", updatePost.UserID)
		fmt.Println("User ID not found ")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var PostFound bool
	for j := 0; j < len(PostData); j++ {
		if PostData[j].ID==updatePost.ID {
			if updatePost.Title != "" {
				PostData[j].Title=updatePost.Title


			}
			if updatePost.Content != "" {
				PostData[j].Content=updatePost.Content
				
			}
			PostFound= true 
			break
			
		}
		
	}
	// Post ID topilmagan holatda
	if  !PostFound {
		fmt.Fprintln(w, "Post ID not found ", updatePost.ID)
		fmt.Println(w, "Post ID not found ", updatePost.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}		

	res,_:=json.Marshal(PostData)
	os.WriteFile("db/posts.json",res , 0)
	

	fmt.Fprintln(w, "Post Updated!")
	fmt.Println("Post updated!")
	w.WriteHeader(http.StatusOK)
	//finished time 13:0   16.01.2024
	}
	
func DeletePost(w http.ResponseWriter, r *http.Request) {
	
	var deletePost models.PostModel
	json.NewDecoder(r.Body).Decode(&deletePost)
	

	var PostData []models.PostModel
	PostByte,_:=os.ReadFile("db/posts.json")
	json.Unmarshal(PostByte, &PostData)

	//----------------------------------------

	var UserData []models.UserModel
	UserByte,_:=os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	var  UserFound bool
	
	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == deletePost.UserID {
			
			UserFound = true
			break
			
		}
	}	

	// User topilmagan holatda ishlaydi
	if !UserFound {
		fmt.Fprintln(w, "User ID can not found ID: ",deletePost.UserID)
		fmt.Println("User ID not found ")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var PostFound bool
	for j := 0; j < len(PostData); j++ {
		if PostData[j].ID==deletePost.ID {
			PostData=append(PostData[:j], PostData[j+1:]...)
			PostFound= true 
			break
			
		}
		
	}
	// Post ID topilmagan holatda
	if  !PostFound {
		fmt.Fprintln(w, "Post ID not found ",deletePost.ID)
		fmt.Println(w, "Post ID not found ",deletePost.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}		

	res,_:=json.Marshal(PostData)
	os.WriteFile("db/posts.json",res , 0)

	fmt.Fprintln(w, "Post deleted!")
	fmt.Println("Post Deleted!")
	w.WriteHeader(http.StatusOK)
		
}
// Create postni boshladim time 20:04 14.01.2024
func PostPost(w http.ResponseWriter, r *http.Request) {

	var NewPost models.PostModel
	json.NewDecoder(r.Body).Decode(&NewPost)

	var PostData []models.PostModel
	PostByte,_:=os.ReadFile("db/posts.json")
	json.Unmarshal(PostByte, &PostData)
	// Post uchun jsondagi malumotni olish tugadi

	// User uchun jsondagi malumotlarni go ga otkazish boshlandi

	var UserData []models.UserModel
	UserByte,_:=os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	// Is it finished 

	// start print 
	var UserFound bool
	for j := 0; j < len(UserData); j++ {
		if UserData[j].ID==NewPost.UserID {
			UserFound =  true 
				break
		}
		
	}
	// Agar berilgan User ID topilmasa shu if ishlashi kerak
		if !UserFound {
			fmt.Fprintln(w,"User ID can not found")
			fmt.Println("User ID can not found")
			w.WriteHeader(http.StatusNotFound)
			// Agar haqiqatdan ID topilmasa Bu funcsiya tugashi kerak shuning uchun <<return>> qoyilgan
			return
		}

		// Agar Berilgan UserID topilsa IF ishlamaydi va to'gridan tog'ri shu kod ishlashni boshlaydi
		NewPost.ID=helper.MaxIdPost(PostData)
			PostData = append(PostData, NewPost)

			res,_:= json.Marshal(PostData)
			os.WriteFile("db/posts.json",res, 0)

			fmt.Fprintln(w, "Post Created!")
			fmt.Println("Post Created with ID ", NewPost.ID)
			json.NewEncoder(w).Encode(NewPost)
			w.WriteHeader(http.StatusCreated)
// finished 21:19

	
}
 