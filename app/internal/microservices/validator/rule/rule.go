package rule

import (
	"fmt"
	"regexp"
	"strconv"
	"tde/internal/microservices/errors/bucket"
	"tde/internal/microservices/errors/detailed"
	"tde/internal/microservices/errors/loggable"
	"tde/internal/microservices/errors/tagged"

	"github.com/google/uuid"
)

var (
	ErrRequiredFieldIsEmpty       = detailed.New("Field can not be empty", "ErrRequiredFieldIsEmpty@Rule")
	ErrWrongTypeInt               = detailed.New("Field value is expected to be an integer.", "ErrWrongTypeInt@Rule")
	ErrWrongTypeUUID              = detailed.New("Field value is expected to be an UUID", "ErrWrongTypeUUID@Rule")
	ErrIntegerNotLessOrEqual      = detailed.New("Field value is bigger than expected", "ErrIntegerNotLessOrEqual@Rule")
	ErrValueDoesNotContainPattern = detailed.New("Field value is different that expected", "ErrValueDoesNotContainPattern@Rule")
)

type Rule []func(string) loggable.Loggable

func New() *Rule {
	return &Rule{}
}

func (r *Rule) Required() *Rule {
	*r = append(*r, func(s string) loggable.Loggable {
		if s == "" {
			return ErrRequiredFieldIsEmpty
		} else {
			return nil
		}
	})
	return r
}

func (r *Rule) UUID() *Rule {
	*r = append(*r, func(s string) loggable.Loggable {
		var _, err = uuid.Parse(s)
		if err != nil {
			return tagged.New(ErrWrongTypeUUID, detailed.AddBase(err, ""))
		}
		return nil
	})
	return r
}

func (r *Rule) LessOrEqual(number int) *Rule {
	*r = append(*r, func(s string) loggable.Loggable {
		var integer, err = strconv.ParseInt(s, 10, 0)
		if err != nil {
			return ErrWrongTypeInt
		}
		if !(int(integer) <= number) {
			return ErrIntegerNotLessOrEqual
		}
		return nil
	})
	return r
}

func (r *Rule) Match(pattern string) *Rule {
	var rgx, err = regexp.Compile(pattern)
	if err != nil {
		panic("could not compile regex for rule: " + pattern)
	}
	*r = append(*r, func(s string) loggable.Loggable {
		if ok := rgx.Match([]byte(s)); !ok {
			return ErrValueDoesNotContainPattern
		}
		return nil
	})
	return r
}

func (r *Rule) MatchWhole(pattern string) *Rule {
	return r.Match(fmt.Sprintf("^%s$", pattern))
}

func (r *Rule) Test(value string) *bucket.Bucket {
	var bucket = &bucket.Bucket{}
	for _, rule := range *r {
		if err := rule(value); err != nil {
			bucket.Add(err)
		}
	}
	return bucket
}
