package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TwiML struct {
	XMLName xml.Name `xml:"Response"`

	Say     string `xml:",omitempty"`
	Play    string `xml:",omitempty"`
	Message string `xml:",omitempty"`
}

func main() {

	http.HandleFunc("/", hello)
	http.HandleFunc("/twiml", twiml)
	http.HandleFunc("/sms", sms)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	//http.ListenAndServe(":4567", nil)

}

type Request struct {
	Ok      bool
	Members []User
}

func sms(w http.ResponseWriter, r *http.Request) {
	// Set initial variables
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	r.ParseForm()
	text := r.PostForm["text"][0]
	bodyArray := strings.Fields(text)
	to_slack_number := bodyArray[0]
	slack_msg := strings.Join(bodyArray[1:], " ")
	fmt.Println(to_slack_number, slack_msg)
	v := url.Values{}
	v.Set("To", to_slack_number)
	v.Set("From", os.Getenv("TWILIO_NUMBER"))
	v.Set("Body", slack_msg)
	rb := *strings.NewReader(v.Encode())

	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
}

func twiml(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	body := r.PostForm["Body"][0]
	from := r.PostForm["From"][0]
	bodyArray := strings.Fields(body)
	bodyArray = append(bodyArray, from)
	slack_channel := "#working-on"
	slack_msg := strings.Join(bodyArray, " ")
	resp, _ := http.Post(os.Getenv("SLACK_WEARHACKS_WEBHOOK_URL"), "text/json", strings.NewReader("{\"text\": \""+slack_msg+"\", \"channel\" : \""+slack_channel+"\"}"))
	fmt.Println(resp.Status)
	msg := "Responding..."
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		msg = "Message successfully sent!!!"

	} else {
		msg = "Message NOT sent!!!"
	}
	twiml := TwiML{Message: msg}
	x, err := xml.Marshal(twiml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Welcome to SlackMS!!!")
}

type User struct {
	Id        string
	Name      string
	Deleted   bool
	Color     string
	Profile   profile
	Is_Admin  bool
	Is_Owner  bool
	Has_2fa   bool
	Has_Files bool
}

type profile struct {
	First_Name string
	Last_Name  string
	Real_Name  string
	Email      string
	Skype      string
	Phone      string
	Image_24   string
	Image_32   string
	Image_48   string
	Image_72   string
	Image_192  string
}
