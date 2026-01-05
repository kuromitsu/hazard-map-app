package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// クライアントから受け取るデータ
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// 外部API（Heartrails）からのレスポンス用構造体
type GeoResponse struct {
	Response struct {
		Location []struct {
			City       string `json:"city"`
			Town       string `json:"town"`
			Prefecture string `json:"prefecture"`
		} `json:"location"`
	} `json:"response"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/check-risk", func(c *gin.Context) {
		var loc Location
		if err := c.BindJSON(&loc); err != nil {
			return
		}

		// 1. 外部APIを叩いて住所を特定（Goから外部へのHTTPリクエスト）
		// Heartrails Geo API (無料・登録不要) を使用
		apiURL := fmt.Sprintf("http://geoapi.heartrails.com/api/json?method=searchByGeo&x=%f&y=%f", loc.Lng, loc.Lat)
		resp, err := http.Get(apiURL)
		
		address := "住所特定不能（海上など）"
		if err == nil {
			defer resp.Body.Close()
			var geoResp GeoResponse
			// JSONをパース
			if err := json.NewDecoder(resp.Body).Decode(&geoResp); err == nil {
				if len(geoResp.Response.Location) > 0 {
					l := geoResp.Response.Location[0]
					address = l.Prefecture + l.City + l.Town
				}
			}
		}

		// 2. 結果を返却
		// 危険度はフロントエンドの地図（レイヤー）で視覚的に確認してもらう設計
		c.JSON(http.StatusOK, gin.H{
			"address": address,
			"message": "ハザードマップレイヤーを確認してください。",
		})
	})

	r.Run(":8080")
}
