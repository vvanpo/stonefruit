package titian

import (
	"testing"
)

func stubContact() *Contact {
	app := New(Config{})
	email, _ := NewEmailAddress(validEmails[0])
	contact, _ := app.contacts.New(email)
	return contact
}

// Tests setting an e-mail address.
func TestContactEmailAddress(t *testing.T) {
	contact := stubContact()
	email2, _ := NewEmailAddress(validEmails[1])

	if contact.EmailAddress().String() != validEmails[0] {
		t.Error("Failure to set initial e-mail address")
	}

	contact.SetPrimaryEmailAddress(email2)

	if contact.EmailAddress().String() != validEmails[1] {
		t.Error("Failure to set primary e-mail address")
	}
}

//
func TestContactVerifyEmailAddress(t *testing.T) {
	contact := stubContact()

	for _, address := range validEmails[1:] {
		email, _ := NewEmailAddress(address)
		err := contact.VerifyEmailAddress(email)

		if err != nil {
			t.Errorf("Failed to verify e-mail address '%v'", email)
		}
	}

	if contact.VerifyEmailAddress(contact.email) == nil {
		t.Fatalf("Failed to identify duplicate e-mail address")
	}

	list := contact.VerifiedEmails()
	if len(list) != len(validEmails) {
		t.Errorf("Verified e-mails don't match inputs:\n\tExpected: %v\n\tActual: %v", validEmails, list)
	}
}

