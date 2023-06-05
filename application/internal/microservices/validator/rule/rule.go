package rule

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

var (
	ErrRequiredFieldIsEmpty       = errors.New("required field is empty")
	ErrWrongTypeInt               = errors.New("type validation: Integer")
	ErrWrongTypeUUID              = errors.New("type validation: UUID")
	ErrIntegerNotLessOrEqual      = errors.New("given value is not less or equal than what is expected")
	ErrValueDoesNotContainPattern = errors.New("value does not contain pattern")
)

type Rule []func(string) error

func New() *Rule {
	return &Rule{}
}

func (vp *Rule) Required() *Rule {
	*vp = append(*vp, func(value string) error {
		if value == "" {
			return ErrRequiredFieldIsEmpty
		} else {
			return nil
		}
	})
	return vp
}

func (vp *Rule) UUID() *Rule {
	*vp = append(*vp, func(value string) error {
		var _, err = uuid.Parse(value)
		if err != nil {
			return errors.Join(err, ErrWrongTypeUUID)
		}
		return nil
	})
	return vp
}

func (vp *Rule) LessOrEqual(number int) *Rule {
	*vp = append(*vp, func(s string) error {
		var integer, err = strconv.ParseInt(s, 10, 0)
		if err != nil {
			return ErrWrongTypeInt
		}
		if !(int(integer) <= number) {
			return ErrIntegerNotLessOrEqual
		}
		return nil
	})
	return vp
}

func (vp *Rule) Match(pattern string) *Rule {
	var rgx, err = regexp.Compile(pattern)
	if err != nil {
		panic("could not compile regex for rule: " + pattern)
	}
	*vp = append(*vp, func(s string) error {
		if ok := rgx.Match([]byte(s)); !ok {
			return ErrValueDoesNotContainPattern
		}
		return nil
	})
	return vp
}

func (vp *Rule) MatchWhole(pattern string) *Rule {
	return vp.Match(fmt.Sprintf("^%s$", pattern))
}

func (vp *Rule) Test(value string) (errs []error) {
	for _, rule := range *vp {
		if err := rule(value); err != nil {
			errs = append(errs, err)
		}
	}
	return
}
