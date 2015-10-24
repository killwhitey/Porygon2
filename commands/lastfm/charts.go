package lastfm

import (
	"fmt"
	"github.com/0x263b/Porygon2"
	"github.com/0x263b/Porygon2/web"
	"strings"
)

const (
	ChartsURL = "http://ws.audioscrobbler.com/2.0/?method=user.gettopartists&user=%s&period=7day&limit=5&api_key=%s&format=json"
)

func charts(command *bot.Cmd, matches []string) (msg string, err error) {
	username := matches[1]

	if username == "" {
		username = checkLastfm(command.Nick)
	}

	if username == "" {
		return "Lastfm not provided, nor on file. Use `-set lastfm <lastfm>` to save", nil
	}

	data := &WeeklyCharts{}
	err = web.GetJSON(fmt.Sprintf(ChartsURL, username, bot.API.Lastfm), data)
	if err != nil || data.Error > 0 {
		return fmt.Sprintf("Could not get charts for %s", username), nil
	}
	if data.Topartists.Attr.Total == "0" {
		return fmt.Sprintf("Could not get charts for %s", username), nil
	}

	var fmtcharts string
	for i := range data.Topartists.Artist[:5] {
		fmtcharts += fmt.Sprintf("%s (%s), ", data.Topartists.Artist[i].Name, data.Topartists.Artist[i].Playcount)
	}
	fmtcharts = strings.TrimSuffix(fmtcharts, ", ")

	output := fmt.Sprintf("Last.fm | Top 5 Weekly artists for %s | %s", username, fmtcharts)

	return output, nil
}

func init() {
	bot.RegisterCommand(
		"^charts(?: (\\S+))?$",
		charts)
}
