package models


type ResultRequest struct {
    Economic      string `json:"economic"`
    Diplomatic    string `json:"diplomatic"`
    Civil         string `json:"civil"`
    Societal      string `json:"societal"`
}

type ResultResponse struct {
    Results []ResultRequest
}
