package main

import (
	"fmt"
	"time"
)

type Book struct {
	Name string
}

type Member struct {
	Name string
}

type Library struct {
	Books []struct {
		Book         *Book
		isAvailable  bool
		isReturned   bool
		CheckedOutAt time.Time
		ReturnedAt   time.Time
		Borrower     *Member
	}
	Members []*Member
}

func (l *Library) toStringAll() string {
	s := ""
	for k, v := range l.Books {
		s += fmt.Sprintf("%d: \n\tBook Name: %s\n\tAvailable: %v\n", k, v.Book.Name, v.isAvailable)
		if v.isAvailable {
			if v.isReturned {
				s += fmt.Sprintf("\tReturnedAt: %s\n", v.ReturnedAt.Format("2006-01-02"))
			}
		} else {
			s += fmt.Sprintf("\tCheckOutAt: %s\n\tBorrower: %s\n", v.CheckedOutAt.Format("2006-01-02"), v.Borrower.Name)
		}
	}
	return s
}

func (l *Library) printAll() {
	fmt.Println(l.toStringAll())
}

func (l *Library) toStringCheckedOut() string {
	s := ""
	for k, v := range l.Books {
		if !v.isAvailable {
			s += fmt.Sprintf("%d: \n\tBook Name: %s\n\tAvailable: %v\n", k, v.Book.Name, v.isAvailable)
			s += fmt.Sprintf("\tCheckOutAt: %s\n\tBorrower: %s\n", v.CheckedOutAt.Format("2006-01-02"), v.Borrower.Name)
		}
	}
	return s
}

func (l *Library) printCheckedOut() {
	fmt.Println(l.toStringCheckedOut())
}

func (l *Library) AddBook(b *Book) {
	l.Books = append(l.Books, struct {
		Book         *Book
		isAvailable  bool
		isReturned   bool
		CheckedOutAt time.Time
		ReturnedAt   time.Time
		Borrower     *Member
	}{
		b, true, false, time.Time{}, time.Time{}, &Member{}})
}

func (l *Library) AddMember(m *Member) {
	l.Members = append(l.Members, m)
}

func (l *Library) CheckOutBook(b *Book, m *Member) {
	for k, v := range l.Books {
		if v.Book == b {
			l.Books[k].Borrower = m
			l.Books[k].CheckedOutAt = time.Now()
			l.Books[k].isReturned = false
			l.Books[k].isAvailable = false
		}
	}
}

func (l *Library) CheckInBook(b *Book, m *Member) {
	for k, v := range l.Books {
		if v.Book == b {
			l.Books[k].isReturned = true
			l.Books[k].ReturnedAt = time.Now()
			l.Books[k].isAvailable = true
		}
	}
}

func main() {
	l := new(Library)

	b1 := Book{Name: "Kitap1"}
	b2 := Book{Name: "Kitap2"}
	b3 := Book{Name: "Kitap3"}
	b4 := Book{Name: "Kitap4"}

	fmt.Println("Add Books")
	l.AddBook(&b1)
	l.AddBook(&b2)
	l.AddBook(&b3)
	l.AddBook(&b4)

	fmt.Println("Print All Books")
	l.printAll()

	m1 := Member{Name: "Member1"}
	m2 := Member{Name: "Member2"}
	m3 := Member{Name: "Member3"}

	fmt.Println("Add Members")
	l.AddMember(&m1)
	l.AddMember(&m2)
	l.AddMember(&m3)

	fmt.Println("Print All Books")
	l.printAll()

	fmt.Println("Check Out a Book")
	l.CheckOutBook(&b2, &m2)

	fmt.Println("Print Checked Out Books")
	l.printCheckedOut()

	fmt.Println("Check In a Book")
	l.CheckInBook(&b2, &m2)

	fmt.Println("Print Checked Out Books")
	l.printCheckedOut()

	fmt.Println("Print All Books")
	l.printAll()

}
