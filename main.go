package main

import (
	"golang.org/x/exp/slog"
	"os"
	"strconv"
	"williamzelesny/release-namer/candies"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	logger.Info("Getting candies")
	candies, err := candies.GetCandies()
	if err != nil {
		logger.Error(err.Error())
	}

	for _, candy := range candies {
		logger.Info(candy)
	}

	logger.Info("Total Candies Returned: " + strconv.Itoa(len(candies)))
}
