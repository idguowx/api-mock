package mock_request

import "api-mock/util"

type PatternField struct {
	Value  string
	Ignore util.NumBool
}

type RequestPattern struct {
	Method         *PatternField
	UriReg         *PatternField
	ContentTypeReg *PatternField
	BodyReg        *PatternField
}

func MakeRequestPattern() *RequestPattern {
	return &RequestPattern{}
}

func (r *RequestPattern) SetMethodPattern(ignore util.NumBool, val string) *RequestPattern {
	r.Method = &PatternField{
		Value:  val,
		Ignore: ignore,
	}

	return r
}

func (r *RequestPattern) SetUriRegPattern(ignore util.NumBool, val string) *RequestPattern {
	r.UriReg = &PatternField{
		Value:  val,
		Ignore: ignore,
	}

	return r
}

func (r *RequestPattern) SetContentTypeRegPattern(ignore util.NumBool, val string) *RequestPattern {
	r.ContentTypeReg = &PatternField{
		Value:  val,
		Ignore: ignore,
	}

	return r
}

func (r *RequestPattern) SetBodyRegPattern(ignore util.NumBool, val string) *RequestPattern {
	r.BodyReg = &PatternField{
		Value:  val,
		Ignore: ignore,
	}

	return r
}
