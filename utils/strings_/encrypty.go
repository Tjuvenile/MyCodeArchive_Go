package strings_

import (
	"MyCodeArchive_Go/utils"
	"MyCodeArchive_Go/utils/logging"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/ssh"
)

func EncryptExample(str string) {
	encryptPasswd, err := RsaEncrypt(str, utils.RsaPublicKey)
	if err != nil {
		logging.Log.Errorf("failed to encrypt Password, err: %v", err)
		return
	}

	dePassword, errDecrypt := RsaDecrypt(encryptPasswd, utils.RsaPrivateKey)
	if errDecrypt != nil {
		logging.Log.Errorf("failed to decrypt Password, err: %v", errDecrypt)
		return
	}

	fmt.Printf("admin, encry: %s, decrypt: %s \n", encryptPasswd, dePassword)
}

/*
************************************************
函数名称:    RsaEncrypt
功能描述:    使用公钥加密
函数入参:    公钥，密码
函数出参:    无
返回结果:    加密后的密码
其他说明:    无
************************************************
*/
func RsaEncrypt(password string, publicKey string) (encryptedPassword string, err error) {
	logging.Log.Infof("start to encrypt password")
	if CheckIsPriKey(password) {
		return password, nil
	}
	// pem解码
	block, _ := pem.Decode([]byte(publicKey))
	// x509解码
	pubKeyInterface, errParsePubKey := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		logging.Log.Errorf("failed to do parse x509 for public key, err: %v", errParsePubKey)
		err = errParsePubKey
		return
	}
	// 对明文进行加密
	cipherText, errEncrypt := rsa.EncryptPKCS1v15(rand.Reader, pubKeyInterface.(*rsa.PublicKey), []byte(password))
	if errEncrypt != nil {
		logging.Log.Errorf("failed to do encrypt password, err: %v", errEncrypt)
		err = errEncrypt
		return
	}
	// base64编码
	encryptedPassword = base64.StdEncoding.EncodeToString(cipherText)
	logging.Log.Infof("end to encrypt password, encrypted password: %v", encryptedPassword)
	return
}

/*
************************************************
函数名称:    RsaDecrypt
功能描述:    使用私钥解密
函数入参:    使用公钥加密后的密码
函数出参:    无
返回结果:    解密后的密码
其他说明:    无
************************************************
*/
func RsaDecrypt(encryptedPassword string, privateKey string) (string, error) {
	logging.Log.Infof("start to decrypt password")
	if CheckIsPriKey(encryptedPassword) {
		return encryptedPassword, nil
	}
	// base64解码
	cipherText, errBase64Decode := base64.StdEncoding.DecodeString(encryptedPassword)
	if errBase64Decode != nil {
		logging.Log.Errorf("failed to decode base64 password, err: %v", errBase64Decode)
		return "", errBase64Decode
	}
	block, _ := pem.Decode([]byte(privateKey))
	// X509解码
	parsedPrivateKey, errParse := x509.ParsePKCS1PrivateKey(block.Bytes)
	if errParse != nil {
		logging.Log.Errorf("failed to parse private key, err: %v", errParse)
		return "", errParse
	}
	// decrypt解密
	decryptedBytes, errDecrypt := rsa.DecryptPKCS1v15(rand.Reader, parsedPrivateKey, cipherText)
	if errDecrypt != nil {
		logging.Log.Errorf("failed to decrypt password, err: %v", errDecrypt)
		return "", errDecrypt
	}
	logging.Log.Infof("end to decrypt password")
	return string(decryptedBytes), nil
}

// 函数名称: CheckPasswordIsPriKey
// 功能描述: 校验是不是标准的私钥
// 函数入参: password
// 函数出参: 无
// 返回结果: bool
// 其他说明:
func CheckIsPriKey(password string) bool {
	logging.Log.Infof("try to parse private key")
	_, err := ssh.ParsePrivateKey([]byte(password))
	if err != nil {
		logging.Log.Infof("unable to parse private key: %v", err)
		return false
	}
	logging.Log.Infof("this is private key")
	return true
}
