//go:build !solution

package game

import (
	pb "datxxx/lab2/quiz/proto"

	"github.com/relab/gorums"
	"google.golang.org/grpc"
)

// Participant structure contains the details of the participant
// name contains the name of the participant
// address is the gorums server address started by the participant,
// quizmaster sends the questions to this server.
// quizDone is used to communicate the completion of the quiz.
// nodeId is received as response to the register and passed as
// input to the MarkReady RPC.
type participant struct {
	name             string
	address          string
	quizDone         chan bool
	nodeId           uint32
	inputChannel     chan int32
	clientConnection *grpc.ClientConn
	client           pb.ParticipantServiceClient
}

// NewParticipant initializes the participant
func NewParticipant(name string) *participant {
	quizDone := make(chan bool)
	return &participant{name: name, quizDone: quizDone}
}

// RegisterParticipant is the first operation of the participant,
// This function perform the following operations.
//  1. Start a gorums server to receive the questions.
//  2. Create a connection to the quizmaster on the serverPort and
//     create a client with this connection.
//  3. Send this server address to the quizmaster with Register RPC.
//  4. Store the node id returned in the register RPC response.
func (p *participant) RegisterParticipant(serverPort int) {
	// TODO(student) implement this function
}

// StartQuiz starts the quiz for the participant and waits until it is completed.
// 1. Call the "MarkReady" RPC with the node id received in Register RPC.
// 2. Wait until the completion of the all the questions (can be done through a channel).
func (p *participant) StartQuiz(serverPort int) {
	// TODO(student) implement this function
}

// Answer function receives the question from quizmaster.
// 1. Display the questions and options to the user
// 2. Waits for ParticipantReadTimeout seconds for user input.
// 3. If the user input is received within the timeout, then the response is sent
// as ParticipantAnswer. If timeout happens then InvalidAnswer is sent as response.
// You can use the ReadOptionFromUser to read the participant's response.
func (p *participant) Answer(ctx gorums.ServerCtx, in *pb.Question) (*pb.ParticipantAnswer, error) {
	// TODO(student) implement this function
	return nil, nil
}

// GetResults fetches the results from the quizmaster.
// 1. Call the "GetResults" RPC to fetch the results of all the participants.
// 2. Waits only ParticipantGetResultTimeout seconds for the RPC completion.
// 3. After the RPC completion, display the results on the STDOUT
func (p *participant) GetResults(serverPort int) {
	// TODO(student) implement this function
}

// ReadOptionFromUser reads the option from an user from the STDIN and pass it to the
// inputChannel. It reads the input within ParticipantReadTimeout seconds
// and if a reply is done by a participant, it sends the reply to the inputChannel.
// In case of timeout, the function should send InvalidAnswer in the inputChannel.
func (p *participant) ReadOptionFromUser() int32 {
	return <-p.inputChannel
}

// CloseConnection closes the client connection at the end of quiz
func (p *participant) CloseConnection() {
	p.clientConnection.Close()
}
