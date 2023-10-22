package main

import (
	"image/jpeg"

	"github.com/fogleman/gg"

	"github.com/abdivasiyev/telegram-bot/pkg/certificator"
)

func main() {
	img, err := certificator.GetCertificate(certificator.Request{
		BgImgPath: "./static/images/certificate.jpeg",
		FontPath:  "./static/fonts/SometypeMono-Bold.ttf",
		FontSize:  36,
		Text:      "Asliddin Abdivasiyev",
	})

	if err != nil {
		panic(err)
	}

	if err = gg.SaveJPG("./static/gen/certificate.jpeg", img, jpeg.DefaultQuality); err != nil {
		panic(err)
	}
}
