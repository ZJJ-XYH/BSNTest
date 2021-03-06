package config

const (
	_MspDir    = "D:/test/bsn-sdk-go/msp"
	_HttpsCert = "D:/github.com/BSNDA/PCNGateway-Go-SDK/test/cert/bsn_gateway_https.crt"
)

func NewMockFabricConfig() (*Config, error) {
	api := "http://beijingnode.bsngate.com:17502"
	userCode := "USER0001202007161739119605411"
	appCode := "app0001202007291443281737652"

	privK := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQg/7RMFXO8U9LyrTJW
EZ3gtdUI5A5K+yPAEb3iiPe7bKegCgYIKoEcz1UBgi2hRANCAASvJdHvty4qiZ2r
xcDYrMrgskyr6vthAy/Tgz/3S6SR/9ERuYVLh+Hzb6ptpIWHo0ek5j05ERh5vSzC
PIXILYkE
-----END PRIVATE KEY-----`

	return NewConfig2(api, userCode, appCode, privK, _MspDir)
}

func NewMockFiscoConfig() (*Config, error) {
	api := "http://beijingnode.bsngate.com:17502"
	userCode := "USER0001202007161739119605411"
	appCode := "app0001202007291443281737652"

	privK := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQg/7RMFXO8U9LyrTJW
EZ3gtdUI5A5K+yPAEb3iiPe7bKegCgYIKoEcz1UBgi2hRANCAASvJdHvty4qiZ2r
xcDYrMrgskyr6vthAy/Tgz/3S6SR/9ERuYVLh+Hzb6ptpIWHo0ek5j05ERh5vSzC
PIXILYkE
-----END PRIVATE KEY-----`

	return NewConfig2(api, userCode, appCode, privK, _MspDir)
}

func NewMockXuperchainConfig() (*Config, error) {
	api := "http://192.168.1.43:17502"
	userCode := "USER0001202010201539390086090"
	appCode := "app0001202010221038364886804"

	privK := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgikcrsA9vRp14VAl0
lhNmOMc7pl0j4xlF9Eno+eJNgSegCgYIKoEcz1UBgi2hRANCAAQuR25rig9+Isir
q7eapxSsDCo8FCdt0qfCa7eGCC7BEdGerbX4qa5j3qljl1/d6I00c7veFAvgGPP8
DFT+md7r
-----END PRIVATE KEY-----`

	return NewConfig2(api, userCode, appCode, privK, _MspDir)
}
func NewMockCitaConfig() (*Config, error) {
	api := "http://192.168.1.43:17502"
	userCode := "USER0001202007151006209656179"
	appCode := "app0001202010281411370741443"

	privK := `-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqBHM9VAYItBHkwdwIBAQQgVykzw4O/mZaJ8kJ8
LXX9bzpPVeXWGggYzVGeMuZM29ugCgYIKoEcz1UBgi2hRANCAASVecisbIxDUFXj
Ra6Rzr88PrjoEekdnqYzcopJhFIpt2mfLcU+QR9IvcyT/BBprDJl3rdWx92Z4/zv
tU4l8jft
-----END PRIVATE KEY-----`
	SetTest()
	return NewConfig2(api, userCode, appCode, privK, _MspDir)
}
