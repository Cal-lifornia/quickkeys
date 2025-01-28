package reader

// Is ripgrep on the system
var rgEnabled bool = false

// The regex which searches for the quickkey extra description
var searchRegex = "[%s] QK: (.*)$"

const lineNumRegex string = `^(?<linenum>\d*)`

var contentRegex string = "[#] QK: (?<content>.*)$"

// Should return a map with the key being the line number and
// the contents being the help description.
// Possibly store the max line number for better sorting later
func grepCommand()
