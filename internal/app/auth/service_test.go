package auth

import (
	"context"
	"errors"

	"encoding/json"
	"testing"

	"github.com/ganesh-kachare-josh/GameON/internal/repository"
	"github.com/ganesh-kachare-josh/GameON/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthServiceTestSuite struct {
	suite.Suite
	service  Service
	authRepo mocks.RepoAuth
}

func (suite *AuthServiceTestSuite) SetupTest() {
	suite.authRepo = mocks.RepoAuth{}
	suite.service = NewService(&suite.authRepo)
}

func (suite *AuthServiceTestSuite) TearDownTest() {
	suite.authRepo.AssertExpectations(suite.T())
}

func TestProfileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}

func (suite *AuthServiceTestSuite) TestLogin() {
	type testCase struct {
		name           string
		user           repository.Login
		setup          func()
		expectedOutput LoginResponse
		expectedError  bool
	}
	testcases := []testCase{
		{
			name: "success",
			user: repository.Login{
				Id:       1,
				Name:     "Ganesh Kachare",
				Email:    "kaganesh12@gmail.com",
				Password: "pass",
			},
			setup: func() {
				suite.authRepo.On("Login", mock.Anything, repository.Login{
					Id:       1,
					Name:     "Ganesh Kachare",
					Email:    "kaganesh12@gmail.com",
					Password: "pass",
				}).Return(repository.LoginResponse{
					Id:    1,
					Email: "kaganesh12@gmail.com",
					Name:  "Ganesh Kachare",
					Token: "aeebc.agwc.c",
				}, nil)
			},
			expectedOutput: LoginResponse{
				Id:    1,
				Email: "kaganesh12@gmail.com",
				Name:  "Ganesh Kachare",
				Token: "aeebc.agwc.c",
			},
			expectedError: false,
		},
		{
			name: "error",
			user: repository.Login{
				Id:       1,
				Name:     "Ganesh Kachare",
				Email:    "kaganesh12@gmail.com",
				Password: "pass",
			},
			setup: func() {
				suite.authRepo.On("Login", mock.Anything, repository.Login{
					Id:       1,
					Name:     "Ganesh Kachare",
					Email:    "kaganesh12@gmail.com",
					Password: "pass",
				}).Return(repository.LoginResponse{}, errors.New("error in login"))
			},
			expectedOutput: LoginResponse{},
			expectedError:  true,
		},
	}
	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			user, err := suite.service.Login(context.Background(), test.user)
			suite.Equal(test.expectedOutput, user)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}

func (suite *AuthServiceTestSuite) TestRegister() {
	type testCase struct {
		name           string
		user           RegisterData
		setup          func()
		expectedOutput RegisterData
		expectedError  bool
	}

	testcases := []testCase{
		{
			name: "success",
			user: RegisterData{
				Id:           1,
				Name:         "Ganesh",
				Email:        "random@gmail.com",
				Password:     "pass",
				Phone_Number: "12344512345",
				Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Badminton": "Intermediate"}); return b }(),
				Created_at:   "",
			},
			setup: func() {
				suite.authRepo.On("Register", mock.Anything, repository.Register{
					Id:           1,
					Name:         "Ganesh",
					Email:        "random@gmail.com",
					Password:     "pass",
					Phone_Number: "12344512345",
					Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Badminton": "Intermediate"}); return b }(),
					Created_at:   "",
				}).Return(repository.Register{
					Id:           1,
					Name:         "Ganesh",
					Email:        "random@gmail.com",
					Password:     "pass",
					Phone_Number: "12344512345",
					Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Badminton": "Intermediate"}); return b }(),
					Created_at:   "",
				}, nil)
			},
			expectedOutput: RegisterData{
				Id:           1,
				Name:         "Ganesh",
				Email:        "random@gmail.com",
				Password:     "pass",
				Phone_Number: "12344512345",
				Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Badminton": "Intermediate"}); return b }(),
				Created_at:   "",
			},
			expectedError: false,
		},

		{
			name: "error",
			user: RegisterData{
				Id:           1,
				Name:         "Ganesh",
				Email:        "random@gmail.com",
				Password:     "pass",
				Phone_Number: "12344512345",
				Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Badminton": "Intermediate"}); return b }(),
				Created_at:   "",
			},
			setup: func() {
				suite.authRepo.On("Register", mock.Anything, repository.Register{
					Id:           1,
					Name:         "Ganesh",
					Email:        "random@gmail.com",
					Password:     "pass",
					Phone_Number: "12344512345",
					Sport:        func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Badminton": "Intermediate"}); return b }(),
					Created_at:   "",
				}).Return(repository.Register{}, errors.New("error while register"))
			},
			expectedOutput: RegisterData{},
			expectedError:  true,
		},
	}
	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			user, err := suite.service.Register(context.Background(), test.user)
			suite.Equal(test.expectedOutput, user)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}
