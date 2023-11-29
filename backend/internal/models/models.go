package models


type ResultRequest struct {
    Id            int    `json:"id"`
    Economic      string `json:"economic"`
    Diplomatic    string `json:"diplomatic"`
    Civil         string `json:"civil"`
    Societal      string `json:"societal"`
    EconomicLabel      string `json:"economicLabel"`
    DiplomaticLabel    string `json:"diplomaticLabel"`
    CivilLabel         string `json:"civilLabel"`
    SocietalLabel      string `json:"societalLabel"`
    ClosestMatch        string `json:"closest"`
}

type ResultResponse struct {
    Results []ResultRequest
}

