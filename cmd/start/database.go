package start
import (
	"orse/common/models"
	"orse/internal/database"
)
func main()  {
	client, _ := database.Open()
	client.AutoMigrate(&models.Menu{})
	client.AutoMigrate(&models.User{},&models.UserDetail{})
	client.AutoMigrate(&models.Order{},&models.OrderAmounts{},&models.OrderSetting{},models.OrderPay{},&models.OrderItem{})


}
