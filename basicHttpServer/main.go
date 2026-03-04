package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

//r : request object ,w : response object

// gin framework similar to express , wont use http package directly
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	_, err := w.Write([]byte("Hello from GO net/http server"))
	if err != nil {
		fmt.Println("Write error:", err)
	}
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	_, err := w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := map[string]any{
		"ok":       true,
		"message":  "Json encode successful",
		"dateTime": time.Now().UTC(),
	}
	_ = json.NewEncoder(w).Encode(res)

}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "Method not allowed",
		})
		return
	}
	defer r.Body.Close()
	var req TestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Invalid JSON",
		})
		return
	}
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		writeJson(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Name is required",
		})
		return
	}
	writeJson(w, http.StatusOK, map[string]any{
		"ok":        true,
		"data":      req,
		"timestamp": time.Now().UTC(),
	})
}

type TestRequest struct {
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/getQuery", queryHandler)
	http.HandleFunc("/ok", successHandler)
	http.HandleFunc("/test", testHandler)

	url := "https://jsonplaceholder.typicode.com/todos"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	if len(bodyBytes) > 500 {
		bodyBytes = bodyBytes[:50] // Truncate for display purposes
	}
	fmt.Println("Response Body:", string(bodyBytes))
	fmt.Println("Response Status:", resp.StatusCode)
	fmt.Println("Status :", resp.Status)




	type CatFactResponseStructure struct {
		Fact   string `json:"fact"`
		Length int    `json:"length"`
	}
	url1 := "https://catfact.ninja/fact"
	resp1, err := http.Get(url1)
	if err != nil {
		fmt.Println("Error fetching cat fact:", err)
		return
	}
	defer resp1.Body.Close()
	bodyBytes1, err := io.ReadAll(resp1.Body)
	if err != nil {
		fmt.Println("Error reading cat fact response body:", err)
		return
	}
	var catFact CatFactResponseStructure
	// unmarshall the json response into go struct
	if err := json.Unmarshal(bodyBytes1, &catFact); err != nil {
		fmt.Println("Error unmarshalling cat fact JSON:", err)
		return
	}
	fmt.Println("Cat Fact:", catFact.Fact)
	fmt.Println("Fact Length:", catFact.Length)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
