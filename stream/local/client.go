package local

import (
	"io"

	"github.com/jacksontj/dataman/stream"
)

func NewClientStream(resultsChan chan stream.Result, errorChan chan error) stream.ClientStream {
	stream := &ClientStream{
		resultsChan: resultsChan,
		errorChan:   errorChan,
	}

	return stream
}

type ClientStream struct {
	resultsChan chan stream.Result
	errorChan   chan error
}

func (s *ClientStream) Close() error {
	return nil
}

func (s *ClientStream) Recv() (stream.Result, error) {
	// TODO: implement this cleaner, its a bit of a mess since we want specific
	// priorities on channel reading
	for {
		// we want to get results first if we have them
		select {
		case result, ok := <-s.resultsChan:
			if ok {
				return result, nil
			} else {
				// if the result chan closed, we want to drain the errors if we have any
				select {
				case err, ok := <-s.errorChan:
					if ok {
						return nil, err
					}
				default:
				}
				return nil, io.EOF
			}
			continue
		default:
		}

		select {
		case result, ok := <-s.resultsChan:
			if ok {
				return result, nil
			} else {
				// if the result chan closed, we want to drain the errors if we have any
				select {
				case err, ok := <-s.errorChan:
					if ok {
						return nil, err
					}
				default:
				}
				return nil, io.EOF
			}
		case err, ok := <-s.errorChan:
			if ok {
				return nil, err
			}
		}
	}
}
