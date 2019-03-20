package wvaulttypes

type KubernetesLoginRequest struct {
	JWT  string `json:"jwt"`
	Role string `json:"role"`
}
