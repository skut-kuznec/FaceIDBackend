package faceapp

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"

	face "github.com/Kagami/go-face"
	"github.com/smart48ru/FaceIDApp/internal/domain"
)

type Recognizer struct {
	modelDir string
}

func New(modelDir string) *Recognizer {
	return &Recognizer{
		modelDir: modelDir,
	}
}

func (r *Recognizer) GetFaceDescriptor(img []byte) (domain.FaceDescriptor, error) {
	img, err := toJPEG(img)
	if err != nil {
		return domain.FaceDescriptor{}, err
	}

	faceRecognizer, err := initRecognizer(r.modelDir)
	if err != nil {
		return domain.FaceDescriptor{}, fmt.Errorf("error to init recognizer: %w", err)
	}

	faces, err := faceRecognizer.Recognize(img)
	if err != nil {
		return domain.FaceDescriptor{}, fmt.Errorf("error while decoding/processing image: %w", err)
	}
	if len(faces) > 1 {
		return domain.FaceDescriptor{}, fmt.Errorf("found multiple faces on image")
	}
	if len(faces) == 0 {
		return domain.FaceDescriptor{}, fmt.Errorf("face not found")
	}

	ret := faces[0]

	return domain.FaceDescriptor(ret.Descriptor), nil
}

func (r *Recognizer) SearchImageID(descr domain.FaceDescriptor, images []domain.Image) (uint64, error) {
	faceRecognizer, err := initRecognizer(r.modelDir)
	if err != nil {
		return 0, fmt.Errorf("error to init recognizer: %w", err)
	}

	setSamples(faceRecognizer, images)

	imgID := faceRecognizer.Classify(face.Descriptor(descr))
	if imgID < 0 {
		return 0, errors.New("face not found")
	}

	return uint64(imgID), nil
}

func initRecognizer(modelDir string) (*face.Recognizer, error) {
	faceRecognizer, err := face.NewRecognizer(modelDir)
	if err != nil {
		return nil, err
	}

	return faceRecognizer, nil
}

func setSamples(faceRecognizer *face.Recognizer, images []domain.Image) {
	samples := make([]face.Descriptor, 0, len(images))
	ids := make([]int32, 0, len(images))
	for _, img := range images {
		samples = append(samples, face.Descriptor(img.FaceDescriptor))
		ids = append(ids, int32(img.ID))
	}

	faceRecognizer.SetSamples(samples, ids)
}

func toJPEG(img []byte) ([]byte, error) {
	contentType := http.DetectContentType(img)

	switch contentType {
	case "image/jpeg":
		return img, nil
	case "image/png":
	default:
		return nil, errors.New("only jpeg and png types supported")
	}

	image, err := png.Decode(bytes.NewReader(img))
	if err != nil {
		return nil, fmt.Errorf("error to decode png image: %w", err)
	}

	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, image, nil); err != nil {
		return nil, fmt.Errorf("unable to encode to jpeg")
	}

	return buf.Bytes(), nil
}
