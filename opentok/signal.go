package opentok

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const maxErrBodyBytes = 256

// SignalData defines the type and data of signal
type SignalData struct {
	// The type of the signal.
	// This is a string value that clients can filter on when listening for signals
	Type string `json:"type"`

	// The data of the signal
	Data string `json:"data"`
}

type SignalDataError struct {
	Code    int
	Message string
}

func (e *SignalDataError) Error() string {
	return fmt.Sprintf("Tokbox error: code: %d; message: %s", e.Code, e.Message)
}

func newSignalDataError(code int, msg string) *SignalDataError {
	return &SignalDataError{Code: code, Message: msg}
}

// SendSessionSignal send signals to all participants in an active OpenTok session.
func (ot *OpenTok) SendSessionSignal(sessionID string, data *SignalData) error {
	return ot.SendSessionSignalContext(context.Background(), sessionID, data)
}

// SendSessionSignalContext uses ctx for HTTP requests.
func (ot *OpenTok) SendSessionSignalContext(ctx context.Context, sessionID string, data *SignalData) error {
	if sessionID == "" {
		return fmt.Errorf("Signal cannot be sent without a session ID")
	}

	jsonStr, _ := json.Marshal(data)

	// Create jwt token
	jwt, err := ot.jwtToken(projectToken)
	if err != nil {
		return err
	}

	endpoint := ot.apiHost + projectURL + "/" + ot.apiKey + "/session/" + sessionID + "/signal"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-OPENTOK-AUTH", jwt)
	req.Header.Add("User-Agent", SDKName+"/"+SDKVersion)

	res, err := ot.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		data := struct {
			Message string `json:"message"`
		}{}

		if err := json.NewDecoder(io.LimitReader(res.Body, maxErrBodyBytes)).Decode(&data); err != nil {
			return fmt.Errorf("Error decoding response from Tokbox: statusCode: %d; %w", res.StatusCode, err)
		}

		return newSignalDataError(res.StatusCode, data.Message)
	}

	return nil
}

// SendConnectionSignal send signals to a specific client in an active OpenTok session.
func (ot *OpenTok) SendConnectionSignal(sessionID, connectionID string, data *SignalData) error {
	return ot.SendConnectionSignalContext(context.Background(), sessionID, connectionID, data)
}

// SendConnectionSignalContext uses ctx for HTTP requests.
func (ot *OpenTok) SendConnectionSignalContext(ctx context.Context, sessionID, connectionID string, data *SignalData) error {
	if sessionID == "" {
		return fmt.Errorf("Signal cannot be sent without a session ID")
	}

	if connectionID == "" {
		return fmt.Errorf("Signal cannot be sent without a connection ID")
	}

	jsonStr, _ := json.Marshal(data)

	// Create jwt token
	jwt, err := ot.jwtToken(projectToken)
	if err != nil {
		return err
	}

	endpoint := ot.apiHost + projectURL + "/" + ot.apiKey + "/session/" + sessionID + "/connection/" + connectionID + "/signal"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-OPENTOK-AUTH", jwt)
	req.Header.Add("User-Agent", SDKName+"/"+SDKVersion)

	res, err := ot.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		data := struct {
			Message string `json:"message"`
		}{}

		if err := json.NewDecoder(io.LimitReader(res.Body, maxErrBodyBytes)).Decode(&data); err != nil {
			return fmt.Errorf("Error decoding response from Tokbox: statusCode: %d; %w", res.StatusCode, err)
		}

		return newSignalDataError(res.StatusCode, data.Message)
	}

	return nil
}
