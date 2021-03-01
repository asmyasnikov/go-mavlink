//go:generate templify register.template
//go:generate templify helpers.template
//go:generate templify errors.template
//go:generate templify encoder.template
//go:generate templify decoder.template
//go:generate templify message.template
//go:generate templify packet.template
//go:generate templify packetV.template
//go:generate templify parser.template
//go:generate templify parserV.template
//go:generate templify version.template
//go:generate templify x25.template
//go:generate templify signature.template

package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

const (
	generatedHeader = `/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/asmyasnikov/go-mavlink/mavgen
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

`
)

var (
	independentTemplates = map[string](func() string){
		"encoder":             encoderTemplate,
		"decoder":             decoderTemplate,
		"errors/errors":       errorsTemplate,
		"helpers/helpers":     helpersTemplate,
		"register/register":   registerTemplate,
		"message/message":     messageTemplate,
		"version/version":     versionTemplate,
		"packet/packet":       packetTemplate,
		"parser/parser":       parserTemplate,
		"signature/signature": signatureTemplate,
		"crc/x25":             x25Template,
	}
	dependentTemplates = map[string](func() string){
		"parser/packetV": packetVTemplate,
		"parser/parserV": parserVTemplate,
	}
)

// helper to remove the extension from the base name
func baseName(s string) string {
	return strings.TrimSuffix(filepath.Base(s), filepath.Ext(s))
}

func generateDialect(dialectPath *string, commonPackage string, schemeFile string, json bool) error {
	d, err := ParseDialect(schemeFile)
	if err != nil {
		return err
	}

	d.FilePath = schemeFile
	d.JSON = json

	baseName := baseName(schemeFile)

	if dialectPath == nil {
		path, err := os.Getwd()
		if err != nil {
			log.Fatal("Getwd(): ", err)
		}
		path = filepath.Join(path, "dialects", baseName)
		dialectPath = &path
	}

	if err := d.generateGo(*dialectPath, baseName, commonPackage); err != nil {
		return err
	}

	return nil
}

func generateCode(dialectDir string, data templateData, templateName string, tmpl string) error {
	t, err := template.New(templateName).Parse(tmpl)

	if err != nil {
		return err
	}

	file, err := os.Create(dialectDir + templateName + ".go")
	if err != nil {
		return err
	}
	defer file.Close()

	n, err := file.Write([]byte(generatedHeader))
	if err != nil {
		return err
	} else if n < len(generatedHeader) {
		return errors.New("couldn't write NO-EDIT header")
	}

	var buffer bytes.Buffer
	if err := t.Execute(&buffer, data); err != nil {
		return err
	}
	formatted, err := format.Source(buffer.Bytes())
	if err != nil {
		for i, s := range strings.Split(buffer.String(), "\n") {
			fmt.Println(i, ": ", s)
		}
		log.Fatal("couldn't format generated "+templateName+".go: ", err)
	}
	n, err = file.Write(formatted)
	if err != nil {
		return err
	} else if n < len(formatted) {
		return errors.New("couldn't write body of " + templateName + ".go")
	}
	return nil
}

func generateCommonPackage(commonPackageURL string) error {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Getwd(): ", err)
	}

	// generate mavlink independent code
	for k, v := range independentTemplates {
		paths := strings.Split(k, string(filepath.Separator))
		path := strings.Join(append([]string{cwd}, paths[:len(paths)-1]...), string(filepath.Separator)) + string(filepath.Separator)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
		if err := generateCode(path, templateData{
			CommonPackageURL: commonPackageURL,
		}, paths[len(paths)-1], v()); err != nil {
			return err
		}
	}

	// generate mavlink dependent code
	for _, m := range []int{1, 2} {
		for k, v := range dependentTemplates {
			paths := strings.Split(k, string(filepath.Separator))
			path := strings.Join(append([]string{cwd}, paths[:len(paths)-1]...), string(filepath.Separator)) + string(filepath.Separator)
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
			if err := generateCode(path, templateData{
				MavlinkVersion:   m,
				CommonPackageURL: commonPackageURL,
			}, paths[len(paths)-1]+strconv.Itoa(m), v()); err != nil {
				return err
			}
		}
	}
	return nil
}
