package errorhandlers

import "fmt"

type UserNotFoundError struct {
	UserID int
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("User with ID %d not found", e.UserID)
}
