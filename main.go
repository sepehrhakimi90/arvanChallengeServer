package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/sepehrhakimi90/arvanChallengeServer/controller"
	"github.com/sepehrhakimi90/arvanChallengeServer/entity"
	"github.com/sepehrhakimi90/arvanChallengeServer/repository"
	"github.com/sepehrhakimi90/arvanChallengeServer/service"
)

var (
	mysqlUsername = os.Getenv("MYSQL_USER")
	mysqlPassword = os.Getenv("MYSQL_PASS")
	mysqlHost     = os.Getenv("MYSQL_HOST")
	mysqlPort     = os.Getenv("MYSQL_PORT")
	mysqlDB       = os.Getenv("MYSQL_DB")
)

func main() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlUsername, mysqlPassword, mysqlHost, mysqlPort, mysqlDB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database\n", err)
	}

	db.AutoMigrate(&entity.Rule{})

	ruleRepo := repository.NewMysqlRuleRepository(db)

	ruleService := service.NewRuleService(ruleRepo)

	redisPublisher := service.NewRedisPublisher(context.Background())

	ruleController := controller.NewController(ruleService, redisPublisher)

	route := NewGinRouter(ruleController)

	route.Run(":8080")



	/*rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	defer rdb.Close()

	ctx := context.Background()
	rule := Rule{
		ID:        12,
		RuleData: RuleData{
			Domain:    "yahoo.com",
			StartTime: time.Now(),
			Suspect: "192.168.1.1",
			TTL:       70,
		},
	}
	data, err := json.Marshal(&rule.RuleData)
	data2, err := json.Marshal(&rule)
	fmt.Println(string(data2))
	if err != nil {
		fmt.Println(err)
	}
	err = rdb.Publish(ctx, "ruleChannel", string(data)).Err()
	if err != nil {
		panic(err)
	}
	*/
}
