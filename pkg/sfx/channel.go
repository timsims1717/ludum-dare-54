package sfx

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/timing"
	"math"
)

type Mode int

const (
	Sequential = iota
	Random
	Repeat
)

type musicSet struct {
	key  string
	set  []string
	curr string
	cId  int
	next string

	mode     Mode
	playNext bool
	fade     float64
	vol      float64

	paused  bool
	stopped bool

	stream beep.StreamSeekCloser
	ctrl   *beep.Ctrl
	volume *effects.Volume
	interV *gween.Tween
	format beep.Format
}

func (s *musicSet) setTracks(keys []string) {
	s.set = keys
	s.cId = 0
	s.next = ""
}

func (s *musicSet) play() {
	if s.stopped {
		s.stopped = false
		s.playNext = true
	} else if s.paused {
		s.pause(false)
	} else {
		s.playNext = true
	}
}

func (s *musicSet) resume() {
	if !s.stopped {
		if s.paused {
			s.pause(false)
		} else {
			s.playNext = true
		}
	}
}

func (s *musicSet) chooseTrack(keys []string) {
	if !s.stopped {
		for _, k := range keys {
			if k == s.curr {
				return
			}
		}
	}
	s.stopped = false
	s.next = keys[random.Intn(len(keys))]
	s.playNext = true
}

func (s *musicSet) setTrack(key string) {
	s.next = key
	s.playNext = s.next != s.curr
	s.stopped = false
}

func (s *musicSet) playTrack(key string) {
	s.pause(true)
	s.next = key
	s.playNext = true
	s.stopped = false
}

func (s *musicSet) pause(pause bool) {
	if pause && s.fade > 0. && s.volume != nil {
		s.interV = gween.New(s.volume.Volume, -8., s.fade, ease.Linear)
	} else if !pause && s.fade > 0. && s.volume != nil {
		s.interV = gween.New(s.volume.Volume, getMusicVolume()+s.vol, s.fade, ease.Linear)
	} else {
		s.interV = nil
	}
	s.paused = pause
}

func (s *musicSet) stop() {
	s.pause(true)
	s.stopped = true
}

func (s *musicSet) setVolume(vol float64) {
	s.vol = vol
	if s.interV != nil && s.fade > 0. && s.volume != nil {
		s.interV = gween.New(s.volume.Volume, getMusicVolume()+s.vol, s.fade, ease.Linear)
	}
}

func (s *musicSet) setFade(fade float64) {
	s.fade = fade
}

func (s *musicSet) update() {
	if s.playNext && len(s.set) > 0 {
		if !MusicPlayer.loading && (s.ctrl == nil || s.volume == nil || s.ctrl.Paused || s.volume.Silent || s.mode == Repeat) {
			if s.next == "" {
				if len(s.set) > 1 {
					switch s.mode {
					case Random:
						t := -1
						for s.next == "" || s.next == s.curr {
							t = random.Intn(len(s.set))
							s.next = s.set[t]
						}
						s.cId = t
					case Sequential:
						s.cId++
						s.cId %= len(s.set)
						s.next = s.set[s.cId]
					}
				} else {
					s.next = s.set[0]
				}
			}
			go MusicPlayer.loadTrack(s)
		} else if !s.paused {
			s.pause(true)
		}
	}
	if s.volume != nil {
		speaker.Lock()
		if s.interV != nil {
			v, fin := s.interV.Update(timing.DT)
			if fin {
				s.volume.Silent = s.paused || getMusicMuted()
				s.volume.Volume = getMusicVolume() + s.vol
				s.ctrl.Paused = s.paused
				s.interV = nil
			} else {
				s.volume.Volume = math.Min(v, getMusicVolume()+s.vol)
				s.volume.Silent = getMusicMuted()
				s.ctrl.Paused = false
			}
		} else {
			if s.paused {
				s.volume.Volume = -8.
			} else {
				s.volume.Volume = getMusicVolume() + s.vol
			}
			s.volume.Silent = s.paused || getMusicMuted()
			s.ctrl.Paused = s.paused
		}
		speaker.Unlock()
	}
}
