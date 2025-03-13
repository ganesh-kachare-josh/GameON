package rating

import (
	"context"
	"log"

	"github.com/ganesh-kachare-josh/GameON/internal/repository"
)

type Service interface {
	GiveRating(ctx context.Context, requestBody RatingRequestBody) (RatingResponse, error)
	GetRatingByUserId(ctx context.Context, user_id int) ([]RatingUserIdResponse, error)
}

type service struct {
	ratingRepo repository.RepoRating
}

func NewService(ratingRepo repository.RepoRating) Service {
	return &service{
		ratingRepo: ratingRepo,
	}
}

func (s *service) GiveRating(ctx context.Context, requestBody RatingRequestBody) (RatingResponse, error) {
	response, err := s.ratingRepo.GiveRating(ctx, repository.RatingRequestBody(requestBody))
	if err != nil {
		log.Println(err)
		return RatingResponse{}, err
	}
	return RatingResponse(response), nil
}

func (s *service) GetRatingByUserId(ctx context.Context, user_id int) ([]RatingUserIdResponse, error) {
	response, err := s.ratingRepo.GetRatingByUserId(ctx, user_id)
	if err != nil {
		log.Println(err)
		return []RatingUserIdResponse{}, err
	}
	var result []RatingUserIdResponse
	for _, r := range response {
		result = append(result, RatingUserIdResponse(r))
	}
	return result, nil
}
