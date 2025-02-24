package httpReceiver

import (
	. "dacV2"
	"errors"
	"strconv"
	"strings"
)

func (SK *HttpReceiver) ReadHeaderRangesRaw() (start int64, end int64, err error) {

	rangeHeader := SK.ReadHeaderRaw("Range")
	if rangeHeader == "" {
		rangeHeader = SK.ReadUrlRaw("Range")
	}

	rangeHeader = strings.Replace(rangeHeader, "bytes=", "", -1)

	ranges := strings.SplitN(rangeHeader, "-", 2)

	if len(ranges) == 2 {

		start, err = strconv.ParseInt(ranges[0], 10, 64)
		if err != nil {
			return
		}

		end, _ = strconv.ParseInt(ranges[1], 10, 64)
		return
	}

	if len(ranges) == 1 {

		start, err = strconv.ParseInt(ranges[0], 10, 64)
		return
	}

	if len(ranges) <= 0 {
		err = errors.New("no se han enviado header de rangos")
		return
	}

	return
}

func (SK *HttpReceiver) WriteHeaderContentRanges(startRange int64, endRange int64, sizeTotal int64) {

	SK.WriteHeaderRaw("Content-Range",
		strings.Join([]string{"bytes ",
			strconv.Itoa(int(startRange)),
			"-",
			strconv.Itoa(int(endRange - 1)),
			"/",
			strconv.Itoa(int(sizeTotal)),
		}, ""))
}

func CalcRange(startRangeHeader int64, fileLen int64, bandwidth int64) (nRange int64, startRange int64, endRange int64) {

	nRange = 1

	if startRangeHeader != 0 {

		nRange = (startRangeHeader / bandwidth) + 1
	}

	startRange = (nRange - 1) * bandwidth

	endRange = nRange * bandwidth
	if endRange > fileLen {
		endRange = fileLen
	}

	return
}

// Esta funcion es un punto final y maneja los errores internamente.
func (SK *HttpReceiver) ServerFileToClientRanges(fileCache *SpaceCacheExpiration, typeContent string, bandwidth int64, dirName ...string) {

	space, err := fileCache.OpenSpaceRange( dirName...)
	if err != nil {
		SK.ErrorStatusInternalServerError(err.Error())
		return
	}

	if space.Size == 0 {

		fileSize , err := space.FileSize()
		if err != nil {
			SK.ErrorStatusInternalServerError(err.Error())
			return
		}
		//Modificacion del space en la cache.
		space.Size = fileSize
	}


	//La primera estapa es solicitar los rangos
	startRangeHeader, _, err := SK.ReadHeaderRangesRaw()
	if err != nil {
		SK.WriteContentLength(space.Size)
		SK.WriteBandwidth(bandwidth)
		SK.WriteContentType(typeContent)
		SK.WriteHeaderCode(206)
		return
	}

	nRange, startRange, endRange := CalcRange(startRangeHeader, space.Size, bandwidth)

	buffer, err := space.GetAtRange( nRange-1 , bandwidth)
	if err != nil {
		SK.ErrorStatusInternalServerError(err.Error())
		return
	}

	if startRangeHeader > startRange {

		buffer = buffer[startRangeHeader-startRange:]
	}

	SK.WriteHeaderContentRanges(startRangeHeader, endRange, space.Size)

	SK.WriteHeaderCode(206)

	err = SK.Write(typeContent, buffer)
	if err != nil {
		SK.ErrorStatusInternalServerError(err.Error())
		return
	}
}

func (SK *HttpReceiver) ClientFileToServerRanges(fileCache *SpaceCacheExpiration, fileSize int64, bandwidth int64, dirName ...string) {

	if SK.ExistUrl("bandwidth") {

		err := SK.WriteInt64(bandwidth)
		if err != nil {
			SK.ErrorStatusBadRequest(err.Error())
			return
		}

		return
	}

	nRange, err := SK.ReadUrlInt64("nRange")
	if err != nil {
		SK.ErrorStatusBadRequest(err.Error())
		return
	}

	body, err := SK.ReadBodyMaxBytes(bandwidth)
	if err != nil {
		SK.ErrorStatusBadRequest(err.Error())
		return
	}

	space, err := fileCache.OpenSpaceRange(dirName...)
	if err != nil {
		SK.ErrorStatusInternalServerError(err.Error())
		return
	}

	err = space.SetAtRange( body,  nRange-1 , bandwidth)
	if err != nil {
		SK.ErrorStatusBadRequest(err.Error())
		return
	}

	err = SK.WriteOk()
	if err != nil {
		SK.ErrorStatusBadRequest(err.Error())
		return
	}
}
