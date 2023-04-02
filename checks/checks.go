package checks

import (
	"net/mail"
	"regexp"

	"github.com/Ba1vo/Proektirovanie/decoder"
)

var symbols = [...]rune{'/', '*', ';', '=', '"', '\'', '-', '\\', '#', '(', ')'}

func InjectCheck(str string) string {
	var strOut []rune
	for _, char := range str {
		for _, symbol := range symbols {
			if char == symbol {
				strOut = append(strOut, '\\')
				break
			}
		}
		strOut = append(strOut, char)
	}
	return string(strOut)
}

func CheckUserReg(User decoder.RegUser) bool {
	return checkName(User.Name) && checkPass(User.Pass) && CheckEmail(User.Email)
}

func CheckUserAuth(User decoder.AuthUser) bool {
	return checkPass(User.Pass) && CheckEmail(User.Email)
}

func CheckAdress(Adress string) bool {
	regexp := regexp.MustCompile(`^[A-яЁё]{1,20}$`)
	return regexp.Match([]byte(Adress))
}

func CheckBook(Book decoder.FullBook) bool {
	for _, v := range Book.Authors {
		if !(checkDispAuthor(v)) {
			return false
		}
	}
	for _, v := range Book.Publishers {
		if !(checkPublisher(v)) {
			return false
		}
	}
	return checkISBN(Book.ISBN) && checkPrice(Book.Price) && checkDesc(Book.Desc)
}

func CheckCode(code string) bool {
	regexp := regexp.MustCompile("[a-z0-9]{6}")
	return regexp.Match([]byte(code))
}

func checkName(Name string) bool {
	regexp := regexp.MustCompile(`^[A-яЁё]{1,20}$`)
	return regexp.Match([]byte(Name))
}

func checkPass(Pass string) bool {
	regexp := regexp.MustCompile(`^(\w){6,20}$`)
	return regexp.Match([]byte(Pass))
}

func CheckEmail(Email string) bool {
	_, err := mail.ParseAddress(Email)
	return err == nil
}

func checkISBN(ISBN string) bool {
	if len(ISBN) != 18 {
		return false
	}
	regexp := regexp.MustCompile(`^[0-9]{3}\-[0-9]{1,8}\-[0-9]{1,8}\-[0-9]{1,8}\-[0-9X]$`)
	return regexp.Match([]byte(ISBN))
}

func checkPrice(Price string) bool {
	regexp := regexp.MustCompile(`^-?\d{1,10}(\.\d{1,2})?$`)
	return regexp.Match([]byte(Price))
}

func checkDesc(Desc string) bool {
	regexp := regexp.MustCompile(`^([A-яЁё\,\(\)\!\?\"]|\d)?$`)
	return regexp.Match([]byte(Desc))
}

func checkPublisher(Publisher string) bool {
	regexp := regexp.MustCompile(`^(p{L}|\w){1,40}$`)
	return regexp.Match([]byte(Publisher))
}

func checkDispAuthor(Author string) bool {
	regexp := regexp.MustCompile(`^[A-яЁё.]{1,20}$`)
	return regexp.Match([]byte(Author))
}

//`^-?\d{1,10}(\.\d{1,2})?$`
