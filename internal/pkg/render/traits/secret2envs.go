package traits

type Secret2Envs struct {
	SecretName string `json:"secretName"`
}

func NewSecret2Envs(secretName string) *Secret2Envs {
	return &Secret2Envs{secretName}
}

func (m *Secret2Envs) GetName() string {
	return "secre"
}

func (m *Secret2Envs) GetType() string {
	return "k-secret2envs"
}
