package store

import (
	"bufio"
	"os"
)

//Store
type Store struct{
	config *Config
	emails []string
}

//New
func New(config *Config) *Store{
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error{
	file, err := os.Open("C:/Users/Artem/Desktop/apischool/iternal/app/store/gmail.txt")

	if err != nil{
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        s.emails = append(s.emails, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return err
    }
	
	return nil
}

func (s *Store) IsEmailPresent(email string) bool {
    for _, e := range s.emails {
        if e == email {
            return true
        }
    }
    return false
}

func (s *Store) AddEmail(email string) error {
    file, err := os.OpenFile("C:/Users/Artem/Desktop/apischool/iternal/app/store/gmail.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    if _, err := file.WriteString(email + "\n"); err != nil {
        return err
    }

    s.emails = append(s.emails, email)

    return nil
}

func (s *Store) GetEmails() ([]string, error) {
	if len(s.emails) == 0 {
		err := s.Open()
		if err != nil {
			return nil, err
		}
	}
	if len(s.emails) == 0 {
		return nil, nil
	}
	return s.emails, nil
}






