package check

// Clippy is the check for the ineffassign command
type Clippy struct {
	Dir       string
	Filenames []string
}

// Name returns the name of the display name of the command
func (g Clippy) Name() string {
	return "Clippy"
}

// Weight returns the weight this check has in the overall average
func (g Clippy) Weight() float64 {
	return 1
}

// Percentage returns the percentage of .go files that pass gofmt
func (g Clippy) Percentage() (float64, []FileSummary, error) {
	return GoTool(g.Dir, g.Filenames, []string{"cargo", "clippy"})
}

// Description returns the description of Clippy
func (g Clippy) Description() string {
	return `<a href="https://rust-lang.github.io/rust-clippy/master/index.html">Clippy Lints</a> detects issues in Rust code.`
}
