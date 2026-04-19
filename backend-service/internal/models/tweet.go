// internal/models/tweet.go
package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"sitter/internal/database/redisclient"
	"sort"
)

type Tweet struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	UserID   string `json:"userid"`
	Likes    int    `json:"like"`
	Shares   int    `json:"share"`
	Slurmed  int    `json:"slurmed"`
	Username string `json:"username"`
}

// Сохранение твита в Redis
func (t *Tweet) Save() error {
	tweetJSON, err := json.Marshal(t)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("tweet:%s", t.ID)
	return redisclient.GetRedisClient().Set(redisclient.GetContext(), key, tweetJSON, 0).Err()
}

// Получение твита по ID
func GetTweet(id string) (*Tweet, error) {
	key := fmt.Sprintf("tweet:%s", id)
	tweetJSON, err := redisclient.GetRedisClient().Get(redisclient.GetContext(), key).Result()
	if err != nil {
		return nil, err
	}

	var tweet Tweet
	err = json.Unmarshal([]byte(tweetJSON), &tweet)
	if err != nil {
		return nil, err
	}

	return &tweet, nil
}

// Лайк твита
func (t *Tweet) Like(userID string) error {
	key := fmt.Sprintf("tweet:%s:likes", t.ID)
	alreadyLiked, err := redisclient.GetRedisClient().SIsMember(redisclient.GetContext(), key, userID).Result()
	if err != nil {
		return err
	}

	// Если лайк уже есть, удалим его
	if alreadyLiked {
		err := redisclient.GetRedisClient().SRem(redisclient.GetContext(), key, userID).Err()
		if err != nil {
			return err
		}

		t.Likes++ // Уменьшаем количество лайков
	} else {
		// Если лайк не был поставлен, добавляем его
		err := redisclient.GetRedisClient().SAdd(redisclient.GetContext(), key, userID).Err()
		if err != nil {
			return err
		}

		t.Likes++ // Увеличиваем количество лайков
	}

	// Сохраняем обновленные данные
	return t.Save()
}

// Слурм твита
func (t *Tweet) Slurm(userID string) error {
	key := fmt.Sprintf("tweet:%s:slurms", t.ID)
	alreadySlurmed, err := redisclient.GetRedisClient().SIsMember(redisclient.GetContext(), key, userID).Result()
	if err != nil {
		return err
	}

	// Если слурм уже есть, удалим его
	if alreadySlurmed {
		err := redisclient.GetRedisClient().SRem(redisclient.GetContext(), key, userID).Err()
		if err != nil {
			return err
		}

		t.Slurmed++ // Уменьшаем количество слурмов
	} else {
		// Если слурм не был поставлен, добавляем его
		err := redisclient.GetRedisClient().SAdd(redisclient.GetContext(), key, userID).Err()
		if err != nil {
			return err
		}

		t.Slurmed++ // Увеличиваем количество слурмов
	}

	// Сохраняем обновленные данные
	return t.Save()
}

// GetSortedTweetsByLikes — получает все твиты и сортирует по лайкам
func GetSortedTweetsByLikes() ([]Tweet, error) {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"), // Подставьте свои настройки Redis
		Username: os.Getenv("REDIS_USERNAME"),   // если legacy auth, можно оставить пустым ""
        Password: os.Getenv("REDIS_PASSWORD"),  
	})

	// Получаем все ключи твитов (например, tweet:*)
	keys, err := client.Keys(ctx, "tweet:*").Result()
	if err != nil {
		return nil, err
	}

	var tweets []Tweet

	// Итерируем по ключам и получаем твиты
	for _, key := range keys {
		data, err := client.Get(ctx, key).Result()
		if err != nil {
			log.Println("Ошибка получения твита:", err)
			continue
		}

		var tweet Tweet
		err = json.Unmarshal([]byte(data), &tweet)
		if err != nil {
			log.Println("Ошибка десериализации твита:", err)
			continue
		}

		tweets = append(tweets, tweet)
	}

	// Сортируем твиты по количеству лайков (по убыванию)
	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i].Likes > tweets[j].Likes
	})

	return tweets, nil
}

// Репост твита
func (t *Tweet) Share() error {
	// Инкрементируем количество репостов
	t.Shares++

	// Сохраняем обновленный твит
	return t.Save()
}

// Получение всех твитов
func GetAllTweets() ([]Tweet, error) {
	// Получаем все ключи, которые начинаются с "tweet:"
	keys, err := redisclient.GetRedisClient().Keys(context.Background(), "tweet:*").Result()
	if err != nil {
		return nil, err
	}

	var tweets []Tweet
	for _, key := range keys {
		// Проверяем тип данных для каждого ключа
		keyType, err := redisclient.GetRedisClient().Type(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}

		// Пропускаем ключи, если тип данных не string
		if keyType != "string" {
			continue
		}

		// Извлекаем значение ключа
		tweetJSON, err := redisclient.GetRedisClient().Get(context.Background(), key).Result()
		if err != nil {
			return nil, err
		}

		var tweet Tweet
		err = json.Unmarshal([]byte(tweetJSON), &tweet)
		if err != nil {
			return nil, err
		}

		tweets = append(tweets, tweet)
	}

	return tweets, nil
}
