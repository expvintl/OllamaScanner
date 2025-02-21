package models

type ModelInfo struct {
	Name    string `json:"name"`
	Size    uint64 `json:"size"`
	Details struct {
		ParameterSize     string `json:"parameter_size"`
		QuantizationLevel string `json:"quantization_level"`
	} `json:"details"`
}
type OllamaInfo struct {
	Host   string      `json:"host"`
	Port   int         `json:"port"`
	Models []ModelInfo `json:"models"`
}
