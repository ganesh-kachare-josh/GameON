package profile

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

type ProfileServiceTestSuite struct {
	suite.Suite
	service     Service
	profileRepo mocks.ProfileRepo
}

func (suite *ProfileServiceTestSuite) SetupTest() {
	suite.profileRepo = mocks.ProfileRepo{}
	suite.service = NewService(&suite.profileRepo)
}

func (suite *ProfileServiceTestSuite) TearDownTest() {
	suite.profileRepo.AssertExpectations(suite.T())
}

func TestProfileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ProfileServiceTestSuite))
}

func (suite *ProfileServiceTestSuite) TestGetUserById() {
	type testCase struct {
		name           string
		user_id        int
		setup          func()
		expectedOutput UserData
		expectedError  bool
	}
	testcases := []testCase{
		{
			name:    "success",
			user_id: 1,
			setup: func() {
				suite.profileRepo.On("GetUserById", mock.Anything, 1).Return(repository.UserData{
					Id:           1,
					Name:         "Ganesh Kachare",
					Email:        "kaganesh12@gmail.com",
					Sports:       func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
					Phone_Number: "7020304393",
					Created_at:   time.Time{},
				}, nil)
			},

			expectedOutput: UserData{
				Id:           1,
				Name:         "Ganesh Kachare",
				Email:        "kaganesh12@gmail.com",
				Sports:       func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
				Phone_Number: "7020304393",
				Created_at:   time.Time{},
			},

			expectedError: false,
		},

		{
			name:    "user not found",
			user_id: 1,
			setup: func() {
				suite.profileRepo.On("GetUserById", mock.Anything, 1).Return(repository.UserData{}, errors.New("user not found"))
			},
			expectedOutput: UserData{},
			expectedError:  true,
		},

		{
			name:    "internal error",
			user_id: 1,
			setup: func() {
				suite.profileRepo.On("GetUserById", mock.Anything, 1).Return(repository.UserData{}, errors.New("Internal server error"))
			},
			expectedOutput: UserData{},
			expectedError:  true,
		},
	}

	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			user, err := suite.service.GetUserById(context.Background(), test.user_id)
			suite.Equal(test.expectedOutput, user)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}

func (suite *ProfileServiceTestSuite) TestUpdateProfile() {
	type testCase struct {
		name           string
		user           UserData
		setup          func()
		expectedOutput UserData
		expectedError  bool
	}

	testcases := []testCase{
		// Successfull Update Testcase.
		{
			name: "success",
			user: UserData{
				Id:           1,
				Name:         "Tushar Kachare", // Updated data
				Email:        "katushar4@gmail.com",
				Sports:       func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Badminton": "Intermediate"}); return b }(),
				Phone_Number: "7823846030",
				Created_at:   time.Time{},
			},
			setup: func() {
				suite.profileRepo.On("UpdateProfile", mock.Anything, repository.UserData{
					Id:           1,
					Name:         "Tushar Kachare", // Updated data
					Email:        "katushar4@gmail.com",
					Sports:       func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Badminton": "Intermediate"}); return b }(),
					Phone_Number: "7823846030",
					Created_at:   time.Time{},
				}).Return(repository.UserData{
					Id:           1,
					Name:         "Tushar Kachare", // Updated data
					Email:        "katushar4@gmail.com",
					Sports:       func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Badminton": "Intermediate"}); return b }(),
					Phone_Number: "7823846030",
					Created_at:   time.Time{},
				}, nil)
			},
			expectedOutput: UserData{
				Id:           1,
				Name:         "Tushar Kachare", // Updated data: Tushar Kachare
				Email:        "katushar4@gmail.com",
				Sports:       func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Badminton": "Intermediate"}); return b }(),
				Phone_Number: "7823846030",
				Created_at:   time.Time{},
			},
			expectedError: false,
		},
		{
			name: "error",
			user: UserData{
				Id:           1,
				Name:         "Tushar Kachare", // Updated data
				Email:        "katushar4@gmail.com",
				Sports:       func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Badminton": "Intermediate"}); return b }(),
				Phone_Number: "7823846030",
				Created_at:   time.Time{},
			},
			setup: func() {
				suite.profileRepo.On("UpdateProfile", mock.Anything, repository.UserData{
					Id:           1,
					Name:         "Tushar Kachare", // Updated data
					Email:        "katushar4@gmail.com",
					Sports:       func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Badminton": "Intermediate"}); return b }(),
					Phone_Number: "7823846030",
					Created_at:   time.Time{},
				}).Return(repository.UserData{}, errors.New("error occured while updating profile"))
			},
			expectedOutput: UserData{},
			expectedError:  true,
		},
	}

	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			user, err := suite.service.UpdateProfile(context.Background(), test.user)
			suite.Equal(test.expectedOutput, user)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}
