package helper

import (
	"Instagram-lite/models"
	
)
// bring MaxID For User
func MaxIDUser(UserArray []models.UserModel) int {
	var maxID =  0
	 for i := 0; i <len(UserArray); i++ {
		if maxID < UserArray[i].ID {
			maxID = UserArray[i].ID
		}
	 }
	 return maxID+1

	 // Bring MaxID for Post
}
func MaxIdPost(PostArray []models.PostModel) int {
	var maxID =  0
	 for i := 0; i <len(PostArray); i++ {
		if maxID < PostArray[i].ID {
			maxID = PostArray[i].ID
		}
	 }
	 return maxID+1

	 // Bring MaxID for Comment
}
func MaxIdComment(CommentArray []models.CommentModel) int {
	var maxID =  0
	 for i := 0; i <len(CommentArray); i++ {
		if maxID < CommentArray[i].Id {
			maxID = CommentArray[i].Id
		}
	 }
	 return maxID+1

}