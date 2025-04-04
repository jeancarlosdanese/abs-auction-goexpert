package auction

import (
	"context"
	"fmt"
	"fullcycle-auction_go/configuration/logger"
	"fullcycle-auction_go/internal/entity/auction_entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (ar *AuctionRepository) StartAuctionExpirationScheduler() {
	ticker := time.NewTicker(time.Minute) // roda a cada 1 minuto
	go func() {
		for range ticker.C {
			ctx := context.Background()
			ar.checkAndExpireAuctions(ctx)
		}
	}()
}

func (ar *AuctionRepository) checkAndExpireAuctions(ctx context.Context) {
	expirationTime := time.Now().Add(-getAuctionInterval()).Unix()
	filter := bson.M{
		"status":    auction_entity.Active,
		"timestamp": bson.M{"$lte": expirationTime},
	}
	update := bson.M{"$set": bson.M{"status": auction_entity.Completed}}

	result, err := ar.Collection.UpdateMany(ctx, filter, update)
	if err != nil {
		logger.Error("Error trying to expire auctions in scheduler", err)
		return
	}

	if result.ModifiedCount > 0 {
		logger.Info(fmt.Sprintf("Scheduler: %d auction(s) expired automatically", result.ModifiedCount))
	}
}
