package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const token = "8F3E6788-7DA2-3F7D-4C11-7F25348E8713"

var finalData = map[string][]RequireData{}

var citymap = map[string]string{
	"6B70D2A3-8061-F37D-73DB-355134A0521F": "上海",
	"8D553380-2180-440A-8E3E-6F926CCB77F6": "北京",
	"834FA3FA-1B9F-4771-B43B-A43F3FE940CC": "臺灣",
	"ECCFF058-9506-4187-536C-EE8A39B5D79D": "香港",
	"38E469DD-34B7-347A-9C77-D0B1F326B0C5": "澳門",
	"E9E7F1CC-2414-4FB1-AC40-282424CC36ED": "廣州",
}

var catMap = map[string]string{
	"A69483F9-8753-D289-22DA-166AE175D7DF": "粵菜",
	"04E94EFB-ED3A-FFD1-F9F3-D2D563767D42": "法國菜",
	"1A502C9D-1CC4-E6AF-A6F3-3DB04971B3F5": "時尚歐陸菜",
	"209258ED-77C7-E705-0DC3-F08CC3F2CDF5": "義大利菜",
	"487110D6-671C-5842-BC22-8409BB78E032": "瑞士",
	"A29EF03F-7E91-F075-DBDD-C5360F6427DA": "時尚法國菜",
	"BA2BC26C-A4FA-4295-80E0-F94A97ACBFF6": "祕魯菜",
	"CC92D10A-9201-4226-9151-14569F510EEE": "北歐菜",
	"D928DAC2-637C-5D24-D450-13E133B83767": "扒房",
	"F3A5E624-BD7B-4DC7-97B6-A163B329CE64": "餃子",
	"FD7625CA-C216-6873-D5DC-10338D2AC325": "美國義大利菜",
	"121FB7A3-2B22-4B56-B32C-28881CCDCFE9": "落林菜",
	"80EA99CA-38DD-403A-8B5F-2361EFA293A9": "里昂菜",
	"A8D610F0-E4C7-4E85-900A-5FF6189A24E7": "海鮮",
	"B2F05785-CFB1-46BD-984A-5C14D540A275": "依善菜",
	"F46081D9-7344-7562-EEDA-02F59C8F61EB": "海鮮",
	"03C659A1-60DE-EB53-E647-81608C5D4227": "臺灣菜",
}

var scoreMap = map[string]string{
	"0590FDA9-9DC0-4ACE-BFEA-83A57FDFB15F": "三星",
	"DC960ECC-E15B-46E2-89F5-B44EC31743F1": "二星",
	"D2CC5438-B44C-436B-9D9C-4BEC4564CE79": "一星",
	"DE23F387-79CE-4992-9398-697AE70F017C": "必比登",
	"D87A87B3-782B-4C02-83FE-C0467D934923": "餐盤",
	"58B8EF0B-024A-4581-831B-0B73AA35058D": "街頭美食",
}

func getCityMap(config Config) {
	for _, cuisine := range config.Cuisines {
		catMap[cuisine.ID] = cuisine.Name

	}
}

func getCatMap(config Config) {
	for _, location := range config.Location {
		citymap[location.ID] = location.NameZh
	}
}

func getConfig() {

	url := "https://api.mzhinan.cn/api/"
	method := "POST"

	payload := strings.NewReader("method=system.token&location=ECCFF058-9506-4187-536C-EE8A39B5D79D&uuid=3BC6EA43-5551-4EB6-8D12-557AD4AD3E3E&language=zh&year=2019&token=" + token + "&server_tag=200")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "MGC/9.5.8 CFNetwork/1121.2.2 Darwin/19.3.0")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-tw")
	req.Header.Add("Cache-Control", "no-cache")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	config := Config{}
	err = json.Unmarshal(body, &config)
	if err != nil {
		panic(err)
	}

	getCatMap(config)
	getCityMap(config)
}
func getInfoFromShareLink(shareLink string) (tel, descZh, openHour, price string) {
	fmt.Println(shareLink)
	res, err := http.Get(shareLink)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div[class=phone-info]").Each(func(i int, selection *goquery.Selection) {
		tel = selection.Text()
		//fmt.Println(tel)
	})

	doc.Find("span[class=article-quote]").Each(func(i int, selection *goquery.Selection) {

		descZh = selection.Text()
	})

	htmlText, _ := doc.Html()
	ss := strings.Split(htmlText, `<div class="opening-hours-item border-bottom">`)

	//get price
	pricess := strings.Split(ss[0], `<div class="price-item border-bottom">`)
	lunchss := strings.Split(pricess[1], `<div class="price-subtitle">晚膳</div>`)
	lunchPricedoc, err := goquery.NewDocumentFromReader(strings.NewReader(lunchss[0]))
	if err != nil {
		panic(err)
	}

	lunchTitle := []string{}
	lunchPricedoc.Find("div[class=price-info]").Each(func(i int, selection *goquery.Selection) {
		lunchTitle = append(lunchTitle, selection.Text())
	})

	lunchInfo := []string{}
	lunchPricedoc.Find("div[class=price-number]").Each(func(i int, selection *goquery.Selection) {
		lunchInfo = append(lunchInfo, selection.Text())
	})

	var dinnerTitle, dinnerInfo []string
	if len(lunchss) > 1 {
		dinerPricedoc, err := goquery.NewDocumentFromReader(strings.NewReader(lunchss[1]))
		if err != nil {
			panic(err)
		}

		dinerPricedoc.Find("div[class=price-info]").Each(func(i int, selection *goquery.Selection) {
			dinnerTitle = append(dinnerTitle, selection.Text())
		})

		dinerPricedoc.Find("div[class=price-number]").Each(func(i int, selection *goquery.Selection) {
			dinnerInfo = append(dinnerInfo, selection.Text())
		})
	}

	price = "lunch\n"
	for i := range lunchTitle {
		price = price + lunchTitle[i] + "\n" + lunchInfo[i] + "\n"
	}
	price = price + "dinner\n"
	for i := range dinnerTitle {
		price = price + dinnerTitle[i] + "\n" + dinnerInfo[i] + "\n"
	}

	sss := strings.Split(ss[1], `<div class="price-subtitle">晚膳</div>`)
	//sss[0] lunch open hour

	domm, err := goquery.NewDocumentFromReader(strings.NewReader(sss[0]))
	if err != nil {
		log.Fatalln(err)
	}
	openLunch := []string{}
	domm.Find("div[class=opening-label-closed]").Each(func(i int, selection *goquery.Selection) {
		openLunch = append(openLunch, strings.TrimSpace(selection.Text()))
	})

	openLunchTime := []string{}
	domm.Find("div[class=opening-info-time]").Each(func(i int, selection *goquery.Selection) {
		exist := strings.Contains(selection.Text(), "-")
		if exist {
			openLunchTime = append(openLunchTime, selection.Text())
		}
	})
	//fmt.Println("Lunch")
	//for i := range openLunchTime {
	//	fmt.Println(openLunch[i], openLunchTime[i])
	//}
	//
	//if len(openLunchTime) == 0 {
	//	fmt.Println(sss[0])
	//}

	ssss := strings.Split(sss[1], `<!-- ICON and TEXT -->`)
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(ssss[0]))
	if err != nil {
		log.Fatalln(err)
	}
	openDiner := []string{}
	dom.Find("div[class=opening-label-closed]").Each(func(i int, selection *goquery.Selection) {
		openDiner = append(openDiner, strings.TrimSpace(selection.Text()))
	})
	openDinerTime := []string{}
	dom.Find("div[class=opening-info-time]").Each(func(i int, selection *goquery.Selection) {
		exist := strings.Contains(selection.Text(), "-")
		if exist {
			openDinerTime = append(openDinerTime, selection.Text())
		}
	})
	//
	//fmt.Println("Dinner")
	//if len(openDiner) == len(openDinerTime) {
	//	for i := range openDiner {
	//		fmt.Println(openDiner[i], openDinerTime[i])
	//	}
	//} else {
	//	fmt.Println(openDiner)
	//	fmt.Println(openDinerTime)
	//	fmt.Println(ssss[0])
	//}
	//if len(openDiner) == 0 {
	//	fmt.Println(ssss[0])
	//}
	OpenDay := "Lunch\n"
	for i := range openLunch {
		OpenDay = OpenDay + openLunch[i] + "\n" + openLunchTime[i] + "\n"
	}

	OpenDay = OpenDay + "Dinner\n"
	for i := range openDiner {
		OpenDay = OpenDay + openDiner[i] + "\n" + openDinerTime[i] + "\n"
	}
	//fmt.Println(OpenDay)
	//fmt.Println(price)

	return tel, descZh, OpenDay, price
}

func getFromEnLink(shareLink string) (enName string, descEn string) {
	//fmt.Println(shareLink)
	res, err := http.Get(shareLink)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("title").Each(func(i int, selection *goquery.Selection) {
		enName = selection.Text()
		//fmt.Println(enName)
	})

	doc.Find("span[class=article-quote]").Each(func(i int, selection *goquery.Selection) {

		descEn = selection.Text()
	})

	return enName, descEn
}

var done = make(chan bool)

func main() {

	getConfig()
	fmt.Println("getConfig success")
	file, err := os.Create("result.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	resaurants := getResaurant()

	for k, v := range resaurants {
		cityName := citymap[k]
		for _, resaurant := range v {
			// Request the HTML page.

			//fmt.Println(resaurant.Name)
			shareLinkZh := "https://www.mzhinan.cn/share/restaurant/" + resaurant.ID + "/zh"
			shareLinkEn := "https://www.mzhinan.cn/share/restaurant/" + resaurant.ID + "/en"

			go func() {
				tel, descZh, open, price := getInfoFromShareLink(shareLinkZh)

				enName, descEn := getFromEnLink(shareLinkEn)

				ls := strings.Split(resaurant.CoordinateAmap, ",")

				finalData[cityName] = append(finalData[cityName], RequireData{
					NameCn:  resaurant.Name,
					NameEn:  enName,
					Address: resaurant.Address,
					Tel:     tel,
					Url:     shareLinkZh,
					Star:    scoreMap[resaurant.Star],
					Cat:     catMap[resaurant.Cuisine],
					DescZh:  descZh,
					DescEn:  descEn,
					Lat:     ls[1],
					Long:    ls[0],
					OpenDay: open,
					Price:   price,
				})
			}()

			<-done
		}
	}

	for k, v := range finalData {
		err := writer.Write([]string{"city ===========" + k})
		if err != nil {
			panic(err)
		}
		for _, val := range v {
			err := writer.Write([]string{
				val.NameCn,
				val.NameEn,
				val.Address,
				val.Tel,
				val.Url,
				val.Star,
				val.Cat,
				val.DescZh,
				val.DescEn,
				val.Lat,
				val.Long,
				val.OpenDay,
				val.Price,
			})
			if err != nil {
				panic(err)
			}
		}
	}

}

func getResaurant() map[string][]Restaurant {
	jsonFile, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	resp := Response{}
	err = json.Unmarshal(byteValue, &resp)
	if err != nil {
		panic(err)
	}

	cityMap := map[string][]Restaurant{}

	for _, restaurant := range resp.Restaurant {
		if _, ok := cityMap[restaurant.City]; !ok {
			cityMap[restaurant.City] = []Restaurant{}
		}
		cityMap[restaurant.City] = append(cityMap[restaurant.City], restaurant)
	}

	return cityMap
}
