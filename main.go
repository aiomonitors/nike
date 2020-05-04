package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/aiomonitors/nike/product"
	"github.com/aiomonitors/nike/types"

	"github.com/aiomonitors/godiscord"
	"github.com/aiomonitors/gologger"
	proxymanager "github.com/aiomonitors/goproxymanager"
	"github.com/fatih/color"
)

type Config struct {
	Webhooks []godiscord.Webhook `json:"webhooks"`
	SKUs     []string            `json:"skus"`
	Color    string              `json:"color"`
	Links    []string            `json:"links"`
}

type Monitor struct {
	Config       Config                    `json:"config"`
	Availability map[string][]string       `json:"availability`
	Manager      proxymanager.ProxyManager `json:"manager"`
	UseProxies   bool                      `json:"use_proxies"`
	Client       http.Client               `json:"client"`
	ConfigPath   string                    `json:"config_path"`
	Products     []string                  `json:"products"`
}

var logger = gologger.Logger{Name: "nike"}
var l = sync.RWMutex{}
var skusLock = sync.RWMutex{}

var headers = map[string]string{
	"Host":            "api.nike.com",
	"user-agent":      "PostmanRuntime/7.24.1",
	"accept-language": "en-US,en;q=0.9",
	"Connection":      "close",
}

func remove(slice []string, s string) []string {
	for k, v := range slice {
		if v == s {
			slice = append(slice[:k], slice[k+1:]...)
		}
	}
	return slice
}

func difference(a, b []string) []string {

	target := map[string]bool{}
	for _, x := range b {
		target[x] = true
	}

	result := []string{}
	for _, x := range a {
		if _, ok := target[x]; !ok {
			result = append(result, x)
		}
	}

	return result
}

func NewMonitor(pathToConfig string, proxyPath string) (*Monitor, error) {
	m := Monitor{}

	//Proxy initialization
	if proxyPath != "" {
		manager, managerErr := proxymanager.NewManager(proxyPath)
		if managerErr == nil {
			m.Manager = *manager
			m.UseProxies = true
			logger.Green("Loaded %v proxies", len(m.Manager.Proxies))
		} else {
			m.UseProxies = false
			color.Red("Error loading proxy file")
		}
	} else {
		m.UseProxies = false
	}
	m.Client = http.Client{Timeout: 5 * time.Second}
	m.UpdateClient()
	//Config initialization
	if pathToConfig != "" {
		file, openErr := ioutil.ReadFile(pathToConfig)
		if openErr != nil {
			return nil, openErr
		}
		var c Config
		unmarshalError := json.Unmarshal(file, &c)
		if unmarshalError != nil {
			return nil, unmarshalError
		}
		m.Config = c
		m.ConfigPath = pathToConfig
	} else {
		return nil, errors.New("Need to provide a config file")
	}
	//Initialize links
	return &m, nil
}

func (m *Monitor) UpdateClient() error {
	if m.UseProxies == true {
		proxy, proxyErr := m.Manager.RandomProxy()
		if proxyErr != nil {
			return proxyErr
		}
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			return err
		}
		defaultTransport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		m.Client.Transport = defaultTransport
	}
	return nil
}

func (m *Monitor) UpdateHooks() {
	i := true
	for i == true {
		time.Sleep(7 * time.Second)
		logger.Yellow("Updating hooks")
		var obj Config
		file, _ := ioutil.ReadFile(m.ConfigPath)
		json.Unmarshal(file, &obj)
		l.Lock()
		m.Config.Webhooks = obj.Webhooks
		l.Unlock()
	}
}

func (m *Monitor) RefreshSKUs() {
	var obj Config
	file, _ := ioutil.ReadFile(m.ConfigPath)
	json.Unmarshal(file, &obj)
	skusLock.Lock()
	m.Config.SKUs = obj.SKUs
	skusLock.Unlock()
}

func (m *Monitor) GetProducts() ([]string, error) {
	client := &http.Client{}
	if m.UseProxies == true {
		proxy, proxyErr := m.Manager.RandomProxy()
		if proxyErr != nil {
			client = &http.Client{}
		}
		proxySplit := strings.Split(proxy, ":")
		proxy = fmt.Sprintf("http://%s:%s@%s:%s", proxySplit[2], proxySplit[3], proxySplit[0], proxySplit[1])
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			logger.Red("Error with proxy %v", err)
			client = &http.Client{}
		}
		defaultTransport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		client = &http.Client{Transport: defaultTransport}
	}

	req, reqErr := http.NewRequest("GET", "https://api.nike.com/product_feed/rollup_threads/v2?filter=marketplace%28US%29&filter=language%28en%29&filter=employeePrice%28true%29&filter=attributeIds%280f64ecc7-d624-4e91-b171-b83a03dd8550%2C16633190-45e5-4830-a068-232ac7aea82c%2C53e430ba-a5de-4881-8015-68eb1cff459f%29&anchor=0&consumerChannelId=d9a5bc42-4b9c-4976-858a-f159cf99c647&count=25", nil)
	if reqErr != nil {
		logger.Red("Error in the request %s", reqErr)
		return nil, reqErr
	}
	req.Close = true

	// set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, resError := client.Do(req)
	if resError != nil {
		logger.Red("Error in the request %s", resError)
		return nil, resError
	}
	defer res.Body.Close()

	var obj types.NewProducts
	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		logger.Red("Error reading body %s", bodyErr)
		return nil, bodyErr
	}
	json.Unmarshal(body, &obj)

	ids := []string{}
	for _, item := range obj.Objects {
		ids = append(ids, item.ProductInfo[0].MerchProduct.StyleColor)
	}
	return ids, nil
}

func (m *Monitor) GetProduct(styleColor string) (types.ProductInfo, error) {
	client := &http.Client{}
	if m.UseProxies == true {
		proxy, proxyErr := m.Manager.RandomProxy()
		if proxyErr != nil {
			client = &http.Client{}
		}
		proxySplit := strings.Split(proxy, ":")
		proxy = fmt.Sprintf("http://%s:%s@%s:%s", proxySplit[2], proxySplit[3], proxySplit[0], proxySplit[1])
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			logger.Red("Error with proxy %v", err)
			client = &http.Client{}
		}
		defaultTransport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		client = &http.Client{Transport: defaultTransport}
	}

	start := time.Now()
	emptyResp := types.ProductInfo{}
	req, reqErr := http.NewRequest("GET", fmt.Sprintf("https://api.nike.com/product_feed/threads/v2?filter=channelId(d9a5bc42-4b9c-4976-858a-f159cf99c647)&filter=marketplace(US)&filter=language(en)&filter=productInfo.merchProduct.styleColor(%s)", styleColor), nil)
	if reqErr != nil {
		logger.Red("Error in the request %s", reqErr)
		return emptyResp, reqErr
	}
	req.Close = true

	// set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	res, resError := client.Do(req)
	if resError != nil {
		logger.Red("Error in the request %s", resError)
		return emptyResp, resError
	}
	defer res.Body.Close()

	var rawObj product.ProductJson
	body, bodyErr := ioutil.ReadAll(res.Body)
	if bodyErr != nil {
		logger.Red("Error reading body %s", bodyErr)
		return emptyResp, bodyErr
	}
	json.Unmarshal(body, &rawObj)
	obj := rawObj.Objects[0]

	Product := types.ProductInfo{}
	Product.Name = obj.ProductInfo[0].ProductContent.FullTitle + " " + obj.ProductInfo[0].ProductContent.ColorDescription
	Product.Style = obj.ProductInfo[0].MerchProduct.StyleColor
	Product.Link = fmt.Sprintf("https://www.nike.com/t/%s/%s", obj.ProductInfo[0].ProductContent.Slug, Product.Style)
	Product.Price = fmt.Sprintf("%v", obj.ProductInfo[0].MerchPrice.FullPrice)
	Product.Image = obj.PublishedContent.Nodes[0].Nodes[0].Properties.SquarishURL

	//SKUS Stuff
	Product.AvailableSizes = make([]string, 0)
	Product.DiscordSKUs = make([]string, 0)

	SKUsAvail := map[string]bool{}   //Map of availability by sku id
	SKUsStock := map[string]string{} // Map of stock by sku id
	for _, sku := range obj.ProductInfo[0].AvailableSkus {
		SKUsAvail[sku.ID] = sku.Available
		SKUsStock[sku.ID] = sku.Level
	}

	for _, sku := range obj.ProductInfo[0].Skus {
		if SKUsAvail[sku.ID] {
			Product.AvailableSizes = append(Product.AvailableSizes, sku.NikeSize)
			spaces := ""
			for i := 5 - len(sku.NikeSize); i >= 0; i-- {
				spaces += " "
			}
			Product.DiscordSKUs = append(Product.DiscordSKUs, fmt.Sprintf("**%s**%s[%s]", sku.NikeSize, spaces, SKUsStock[sku.ID]))
		}
	}

	Product.Exec = fmt.Sprintf("%v", time.Since(start))
	return Product, nil
}

func (m *Monitor) Initialize() {
	s, sErr := m.GetProducts()
	if sErr != nil {
		logger.Red("Error initializing, retrying in .5 secs")
		time.Sleep(500 * time.Millisecond)
		m.Initialize()
	}
	m.Products = s
	logger.Green("Initialized monitor with %v products", len(m.Products))
	return
}

func (m *Monitor) InitializeSKUs() {
	var wg sync.WaitGroup
	wg.Add(len(m.Config.SKUs))
	m.Availability = map[string][]string{}

	start := time.Now()

	for _, sku := range m.Config.SKUs {
		go func(sku string) {
			defer wg.Done()
			s, sErr := m.GetProduct(sku)
			if sErr != nil {
				logger.Red("Error initializing %s", sku)
				m.Config.SKUs = remove(m.Config.SKUs, sku)
				return
			}
			l.Lock()
			m.Availability[sku] = s.AvailableSizes
			l.Unlock()
			logger.Green("[%s] Initialized %s", s.Exec, s.Name)
			return
		}(sku)
	}
	wg.Wait()
	logger.Green("[%v] Initialized %v skus", time.Since(start), len(m.Config.SKUs))
}

func (m *Monitor) InitializeSKU(sku string) {
	s, sErr := m.GetProduct(sku)
	if sErr != nil {
		logger.Red("Error initializing %s", sku)
		m.Config.SKUs = remove(m.Config.SKUs, sku)
		return
	}
	m.Availability[sku] = s.AvailableSizes
	logger.Green("[%s] Initialized %s", s.Exec, s.Name)
	return
}

func (m *Monitor) MonitorSKU(sku string) {
	s, sErr := m.GetProduct(sku)
	if sErr != nil {
		logger.Red("Error monitoring %s")
		return
	}
	l.Lock()
	diff := difference(s.AvailableSizes, m.Availability[sku])
	l.Unlock()
	if len(diff) > 0 {
		logger.Green("%s Restocked [S:%s]", s.Name, sku)
		s.Notification = "Restock"
		go m.SendToDiscord(s)
	}
	l.Lock()
	m.Availability[sku] = s.AvailableSizes
	l.Unlock()
	logger.Blue("[%s] Monitored %s", s.Exec, s.Name)
	return
}

func (m *Monitor) MonitorNew() {
	i := true
	for i == true {
		time.Sleep(750 * time.Millisecond)
		go func() {
			start := time.Now()
			s, sErr := m.GetProducts()
			if sErr != nil {
				logger.Red("Error monitoring products: %v", sErr)
				return
			}
			diff := difference(s, m.Products)
			if len(diff) > 1 {
				logger.Green("%v New Products found", len(diff))
				for _, product := range diff {
					go m.NewProduct(product)
				}
				l.Lock()
				m.Products = append(m.Products, diff...)
				l.Unlock()
			}
			logger.Blue("[%v] Monitored", time.Since(start))
			return
		}()
	}
}

func (m *Monitor) MonitorSKUs() {
	i := true
	var wg sync.WaitGroup
	for i == true {
		m.RefreshSKUs()
		rand.Seed(time.Now().UnixNano())
		wg.Add(len(m.Config.SKUs))
		skusLock.Lock()
		for _, sku := range m.Config.SKUs {
			go func(sku string) {
				defer wg.Done()
				time.Sleep(750 * time.Millisecond)
				if _, exists := m.Availability[sku]; exists {
					m.MonitorSKU(sku)
				} else {
					m.InitializeSKU(sku)
				}
				return
			}(sku)
		}
		skusLock.Unlock()
		wg.Wait()
		continue
	}
}

func (m *Monitor) NewProduct(styleCode string) {
	p, pErr := m.GetProduct(styleCode)
	if pErr != nil {
		logger.Red("Erorr fetching %s [NEW]", styleCode)
		return
	}
	p.Notification = "New"
	m.SendToDiscord(p)
}

func (m *Monitor) SendToDiscord(p types.ProductInfo) {
	for _, webhook := range m.Config.Webhooks {
		go func(p *types.ProductInfo, webhook *godiscord.Webhook) {
			emb := godiscord.NewEmbed(p.Name, "", p.Link)
			emb.AddField("Price", fmt.Sprintf("**$%s**", p.Price), true)
			emb.AddField("Type", p.Notification, true)
			emb.AddField("Style", p.Style, false)
			if len(p.DiscordSKUs) > 6 {
				emb.AddField("Sizes", strings.Join(p.DiscordSKUs[:len(p.DiscordSKUs)/2], "\n"), true)
				emb.AddField("Sizes", strings.Join(p.DiscordSKUs[len(p.DiscordSKUs)/2:], "\n"), true)
			} else if len(p.DiscordSKUs) > 0 {
				emb.AddField("Sizes", strings.Join(p.DiscordSKUs, "\n"), false)
			}
			emb.SetThumbnail(p.Image)
			if len(webhook.Color) > 0 {
				emb.SetColor(webhook.Color)
			} else {
				emb.SetColor("#7f5af0")
			}
			emb.SetAuthor("Nike US", "", "")
			emb.SetFooter(webhook.Text, webhook.IconURL)
			emb.SendToWebhook(webhook.URL)
		}(&p, &webhook)
	}
}

func (m *Monitor) Start() {
	logger.Green("Starting monitor")
	go m.Initialize()
	m.InitializeSKUs()
	go m.RefreshSKUs()
	//go m.UpdateHooks()
	go m.MonitorNew()
	m.MonitorSKUs()
}

func main() {
	m, mErr := NewMonitor("config.json", "./proxies.txt")
	if mErr != nil {
		panic(mErr)
	}

	m.Start()
	// p, pErr := m.GetProduct("852542-011")
	// if pErr != nil {
	// 	panic(pErr)
	// }
	// fmt.Println(p)
	// p.Notification = "Restock"
	// m.SendToDiscord(p)
	// 	time.Sleep(2000 * time.Millisecond)
	// 	skusLock.Lock()
	// 	m.Availability["852542-011"] = m.Availability["852542-011"][2:]
	// 	skusLock.Unlock()
	// 	logger.Red("Changed")
	// 	time.Sleep(3000 * time.Millisecond)
	// }
}
