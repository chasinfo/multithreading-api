package dtos

type BrasilApiCep struct {
	Cep        string `json:"cep"`
	UF         string `json:"state"`
	Localidade string `json:"city"`
	Bairro     string `json:"neighborhood"`
	Logradouro string `json:"street"`
	Service    string `json:"service"`
}
