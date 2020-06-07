package privacy

import "github.com/iot-bc/eer-privacy/src/privacy"

func AESEncryptData(data string) string {
	key := privacy.GenerateRandomKey()
	dataEncrypted := privacy.AESEncryptData(key, data)
	return dataEncrypted
}

func AESDecryptData(data string) string {
	key := privacy.GenerateRandomKey()
	dataDecrypted := privacy.AESDecryptData(key, data)
	return dataDecrypted
}

func AnonymizeDevice(deviceName string) string {
	fakerDeviceName := privacy.AnonymizeDevice(deviceName)
	return fakerDeviceName
}
