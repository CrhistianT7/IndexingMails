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
	"time"

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
		t0 := time.Now()
		Index(args[1])
		fmt.Printf("Total time is: %v\n", time.Since(t0))
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
		index_per_user(database_path, user, folders_per_user, &contador)
	}
}

func index_per_user(database_path string, user string, folders_per_user []string, contador *int) {
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
			*contador++
			index_data(parse_data(lines, *contador))
			mail_content.Close()
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

func get_key_value(key string, current_line string) string {
	index := len(key) + 1
	if index <= len(current_line) {
		return current_line[index:]
	}
	return ""
}

func parse_data(data_lines *bufio.Scanner, id int) models.Email {
	var data models.Email

	for data_lines.Scan() {
		data.ID = strconv.Itoa(id)
		switch {
		case strings.Contains(data_lines.Text(), "Message-ID:"):
			data.Message_ID = get_key_value("Message-ID:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "Date:"):
			data.Date = get_key_value("Date:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "From:") && !strings.Contains(data_lines.Text(), "X-From:"):
			data.From = get_key_value("From:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "To:") && !strings.Contains(data_lines.Text(), "X-To:"):
			data.To = get_key_value("To:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "Subject:"):
			data.Subject = get_key_value("Subject:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "Cc:"):
			data.Cc = get_key_value("Cc:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "Mime-Version:"):
			data.Mime_Version = get_key_value("Mime-Version:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "Content-Type:"):
			data.Content_Type = get_key_value("Content-Type:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "Content-Transfer-Encoding:"):
			data.Content_Transfer_Encoding = get_key_value("Content-Transfer-Encoding:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "X-From:"):
			data.X_From = get_key_value("X-From:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "X-To:"):
			data.X_To = get_key_value("X-To:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "X-cc:"):
			data.X_cc = get_key_value("X-cc:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "X-bcc:"):
			data.X_bcc = get_key_value("X-bcc:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "X-Folder:"):
			data.X_Folder = get_key_value("X-Folder:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "X-Origin:"):
			data.X_Origin = get_key_value("X-Origin:", data_lines.Text())
		case strings.Contains(data_lines.Text(), "X-FileName:"):
			data.X_FileName = get_key_value("X-FileName:", data_lines.Text())
		default:
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
