package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io"
	"net/http"
)

const (
	BaseURL = "https://groupietrackers.herokuapp.com/api"
)

// Client gère les requêtes à l'API
type Client struct {
	baseURL string
	client  *http.Client
}

// NewClient crée un nouveau client API
func NewClient() *Client {
	return &Client{
		baseURL: BaseURL,
		client:  &http.Client{},
	}
}

// GetArtists récupère tous les artistes
func (c *Client) GetArtists() ([]models.Artist, error) {
	resp, err := c.client.Get(c.baseURL + "/artists")
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des artistes: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erreur HTTP: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture de la réponse: %w", err)
	}

	var artists []models.Artist
	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, fmt.Errorf("erreur lors du parsing JSON: %w", err)
	}

	return artists, nil
}

// GetLocations récupère tous les lieux
func (c *Client) GetLocations() ([]models.Location, error) {
	resp, err := c.client.Get(c.baseURL + "/locations")
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des lieux: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var locationData struct {
		Index []models.Location `json:"index"`
	}

	if err := json.Unmarshal(body, &locationData); err != nil {
		return nil, err
	}

	return locationData.Index, nil
}

// GetDates récupère toutes les dates
func (c *Client) GetDates() ([]models.Date, error) {
	resp, err := c.client.Get(c.baseURL + "/dates")
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des dates: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dateData struct {
		Index []models.Date `json:"index"`
	}

	if err := json.Unmarshal(body, &dateData); err != nil {
		return nil, err
	}

	return dateData.Index, nil
}

// GetRelations récupère toutes les relations
func (c *Client) GetRelations() ([]models.Relation, error) {
	resp, err := c.client.Get(c.baseURL + "/relation")
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la récupération des relations: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var relationData struct {
		Index []models.Relation `json:"index"`
	}

	if err := json.Unmarshal(body, &relationData); err != nil {
		return nil, err
	}

	return relationData.Index, nil
}

// LoadAllData charge toutes les données de l'API
func (c *Client) LoadAllData() (*models.APIData, error) {
	data := &models.APIData{}

	artists, err := c.GetArtists()
	if err != nil {
		return nil, fmt.Errorf("erreur chargement artistes: %w", err)
	}
	data.Artists = artists

	relations, err := c.GetRelations()
	if err != nil {
		return nil, fmt.Errorf("erreur chargement relations: %w", err)
	}
	data.Relations = relations

	locations, err := c.GetLocations()
	if err != nil {
		// Non bloquant
		data.Locations = []models.Location{}
	} else {
		data.Locations = locations
	}

	dates, err := c.GetDates()
	if err != nil {
		// Non bloquant
		data.Dates = []models.Date{}
	} else {
		data.Dates = dates
	}

	return data, nil
}
