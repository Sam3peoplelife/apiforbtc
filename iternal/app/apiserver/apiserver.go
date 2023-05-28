package apiserver

import (
	"apischool/iternal/app/store"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"io/ioutil"
	"net/http"
	//"net/smtp"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer
type APIServer struct{
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

func NEw(config *Config) * APIServer{
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),	
	}
}

func (s *APIServer) Start() error{
	if err := s.configureLogger(); err != nil {
		return err

	}
	s.configRouter()

	
	if err := s.configureStore(); err != nil {
		return err
	}	
	
	s.logger.Info("starting api server")
	

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error{
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil{
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configRouter(){
	s.router.HandleFunc("/subscribe", s.Subscribe())
	s.router.HandleFunc("/get", s.GetExchangeRate())
	//s.router.HandleFunc("/sendEmails", s.sendEmails())
}
func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}
func (s *APIServer) Subscribe() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Error reading request body", http.StatusInternalServerError)
            return
        }

        email := string(body)

        if s.store.IsEmailPresent(email) {
            http.Error(w, "Email already subscribed", http.StatusConflict)
            return
        }

        if err := s.store.AddEmail(email); err != nil {
            http.Error(w, "Error adding email", http.StatusInternalServerError)
            return
        }
    }
}
var BTCtoUAH float64 
func (s *APIServer) GetExchangeRate() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        resp, err := http.Get("https://api.coinbase.com/v2/exchange-rates?currency=BTC")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        var result map[string]interface{}
        json.NewDecoder(resp.Body).Decode(&result)

        rates := result["data"].(map[string]interface{})["rates"].(map[string]interface{})
        BTCtoUAHStr, ok := rates["UAH"].(string)
        if !ok {
            http.Error(w, "Error parsing exchange rate", http.StatusInternalServerError)
            return
        }

        BTCtoUAH, err := strconv.ParseFloat(BTCtoUAHStr, 64)
        if err != nil {
            http.Error(w, "Error converting exchange rate to float64", http.StatusInternalServerError)
            return
        }

        io.WriteString(w, fmt.Sprintf("Current exchange rate from BTC to UAH: %f", BTCtoUAH))
    }
}
/*func (s *APIServer) sendEmails() http.HandlerFunc {
	from := "DFTandgameschanel@tmail.com"
	pass := "skotina."
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	emails, err := s.store.GetEmails()
	if err != nil {
		s.logger.Error("Failed to retrieve emails from store: ", err)
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "Failed to retrieve emails from store", http.StatusInternalServerError)
		}
	}

	for _, email := range emails {
		message := "Current Currency Rate\n\n" + strconv.FormatFloat(BTCtoUAH, 'f', -1, 64)

		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(message))
		if err != nil {
			s.logger.Error("Failed to send email: ", err)
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Emails sent successfully"))
	}
}
*/





