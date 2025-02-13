package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	. "dacV2"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	keyPrivate int64 = iota + 1 // 1

)

type SpaceCipher struct {
	key   []byte
	block cipher.Block
}

func InitToken(dirPath ...string) (spaceCipher *SpaceCipher , err error) {


	newCipher := &SpaceCipher{}
	
	mapFields, sizeField,  err := CreateMap([]SpaceList{
		{Name: keyPrivate, Len: 32},
	})
	if err != nil {
		return nil , err
	}

	gCipher, err := NewSpace(mapFields, sizeField, nil, 0, dirPath...)
	if err != nil {
		return nil , err
	}

	data, err := gCipher.GetFieldRaw(keyPrivate)
	if err != nil {
		return nil , err
	}

	if len(data) == 0 {

		bufferBytes := make([]byte, 32)
		_, err := rand.Read(bufferBytes)
		if err != nil {
			return nil , err
		}

		newCipher.key = bufferBytes
		gCipher.SetFieldRaw(keyPrivate, bufferBytes)

	} else {

		newCipher.key, err = gCipher.GetFieldRaw(keyPrivate)
		if err != nil {
			return nil , err
		}
	}

	//Crea una nueva clave de cifrado
	newCipher.block, err = aes.NewCipher(newCipher.key)
	if err != nil {
		return nil , err
	}

	return newCipher , nil
}

func (spaceCipher *SpaceCipher)Encripter(bufferBytes []byte) ([]byte , error) {

	if len(bufferBytes) == 0 {
		bufferBytes = nil
		return nil , errors.New("tamaño de buffer 0")
	}

	//El texto debe ser cifrado en bloques de 16 bytes
	//Añadimos padding hasta completar.
	if count := len(bufferBytes) % aes.BlockSize; count != 0 {

		count = aes.BlockSize - count

		padding := len(bufferBytes) + count

		bufferBytes = SpacePadding(bufferBytes, [3]int64{0, int64(padding), int64(padding)})

	}

	//Creamos un array con el tamaño de bloque + el texto
	ciphertext := make([]byte, aes.BlockSize+len(bufferBytes))

	//Añadimos bytes aleatorios al principio del texto
	_, err := rand.Read(ciphertext[:aes.BlockSize])
	if err != nil {
		return nil , err
	}

	mode := cipher.NewCBCEncrypter(spaceCipher.block, ciphertext[:aes.BlockSize])

	mode.CryptBlocks(ciphertext[aes.BlockSize:], bufferBytes)

	bufferBytes = ciphertext

	return bufferBytes , nil
}

func (spaceCipher *SpaceCipher)Dencripter(ciphertext []byte)([]byte , error){

	if len(ciphertext) < aes.BlockSize {
		return nil , errors.New("el tamaño del texto es menos que el tamaño de bloque de aes")
	}

	vectorInit := (ciphertext)[:aes.BlockSize]

	ciphertext = (ciphertext)[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {

		return nil , errors.New("el tamaño del texto de cifrado no es divisible entre el bloque de aes")
	}

	mode := cipher.NewCBCDecrypter(spaceCipher.block, vectorInit)

	mode.CryptBlocks(ciphertext, ciphertext)

	ciphertext = SpaceTrimNull(ciphertext)

	return ciphertext , nil
}

func (spaceCipher *SpaceCipher)EncodeCipherTokenBase64(params ...string) (cipherJWT string , err error) {

	if len(params) == 0 {
		return
	}

	for ind, param := range params {
		if param == "" {
			return
		}
		params[ind] = base64.StdEncoding.EncodeToString([]byte(param))
	}

	var JWT []byte
	if len(params) > 1 {

		JWT = []byte(strings.Join(params, "."))
	} else {

		JWT = []byte(params[0])
	}

	if len(JWT) == 0 {
		return "" , errors.New("lista de parametros vacia")
	}

	cipherBytes , err := spaceCipher.Encripter(JWT)
	if err != nil{
		return
	}

	return base64.StdEncoding.EncodeToString(cipherBytes) , nil
}


func (spaceCipher *SpaceCipher)DecodeCipherTokenBase64(JWT string) (params []string , err error) {

	if len(JWT) == 0 {
		return
	}

	cipherJWT, err := base64.StdEncoding.DecodeString(JWT)
	if err != nil {
		return
	}

	cipherJWT , err = spaceCipher.Dencripter(cipherJWT)
	if err != nil{
		return
	}

	params = strings.Split(string(cipherJWT), ".")

	for ind, param := range params {

		if param == "" {
			continue
		}
		decode, err := base64.StdEncoding.DecodeString(param)
		if err != nil {
			continue
		}

		params[ind] = string(decode)

	}

	if(len(params) == 0 ){
		return nil , errors.New("no se decodificaron parametros")
	}

	return params , nil
}

func (spaceCipher *SpaceCipher)NewToken(line int64, typeToken string, timeNow time.Time) (JWT string , err error) {

	if line < 0 {
		return
	}

	return spaceCipher.EncodeCipherTokenBase64(strconv.FormatInt(int64(line), 10),
		typeToken,
		strconv.FormatInt(timeNow.UnixMilli(), 10))

}

func (spaceCipher *SpaceCipher)DecodeToken(JWT string) (Line int64, typeToken string, timeToken64 int64, err error) {

	params , err := spaceCipher.DecodeCipherTokenBase64(JWT)
	if err != nil{
		return
	}

	if len(params) == 0 {
		err = errors.New("tokenInvalido")
		return
	}

	Line, err = strconv.ParseInt(params[0], 10, 64)
	if err != nil {
		err = errors.New("tokenInvalido")
		return
	}

	typeToken = params[1]

	timeToken64, err = strconv.ParseInt(params[2], 10, 64)
	if err != nil {
		err = errors.New("tokenInvalido")
		return
	}

	return
}


func (spaceCipher *SpaceCipher) EncodeCipherToken(data interface{}) (cipherBytes []byte , err error) {

	// Convertir la estructura a JSON
	JWT, err := json.Marshal(data)
	if err != nil {
		return
	}

	// Verificar si el JSON está vacío
	if len(JWT) == 0 {
		return nil , errors.New("se envio un json vacio")
	}

	// Cifrar los bytes JSON
	cipherBytes, err = spaceCipher.Encripter(JWT)
	if err != nil {
		return
	}

	return 
}

func (spaceCipher *SpaceCipher) DecodeCipherToken(cipherJWT []byte , out interface{}) (err error) {
	
	// Verificar si el input es válido
	if len(cipherJWT) == 0 {
		return errors.New("el token cifrado está vacío")
	}

	// Desencriptar los bytes
	decryptedBytes, err := spaceCipher.Dencripter(cipherJWT)
	if err != nil {
		return
	}

	// Convertir JSON a la estructura de salida
	err = json.Unmarshal(decryptedBytes, out)
	if err != nil {
		return
	}

	return nil
}
