package loki

type Response struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type Data struct {
	ResultType string   `json:"resultType"`
	Result     []Result `json:"result"`
}

type Result struct {
	Stream Stream     `json:"stream"`
	Values [][]string `json:"values"`
}

type Stream struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}
