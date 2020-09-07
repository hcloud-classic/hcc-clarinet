package errors

import (
	"errors"
	"log"
	"strconv"
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
	undefinedError
	argumentError
	jsonMarshalError
	jsonUnmarshalError
	requestError  // send Request fail
	responseError // get Response fail or has error
	sendError     // send error to client
	receiveError  // get error as result from server
	parsingError
	tokenExpired
	dbOperationFail
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
	"DB operationfail -> ",
}

var errlogger *log.Logger

func SetErrLogger(l *log.Logger) {
	errlogger = l
}

/*    HCCERROR    */

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

func (e HccError) Text() string {
	return e.ErrText
}

func (e HccError) ToString() string {
	m := e.ErrCode / 10000
	f := e.ErrCode % 10000 / 1000
	a := e.ErrCode % 1000

	return "[" + middleWareList[m] + "] Code :" + strconv.FormatUint(e.ErrCode, 10) + " (" + functionList[f] + ") " + actionList[a] + " " + e.ErrText
}

func (e HccError) Println() {
	errlogger.Println(e.ToString())
}

func (e HccError) Fatal() {
	errlogger.Fatal(e.ToString())
}

/*    HCCERRORSTACK    */

type HccErrorStack []HccError

func NewHccErrorStack(errList ...*HccError) *HccErrorStack {
	es := HccErrorStack{HccError{ErrCode: 0, ErrText: ""}}

	for _, err := range errList {
		es.Push(err)
	}
	return &es
}

func (es *HccErrorStack) Len() int {
	return es.len() - 1
}

func (es *HccErrorStack) len() int {
	return len(*es)
}

func (es *HccErrorStack) Pop() *HccError {
	l := es.len()
	if l > 1 {
		err := (*es)[l-1]
		*es = (*es)[:l-1]
		return &err
	}
	return nil
}

func (es *HccErrorStack) Push(err *HccError) {
	*es = append(*es, *err)
}

// Dump() will clean stack
func (es *HccErrorStack) Dump() *HccError {
	var firstErr *HccError = nil
	if es.Len() == 0 {
		return nil
	}
	errlogger.Printf("------ [Dump Error Stack] ------\n")
	errlogger.Printf("Stack Size : %v\n", es.Len())
	firstErr = es.Pop()
	firstErr.Println()
	for err := es.Pop(); err != nil; err = es.Pop() {
		err.Println()
	}
	errlogger.Println("--------- [ End Here ] ---------")
	return firstErr
}

func (es *HccErrorStack) ConvertReportForm() *HccErrorStack {
	for idx := 1; idx < es.len(); idx++ {
		(*es)[idx].ErrText = (*es)[idx].ToString()
	}
	return es
}
