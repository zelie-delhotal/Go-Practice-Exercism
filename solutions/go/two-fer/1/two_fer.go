// Package twofer contains a tool to converse with people you share bakery cookies with.
package twofer

// ShareWith returns a string containing what you need to say to the person, depending
// on wether you know their name.
func ShareWith(name string) string {
    if name == "" {
        name = "you"
    }
	return "One for " + name + ", one for me."
}
