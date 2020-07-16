package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var gatewayurl string

// function used to create the shards object and to fetch the shard amount
func resetShardCount(shardCount int, reset bool) {
	if shardCount == 0 {
		// fetch recommended shard amount from discord api
		client := &http.Client{}
		var response gatewayresponse
		req, err := http.NewRequest("GET", "https://discordapp.com/api/gateway/bot", nil)
		if err != nil {
			log.Error(err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bot %v", bottoken))
		resp, err := client.Do(req)
		if err != nil {
			log.Error(err)
		}
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			log.Error(err)
		}
		shardCount = response.Shards
		gatewayurl = response.URL
	}

	if len(shards) == 0 || reset {
		// create list with shard objects
		shards = make([]shard, shardCount)

		// set shardIDs and roles
		for shardID := 0; shardID < shardCount; shardID++ {
			shards[shardID].ShardID = shardID
		}
	} else {
		// Create temporary buffers for new shards
		newShards := make([]shard, shardCount)

		for shardID := 0; shardID < shardCount; shardID++ {
			newShards[shardID].ShardID = shardID
			if shardID < len(shards) {
				newShards[shardID] = shards[shardID]
			}
		}

		// Update shards
		shards = newShards
	}
}