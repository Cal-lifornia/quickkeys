package reader

import (
	"regexp"
)

// Is ripgrep on the system
var rgEnabled bool = false

// The regex which searches for the quickkey extra description
const (
	searchRegex  string = "[%s] QK: (.*)$"
	contentRegex string = "[%s] QK: (?<content>.*)$"
)

var lineNumRegex = regexp.MustCompile(`^(?<linenum>\d*)`)

// Should return a map with the key being the line number and
// the contents being the help description.
// Possibly store the max line number for better sorting later
// func grepComments(commentToken string, path string) (map[int]string, error) {
// 	rgx := fmt.Sprintf(searchRegex, commentToken)
// 	var cmd *exec.Cmd
// 	if rgEnabled {
// 		cmd = exec.Command("rg", "-n", rgx, path)
// 	} else {
// 		cmd = exec.Command("grep", "-En", rgx, path)
// 	}
// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		logger.Error("error running grep command")
// 		return nil, err
// 	}
// 	err = cmd.Run()
// 	if err != nil {
// 		logger.Error("error running grep command")
// 		return nil, err
// 	}

// 	return nil, nil
// }
