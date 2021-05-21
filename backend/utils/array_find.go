package utils

// simple helper function, takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1
func Find(slice []string, val string) int {
    for i, item := range slice {
        if item == val {
            return i
        }
    }
    return -1
}
