package main
import (
	"fmt"
	"math/big"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"strconv"
)

var router *gin.Engine
var webDirectory = "./web"
var apiPath = "/api"
var port = ":80"

func main() {
	router = gin.Default()

	router.Use(static.Serve("/", static.LocalFile(webDirectory, true)))

	api := router.Group(apiPath)
	{
		api.GET("/GetPreviousPrime", func(ctx *gin.Context) {
			
			ctx.JSON(200, gin.H{
				"previousPrime": GetAndPrintClosestPrimeNumber(ctx.Request.URL.Query().Get("number")),
			})
		})
	}
	router.Run(port)
}

/*
Endpoint for the /api/GetPreviousPrime method
Responds with the closest previous prime to the client, or with "0" in case there is an error
*/
func GetAndPrintClosestPrimeNumber(inputNumber string) string{
	i, err := strconv.ParseInt(inputNumber, 10, 64)
	if i < 2 || err != nil{
		return "0"
	}

	var bignum, ok = new(big.Int).SetString(inputNumber, 0)
	if ok == false{
		return "0";
	}
	primeNumber := GetPrimeLowerThan(bignum)
	return fmt.Sprintf("%d", primeNumber)
}

/*
The big.int argument is reduced by 1, and then checked with ProbablyPrime with the highest accuracy until the closest previous prime is found, then returns than number
*/
func GetPrimeLowerThan(number *big.Int) *big.Int {
    isPrime := false
	for isPrime == false{
		number.Sub(number,big.NewInt(1))
		isPrime = number.ProbablyPrime(1)
	}
    return number
}