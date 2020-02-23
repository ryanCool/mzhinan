package main

type RestaurantInfo struct {
	Code       string     `json:"code"`
	Message    string     `json:"message"`
	Restaurant DetailInfo `json:"restaurant"`
}

type Config struct {
	Code            string     `json:"code"`
	Message         string     `json:"message"`
	Token           string     `json:"token"`
	Areas           []CatInfo  `json:"areas"`
	APIVersion      string     `json:"api_version"`
	Location        []Location `json:"location"`
	CuisineCat      []CatSet   `json:"cuisine_cat"`
	AzureBlobConfig struct {
		AccessKeyID     string `json:"accessKeyId"`
		AccessKeySecret string `json:"accessKeySecret"`
		Bucket          string `json:"bucket"`
		OuterInternet   string `json:"outerInternet"`
		InnerInternet   string `json:"innerInternet"`
	} `json:"azure_blob_config"`
	AzureImgURL         string `json:"azure_img_url"`
	AzureBucket         string `json:"azure_bucket"`
	SwitchLocationLogin int    `json:"switch_location_login"`
	Taste               string `json:"taste"`
	UpgradeInfo         string `json:"upgrade_info"`
	ShowLocationMore    int    `json:"show_location_more"`
	AccountBg           []struct {
		LocationID string `json:"location_id"`
		Img        string `json:"img"`
	} `json:"account_bg"`
	ShowAndroidUpdate int    `json:"show_android_update"`
	AndroidVersion    string `json:"android_version"`
	ShowIosUpdate     int    `json:"show_ios_update"`
	IosVersion        string `json:"ios_version"`
	LocationMore      struct {
		ID     string `json:"id"`
		NameEn string `json:"name_en"`
		NameZh string `json:"name_zh"`
		Image  string `json:"image"`
	} `json:"location_more"`
	OssBucket   string  `json:"oss_bucket"`
	Comfort     []Icon  `json:"comfort"`
	AllComfort  Icon    `json:"all_comfort"`
	AllSymbols  Icon    `json:"all_symbols"`
	AllStar     Icon    `json:"all_star"`
	AllCuisines CatInfo `json:"all_cuisines"`
	AllAreas    CatInfo `json:"all_areas"`
	Symbols     []Icon  `json:"symbols"`
	Star        []struct {
		ID    string `json:"id"`
		Index string `json:"index"`
		Icon  string `json:"icon"`
		Name  string `json:"name"`
	} `json:"star"`
	Cuisines       []CatInfo `json:"cuisines"`
	BaseURLImage   string    `json:"base_url_image"`
	InvitationCode string    `json:"invitation_code"`
	InvitationMode int       `json:"invitation_mode"`
	ArticleCat     []struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Children []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"children"`
	} `json:"article_cat"`
}

type DetailInfo struct {
	ID                 string      `json:"id"`
	Ga                 string      `json:"ga"`
	Name               string      `json:"name"`
	Cuisine            string      `json:"cuisine"`
	Tel                string      `json:"tel"`
	Star               string      `json:"star"`
	Comfort            interface{} `json:"comfort"`
	Website            string      `json:"website"`
	IsRed              string      `json:"is_red"`
	Price              Price       `json:"price"`
	Carousel           []string    `json:"carousel"`
	City               string      `json:"city"`
	District           string      `json:"district"`
	Area               string      `json:"area"`
	Address            string      `json:"address"`
	ShortAddress       string      `json:"short_address"`
	CoordinateAmap     string      `json:"coordinate_amap"`
	CoordinateOri      string      `json:"coordinate_ori"`
	OpeningHoursLunch  []OpenHour  `json:"opening_hours_lunch"`
	OpeningHoursDinner []OpenHour  `json:"opening_hours_dinner"`
	Holidays           string      `json:"holidays"`
	Desc               string      `json:"desc"`
	Preorder           string      `json:"preorder"`
	Refuse             string      `json:"refuse"`
	KoubeiID           string      `json:"koubei_id"`
	KoubeiURL          string      `json:"koubei_url"`
	AeURL              interface{} `json:"ae_url"`
	AeText             interface{} `json:"ae_text"`
	Owner              string      `json:"owner"`
	OwnerSince         string      `json:"owner_since"`
	OwnerExp           string      `json:"owner_exp"`
	Manager            string      `json:"manager"`
	ManagerSince       string      `json:"manager_since"`
	ManagerExp         string      `json:"manager_exp"`
	Chef               string      `json:"chef"`
	ChefSince          string      `json:"chef_since"`
	ChefExp            string      `json:"chef_exp"`
	Symbols            []string    `json:"symbols"`
	Capacity           string      `json:"capacity"`
	Nearby             []struct {
		ID        string `json:"id"`
		Distrance string `json:"distrance"`
	} `json:"nearby"`
	Articles  []interface{} `json:"articles"`
	ShareURL  string        `json:"share_url"`
	Set       []interface{} `json:"set"`
	HasSet    bool          `json:"has_set"`
	PhotoList []interface{} `json:"photo_list"`
	Title     string        `json:"title"`
}

type Restaurant struct {
	ID                 string      `json:"id"`
	Name               string      `json:"name"`
	Cuisine            string      `json:"cuisine"`
	Star               string      `json:"star"`
	Comfort            string      `json:"comfort"`
	IsRed              string      `json:"is_red"`
	DinnerMenuMin      string      `json:"dinner_menu_min"`
	DinnerMenuMax      string      `json:"dinner_menu_max"`
	DinnerSetMin       string      `json:"dinner_set_min"`
	DinnerSetMax       string      `json:"dinner_set_max"`
	LunchMenuMin       string      `json:"lunch_menu_min"`
	LunchMenuMax       string      `json:"lunch_menu_max"`
	LunchSetMin        string      `json:"lunch_set_min"`
	LunchSetMax        string      `json:"lunch_set_max"`
	CoordinateAmap     string      `json:"coordinate_amap"`
	Address            string      `json:"address"`
	ShortAddress       string      `json:"short_address"`
	City               string      `json:"city"`
	Area               string      `json:"area"`
	Thumbnail          string      `json:"thumbnail"`
	Symbols            []string    `json:"symbols"`
	RelatedArticleTime interface{} `json:"related_article_time"`
	HasSet             bool        `json:"has_set"`
	Desc               string      `json:"desc
"`
}

type Response struct {
	Code       string       `json:"code"`
	Message    string       `json:"message"`
	Restaurant []Restaurant `json:"restaurant"`
}

type CatInfo struct {
	ID    string `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}

type Icon struct {
	ID   string `json:"id"`
	Icon string `json:"icon"`
	Name string `json:"name"`
}

type CatSet struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Image    string    `json:"image"`
	Children []CatInfo `json:"children"`
}

type Location struct {
	ID         string `json:"id"`
	Code       string `json:"code"`
	NameZh     string `json:"name_zh"`
	NameEn     string `json:"name_en"`
	Image      string `json:"image"`
	Index      string `json:"index"`
	RegionCode string `json:"region_code"`
}

type RequireData struct {
	NameCn    string
	NameEn    string
	Address   string
	Tel       string
	Url       string
	Star      string
	Cat       string
	DescZh    string
	DescEn    string
	Lat       string
	Long      string
	OpenLunch string
	OpenDiner []string
	Price     Price
}

type OpenHour struct {
	Range string `json:"range"`
	Open  string `json:"open"`
	End   string `json:"end"`
	Lo    string `json:"lo"`
}

type Price struct {
	LunchSetMin   string `json:"lunch_set_min"`
	LunchSetMax   string `json:"lunch_set_max"`
	LunchMenuMin  string `json:"lunch_menu_min"`
	LunchMenuMax  string `json:"lunch_menu_max"`
	DinnerSetMin  string `json:"dinner_set_min"`
	DinnerSetMax  string `json:"dinner_set_max"`
	DinnerMenuMin string `json:"dinner_menu_min"`
	DinnerMenuMax string `json:"dinner_menu_max"`
}
