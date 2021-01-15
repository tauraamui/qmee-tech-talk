rtspAddr := "rtsp://{username}:{password}@192.168.1.164//"
goroutine 1 => {
	videoCapture, err := gocv.OpenVideoCapture(rtspAddr)
	imageMatrix := gocv.NewMat()
	videoCapture.Read(&imageMatrix)
	captureChannel <- imageMatrix
}