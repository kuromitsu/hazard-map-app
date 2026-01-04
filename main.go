package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Location struct {
    Lat float64 `json:"lat"`
    Lng float64 `json:"lng"`
}

func main() {
    r := gin.Default()

    // HTMLを表示する設定
    r.LoadHTMLGlob("templates/*")
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    // 危険度判定API
    r.POST("/check-risk", func(c *gin.Context) {
        var loc Location
        if err := c.BindJSON(&loc); err != nil {
            return
        }

        // ここに判定ロジック（簡易版）
        riskLevel := "安全"
        details := "特に危険な情報は報告されていません。"

        // 例えば特定の座標なら「危険」とするなどのロジックを入れる
        // 今回はデモ用にランダムや固定値でもOK

        c.JSON(http.StatusOK, gin.H{
            "risk": riskLevel,
            "details": details,
        })
    })

    r.Run(":8080")
}