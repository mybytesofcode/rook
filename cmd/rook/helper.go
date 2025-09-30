package main

import (
	"path/filepath"

	bashpkg "github.com/mybytesofcode/rook/internal/bash"
	mergepkg "github.com/mybytesofcode/rook/internal/merge"
	templatepkg "github.com/mybytesofcode/rook/internal/template"
	valuespkg "github.com/mybytesofcode/rook/internal/values"
)

func Render(valuesPaths []string, scriptsPaths []string) ([]string, error) {
	values := make(valuespkg.Values)
	for _, path := range valuesPaths {
		valuesDst, err := valuespkg.ValuesRead(path)
		if err != nil {
			return nil, err
		}

		values = mergepkg.Merge(values, valuesDst)
	}

	result := []string{}
	for _, path := range scriptsPaths {
		lines, err := templatepkg.TemplateRender(path, values)
		if err != nil {
			return nil, err
		}

		tmpPath, err := bashpkg.BashPreprocess(filepath.Dir(path), lines, values)
		if err != nil {
			return nil, err
		}

		result = append(result, tmpPath)
	}

	return result, nil
}
