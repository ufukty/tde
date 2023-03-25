package slot_manager

import (
	"tde/internal/folders/copy_module"
	"tde/internal/folders/types"
	"tde/internal/utilities"
	"tde/models/in_program_models"

	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Config struct {
	OriginalModule types.AbsolutePath // copy of original module that is filtered from unnecessary dirs and tester_package mounted into tested package

	Package       types.InModulePath // eg. .../examples/word_reverse
	PackageImport string             // eg. examples/word_reverse/word_reverse

	ImplementationFile types.InModulePath // eg. .../examples/word_reverse/word_reverse.go
	TestFile           types.InModulePath // eg. .../examples/word_reverse/word_reverse_tde.go
	TestName           string             // eg. TDE_WordReverse
}

type slots struct {
	free     []types.SlotPath
	assigned map[in_program_models.CandidateID]types.SlotPath
}

type Session struct {
	config *Config
	tmp    types.TempPath // reserved in instantiation. all
	slots  slots
}

// returns eg. 65/36/f1/24/b8/56/5a/ad/8c/cc/22/ea/c3/7d/8e/63
func (s *Session) genNewSlotPath() (types.SlotPath, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", errors.New("can't create a uuid")
	}
	basename := strings.Join(utilities.StringFold(strings.ReplaceAll(uuid.String(), "-", ""), 2), "/")
	path := filepath.Join(string(s.tmp), basename)
	return types.SlotPath(path), nil
}

func (s *Session) createMainFolder() error {
	path, err := os.MkdirTemp(os.TempDir(), "tde.runner-folders.*")
	if err != nil {
		return errors.New("failed to create main folder for slot_manager in temp directory")
	}
	s.tmp = types.TempPath(path)
	return nil
}

func (s *Session) createModuleDuplicate() error {
	newSlotPath, err := s.genNewSlotPath()
	if err != nil {
		return errors.Wrap(err, "genNewtypes.SlotPath")
	}

	// fmt.Println("Original module duplicated:", path)
	err = copy_module.Module(string(s.config.OriginalModule), string(newSlotPath), true, copy_module.DefaultSkipDirs)
	if err != nil {
		return errors.Wrap(err, "copy_module.Module")
	}

	s.slots.free = append(s.slots.free, newSlotPath)
	return nil
}

func (s *Session) assignCandidateToASlot(candidateID in_program_models.CandidateID) (slot types.SlotPath) {
	s.slots.free, slot = slicePop(s.slots.free)
	s.slots.assigned[candidateID] = slot
	// fmt.Println("assigning", candidateID, "to folder", choosen)
	return
}

func (s *Session) printToFile(candidate *in_program_models.Candidate) {
	slot := s.slots.assigned[candidate.UUID]
	s.tmp.FindInModulePath(slot, s.config.ImplementationFile)
}

func (s *Session) placeCandidate(candidate *in_program_models.Candidate) {
	if len(s.slots.free) == 0 {
		s.createModuleDuplicate()
	}
	s.assignCandidateToASlot(candidate.UUID)
	s.printToFile(candidate)
}

func NewSession(config *Config) *Session {
	s := Session{config: config}
	s.createMainFolder()
	return &s
}

func (s *Session) PlaceCandidatesIntoSlots(candidates []*in_program_models.Candidate) {
	for _, candidate := range candidates {
		s.placeCandidate(candidate)
	}
}
