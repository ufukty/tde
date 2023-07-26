package list

import (
	"time"

	"golang.org/x/exp/maps"
)

type Module struct {
	Path       string       // module path
	Query      string       // version query corresponding to this version
	Version    string       // module version
	Versions   []string     // available module versions
	Replace    *Module      // replaced by this module
	Time       *time.Time   // time version was created
	Update     *Module      // available update (with -u)
	Main       bool         // is this the main module?
	Indirect   bool         // module is only indirectly needed by main module
	Dir        string       // directory holding local copy of files, if any
	GoMod      string       // path to go.mod file describing module, if any
	GoVersion  string       // go version used in module
	Retracted  []string     // retraction information, if any (with -retracted or -u)
	Deprecated string       // deprecation message, if any (with -u)
	Error      *ModuleError // error loading module
	Origin     any          // provenance of module
	Reuse      bool         // reuse of old module info is safe
}

type ModuleError struct {
	Err string // the error itself
}

type PackageError struct {
	ImportStack []string
	Pos         string
	Err         string
}

type Package struct {
	Dir           string
	ImportPath    string
	ImportComment string
	Name          string
	Doc           string
	Target        string
	Shlib         string
	Goroot        bool
	Standard      bool
	Stale         bool
	StaleReason   string
	Root          string
	ConflictDir   string
	BinaryOnly    bool
	ForTest       string
	Export        string
	BuildID       string
	Module        *Module
	Match         []string
	DepOnly       bool

	// Source files
	GoFiles           []string
	CgoFiles          []string
	CompiledGoFiles   []string
	IgnoredGoFiles    []string
	IgnoredOtherFiles []string
	CFiles            []string
	CXXFiles          []string
	MFiles            []string
	HFiles            []string
	FFiles            []string
	SFiles            []string
	SwigFiles         []string
	SwigCXXFiles      []string
	SysoFiles         []string
	TestGoFiles       []string
	XTestGoFiles      []string

	// Embedded files
	EmbedPatterns      []string
	EmbedFiles         []string
	TestEmbedPatterns  []string
	TestEmbedFiles     []string
	XTestEmbedPatterns []string
	XTestEmbedFiles    []string

	// Cgo directives
	CgoCFLAGS    []string
	CgoCPPFLAGS  []string
	CgoCXXFLAGS  []string
	CgoFFLAGS    []string
	CgoLDFLAGS   []string
	CgoPkgConfig []string

	// Dependency information
	Imports      []string
	ImportMap    map[string]string
	Deps         []string
	TestImports  []string
	XTestImports []string

	// Error information
	Incomplete bool
	Error      *PackageError
	DepsErrors []*PackageError
}

type Packages map[string]*Package

func (p Packages) First() *Package {
	var keys = maps.Keys(p)
	return p[keys[0]]
}
