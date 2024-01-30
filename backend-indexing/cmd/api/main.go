package main

import (
	"backend/internal/models"
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type application struct {
	Domain          string
	Host            string
	ZincsearchHost  string
	ZincsearchIndex string
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("There is no port defined in the env variables")
	}

	// set application config
	var app application

	// connect to the database
	app.Domain = "example.com"
	app.ZincsearchHost = "http://localhost:4080"
	app.ZincsearchIndex = "enronMailCT"

	args := os.Args

	if len(args) > 2 && args[2] == "index" {
		Index(args[1])
	}

	server := &http.Server{
		Handler: app.routes(),
		Addr:    ":" + port,
	}

	log.Println("Starting application on port", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func Index(database string) {
	fmt.Println("Indexing...")
	database_path := "../../" + database + "/maildir/"
	userList := listSubFolders(database_path)
	contador := 0
	for _, user := range userList {
		folders_per_user := listSubFolders(database_path + user)
		fmt.Println(user)
		for _, folder := range folders_per_user {
			mails_list, err := list_mails(database_path + user + "/" + folder + "/")
			if err != nil {
				continue
			}
			for _, mail_file := range mails_list {
				mail_content, err := os.Open(database_path + user + "/" + folder + "/" + mail_file)
				if err != nil {
					continue
				}
				lines := bufio.NewScanner(mail_content)
				contador++
				index_data(parse_data(lines, contador))
				mail_content.Close()
			}
		}
	}
}

func listSubFolders(data_base_name string) []string {
	users, err := os.ReadDir(data_base_name)
	if err != nil {
		log.Fatal("Unable to read the database because of ", err)
	}

	var list_users []string
	for _, user := range users {
		if user.Name() != ".DS_Store" {
			list_users = append(list_users, user.Name())
		}
	}

	return list_users
}

func list_mails(folder_name string) ([]string, error) {
	files, err := os.ReadDir(folder_name)
	if err != nil {
		return []string{}, err
	}

	var file_names []string
	for _, file := range files {
		file_names = append(file_names, file.Name())
	}
	return file_names, nil
}

func parse_data(data_lines *bufio.Scanner, id int) models.Email {
	var data models.Email

	for data_lines.Scan() {
		data.ID = strconv.Itoa(id)
		if strings.Contains(data_lines.Text(), "Message-ID:") {
			if 12 <= len(data_lines.Text()) {
				data.Message_ID = data_lines.Text()[12:]
			} else {
				data.Message_ID = ""
			}
		} else if strings.Contains(data_lines.Text(), "Date:") {
			if 6 <= len(data_lines.Text()) {
				data.Date = data_lines.Text()[6:]
			} else {
				data.Date = ""
			}
		} else if strings.Contains(data_lines.Text(), "From:") && !strings.Contains(data_lines.Text(), "X-From:") {
			if 6 <= len(data_lines.Text()) {
				data.From = data_lines.Text()[6:]
			} else {
				data.From = ""
			}
		} else if strings.Contains(data_lines.Text(), "To:") && !strings.Contains(data_lines.Text(), "X-To:") {
			if 4 <= len(data_lines.Text()) {
				data.To = data_lines.Text()[4:]
			} else {
				data.To = ""
			}
		} else if strings.Contains(data_lines.Text(), "Subject:") {
			if 9 <= len(data_lines.Text()) {
				data.Subject = data_lines.Text()[9:]
			} else {
				data.Subject = ""
			}
		} else if strings.Contains(data_lines.Text(), "Cc:") {
			if 4 <= len(data_lines.Text()) {
				data.Cc = data_lines.Text()[4:]
			} else {
				data.Cc = ""
			}
		} else if strings.Contains(data_lines.Text(), "Mime-Version:") {
			if 14 <= len(data_lines.Text()) {
				data.Mime_Version = data_lines.Text()[14:]
			} else {
				data.Mime_Version = ""
			}
		} else if strings.Contains(data_lines.Text(), "Content-Type:") {
			if 14 <= len(data_lines.Text()) {
				data.Content_Type = data_lines.Text()[14:]
			} else {
				data.Content_Type = ""
			}
		} else if strings.Contains(data_lines.Text(), "Content-Transfer-Encoding:") {
			if 27 <= len(data_lines.Text()) {
				data.Content_Transfer_Encoding = data_lines.Text()[27:]
			} else {
				data.Content_Transfer_Encoding = ""
			}
		} else if strings.Contains(data_lines.Text(), "X-From:") {
			if 8 <= len(data_lines.Text()) {
				data.X_From = data_lines.Text()[8:]
			} else {
				data.X_From = ""
			}
		} else if strings.Contains(data_lines.Text(), "X-To:") {
			if 6 <= len(data_lines.Text()) {
				data.X_To = data_lines.Text()[6:]
			} else {
				data.X_To = ""
			}
		} else if strings.Contains(data_lines.Text(), "X-cc:") {
			if 6 <= len(data_lines.Text()) {
				data.X_cc = data_lines.Text()[6:]
			} else {
				data.X_cc = ""
			}
		} else if strings.Contains(data_lines.Text(), "X-bcc:") {
			if 7 <= len(data_lines.Text()) {
				data.X_bcc = data_lines.Text()[7:]
			} else {
				data.X_bcc = ""
			}
		} else if strings.Contains(data_lines.Text(), "X-Folder:") {
			if 10 <= len(data_lines.Text()) {
				data.X_Folder = data_lines.Text()[10:]
			} else {
				data.X_Folder = ""
			}
		} else if strings.Contains(data_lines.Text(), "X-Origin:") {
			if 10 <= len(data_lines.Text()) {
				data.X_Origin = data_lines.Text()[10:]
			} else {
				data.X_Origin = ""
			}
		} else if strings.Contains(data_lines.Text(), "X-FileName:") {
			if 12 <= len(data_lines.Text()) {
				data.X_FileName = data_lines.Text()[12:]
			} else {
				data.X_FileName = ""
			}
		} else {
			data.Body = data.Body + data_lines.Text()
		}
	}
	return data
}

func index_data(data models.Email) {
	auth := "admin:Complexpass#123"
	base64encoded_creds := base64.StdEncoding.EncodeToString([]byte(auth))
	index := "enronMailCT"
	zincsearch_host := "http://localhost:4080"
	zincsearch_url := zincsearch_host + "/api/" + index + "/_doc"
	jsonData, _ := json.MarshalIndent(data, "", "   ")

	req, err := http.NewRequest("POST", zincsearch_url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error reading request", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+base64encoded_creds)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}
