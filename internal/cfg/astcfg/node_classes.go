package astcfg

type NodeTypeClass int

const (
	Expression = NodeTypeClass(iota)
	Statement
	Declaration
	Specification
	TypeDeclaration
)
