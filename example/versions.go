package main

import (
	"log"

	"github.com/sailorvii/goav/avcodec"
	"github.com/sailorvii/goav/avdevice"
	"github.com/sailorvii/goav/avfilter"
	"github.com/sailorvii/goav/avformat"
	"github.com/sailorvii/goav/avutil"
	"github.com/sailorvii/goav/swresample"
	"github.com/sailorvii/goav/swscale"
)

func main() {

	// Register all formats and codecs
	avformat.AvRegisterAll()
	avcodec.AvcodecRegisterAll()

	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())

}
