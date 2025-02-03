package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Define the structure for a valid JSON response
type Resp struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

// Define the structure for error response
type ErrorResp struct {
	Number string `json:"alphabet"`
	Error  bool   `json:"error"`
}

// Define the entrypoint
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/classify-number", numberHandler).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func numberHandler(w http.ResponseWriter, r *http.Request) {
	numberParam := r.URL.Query().Get("number")
	if numberParam == "" {
		sendErrorResponse(w, "invalid-input", http.StatusBadRequest)
		return
	}

	number, err := strconv.Atoi(numberParam)
	if err != nil {
		sendErrorResponse(w, numberParam, http.StatusBadRequest)
		return
	}

	funFact, err := fetchFunFact(number)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	response := getResponse(number, funFact)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, numParam string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResp{
		Number: numParam,
		Error:  true,
	})
}

func fetchFunFact(n int) (string, error) {
	url := fmt.Sprintf("http://numbersapi.com/%d/math", n)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch fun fact")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func getResponse(n int, funFact string) Resp {
	props := []string{}
	if isArmstrong(n) {
		props = append(props, "armstrong")
	}

	if n%2 == 0 {
		props = append(props, "even")
	} else {
		props = append(props, "odd")
	}

	return Resp{
		Number:     n,
		IsPrime:    isPrime(n),
		IsPerfect:  isPerfect(n),
		Properties: props,
		DigitSum:   digitSum(n),
		FunFact:    funFact,
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPerfect(n int) bool {
	if n <= 1 {
		return false
	}
	sum := 1
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			sum += i
			if otherNum := n / i; otherNum != i {
				sum += otherNum
			}
		}
	}
	return sum == n
}

func digitSum(n int) int {
	sum := 0
	n = int(math.Abs(float64(n)))
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func isArmstrong(n int) bool {
	original := n
	length := len(strconv.Itoa(int(math.Abs(float64(n)))))
	sum := 0
	temp := int(math.Abs(float64(n)))
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), (float64(length))))
		temp /= 10
	}
	return sum == original
}
