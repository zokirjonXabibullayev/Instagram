package Hendler

import (
	"Instagram-lite/helper"
	"Instagram-lite/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func UserHendler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllUSers(w, r)

	case "POST":
		CreateUser(w, r)
	case "PUT":
		UpdateUser(w, r)
	case "DELETE":
		DeleteUser(w, r)
	}

}

func GetAllUSers(w http.ResponseWriter, r *http.Request) {

	// userni malumotlarini jsondan yechib olish
	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	var PostData []models.PostModel
	PostByte, _ := os.ReadFile("db/posts.json")
	json.Unmarshal(PostByte, &PostData)

	// USERNI malumotlarini postmanda chiqarish
	for i := 0; i < len(UserData); i++ {
		fmt.Fprintln(w, "User's ID: ", UserData[i].ID)
		fmt.Fprintln(w, "User's Firstname: ", UserData[i].Firsname)
		fmt.Fprintln(w, "User's Lastname; ", UserData[i].Lastname)
		fmt.Fprintln(w, UserData[i].Firsname, "'s posts...")
		fmt.Fprintln(w, "----------------------------------")
		for l := 0; l < len(PostData); l++ {
			if UserData[i].ID == PostData[l].ID {
				fmt.Fprintln(w, "  Post's ID:", PostData[l].ID)
				fmt.Fprintln(w, "  User ID:", PostData[l].UserID)
				fmt.Fprintln(w, "  Post's Title ", PostData[l].Title)
				fmt.Fprintln(w, "  Post's Content", PostData[l].Content)
				fmt.Fprintln(w, "  Post's Like count:", PostData[l].Likes)
				fmt.Fprintln(w, "  Post's comments...")
				fmt.Fprintln(w, " --------------------------------------")

			}

		}

	}
	// res, _:=json.Marshal(UserData)
	// os.Writefile("db/user.json", res, 0)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var UpdateUser models.UserModel
	json.NewDecoder(r.Body).Decode(&UpdateUser)

	var NewData []models.UserModel
	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &NewData)

	var UserFound bool
	// Update qilinyayotgan userni ID boyicha topih uchun for dan foydalanish
	for i := 0; i < len(NewData); i++ {
		if UpdateUser.ID == NewData[i].ID {

			// agar updata qilinyayotgan Userni Firsname  yangilanishni hohlama Bo'sh string keladi va shunga
			// tekshiriladi agar bo'sh bolmasa o'zgartiradi
			if UpdateUser.Firsname != "" {
				NewData[i].Firsname = UpdateUser.Firsname

			}
			// bu if agar Lastname ni o'zgartirmoqchi bolmasa
			if UpdateUser.Lastname != "" {
				NewData[i].Lastname = UpdateUser.Lastname
			}
			UserFound = true
			break
		}

	}

	// Agar Berilgan ID topilmagan bolsa quyidagi IF ishlaydi
	if !UserFound {
		fmt.Fprintln(w, "User can not found with ID: ", UpdateUser.ID)
		fmt.Println("User can not found with ID: ", UpdateUser.ID)
		w.WriteHeader(http.StatusNotFound)
		return

	}

	res, _ := json.Marshal(NewData)
	os.WriteFile("db/user.json", res, 0)
	json.NewEncoder(w).Encode(UpdateUser)
	// ekranga chiqarish uchun
	fmt.Println("User Created ", UpdateUser.ID)
	fmt.Fprintln(w, "User Created ")

}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var DeleteUser models.UserModel

	var UserData []models.UserModel
	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &UserData)

	var UserFound bool

	for i := 0; i < len(UserData); i++ {
		if UserData[i].ID == DeleteUser.ID {
			UserData = append(UserData[:i], UserData[i+1:]...)
			UserFound = true
		}

	}

	// endi Berilgan ID topilmagan holarni korib chiqamiz

	if UserFound {
		fmt.Println("User deleted with ID: ", DeleteUser.ID)
		fmt.Fprintln(w, "User deleted with ID: ", DeleteUser.ID)
		w.WriteHeader(http.StatusOK)

	} else {
		fmt.Println("User can not found with ID ", DeleteUser.ID)
		fmt.Fprintln(w, "User can nor found with ID: ", DeleteUser.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// delete bolgan arrayni endi jsonga o'tkazish kerak
	res, _ := json.Marshal(UserData)
	os.WriteFile("db.user.json", res, 0)

	// ekranga User delete qilinganligi haqidagi malumot chiqarish

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var NewUser models.UserModel
	json.NewDecoder(r.Body).Decode(&NewUser)

	var NewData []models.UserModel
	UserByte, _ := os.ReadFile("db/user.json")
	json.Unmarshal(UserByte, &NewData)
	// Yangi yaratilayotgan Userga ID berish
	NewUser.ID = helper.MaxIDUser(NewData)
	// Yangi yaratilgan Userni qolgan Userlar ga qoshib qoyish
	NewData = append(NewData, NewUser)

	res, _ := json.Marshal(NewData)
	os.WriteFile("db/user.json", res, 0)
	// ekranga chiqarish uchun
	fmt.Println("User Created ", NewUser.ID)
	fmt.Fprintln(w, "User Created ")
	json.NewEncoder(w).Encode(NewUser)
}

// Userda qilinadigan ish qolmadi
// Yana qaytish mumkin
