// Package techpalace contains tools to format text message for a billboard
package techpalace

import (
    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border of n stars to a welcome message.
func AddBorder(welcomeMsg string, n int) string {
    var border = strings.Repeat("*", n)
    return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    oldMsg = strings.ReplaceAll(oldMsg, "*", "")
    return strings.TrimSpace(oldMsg)
}
