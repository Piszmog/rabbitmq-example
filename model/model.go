package model

type CloudFoundryEnvironment struct {
    CloudAMQP []CloudAMQP `json:"cloudamqp"`
}

type CloudAMQP struct {
    Credentials Credentials `json:"credentials"`
}

type Credentials struct {
    Uri string `json:"uri"`
}
