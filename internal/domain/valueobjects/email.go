package valueobjects

import (
    "errors"
    "regexp"
    "strings"
)

type Email struct {
    value string
}

func NewEmail(email string) (*Email, error) {
    email = strings.TrimSpace(strings.ToLower(email))
    
    if email == "" {
        return nil, errors.New("email cannot be empty")
    }
    
    // Simple email validation
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(email) {
        return nil, errors.New("invalid email format")
    }
    
    return &Email{value: email}, nil
}

func (e Email) String() string {
    return e.value
}

func (e Email) Domain() string {
    parts := strings.Split(e.value, "@")
    if len(parts) != 2 {
        return ""
    }
    return parts[1]
}

// Value objects are compared by value, not identity
func (e Email) Equals(other Email) bool {
    return e.value == other.value
}
