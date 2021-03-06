package processing

import (
	"errors"

	"github.com/move-ton/ton-client-go/domain"
)

type processing struct {
	config domain.Config
	client domain.ClientGateway
}

// NewProcessing ...
func NewProcessing(
	config domain.Config,
	client domain.ClientGateway,
) domain.ProcessingUseCase {
	return &processing{
		config: config,
		client: client,
	}
}

//change callback??????

// SendMessage method processing.send_message
func (p *processing) SendMessage(pOSM domain.ParamsOfSendMessage, callback domain.EventCallback) (*domain.ResultOfSendMessage, error) {
	if pOSM.SendEvents && callback == nil {
		return nil, errors.New("Don't find callback")
	}

	responses, err := p.client.Request("processing.send_message", pOSM)
	if err != nil {
		return nil, err
	}

	if pOSM.SendEvents {
		responses = domain.DynBufferForResponses(responses)
	}

	result := &domain.ResultOfSendMessage{}
	return result, domain.HandleEvents(responses, callback, result)
}

// WaitForTransaction method processing.wait_for_transaction
func (p *processing) WaitForTransaction(pOWFT domain.ParamsOfWaitForTransaction, callback domain.EventCallback) (*domain.ResultOfProcessMessage, error) {
	if pOWFT.SendEvents && callback == nil {
		return nil, errors.New("Don't find callback")
	}

	responses, err := p.client.Request("processing.wait_for_transaction", pOWFT)
	if err != nil {
		return nil, err
	}

	if pOWFT.SendEvents {
		responses = domain.DynBufferForResponses(responses)
	}

	result := &domain.ResultOfProcessMessage{}
	return result, domain.HandleEvents(responses, callback, result)
}

// ProcessMessage method processing.process_message
func (p *processing) ProcessMessage(pOPM domain.ParamsOfProcessMessage, callback domain.EventCallback) (*domain.ResultOfProcessMessage, error) {
	if pOPM.SendEvents && callback == nil {
		return nil, errors.New("Don't find callback")
	}

	responses, err := p.client.Request("processing.process_message", pOPM)
	if err != nil {
		return nil, err
	}

	if pOPM.SendEvents {
		responses = domain.DynBufferForResponses(responses)
	}

	result := &domain.ResultOfProcessMessage{}
	return result, domain.HandleEvents(responses, callback, result)
}
