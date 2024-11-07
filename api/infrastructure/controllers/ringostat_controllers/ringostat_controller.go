package ringostat_controllers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const ringostatAPIToken = "BICavaUTsaqnfxwSNfBuOfI1zc5aDqAx"
const ringostatAPIURL = "https://api.ringostat.net/api/v1/callback"

type CallbackRequest struct {
	Phone      string `json:"phone"`       // Телефон клиента
	OperatorID string `json:"operator_id"` // ID оператора, к которому нужно соединить звонок
	ProjectID  int    `json:"project_id"`  // ID проекта Ringostat
}

//type CallbackRequest struct {
//	From      string `json:"from" binding:"required"`
//	To        string `json:"to" binding:"required"`
//	ProjectID string `json:"project_id" binding:"required"`
//}

func Call(c *gin.Context) {
	var callbackReq CallbackRequest
	if err := c.ShouldBindJSON(&callbackReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Подготовка данных для Ringostat
	requestData, err := json.Marshal(callbackReq)
	if err != nil {
		log.Printf("Error marshaling request data: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Настройка HTTP-запроса
	req, err := http.NewRequest("POST", ringostatAPIURL, bytes.NewBuffer(requestData))
	if err != nil {
		log.Printf("Error creating HTTP request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Добавление заголовков для авторизации и типа контента
	req.Header.Set("Authorization", "Bearer "+ringostatAPIToken)
	req.Header.Set("Content-Type", "application/json")

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to Ringostat: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer resp.Body.Close()

	// Обработка ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Проверка успешности запроса
	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": fmt.Sprintf("Ringostat API error: %s", body)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Callback initiated successfully"})
}

func Test(c *gin.Context) {
	authKey := "BICavaUTsaqnfxwSNfBuOfI1zc5aDqAx" // Замените на ваш реальный ключ авторизации
	apiURL := "https://api.ringostat.net/callback/outward_call"

	// Параметры запроса
	extension := c.Query("extension")
	destination := c.Query("destination")

	if extension == "" || destination == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметры 'extension' и 'destination' обязательны"})
		return
	}

	// Формирование тела запроса
	data := url.Values{}
	data.Set("extension", extension)
	data.Set("destination", destination)

	// Создание HTTP-запроса
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании запроса"})
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Auth-key", authKey)

	// Отправка запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при выполнении запроса"})
		return
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при чтении ответа"})
		return
	}

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		c.JSON(http.StatusOK, gin.H{"message": "Звонок успешно инициирован", "response": string(body)})
	} else {
		c.JSON(resp.StatusCode, gin.H{"error": "Ошибка при создании звонка", "response": string(body)})
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var operatorConn *websocket.Conn // WebSocket-соединение для оператора
var clientConn *websocket.Conn   // WebSocket-соединение для клиента

func Signal(c *gin.Context) {
	userID := c.Query("user")
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set WebSocket upgrade:", err)
		return
	}
	defer ws.Close()

	// Сохраняем соединения для оператора и клиента
	if userID == "operator" {
		operatorConn = ws
		fmt.Println("Operator connected")
	} else if userID == "client" {
		clientConn = ws
		fmt.Println("Client connected")
	}

	// Обработка входящих сообщений
	for {
		var msg map[string]interface{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println("WebSocket read error:", err)
			break
		}

		fmt.Printf("Received message from %s: %v\n", userID, msg)

		if msg["type"] == "offer" && userID == "client" {
			// Передаем offer оператору
			if operatorConn != nil {
				operatorConn.WriteJSON(msg)
				fmt.Println("Sent offer to operator")
			}
		} else if msg["type"] == "answer" && userID == "operator" {
			// Передаем answer клиенту
			if clientConn != nil {
				clientConn.WriteJSON(msg)
				fmt.Println("Sent answer to client")
			}
		} else if msg["type"] == "candidate" {
			// Передаем ICE-кандидаты между клиентом и оператором
			if userID == "client" && operatorConn != nil {
				operatorConn.WriteJSON(msg)
				fmt.Println("Sent ICE candidate to operator")
			} else if userID == "operator" && clientConn != nil {
				clientConn.WriteJSON(msg)
				fmt.Println("Sent ICE candidate to client")
			}
		}
	}
}
