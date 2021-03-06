package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// MAL ...
func MAL(ctx *aero.Context) string {
	return editorList(
		ctx,
		"Anime without MAL mappings",
		func(anime *arn.Anime) bool {
			return anime.GetMapping("myanimelist/anime") == ""
		},
		func(anime *arn.Anime) string {
			return "https://myanimelist.net/anime.php?q=" + anime.Title.Canonical
		},
	)
}
