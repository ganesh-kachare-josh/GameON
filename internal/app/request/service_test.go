package request

import (
	"context"
	"encoding/json"
	"errors"

	"testing"
	"time"

	"github.com/ganesh-kachare-josh/GameON/internal/repository"
	"github.com/ganesh-kachare-josh/GameON/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RequestServiceTestSuite struct {
	suite.Suite
	service     Service
	requestRepo mocks.RepoPerson
}

func (suite *RequestServiceTestSuite) SetupTest() {
	suite.requestRepo = mocks.RepoPerson{}
	suite.service = NewService(&suite.requestRepo)
}

func (suite *RequestServiceTestSuite) TearDownTest() {
	suite.requestRepo.AssertExpectations(suite.T())
}

func TestProfileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RequestServiceTestSuite))
}

func (suite *RequestServiceTestSuite) TestGetRequestById() {
	type testCase struct {
		name           string
		user_id        int
		setup          func()
		expectedOutput Request
		expectedError  bool
	}
	testcases := []testCase{
		{
			name:    "success",
			user_id: 1,
			setup: func() {
				suite.requestRepo.On("GetRequestById", mock.Anything, 1).Return(repository.Request{
					Id:           1,
					User_id:      1,
					Name:         "Ganesh",
					Email:        "pass@gmail.com",
					Phone_Number: "1234512345",
					Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
					Location:     "balewadi Stadium",
					Time:         time.Time{},
					CourtPrice:   200,
				}, nil)
			},
			expectedOutput: Request{
				Id:           1,
				User_id:      1,
				Name:         "Ganesh",
				Email:        "pass@gmail.com",
				Phone_Number: "1234512345",
				Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
				Location:     "balewadi Stadium",
				Time:         time.Time{},
				CourtPrice:   200,
			},
			expectedError: false,
		},
		{
			name:    "error",
			user_id: 1,
			setup: func() {
				suite.requestRepo.On("GetRequestById", mock.Anything, 1).Return(repository.Request{}, errors.New("Request not found"))
			},
			expectedOutput: Request{},
			expectedError:  true,
		},
	}
	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			user, err := suite.service.GetRequestById(context.Background(), test.user_id)
			suite.Equal(test.expectedOutput, user)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}

func (suite *RequestServiceTestSuite) TestGetAllRequests() {
	type testCase struct {
		name           string
		setup          func()
		expectedOutput []Request
		expectedError  bool
	}
	testcases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.requestRepo.On("GetAllRequests", mock.Anything).Return([]repository.Request{
					{
						Id:           1,
						User_id:      1,
						Name:         "Ganesh",
						Email:        "pass@gmail.com",
						Phone_Number: "1234512345",
						Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
						Location:     "balewadi Stadium",
						Time:         time.Time{},
						CourtPrice:   200,
					},
				}, nil)
			},
			expectedOutput: []Request{
				{
					Id:           1,
					User_id:      1,
					Name:         "Ganesh",
					Email:        "pass@gmail.com",
					Phone_Number: "1234512345",
					Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
					Location:     "balewadi Stadium",
					Time:         time.Time{},
					CourtPrice:   200,
				},
			},
			expectedError: false,
		},
		{
			name: "error",
			setup: func() {
				suite.requestRepo.On("GetAllRequests", mock.Anything).Return([]repository.Request{}, errors.New("Request not found"))
			},
			expectedOutput: []Request{},
			expectedError:  true,
		},
	}
	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			user, err := suite.service.GetAllRequests(context.Background())
			suite.Equal(test.expectedOutput, user)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}

func (suite *RequestServiceTestSuite) TestGetAllParticipants() {
	type testCase struct {
		name           string
		request_id     int
		setup          func()
		expectedOutput []ParticipantData
		expectedError  bool
	}
	testcases := []testCase{
		{
			name:       "success",
			request_id: 1,
			setup: func() {
				suite.requestRepo.On("GetAllParticipants", mock.Anything, 1).Return([]repository.ParticipantData{
					{
						Id:     1,
						UserId: 1,
						Name:   "Ganesh",
						Status: "Open",
					},
				})
			},
			expectedOutput: []ParticipantData{
				{
					Id:     1,
					UserId: 1,
					Name:   "Ganesh",
					Status: "Open",
				},
			},
			expectedError: false,
		},
	}

	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			user := suite.service.GetAllParticipants(context.Background(), test.request_id)
			suite.Equal(test.expectedOutput, user)
		})
	}
}

func (suite *RequestServiceTestSuite) TestAcceptRequest() {
	type testCase struct {
		name            string
		request         AcceptRequestBody
		setup           func()
		expectedOutput1 AcceptRequestData
		expectedOutput2 ResponseForEmail
		expectedError   bool
	}

	testcases := []testCase{
		{
			name: "success",
			request: AcceptRequestBody{
				Request_id: 1,
				User_id:    2,
			},
			setup: func() {
				suite.requestRepo.On("AcceptRequest", mock.Anything, repository.AcceptRequestBody{
					Request_id: 1,
					User_id:    2,
				}).Return(repository.AcceptRequestData{
					Id:         1,
					Request_id: 1,
					User_id:    1,
					Status:     "Open",
				}, repository.ResponseForEmail{
					CreatorName:     "ganesh",
					ParticipantName: "Tushar",
					Sport:           "Cricket",
					Email:           "pass@gmail.com",
				}, nil)
			},
			expectedOutput1: AcceptRequestData{
				Id:         1,
				Request_id: 1,
				User_id:    1,
				Status:     "Open",
			},
			expectedOutput2: ResponseForEmail{
				CreatorName:     "ganesh",
				ParticipantName: "Tushar",
				Sport:           "Cricket",
				Email:           "pass@gmail.com",
			},
			expectedError: false,
		},

		{
			name: "error",
			request: AcceptRequestBody{
				Request_id: 1,
				User_id:    2,
			},
			setup: func() {
				suite.requestRepo.On("AcceptRequest", mock.Anything, repository.AcceptRequestBody{
					Request_id: 1,
					User_id:    2,
				}).Return(repository.AcceptRequestData{}, repository.ResponseForEmail{}, errors.New("Error in accepting the request"))
			},
			expectedOutput1: AcceptRequestData{},
			expectedOutput2: ResponseForEmail{},
			expectedError:   true,
		},
	}
	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			o1, o2, err := suite.service.AcceptRequest(context.Background(), test.request)
			suite.Equal(test.expectedOutput1, o1)
			suite.Equal(test.expectedOutput2, o2)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}

func (suite *RequestServiceTestSuite) TestConfirmRequest() {
	type testCase struct {
		name            string
		request         AcceptRequestBody
		setup           func()
		expectedOutput1 AcceptRequestData
		expectedOutput2 ResponseForEmail
		expectedError   bool
	}

	testcases := []testCase{
		{
			name: "success",
			request: AcceptRequestBody{
				Request_id: 1,
				User_id:    2,
			},
			setup: func() {
				suite.requestRepo.On("ConfirmRequest", mock.Anything, repository.AcceptRequestBody{
					Request_id: 1,
					User_id:    2,
				}).Return(repository.AcceptRequestData{
					Id:         1,
					Request_id: 1,
					User_id:    1,
					Status:     "Confirmed",
				}, repository.ResponseForEmail{
					CreatorName:     "ganesh",
					ParticipantName: "Tushar",
					Sport:           "Cricket",
					Email:           "pass@gmail.com",
				}, nil)
			},
			expectedOutput1: AcceptRequestData{
				Id:         1,
				Request_id: 1,
				User_id:    1,
				Status:     "Confirmed",
			},
			expectedOutput2: ResponseForEmail{
				CreatorName:     "ganesh",
				ParticipantName: "Tushar",
				Sport:           "Cricket",
				Email:           "pass@gmail.com",
			},
			expectedError: false,
		},

		{
			name: "error",
			request: AcceptRequestBody{
				Request_id: 1,
				User_id:    2,
			},
			setup: func() {
				suite.requestRepo.On("ConfirmRequest", mock.Anything, repository.AcceptRequestBody{
					Request_id: 1,
					User_id:    2,
				}).Return(repository.AcceptRequestData{}, repository.ResponseForEmail{}, errors.New("Error in confirming the request"))
			},
			expectedOutput1: AcceptRequestData{},
			expectedOutput2: ResponseForEmail{},
			expectedError:   true,
		},
	}
	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			o1, o2, err := suite.service.ConfirmRequest(context.Background(), test.request)
			suite.Equal(test.expectedOutput1, o1)
			suite.Equal(test.expectedOutput2, o2)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}

func (suite *RequestServiceTestSuite) TestRejectParticipant() {
	type testCase struct {
		name           string
		participant_id int
		setup          func()
		expectedOutput ResponseForEmail
		expectedError  bool
	}

	testcases := []testCase{
		{
			name:           "success",
			participant_id: 1,
			setup: func() {
				suite.requestRepo.On("RejectParticipant", mock.Anything, 1).Return(repository.ResponseForEmail{
					CreatorName:     "ganesh",
					ParticipantName: "tushar",
					Sport:           "chess",
					Email:           "pass@gmail.com",
				}, nil)
			},
			expectedOutput: ResponseForEmail{
				CreatorName:     "ganesh",
				ParticipantName: "tushar",
				Sport:           "chess",
				Email:           "pass@gmail.com",
			},
			expectedError: false,
		},

		{
			name:           "error",
			participant_id: 1,
			setup: func() {
				suite.requestRepo.On("RejectParticipant", mock.Anything, 1).Return(repository.ResponseForEmail{}, errors.New("error while rejecting participant"))
			},
			expectedOutput: ResponseForEmail{},
			expectedError:  true,
		},
	}

	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			user, err := suite.service.RejectParticipant(context.Background(), test.participant_id)
			suite.Equal(test.expectedOutput, user)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}

func (suite *RequestServiceTestSuite) TestCreateRequest() {
	type testCase struct {
		name           string
		request        Request
		setup          func()
		expectedOutput Request
		expectedError  bool
	}
	testcases := []testCase{
		{
			name: "success",
			request: Request{
				Id:           1,
				User_id:      1,
				Name:         "Ganesh",
				Email:        "pass@gmail.com",
				Phone_Number: "1234512345",
				Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
				Location:     "Balewadi Stadium",
				Time:         time.Time{},
				CourtPrice:   100,
			},
			setup: func() {
				suite.requestRepo.On("CreateRequest", mock.Anything, repository.Request{
					Id:           1,
					User_id:      1,
					Name:         "Ganesh",
					Email:        "pass@gmail.com",
					Phone_Number: "1234512345",
					Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
					Location:     "Balewadi Stadium",
					Time:         time.Time{},
					CourtPrice:   100,
				}).Return(repository.Request{
					Id:           1,
					User_id:      1,
					Name:         "Ganesh",
					Email:        "pass@gmail.com",
					Phone_Number: "1234512345",
					Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
					Location:     "Balewadi Stadium",
					Time:         time.Time{},
					CourtPrice:   100,
				}, nil)
			},
			expectedOutput: Request{
				Id:           1,
				User_id:      1,
				Name:         "Ganesh",
				Email:        "pass@gmail.com",
				Phone_Number: "1234512345",
				Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
				Location:     "Balewadi Stadium",
				Time:         time.Time{},
				CourtPrice:   100,
			},
			expectedError: false,
		},

		{
			name: "error",
			request: Request{
				Id:           1,
				User_id:      1,
				Name:         "Ganesh",
				Email:        "pass@gmail.com",
				Phone_Number: "1234512345",
				Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
				Location:     "Balewadi Stadium",
				Time:         time.Time{},
				CourtPrice:   100,
			},
			setup: func() {
				suite.requestRepo.On("CreateRequest", mock.Anything, repository.Request{
					Id:           1,
					User_id:      1,
					Name:         "Ganesh",
					Email:        "pass@gmail.com",
					Phone_Number: "1234512345",
					Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
					Location:     "Balewadi Stadium",
					Time:         time.Time{},
					CourtPrice:   100,
				}).Return(repository.Request{}, errors.New("Error in creating the request."))
			},
			expectedOutput: Request{},
			expectedError:  true,
		},
	}
	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			user, err := suite.service.CreateRequest(context.Background(), test.request)
			suite.Equal(test.expectedOutput, user)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}

func (suite *RequestServiceTestSuite) TestGetJoinedRequestById() {
	type testCase struct {
		name           string
		user_id        int
		setup          func()
		expectedOutput []int
		expectedError  bool
	}
	testcases := []testCase{
		{
			name:    "success",
			user_id: 1,
			setup: func() {
				suite.requestRepo.On("GetJoinedRequestById", mock.Anything, 1).Return([]int{1, 2, 3}, nil)
			},
			expectedOutput: []int{1, 2, 3},
			expectedError:  false,
		},
	}

	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			user, err := suite.service.GetJoinedRequestById(context.Background(), test.user_id)
			suite.Equal(test.expectedOutput, user)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}
