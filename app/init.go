package app

import (
	"fmt"

	. "github.com/AnnatarHe/exam-online-be/app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/revel/revel"
)

// the database instance
var Gorm *gorm.DB

// init the database
func InitDB() {

	config := revel.Config

	username, _ := config.String("db.username")
	pwd, _ := config.String("db.pwd")
	dbname, _ := config.String("db.dbname")
	connstring := fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable", username, pwd, dbname)
	var err error
	Gorm, err = gorm.Open("postgres", connstring)
	defer Gorm.Close()

	if err != nil {
		revel.INFO.Println("DB error: ", err)
	}

	Gorm.AutoMigrate(&Course{}, &User{}, &News{}, &Paper{}, &Question{})
	if !Gorm.HasTable(&Course{}) {
		revel.INFO.Println("there is user")
	}

	revel.INFO.Println("DB Connected")
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
