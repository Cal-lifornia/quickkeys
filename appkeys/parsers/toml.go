package parsers

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

// Majority of this code comes from the Participle TOML example
// found at https://github.com/alecthomas/participle/blob/master/_examples/toml/main.go
// with edits to also parse Comments

type TOML struct {
	Pos     lexer.Position
	Entries []*Entry `@@*`
}

type Entry struct {
	Field   *Field   `  @@`
	Section *Section `| @@`
}

// [TODO] Currently there is no way to know if a comment is inline or not as an inline comment appears
// the same as a normal comment to the Lexer
type Field struct {
	// Field can either be a key value or a comment so that in the resulting
	// it's easy to know what comment belongs to what field by the order.
	Comment  *string   `@Comment`
	KeyValue *KeyValue `| @@`
}

type KeyValue struct {
	Key   string `@Ident "="`
	Value *Value `@@`
}

type Value struct {
	String   *string  `  @String`
	DateTime *string  `| @DateTime`
	Date     *string  `| @Date`
	Time     *string  `| @Time`
	Bool     *bool    `| (@"true" | "false")`
	Number   *float64 `| @Number`
	List     []*Value `| ("[" ( @@ ","? ( @@  ","?)* )? "]")`
}

type Section struct {
	Name   string   `"[" @(Ident ( "." Ident )*) "]"`
	Fields []*Field `@@*`
}

var (
	tomlLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"DateTime", `\d\d\d\d-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)?(-\d\d:\d\d)?`},
		{"Date", `\d\d\d\d-\d\d-\d\d`},
		{"Time", `\d\d:\d\d:\d\d(\.\d+)?`},
		{"Ident", `[a-zA-Z_\-][a-zA-Z_0-9\-]*`},

		{"String", `"[^"]*"`},
		{"Number", `[-+]?[.0-9]+\b`},
		{"Punct", `\[|]|[-!()+/*=,]`},
		// Comment is capitalised so that the Lexer properly picks it up
		{"Comment", `#[^\n]+`},
		{"whitespace", `\s+`},
	})
	TomlParser = participle.MustBuild[TOML](
		participle.Lexer(tomlLexer),
		participle.Elide("Comment"),
		participle.Unquote("String"),
	)
)
