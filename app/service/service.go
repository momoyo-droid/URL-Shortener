package service

import (
	"fmt"
	"net/url"
	"projects/url-shortener/app/handler/model"
	"projects/url-shortener/app/repository"
	"projects/url-shortener/settings"
)

func ShortenURL(request *model.Request) (*model.Request, error) {
	if request.URL == "" {
		return nil, fmt.Errorf("URL is empty. Please provide a valid URL")
	}

	if err := validateURL(request.URL); err != nil {
		return nil, fmt.Errorf("invalid URL. Please provide a valid URL")
	}

	if request.CustomAlias == "" {
		request.CustomAlias = generateCustomAlias()
	}

	shortenedURL := fmt.Sprintf(settings.LocalHost, request.CustomAlias)

	if err := repository.CreateShortenedURL(request.CustomAlias, shortenedURL); err != nil {
		return nil, fmt.Errorf("failed to create shortened URL. Original URL: %s, Custom Alias: %s, Shortened URL: %s",
			request.URL, request.CustomAlias, shortenedURL)
	}

	return request, nil
}

func validateURL(URL string) error {
	if _, err := url.ParseRequestURI(URL); err != nil {
		return err
	}
	return nil
}

func generateCustomAlias() string {
	customAlias := "alias"
	return customAlias
}
