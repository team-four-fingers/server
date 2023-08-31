package config

import "server/pkg/env"

type Setting struct {
	PortNumber      string
	KakaoRESTAPIKey string
}

func NewSetting() *Setting {
	return &Setting{
		PortNumber:      env.MustGetEnvString("PORT", "8080"),
		KakaoRESTAPIKey: env.MustGetEnvString("KAKAO_REST_API_KEY", ""),
	}
}
