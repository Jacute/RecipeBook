package utils

var validImageMimeTypes = []string{
	"image/jpeg",
	"image/png",
	"image/webp",
}

var validImageExtensions = []string{
	".jpg",
	".jpeg",
	".png",
	".webp",
}

func IsValidImageMimeType(mimeType string) bool {
	for _, imageMimeType := range validImageMimeTypes {
		if mimeType == imageMimeType {
			return true
		}
	}
	return false
}

func IsValidImageExtension(extension string) bool {
	for _, imageExtension := range validImageExtensions {
		if extension == imageExtension {
			return true
		}
	}
	return false
}

func IsValidImage(mimeType, extension string) bool {
	return IsValidImageMimeType(mimeType) && IsValidImageExtension(extension)
}
