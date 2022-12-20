//go:build !solution

package game

import (
	"context"
	"fmt"
	"log"
	"sync"

	"datxxx/lab2/quiz"
	pb "datxxx/lab2/quiz/proto"
	"datxxx/lab2/quiz/utils"

	"github.com/relab/gorums"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Quizmaster structure holds the data to conduct quiz
// participants map stores the nodeID and the participant
// serverPort stores the port on which the server is running
// correctAnswers map stores the question id and the correct answer for that question
// configuration is the nodes to send the question and can change for each question
type quizmaster struct {
	sync.Mutex
	pb.UnimplementedParticipantServiceServer
	mgr           *pb.Manager
	participants  map[uint32]*pb.Participant
	serverPort    int
	configuration *pb.Configuration
}

// NewQuizMaster initializes the quizmaster structure
func NewQuizMaster(port int) *quizmaster {
	participants := make(map[uint32]*pb.Participant)

	// init gorums manager
	mgr := pb.NewManager(
		gorums.WithDialTimeout(quiz.QuizmasterParticipantConnect),
		gorums.WithGrpcDialOptions(
			grpc.WithInsecure(), // disable TLS
			grpc.WithBlock(),    // block until connections are made
		),
	)
	// TODO(student) Add any other necessary information to the quizmaster
	return &quizmaster{
		mgr:          mgr,
		serverPort:   port,
		participants: participants,
	}
}

// Register handles the Register RPC called from the participants.
// 1. Check the contents of the pb.Participant sent from the participants.
// 2. If the details are incorrect, send error
// 3. If all the details are correct, then store the participant
// in the quizmaster structure.
func (q *quizmaster) Register(_ context.Context, p *pb.Participant) (*pb.RegisterResponse, error) {
	// TODO(student) Implement this function
	return nil, nil
}

// MarkReady handles the "MarkReady" RPC from the participants,
// This marks the participants are ready to participate in the
// Quiz.
// Caller Participant is included in the configuration of the next question.
func (q *quizmaster) MarkReady(ct context.Context,
	resp *pb.RegisterResponse) (*emptypb.Empty, error) {
	// TODO(student) Implement this function
	return nil, nil
}

// GetResults handles the "GetResults" RPC from the participants.
// pb.Result should contain the score of all the registered participants.
// If no participants are available return error.
func (q *quizmaster) GetResults(ct context.Context, e *emptypb.Empty) (res *pb.Result, err error) {
	// TODO(student) Implement this function
	return nil, nil
}

// AnswerQF is the quorum function for the "Answer" RPC. It handles the
// responses for the question sent from the quizmaster.
// 1. It checks the answer for the question, if the answer is correct then
// updates the score of the participant.
// 2. Once all the responses are received, then return true,
// otherwise return false to further process the replies.
// P.S. This function is called multiple times with "replies" containing both
// old and new responses from the participants. While updating the scores,
// one should take care of not processing the reply twice.
func (q *quizmaster) AnswerQF(in *pb.Question, replies map[uint32]*pb.ParticipantAnswer) (*pb.ParticipantAnswer, bool) {
	// TODO(student) Implement this function
	return nil, false
}

// sendQuestions sends the questions from the questions variable from
// the questions.go.
// 1. Before sending the question, a new gorums client should be formed
// with the configuration of ready and live participants.
// A dead participant should not be included in the configuration.
// 2. It should call the "Answer" RPC on the participants
// with the question and wait for QuizmasterQuestionTimeout seconds.
// 3. After receiving the responses for the question,
// it should take user input from the user to "send next question" or "show results".
// This function returns after sending all questions.
func (q *quizmaster) sendQuestions() {
	// TODO(student) Implement this function
}

// Checks whether the participant is still live by dialing
func isLive(address string) bool {
	// TODO(student) Implement this function
	return false
}

// getLiveParticipants returns the list of addresses of ready and live participants
// This function should check if a registered participant is ready a alive.
// You can use the isLive function to perform the latter.
// It should remove participants that are not alive.
func (q *quizmaster) getLiveParticipants() (addresses []string) {
	// TODO(student) Implement this function
	return addresses
}

// runGorumsClientForEachQuestion creates a gorums manager
// and create a configuration with ready and live participants (you can use the getLiveParticipants)
// store this configuration object in quizmaster to use in Vote RPC.
func (q *quizmaster) runGorumsClientForEachQuestion() {
	// TODO(student) Implement this function
}

// RunGrpcServer starts the GRPC server to receive the RPC from participants.
// 1. Starts a GRPC server with the user input port
// 2. register the quizmaster as the handler for the RPC of the participants.
func (q *quizmaster) RunGrpcServer() {
	// TODO(student) Implement this function
}

// Print the results of the participants
func (q *quizmaster) showResults() {
	q.Lock()
	defer q.Unlock()
	for _, participant := range q.participants {
		log.Printf("Participant %s score is %d\n\n", participant.Name, participant.Score)
	}
}

// startQuiz: Take the user input to start the quiz
// Wait until the participants are ready and live.
// If at least one live participant is ready then start
// the quiz by sending the questions.
// After sending all the questions, show options of
// "ShowResults" and "Exit".
// If "ShowResults" is select print the results and exit
// if the "Exit" option is selected.
func (q *quizmaster) StartQuiz() {
	fmt.Println("Press Enter to start the Quiz")
	if !utils.ReadEnter() {
		log.Fatal("unable to read input")
		return
	}
	// TODO(student) You should wait for ready and alive participants
	// before start to send questions.
	q.sendQuestions()
	log.Println("quiz is completed")
	for {
		expectedCommands := []string{"ShowResults", "Exit"}
		option := utils.ReadCommand(expectedCommands)
		if option == 0 {
			q.showResults()
		} else {
			q.reset()
			return
		}
	}
}

func (q *quizmaster) reset() {
	q.participants = make(map[uint32]*pb.Participant)
}
