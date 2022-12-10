package traffic

import (
    "gogc/src/common/timer"
    "gogc/src/common/kvs"
    "time"
    "os"
    "fmt"
    "strconv"
)

// Traffic Data
var trafficData map[string]int

func Init() {

    go Collect()
}

func Collect() {

    nowTime := timer.Now()
    startTime := nowTime.Format(TimeFormat)
    nfType := os.Getenv("ENV_NF_TYPE")
    key := fmt.Sprintf(KvsFormat, nfType)
    expire, _ := strconv.Atoi(os.Getenv("ENV_TRAFFIC_EXPIRE"))
    if expire == 0 {
        expire = DefaultExpire
    }

    for {

        nowTime = timer.Now()
        periodTime := nowTime.Format(TimeFormat)

        if startTime != periodTime {
            // Set Kvs
            kvs.HSet(key, startTime, trafficData)

            // Init Data
            trafficData = map[string]int{}
            startTime = periodTime
        }

        time.Sleep(time.Second * 1)
    }
}

func SetTraffic(trafficInfo TrafficInfo) {

    signal := fmt.Sprintf("%s[%v]", trafficInfo.Signal, trafficInfo.Status)
    trafficData[signal]++
}