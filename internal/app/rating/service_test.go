package rating

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

type RatingServiceTestSuite struct {
	suite.Suite
	service    Service
	ratingRepo mocks.RepoRating
}

func (suite *RatingServiceTestSuite) SetupTest() {
	suite.ratingRepo = mocks.RepoRating{}
	suite.service = NewService(&suite.ratingRepo)
}

func (suite *RatingServiceTestSuite) TearDownTest() {
	suite.ratingRepo.AssertExpectations(suite.T())
}

func TestProfileServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RatingServiceTestSuite))
}

func (suite *RatingServiceTestSuite) TestGiveRating() {
	type testCase struct {
		name           string
		rating         RatingRequestBody
		setup          func()
		expectedOutput RatingResponse
		expectedError  bool
	}

	testcases := []testCase{
		{
			name: "success",
			rating: RatingRequestBody{
				GivenBy:    1,
				GivenTo:    2,
				Request_id: 1,
				Rating:     5,
				Feedback:   "Excellent Player",
			},
			setup: func() {
				suite.ratingRepo.On("GiveRating", mock.Anything, repository.RatingRequestBody{
					GivenBy:    1,
					GivenTo:    2,
					Request_id: 1,
					Rating:     5,
					Feedback:   "Excellent Player",
				}).Return(repository.RatingResponse{
					Message: "Rating Given Successfully.",
				}, nil)
			},
			expectedOutput: RatingResponse{
				Message: "Rating Given Successfully.",
			},
			expectedError: false,
		},

		{
			name: "error",
			rating: RatingRequestBody{
				GivenBy:    1,
				GivenTo:    2,
				Request_id: 1,
				Rating:     5,
				Feedback:   "Excellent Player",
			},
			setup: func() {
				suite.ratingRepo.On("GiveRating", mock.Anything, repository.RatingRequestBody{
					GivenBy:    1,
					GivenTo:    2,
					Request_id: 1,
					Rating:     5,
					Feedback:   "Excellent Player",
				}).Return(repository.RatingResponse{}, errors.New("Error occured while giving rating.")) // Return error here
			},
			expectedOutput: RatingResponse{},
			expectedError:  true,
		},
	}

	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			user, err := suite.service.GiveRating(context.Background(), test.rating)
			suite.Equal(test.expectedOutput, user)
			suite.Equal(test.expectedError, err != nil)
		})
	}

}

func (suite *RatingServiceTestSuite) TestGetRatingByUserId() {
	type testCase struct {
		name           string
		user_id        int
		setup          func()
		expectedOutput []RatingUserIdResponse
		expectedError  bool
	}

	testcases := []testCase{
		{
			name:    "success",
			user_id: 1,
			setup: func() {
				suite.ratingRepo.On("GetRatingByUserId", mock.Anything, 1).Return([]repository.RatingUserIdResponse{
					{
						Id:         1,
						Given_by:   2,
						Name:       "Tushar Kachare",
						Given_to:   1,
						Request_id: 1,
						Sport:      func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
						Rating:     5,
						Feedback:   "Excellent Player",
						Created_at: time.Time{},
					},
				}, nil)
			},
			expectedOutput: []RatingUserIdResponse{
				{
					Id:         1,
					Given_by:   2,
					Name:       "Tushar Kachare",
					Given_to:   1,
					Request_id: 1,
					Sport:      func() json.RawMessage { b, _ := json.Marshal(map[string]string{"Chess": "Intermediate"}); return b }(),
					Rating:     5,
					Feedback:   "Excellent Player",
					Created_at: time.Time{},
				},
			},
			expectedError: false,
		},
		{
			name:    "error",
			user_id: 1,
			setup: func() {
				suite.ratingRepo.On("GetRatingByUserId", mock.Anything, 1).Return([]repository.RatingUserIdResponse{}, errors.New("error in getting rating"))
			},
			expectedOutput: []RatingUserIdResponse{},
			expectedError:  true,
		},
	}
	for _, test := range testcases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			user, err := suite.service.GetRatingByUserId(context.Background(), test.user_id)
			suite.Equal(test.expectedOutput, user)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}
