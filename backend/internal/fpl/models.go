package fpl

type Player struct {
	ID      int	   `json:"id"`
    WebName string `json:"web_name"`
    Team    int    `json:"team"`
}

type Team struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Code      int    `json:"code"` // Matches to SVG path from official FPL API to extract the logo
}

type GameWeek struct {
	ID           int        `json:"id"`
	Name         string     `json:"name"`
	DeadlineTime string     `json:"deadline_time"`
	MostPoints   Player `json:"top_element"`
	MostSelected Player `json:"most_selected"`
}

