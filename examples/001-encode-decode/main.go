package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Profile struct {
	Name      string
	Region    string
	Raiting   string
	Interests string
}

func (p *Profile) UnmarshalJSON(data []byte) error {
	type ProfileJSON struct {
		Name      string   `json:"name"`
		Region    string   `json:"region"`
		Raiting   string   `json:"raiting"`
		Interests []string `json:"interests"`
	}

	var pj ProfileJSON

	if err := json.Unmarshal(data, &pj); err != nil {
		return fmt.Errorf("failed to unmarshal profile: %w", err)
	}

	*p = Profile{
		Name:      pj.Name,
		Region:    pj.Region,
		Raiting:   pj.Raiting,
		Interests: strings.Join(pj.Interests, ";"),
	}

	return nil
}

func (p Profile) MarshalJSON() ([]byte, error) {
	type ProfileJSON struct {
		Name      string   `json:"name"`
		Region    string   `json:"region"`
		Raiting   string   `json:"raiting"`
		Interests []string `json:"interests"`
	}

	pj := ProfileJSON{
		Name:      p.Name,
		Region:    p.Region,
		Raiting:   p.Raiting,
		Interests: strings.Split(p.Interests, ";"),
	}

	return json.Marshal(pj)
}
