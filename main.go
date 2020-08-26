package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"github.com/acarl005/stripansi"
)

var wg sync.WaitGroup

func main() {
	var oneLine, verboseMode bool
	var ApiToken, lines string
	flag.StringVar(&ApiToken, "u", "", "Telegram ApiToken")
	flag.BoolVar(&oneLine, "1", false, "Send message line-by-line")
	flag.BoolVar(&verboseMode, "v", false, "Verbose mode")
	flag.Parse()

	apitokenEnv := os.Getenv("TELEGRAM_API_TOKEN")
	TelegramEnv := "https://api.telegram.org/bot"+apitokenEnv+"/sendMessage"
	chatid := os.Getenv("TELEGRAM_CHAT_ID")
	if TelegramEnv != "" {
		ApiToken = TelegramEnv
	} else {
		if ApiToken == "" {
			if verboseMode {
				fmt.Println("Telegram ApiToken not set!")
			}
		}
	}

	if !isStdin() {
		os.Exit(1)
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()

		fmt.Println(line)
		if oneLine {
			if ApiToken != "" {
				wg.Add(1)
				go teleman(ApiToken, line,chatid)
			}
		} else {
			lines += line
			lines += "\n"
		}
	}

	if !oneLine {
		wg.Add(1)
		go teleman(ApiToken, lines,chatid)
	}
	wg.Wait()
}

func isStdin() bool {
	f, e := os.Stdin.Stat()
	if e != nil {
		return false
	}

	if f.Mode()&os.ModeNamedPipe == 0 {
		return false
	}

	return true
}

type data struct {
	Chat_id string `json:"chat_id"`
	Text string `json:"text"`
}

func teleman(url string, line string,chatid string) {
	data, _ := json.Marshal(data{Chat_id:chatid,Text: stripansi.Strip(line)})
	http.Post(url, "application/json", strings.NewReader(string(data)))
	wg.Done()
}

