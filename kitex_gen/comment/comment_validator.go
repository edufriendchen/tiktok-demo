// Code generated by Validator v0.1.4. DO NOT EDIT.

package comment

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *Comment) IsValid() error {
	if p.User != nil {
		if err := p.User.IsValid(); err != nil {
			return fmt.Errorf("filed User not valid, %w", err)
		}
	}
	return nil
}
func (p *ActionRequest) IsValid() error {
	_src := []string{string("1"), string("2")}
	var _exist bool
	for _, src := range _src {
		if p.ActionType == src {
			_exist = true
			break
		}
	}
	if !_exist {
		return fmt.Errorf("field ActionType in rule failed, current value: %v", p.ActionType)
	}
	return nil
}
func (p *ActionResponse) IsValid() error {
	if p.Comment != nil {
		if err := p.Comment.IsValid(); err != nil {
			return fmt.Errorf("filed Comment not valid, %w", err)
		}
	}
	return nil
}
func (p *CommentRequest) IsValid() error {
	return nil
}
func (p *CommentResponse) IsValid() error {
	return nil
}
