package errors

import (
	"errors"
	"log"
	"strconv"

	"github.com/golang-collections/collections/stack"
)

/*** Match Enum squence with xxxList ***/
const (
	// code for MiddleWare
	cello uint64 = (1 + iota) * 10000
	clarinet
	flute
	harp
	oboe
	piano
	piccolo
	viola
	violin
	violinNoVNC
	violinScheduler
)

var middleWareList = [...]string{"", "Cello", "Clarinet", "Flute", "Harp", "Oboe", "Piano", "Piccolo", "Viola", "Violin", "NoVNC", "Scheduler"}

const (
	internal uint64 = (1 + iota) * 1000 // lib
	driver                              // driver
	graphql                             // action
	grpc
	sql
	rabbitmq
)

var functionList = [...]string{"", "Internal", "Driver", "GraphQL", "Grpc", "SQL", "RabbitMQ"}

const (
	initFail uint64 = 1 + iota
	connectionFail
	argumentError
	jsonMarshalError
	jsonUnmarshalError
	requestError  // send Request fail
	responseError // get Response fail or has error
	sendError     // send error to client
	receiveError  // get error as result from server
	parsingError
	tokenExpired
)

var actionList = [...]string{
	"",
	"Initialize fail -> ",
	"Connection fail -> ",
	"Argumnet error -> ",
	"JSON marshal fail -> ",
	"JSON unmarshal fail -> ",
	"Request error -> ",
	"Response error -> ",
	"Send error -> ",
	"Receive error -> ",
	"Parsing error -> ",
	"Token Expired -> ",
}

var errlogger *log.Logger

func SetErrLogger(l *log.Logger) {
	errlogger = l
}

type HccError struct {
	ErrCode uint64 // decimal error code
	ErrText string // error string
}

func NewHccError(errorCode uint64, errorText string) *HccError {
	return &HccError{
		ErrText: errorText,
		ErrCode: errorCode,
	}
}

func (e HccError) New() error {
	return errors.New(e.ToString())
}

func (e HccError) Error() string {
	return e.ToString()
}

func (e HccError) Code() uint64 {
	return e.ErrCode
}

func (e HccError) ToString() string {
	m := e.ErrCode / 10000
	f := e.ErrCode % 10000 / 1000
	a := e.ErrCode % 1000

	return "[" + middleWareList[m] + "] " + functionList[f] + ": " + actionList[a] + strconv.FormatUint(e.ErrCode, 10) + " " + e.ErrText
}

func (e HccError) Println() {
	errlogger.Println(e.ToString())
}

func (e HccError) Fatal() {
	errlogger.Fatal(e.ToString())
}

type HccErrorStack struct {
	errStack *stack.Stack
}

func NewHccErrorStack(errList ...*HccError) *HccErrorStack {
	es := HccErrorStack{errStack: stack.New()}

	for _, err := range errList {
		es.Push(err)
	}
	return &es
}

func (es HccErrorStack) Len() int {
	return es.errStack.Len()
}

func (es HccErrorStack) Pop() *HccError {
	if err := es.errStack.Pop(); err != nil {
		return NewHccError(err.(HccError).ErrCode, err.(HccError).ErrText)
	}
	return NewHccError(0, "")
}

func (es HccErrorStack) Push(err *HccError) {
	es.errStack.Push(*err)
}

func (es HccErrorStack) Dump() *HccError {
	var firstErr *HccError = nil
	if es.Len() == 0 {
		return nil
	}
	errlogger.Printf("------ [Dump Error Stack] ------\n")
	errlogger.Printf("Stack Size : %v\n", es.Len())
	firstErr = es.Pop()
	firstErr.Println()
	for err := es.Pop(); err.Code() != 0; err = es.Pop() {
		err.Println()
	}
	errlogger.Println("--------- [ End Here ] ---------")
	return firstErr
}
