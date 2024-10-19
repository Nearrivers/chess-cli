package piece

type EmptySquare struct{}

func (es EmptySquare) Display() string {
	return ""
}
