package entities

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("123", "john@example.com", "john_doe")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if user.EmailString() != "john@example.com" {
		t.Errorf("Expected email john@example.com, got %s", user.EmailString())
	}

	if user.Role() != RoleReader {
		t.Errorf("Expected default role reader, got %s", user.Role())
	}
}

func TestNewUserWithInvalidEmail(t *testing.T) {
	_, err := NewUser("123", "invalid-email", "john_doe")

	if err == nil {
		t.Error("Expected error for invalid email format")
	}
}

func TestUserEmailDomainLogic(t *testing.T) {
	user, _ := NewUser("123", "john@company.com", "john_doe")

	if !user.IsFromDomain("company.com") {
		t.Error("User should be from company.com domain")
	}

	if user.IsFromDomain("other.com") {
		t.Error("User should not be from other.com domain")
	}
}
