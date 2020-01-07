package utils

const STAVUrlProd = "https://api.st-av.net/v1"
const STAVUrlStag = "https://api.s.st-av.net/v1"
const STAVUrlAccept = "https://api.a.st-av.net/v1"
const STAVUrlDev = "https://api.d.st-av.net/v1"

func GetAVPlatformUrl(env string) string {

	switch env {
	case "development":
		return STAVUrlDev
	case "staging":
		return STAVUrlStag
	case "acceptance":
		return STAVUrlAccept
	case "production":
		return STAVUrlProd

	default:
		return ""
	}
}
