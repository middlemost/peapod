package mock

import (
	"context"
	"io"
	"net/url"

	"github.com/middlemost/peapod"
)

var _ peapod.TrackService = &TrackService{}

type TrackService struct {
	FindTrackByIDFn func(ctx context.Context, id int) (*peapod.Track, error)
	CreateTrackFn   func(ctx context.Context, track *peapod.Track) error
}

func (s *TrackService) FindTrackByID(ctx context.Context, id int) (*peapod.Track, error) {
	return s.FindTrackByIDFn(ctx, id)
}

func (s *TrackService) CreateTrack(ctx context.Context, track *peapod.Track) error {
	return s.CreateTrackFn(ctx, track)
}

var _ peapod.URLTrackGenerator = &URLTrackGenerator{}

type URLTrackGenerator struct {
	GenerateTrackFromURLFn func(ctx context.Context, url url.URL) (*peapod.Track, io.ReadCloser, error)
}

func (g *URLTrackGenerator) GenerateTrackFromURL(ctx context.Context, url url.URL) (*peapod.Track, io.ReadCloser, error) {
	return g.GenerateTrackFromURLFn(ctx, url)
}
