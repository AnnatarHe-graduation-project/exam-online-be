package app

import (
	. "github.com/AnnatarHe/exam-online-be/app/models"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
	"gopkg.in/redis.v5"
)

// Gorm: the database instance
var Gorm *gorm.DB

// Redis: the redis instance
var Redis *redis.Client

// init the database
func initDatabase() {
	var err error

	config := revel.Config

	username, _ := config.String("db.username")
	pwd, _ := config.String("db.pwd")
	dbname, _ := config.String("db.dbname")
	connstring := fmt.Sprintf("host=db user=%s password=%s dbname=%s sslmode=disable", username, pwd, dbname)

	Gorm, err = gorm.Open("postgres", connstring)
	// defer Gorm.Close()
	Gorm.LogMode(true)

	if err != nil {
		revel.INFO.Println("DB error: ", err)
	}

	if err = Gorm.AutoMigrate(
		&Course{}, &User{}, &News{}, &Paper{}, &Question{}, &StudentPaper{},
	).Error; err != nil {
		revel.INFO.Printf("No error should happen when create table, but got %+v", err)
	}

	revel.INFO.Println("DB Connected")
}

// init redis server
func initRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	pong, err := Redis.Ping().Result()

	revel.INFO.Println("redis: ", pong, err)
}

// InitDB: 初始化数据库
func InitDB() {
	initDatabase()
	initRedis()
}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
