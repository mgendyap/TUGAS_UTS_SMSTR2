package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "os/exec"
	// "time"
	"utsstrukdat/controller/message"
	"utsstrukdat/controller/post"
	"utsstrukdat/controller/user"
	"utsstrukdat/db"
)

// func clear() {
// 	cmd := exec.Command("cmd", "/c", "cls")
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	for {
		
		var username string
		var password string
		var title string
		var body string
		var category string
		var pilihCategory int


		menu:
		fmt.Println("Menu")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Println("Masukan pilihan : ")

		var inputMenu int
		fmt.Scan(&inputMenu)

		if inputMenu == 1 {
			
			fmt.Println("Register")
			fmt.Println("Masukan Username : ")

			fmt.Scan(&username)

			fmt.Println("Masukan password : ")
			fmt.Scan(&password)

			fmt.Println("Masukan verifikasi password : ")
			var verPassword string
			fmt.Scan(&verPassword)

			dataUser := &db.FieldUser{
				Username: username,
				Password: password,
			}

			kode := userController.Register(dataUser, &verPassword)

			if kode == 200 {
				fmt.Println("Akun sudah di buat")
			} else if kode == 400 {
				fmt.Println("Password tidak sama")
			} else if kode == 409 {
				fmt.Println("Username sudah di gunakan")
			}
		} else if inputMenu == 2 {
			
			fmt.Println("Login")
			fmt.Println("Masukan username : ")
			fmt.Scan(&username)

			fmt.Println("Masukan password : ")
			fmt.Scan(&password)

			dataUser := &db.FieldUser{
				Username: username,
				Password: password,
			}

			token := userController.Login(dataUser)
			
			if token != nil {
				for {
					dashboard:

					fmt.Println("Dashboard")
					fmt.Println("1. Beranda")
					fmt.Println("2. Pesan")
					fmt.Println("3. Buat cerita")
					fmt.Println("4. Edit cerita")
					fmt.Println("5. Hapus cerita")
					fmt.Println("6. Cari akun")
					fmt.Println("7. Cari Kategori")
					fmt.Println("8. Logout")

					var inputDashboard int
					fmt.Scan(&inputDashboard)

					if inputDashboard == 1 {
						
						post := postController.ShowPost()

						for _, val := range *post {
							fmt.Println("Author : ", val.Author)
							fmt.Println("Kategori : ", val.Category)
							fmt.Println("Title : ", val.Title)
							fmt.Println("Body : ", val.Body)
						}

					} else if inputDashboard == 2 {

						for {
							
							var pilihPesan int
							fmt.Println("1. Lihat Pesan")
							fmt.Println("2. Kirim Pesan")
							fmt.Println("3. Kembali ke Dashboard")
							fmt.Scan(&pilihPesan)

							if pilihPesan == 1 {
								

								message := messageController.ShowMessage(&token.Username)

								for _, val := range *message {
									fmt.Println("From : ", val.From)
									fmt.Println("Message : ", val.Message)
								}

							} else if pilihPesan == 2 {
								
								var penerima string
								fmt.Println("Masukan username yang mau anda kirimkan pesan")
								fmt.Scan(&penerima)

								var pesan string
								fmt.Println("Masukan Pesan")
								fmt.Scan(&pesan)

								check := messageController.SendMessage(&penerima, &token.Username, &pesan)

								if check == 200 {
									fmt.Println("Pesan anda sudah terkirim ke ", penerima)
								} else {
									fmt.Println("Username penerima tidak di temukan")
								}
							} else if pilihPesan == 3 {
								goto dashboard
							} else {
								fmt.Println("Format tidak di temukan")
							}
						}

					} else if inputDashboard == 3 {
						
						fmt.Println("1. Blog")
						fmt.Println("2. Programming")
						fmt.Println("Masukan kategori")
						fmt.Scan(&pilihCategory)
						if pilihCategory == 1 {
							category = "Blog"
						} else {
							category = "Programming"
						}

						fmt.Println("Masukan title")
						fmt.Scan(&title)

						fmt.Println("Masukan body")
						fmt.Scan(&body)

						dataPost := &db.FieldPost{
							Author:   token.Username,
							Category: category,
							Title:    title,
							Body:     body,
						}

						kode := postController.InsertPost(dataPost)

						if kode == 200 {
							fmt.Println("postingan telah di buat")
						} else {
							fmt.Println("Judul telah di gunakan")
						}
						
					} else if inputDashboard == 4 {
						
						fmt.Println("Masukan judul post yang mau di ubah")
						fmt.Scan(&title)

						fmt.Println("Masukan carita baru")
						fmt.Scan(&body)
						kode := postController.UpdatePost(&title, &body, &token.Username)

						if kode == 200 {
							fmt.Println("Postingan sudah di update")
						} else if kode == 403 {
							fmt.Println("Postingan tersebut bukan milik anda")
						} else {
							fmt.Println("Postingan tidak di temukan")
						}
					} else if inputDashboard == 5 {
						
						fmt.Println("Masukan judul post yang mau di hapus")
						fmt.Scan(&title)

						kode := postController.DeletePost(&title, &token.Username)

						if kode == 200 {
							fmt.Println("Postingan sudah di hapus")
						} else if kode == 403 {
							fmt.Println("Postingan tersebut bukan milik anda")
						} else {
							fmt.Println("Postingan tidak di temukan")
						}
					} else if inputDashboard == 6 {
						
						fmt.Println("Masukan username")
						fmt.Scan(&username)

						akun := userController.ShowPostByAccount(&username)

						if akun != nil {
							fmt.Println("Username : ", username, "\n")
							for _, val := range *akun {
								fmt.Println("Author : ", val.Author)
								fmt.Println("Category : ", val.Category)
								fmt.Println("Title : ", val.Title)
								fmt.Println("Body : ", val.Body)
							}
						} else {
							fmt.Println("Akun tersebut belum memposting apapun")
						}
					} else if inputDashboard == 7 {
						
						fmt.Println("Masukan Kategori yang mau anda cari")
						fmt.Println("1. Blog")
						fmt.Println("2. Programming")
						fmt.Scan(&pilihCategory)

						if pilihCategory == 1 || pilihCategory == 2 {
							if pilihCategory == 1 {
								category = "Blog"
							} else if pilihCategory == 2 {
								category = "Programming"
							}
							post := postController.ShowByCategory(&category)

							for _, val := range *post {
								fmt.Println("Author : ", val.Author)
								fmt.Println("Kategori : ", val.Category)
								fmt.Println("Title : ", val.Title)
								fmt.Println("Body : ", val.Body)
							}
						} else {
							fmt.Println("Kategori tidak di temukan")
						}
					} else if inputDashboard == 8 {
						goto menu
					}
				}
			} else {
				fmt.Println("Username atau password salah")
			}
		} else {
			break
		}
	}
}
