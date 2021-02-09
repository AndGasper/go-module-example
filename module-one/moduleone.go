package moduleone

import "fmt"
// ModuleOne top level
func ModuleOne(name string) string {
	message:= fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}