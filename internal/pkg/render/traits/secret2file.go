package traits

type Secret2File struct {
	SecretName string `json:"secretName"`
}

func NewSecret2File(secretName string) *Secret2File {
	return &Secret2File{secretName}
}

func (m *Secret2File) GetName() string {
	return "secre"
}

func (m *Secret2File) GetType() string {
	return "k-secret2file"
}
