package reader

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"

	"go.uber.org/zap"
)

// Is ripgrep on the system
var rgEnabled bool = false

// The regex which searches for the quickkey extra description
const (
	searchRegex string = `[%s] \[QK\] (.*)$`
	// contentRegex string = `[%s] \[QK\] (?<content>.*)$`
	// lineRegex    string = `^(?<linenum>\d*)`
	grepLineRegex string = `^(?<linenum>\d*):.*%s+ \[QK\] (?<content>.*)$`
)

// Example results
// 31:# QK: Save all open buffers
// 33:C-j = "save_selection"          # QK: Save open buffer

var lineNumRegex = regexp.MustCompile(`^(?<linenum>\d*)`)

// Should return a map with the key being the line number and
// the contents being the help description.
// Possibly store the max line number for better sorting later
func grepComments(commentToken string, path string) (map[int]string, error) {
	localLogger := logger.With(
		zap.String("area", "grepping QK comments"),
		zap.String("filepath", path),
	)

	rgx := fmt.Sprintf(searchRegex, commentToken)
	var cmd *exec.Cmd
	if rgEnabled {
		cmd = exec.Command("rg", "-n", rgx, path)
	} else {
		cmd = exec.Command("grep", "-En", rgx, path)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		localLogger.Error("error declaring StdoutPipe for grep command")
		return nil, err
	}

	scanner := bufio.NewScanner(stdout)

	err = cmd.Start()
	if err != nil {
		localLogger.Error("error running grep command")
		return nil, err
	}

	localLogger.Info("running grep command to look for QK comments")

	results := make(map[int]string)

	for scanner.Scan() {
		lineRegex := fmt.Sprintf(grepLineRegex, commentToken)
		re := regexp.MustCompile(lineRegex)
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		for _, match := range matches {
			lineNumber, err := strconv.Atoi(match[1])
			if err != nil {
				localLogger.Warn("failed to convert string to int")
				continue
			}
			results[lineNumber] = match[2]
			localLogger.Debug("matched QK comment with grep",
				zap.Int("line_number", lineNumber),
				zap.String("comment", match[2]),
			)
		}
	}

	localLogger.Info("successfully parsed file for QK comments")
	return results, nil
}
