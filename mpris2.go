package main

import (
	"github.com/go-music-players/mpris"
	"github.com/wildeyedskies/go-mpv/mpv"
)

// MPRIS Player interface implementation
// Player directly implements mpris.Player interface

// Play starts playback (implements mpris.Player)
func (p *Player) Play() error {
	isPaused, err := p.IsPaused()
	if err != nil {
		return err
	}
	if isPaused {
		return p.Instance.SetProperty("pause", mpv.FORMAT_FLAG, false)
	}
	return nil
}

// Pause pauses playback (implements mpris.Player)
func (p *Player) Pause() error {
	isPaused, err := p.IsPaused()
	if err != nil {
		return err
	}
	if !isPaused {
		return p.Instance.SetProperty("pause", mpv.FORMAT_FLAG, true)
	}
	return nil
}

// Next plays next track (implements mpris.Player)
func (p *Player) Next() error {
	p.PlayNextTrack()
	return nil
}

// Previous plays previous track (implements mpris.Player)
func (p *Player) Previous() error {
	// stmp doesn't support previous
	return nil
}

// GetPlaybackStatus returns current playback status (implements mpris.Player)
func (p *Player) GetPlaybackStatus() (mpris.PlaybackStatus, error) {
	state, err := p.State()
	if err != nil {
		return mpris.StatusStopped, err
	}

	switch state {
	case PlayerPlaying:
		return mpris.StatusPlaying, nil
	case PlayerPaused:
		return mpris.StatusPaused, nil
	default:
		return mpris.StatusStopped, nil
	}
}

// GetMetadata returns track metadata (implements mpris.Player)
func (p *Player) GetMetadata() (*mpris.Metadata, error) {
	if len(p.Queue) == 0 {
		return nil, nil
	}

	track := p.Queue[0]

	metadata := &mpris.Metadata{
		TrackID: track.Id,
		Title:   track.Title,
		Artist:  []string{track.Artist},
	}

	return metadata, nil
}

// CanPlay returns true if can play (implements mpris.Player)
func (p *Player) CanPlay() bool {
	return true
}

// CanPause returns true if can pause (implements mpris.Player)
func (p *Player) CanPause() bool {
	return true
}

// CanGoNext returns true if can go to next track (implements mpris.Player)
func (p *Player) CanGoNext() bool {
	return true
}

// CanGoPrevious returns true if can go to previous track (implements mpris.Player)
func (p *Player) CanGoPrevious() bool {
	return false // stmp doesn't support previous
}

// CanSeek returns true if can seek (implements mpris.Player)
func (p *Player) CanSeek() bool {
	return false // stmp doesn't support seeking via MPRIS
}

// CanControl returns true if player can be controlled (implements mpris.Player)
func (p *Player) CanControl() bool {
	return true
}
