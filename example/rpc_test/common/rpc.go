package common

import "errors"

type Args struct {
	A float32
	B float32
}

type Result struct {
	Value float32
}

type MathService struct {
}

func (s *MathService) Sum(args *Args, result *Result) error {
	result.Value = args.A + args.B
	return nil
}

func (s *MathService) Divide(args *Args, result *Result) error {
	if args.B == 0 {
		return errors.New("args b is 0")
	}

	result.Value = args.A / args.B
	return nil
}
