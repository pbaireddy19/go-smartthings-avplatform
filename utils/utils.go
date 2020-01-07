package utils

const STAVUrlProd = "https://api.st-av.net/v1"
const STAVUrlStag = "https://api.s.st-av.net/v1"
const STAVUrlAccept = "https://api.a.st-av.net/v1"
const STAVUrlDev = "https://api.d.st-av.net/v1"

func GetAVPlatformUrl(env string) string {

	switch env {
	case "DEV":
		return STAVUrlDev
	case "ACCEPT":
		return STAVUrlAccept
	case "PROD":
		return STAVUrlProd
	case "STAG":
		return STAVUrlStag
	default:
		return ""
	}
}
