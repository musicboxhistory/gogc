package traffic

type TrafficInfo struct {
    Signal string
    Status int32
}

// Format
const (
    TimeFormat = "2006/01/02 15:04"
    KvsFormat = "%s_traffic_info"
)

// Default
const (
    DefaultExpire = 24
)

// signal
const (
    // 5geir
    EquipmentStatus  = "equipmentstatus"

    // Af

    // Amf

    // Ausf

    // Lmf

    // Nrf

    // Nssf

    // Pcf

    // Smf

    // Smsf

    // Udm

    // Udr
)