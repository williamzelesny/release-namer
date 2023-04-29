package main

import (
	"golang.org/x/exp/slog"
	"os"
	"strconv"
	"williamzelesny/release-namer/scrapers"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	logger.Info("Getting candies")
	candies, err := scrapers.GetCandies()
	if err != nil {
		logger.Error(err.Error())
	}

	for _, candy := range candies {
		logger.Info(candy)
	}

	logger.Info("Total Candies Returned: " + strconv.Itoa(len(candies)))
}
