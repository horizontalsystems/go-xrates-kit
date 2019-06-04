package handler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/horizontalsystems/xrates-kit/pkg/config"
	httputil "github.com/horizontalsystems/xrates-kit/pkg/util/http"
)

//Ipfs handler object
type Ipfs struct {
	Conf *config.IpfsConfig
}

func (ipfs *Ipfs) GetLatestXRatesAsJSON(dCcy string, fCcy string, exchange string) (string, error) {

	respStr, err := httputil.DoGet(
		6, ipfs.Conf.URL, "ipns/"+ipfs.Conf.IpnsID+"/xrates/latest/"+fCcy+"/index.json", "")

	if err != nil {
		if strings.Contains(err.Error(), "Client.Timeout") {
			respStr, err := httputil.DoGet(
				6, ipfs.Conf.PublicURL, "ipns/"+ipfs.Conf.IpnsID+"/xrates/latest/"+fCcy+"/index.json", "")

			return respStr, err
		}
	}

	return respStr, err
}

func (ipfs *Ipfs) GetXRatesAsJSON(dCcy string, fCcy string, exchange string, epochSec *int64) (string, error) {

	var uriPathMinute, uriPathDay string

	timeSecObj := time.Unix(*epochSec, 0).UTC()
	year, month, day := timeSecObj.Date()
	hour := timeSecObj.Hour()

	uriPathDay = "ipns/" + ipfs.Conf.IpnsID +
		"/xrates/historical/" + dCcy + "/" + fCcy + "/" +
		strconv.Itoa(year) + "/" +
		fmt.Sprintf("%02d", int(month)) + "/" +
		fmt.Sprintf("%02d", day)

	uriPathMinute = uriPathDay + "/" + fmt.Sprintf("%02d", hour) + "/index.json"

	respStr, err := httputil.DoGet(6, ipfs.Conf.URL, uriPathMinute, "")

	if err != nil {

		if strings.Contains(err.Error(), "Client.Timeout") {
			respStr1, err := httputil.DoGet(6, ipfs.Conf.PublicURL, uriPathDay, "")

			if err != nil {
			}

			respStr = respStr1

		} else {
			// if Minute data is not found try By DAY.
			uriPathDay += "/index.json"

			respStr1, err := httputil.DoGet(6, ipfs.Conf.URL, uriPathDay, "")

			if err != nil {
			}

			respStr = respStr1
		}
	}

	dataS := reformatIPFSResponse(respStr, dCcy, fCcy, &timeSecObj)

	return dataS, err
}

//Change response json to XRates compatible format
func reformatIPFSResponse(jsonData string, dCcy string, fCcy string, timeObj *time.Time) string {

	var rate string

	if jsonData == "" {
		return ""
	}

	data := map[string]string{}
	minute := fmt.Sprintf("%02d", timeObj.Minute())
	err := json.Unmarshal([]byte(jsonData), &data)

	if err != nil {
		rate = strings.ReplaceAll(jsonData, "\"", "")
	} else {
		rate = data[minute]
	}

	rates := map[string]XRatesData{}
	rates[dCcy] = XRatesData{fCcy, rate, timeObj.Unix()}

	bytes, err := json.Marshal(rates)
	if err != nil {
	}

	return string(bytes)

}
