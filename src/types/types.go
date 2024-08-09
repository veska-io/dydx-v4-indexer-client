package types

type DYDXError struct {
	Value    string `json:"value"`
	Msg      string `json:"msg"`
	Param    string `json:"param"`
	Location string `json:"location"`
}
