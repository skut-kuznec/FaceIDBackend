package facerecognizer

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
	faceRecognizer *face.Recognizer
}

func New(modelDir string, images []domain.Image) (*Recognizer, error) {
	faceRecognizer, err := face.NewRecognizer(modelDir)
	if err != nil {
		return nil, err
	}

	samples := make([]face.Descriptor, 0, len(images))
	ids := make([]int32, 0, len(images))
	for _, img := range images {
		samples = append(samples, face.Descriptor(img.FaceDescriptor))
		ids = append(ids, int32(img.ID))
	}

	faceRecognizer.SetSamples(samples, ids)
	return &Recognizer{
		faceRecognizer: faceRecognizer,
	}, nil
}

func (r *Recognizer) GetFaceDescriptor(img []byte) (domain.FaceDescriptor, error) {
	descr, err := r.getFaceDescriptor(img)
	if err != nil {
		return domain.FaceDescriptor{}, err
	}

	return domain.FaceDescriptor(descr), nil
}

func (r *Recognizer) SearchImageID(img []byte) (uint64, error) {
	descr, err := r.getFaceDescriptor(img)
	if err != nil {
		return 0, err
	}

	imgID := r.faceRecognizer.Classify(descr)
	if imgID < 0 {
		return 0, errors.New("face not found")
	}

	return uint64(imgID), nil
}

func (r *Recognizer) getFaceDescriptor(img []byte) (face.Descriptor, error) {
	img, err := toJPEG(img)
	if err != nil {
		return face.Descriptor{}, err
	}

	faces, err := r.faceRecognizer.Recognize(img)
	if err != nil {
		return face.Descriptor{}, fmt.Errorf("error while decoding/processing image: %w", err)
	}
	if len(faces) > 1 {
		return face.Descriptor{}, fmt.Errorf("found multiple faces on image")
	}
	if len(faces) == 0 {
		return face.Descriptor{}, fmt.Errorf("face not found")
	}

	ret := faces[0]

	return ret.Descriptor, nil
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
