package gamecore

type BackendProcessorRequest struct {
	Msg      string `json:"msg"`
	PlayerId string `json:"player-id"`
}
