package services

import (
	"fmt"
	"groupie-tracker/models"
	"strings"
)

// SearchService gère toutes les recherches
type SearchService struct {
	data *models.APIData
}

// NewSearchService crée un nouveau service de recherche
func NewSearchService(data *models.APIData) *SearchService {
	return &SearchService{data: data}
}

// SearchArtists recherche des artistes par nom
func (s *SearchService) SearchArtists(query string) []models.Artist {
	if s.data == nil {
		return nil
	}

	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return s.data.Artists
	}

	var results []models.Artist
	for _, artist := range s.data.Artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			results = append(results, artist)
		}
	}
	return results
}

// SearchByMember recherche des artistes par membre
func (s *SearchService) SearchByMember(memberName string) []models.Artist {
	if s.data == nil {
		return nil
	}

	memberName = strings.ToLower(strings.TrimSpace(memberName))
	if memberName == "" {
		return nil
	}

	var results []models.Artist
	for _, artist := range s.data.Artists {
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), memberName) {
				results = append(results, artist)
				break
			}
		}
	}
	return results
}

// SearchByLocation recherche des concerts par lieu
func (s *SearchService) SearchByLocation(location string) []models.Concert {
	if s.data == nil {
		return nil
	}

	location = strings.ToLower(strings.TrimSpace(location))
	if location == "" {
		return nil
	}

	var concerts []models.Concert
	for i, artist := range s.data.Artists {
		if i < len(s.data.Relations) {
			relation := s.data.Relations[i]
			for loc, dates := range relation.DatesLocations {
				if strings.Contains(strings.ToLower(loc), location) {
					concerts = append(concerts, models.Concert{
						ArtistID:   artist.ID,
						ArtistName: artist.Name,
						Location:   loc,
						Dates:      dates,
					})
				}
			}
		}
	}
	return concerts
}

// SearchByAlbumDate recherche par date de premier album
func (s *SearchService) SearchByAlbumDate(date string) []models.Artist {
	if s.data == nil {
		return nil
	}

	date = strings.TrimSpace(date)
	if date == "" {
		return nil
	}

	var results []models.Artist
	for _, artist := range s.data.Artists {
		if strings.Contains(artist.FirstAlbum, date) {
			results = append(results, artist)
		}
	}
	return results
}

// SearchByCreationDate recherche par année de création
func (s *SearchService) SearchByCreationDate(year int) []models.Artist {
	if s.data == nil {
		return nil
	}

	var results []models.Artist
	for _, artist := range s.data.Artists {
		if artist.CreationDate == year {
			results = append(results, artist)
		}
	}
	return results
}

// UniversalSearch effectue une recherche globale
func (s *SearchService) UniversalSearch(query string) []models.SearchResult {
	if s.data == nil {
		return nil
	}

	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return nil
	}

	var results []models.SearchResult
	seen := make(map[string]bool)

	// Recherche d'artistes
	for _, artist := range s.data.Artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			key := fmt.Sprintf("artist-%s", artist.Name)
			if !seen[key] {
				results = append(results, models.SearchResult{
					Type:   "artist",
					Value:  artist.Name,
					Artist: &artist,
				})
				seen[key] = true
			}
		}

		// Recherche de membres
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				key := fmt.Sprintf("member-%s-%s", member, artist.Name)
				if !seen[key] {
					results = append(results, models.SearchResult{
						Type:   "member",
						Value:  fmt.Sprintf("%s (membre de %s)", member, artist.Name),
						Artist: &artist,
					})
					seen[key] = true
				}
			}
		}

		// Recherche par date d'album
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			key := fmt.Sprintf("album-%s", artist.Name)
			if !seen[key] {
				results = append(results, models.SearchResult{
					Type:   "album",
					Value:  fmt.Sprintf("%s - Premier album: %s", artist.Name, artist.FirstAlbum),
					Artist: &artist,
				})
				seen[key] = true
			}
		}
	}

	// Recherche de lieux
	for i, artist := range s.data.Artists {
		if i < len(s.data.Relations) {
			relation := s.data.Relations[i]
			for location := range relation.DatesLocations {
				if strings.Contains(strings.ToLower(location), query) {
					key := fmt.Sprintf("location-%s-%s", location, artist.Name)
					if !seen[key] {
						results = append(results, models.SearchResult{
							Type:   "location",
							Value:  fmt.Sprintf("%s - Concert à %s", artist.Name, FormatLocation(location)),
							Artist: &artist,
						})
						seen[key] = true
					}
				}
			}
		}
	}

	return results
}

// FilterByMemberCount filtre les artistes par nombre de membres
func (s *SearchService) FilterByMemberCount(min, max int) []models.Artist {
	if s.data == nil {
		return nil
	}

	var results []models.Artist
	for _, artist := range s.data.Artists {
		count := len(artist.Members)
		if count >= min && count <= max {
			results = append(results, artist)
		}
	}
	return results
}

// FilterByCreationYear filtre par année de création
func (s *SearchService) FilterByCreationYear(minYear, maxYear int) []models.Artist {
	if s.data == nil {
		return nil
	}

	var results []models.Artist
	for _, artist := range s.data.Artists {
		if artist.CreationDate >= minYear && artist.CreationDate <= maxYear {
			results = append(results, artist)
		}
	}
	return results
}

// GetConcertsByArtistID récupère tous les concerts d'un artiste
func (s *SearchService) GetConcertsByArtistID(artistID int) []models.Concert {
	if s.data == nil {
		return nil
	}

	var concerts []models.Concert
	for i, artist := range s.data.Artists {
		if artist.ID == artistID && i < len(s.data.Relations) {
			relation := s.data.Relations[i]
			for location, dates := range relation.DatesLocations {
				concerts = append(concerts, models.Concert{
					ArtistID:   artist.ID,
					ArtistName: artist.Name,
					Location:   location,
					Dates:      dates,
				})
			}
		}
	}
	return concerts
}

// FormatLocation formate un nom de lieu
func FormatLocation(location string) string {
	location = strings.ReplaceAll(location, "-", ", ")
	location = strings.ReplaceAll(location, "_", " ")
	parts := strings.Split(location, ", ")
	for i, part := range parts {
		parts[i] = strings.Title(strings.ToLower(part))
	}
	return strings.Join(parts, ", ")
}
