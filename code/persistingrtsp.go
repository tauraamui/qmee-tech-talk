img := <-camera.streamBuffer
writer, _ := gocv.VideoWriterFile("output.mp4", "avc1.4d001e", 30, img.Cols(), img.Rows(), true)
defer writer.Close()

var framesWritten uint
for framesWritten = 0; framesWritten < 30*uint(camera.secondsPerClip); framesWritten++ {
	img = <-camera.streamBuffer
	defer img.Close()

	if writer.IsOpened() {
		writer.Write(*img)
	}
}
