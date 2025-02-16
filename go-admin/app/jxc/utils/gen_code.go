package utils

import (
	"fmt"
	"time"
)

// GenerateSalesOrderNumber 生成销售单号
func GenerateSalesOrderNumber(orderCount int64) string {
	orderCount += 1
	fixedChar := "XS"
	today := time.Now().Format("20060102")                               // 格式化为 YYYYMMDD
	orderNumber := fmt.Sprintf("%s%s%04d", fixedChar, today, orderCount) // 将数量格式化为四位数
	return orderNumber
}

// GeneratePurchaseOrderNumber 生成进货单号
func GeneratePurchaseOrderNumber(orderCount int64) string {
	orderCount += 1
	fixedChar := "JH"
	today := time.Now().Format("20060102")                               // 格式化为 YYYYMMDD
	orderNumber := fmt.Sprintf("%s%s%04d", fixedChar, today, orderCount) // 将数量格式化为四位数
	return orderNumber
}
