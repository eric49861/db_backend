package gogpt

type ModelRequest struct {
	Model  string `json:"model"`   // 要调用的模型名称
	Key    string `json:"api_key"` // api_key
	Prompt string `json:"prompt"`  // 对话内容
}

