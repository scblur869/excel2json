package _models

type CloudRoomOverview struct {
	Id                 int     `json:"id"`
	CloudId            string  `json:"cloudid"`
	Name               string  `json:"name"`
	Provider           string  `json:"provider"`
	Tenant             string  `json:"tenant"`
	CreateDate         string  `json:"createdate"`
	State              string  `json:"state"`
	ICTObject          string  `json:"ictobject"`
	AppName            string  `json:"appname"`
	AppId              string  `json:"appid"`
	ITResp             string  `json:"itresp"`
	CloudroomResp      string  `json:"cloudroomresp"`
	OrgResp            string  `json:"orgresp"`
	MainDept           string  `json:"maindept"`
	CC                 string  `json:"cc"`
	PSP                string  `json:"psp"`
	BMWOrg             string  `json:"bmworg"`
	FS                 string  `json:"fs"`
	SecondaryResp      string  `json:"secondaryresp"`
	ManagerResp        string  `json:"managerresp"`
	InvQuality         string  `json:"invquality"`
	ProductId          string  `json:"productid"`
	ProductOwner       string  `json:"productowner"`
	AppRespEmail       string  `json:"apprespemail"`
	PrimaryRespEmail   string  `json:"primaryrespemail"`
	SecondaryRespEmail string  `json:"secondaryrespemail"`
	ManagerRespEmail   string  `json:"managerrespemail"`
	LSV                float32 `json:"lsv"`  // used for combined report
	Cost               float32 `json:"cost"` // used for combined report
	Created            string  `json:"created"`
}

type CloudCostOverview struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	State         string  `json:"state"`
	CSP           string  `json:"csp"`
	AppId         string  `json:"appid"`
	AYTD          float32 `json:"ayod"`
	Periods       int     `json:"periods"`
	AMA           float32 `json:"ama"`
	ALM           float32 `json:"alm"`
	FCP           float32 `json:"fcp"`
	FFY           float32 `json:"ffy"`
	Currency      string  `json:"currency"`
	ITResp        string  `json:"itresp"`
	CloudroomResp string  `json:"cloudroomresp"`
	Created       string  `json:"created"`
}

type LSVData struct {
	Id            int    `json:"id"`
	Prime         string `json:"prime"`
	PrimeHal      string `json:"primehal"`
	MarketOrPlant string `json:"marketorplant"`
	FGXReporting  string `json:"fgxreporting"`
	EOY           int    `json:"eoy"`
	Status        string `json:"status"`
	EndDate       string `json:"enddate"`
	FMM           string `json:"fmm"`
	TMM           string `json:"tmm"`
	VM            string `json:"vm"`
	Application   string `json:"application"`
	AppName       string `json:"appname"`
	CMDBID        string `json:"cmdbid"`
	AppStatus     string `json:"appstatus"`
	Migrated      string `json:"migrated"`
	Dept          string `json:"dept"`
	Consumption   string `json:"consumption"`
	Created       string `json:"created"`
}
