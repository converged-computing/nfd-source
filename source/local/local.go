/*
Copyright 2018-2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package local

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/exp/slog"

	nfdv1alpha1 "github.com/converged-computing/nfd-source/pkg/apis/nfd/v1alpha1"
	"github.com/converged-computing/nfd-source/pkg/utils"
	"github.com/converged-computing/nfd-source/source"
)

// Name of this feature source
const Name = "local"

// LabelFeature of this feature source
const LabelFeature = "label"

// RawFeature of this feature source
const RawFeature = "feature"

const (
	// ExpiryTimeKey is the key of this feature source indicating
	// when features should be removed.
	DirectiveExpiryTime = "expiry-time"

	// NoLabel indicates whether the feature should be included
	// in exposed labels or not.
	DirectiveNoLabel = "no-label"

	// NoFeature indicates whether the feature should be included
	// in exposed raw features or not.
	DirectiveNoFeature = "no-feature"
)

// DirectivePrefix defines the prefix of directives that should be parsed
const DirectivePrefix = "# +"

// MaxFeatureFileSize defines the maximum size of a feature file size
const MaxFeatureFileSize = 65536

// Config
var (
	featureFilesDir = "/etc/kubernetes/node-feature-discovery/features.d/"
	hookDir         = "/etc/kubernetes/node-feature-discovery/source.d/"
)

// localSource implements the FeatureSource and LabelSource interfaces.
type localSource struct {
	features *nfdv1alpha1.Features
	config   *Config
}

type Config struct {
	HooksEnabled bool `json:"hooksEnabled,omitempty"`
}

// parsingOpts contains options used for directives parsing
type parsingOpts struct {
	ExpiryTime  time.Time
	SkipLabel   bool
	SkipFeature bool
}

// Singleton source instance
var (
	src                           = localSource{config: newDefaultConfig()}
	_   source.FeatureSource      = &src
	_   source.LabelSource        = &src
	_   source.ConfigurableSource = &src
)

// Name method of the LabelSource interface
func (s *localSource) Name() string { return Name }

// NewConfig method of the LabelSource interface
func (s *localSource) NewConfig() source.Config { return newDefaultConfig() }

// GetConfig method of the LabelSource interface
func (s *localSource) GetConfig() source.Config { return s.config }

// SetConfig method of the LabelSource interface
func (s *localSource) SetConfig(conf source.Config) {
	switch v := conf.(type) {
	case *Config:
		s.config = v
	default:
		panic(fmt.Sprintf("invalid config type: %T", conf))
	}
}

// Priority method of the LabelSource interface
func (s *localSource) Priority() int { return 20 }

// GetLabels method of the LabelSource interface
func (s *localSource) GetLabels() (source.FeatureLabels, error) {
	labels := make(source.FeatureLabels)
	features := s.GetFeatures()

	for k, v := range features.Attributes[LabelFeature].Elements {
		labels[k] = v
	}
	return labels, nil
}

// newDefaultConfig returns a new config with pre-populated defaults
func newDefaultConfig() *Config {
	return &Config{
		HooksEnabled: false,
	}
}

// Discover method of the FeatureSource interface
func (s *localSource) Discover() error {
	s.features = nfdv1alpha1.NewFeatures()

	featuresFromFiles, labelsFromFiles, err := getFeaturesFromFiles()
	if err != nil {
		slog.Any("failed to read feature files", err)
	}

	if s.config.HooksEnabled {

		slog.Info("starting hooks...")
		slog.Info("NOTE: hooks are deprecated and will be completely removed in a future release.")

		featuresFromHooks, labelsFromHooks, err := getFeaturesFromHooks()
		if err != nil {
			slog.Any("failed to run hooks", err)
		}

		// Merge features from hooks and files
		for k, v := range featuresFromHooks {
			if old, ok := featuresFromFiles[k]; ok {
				slog.Info("overriding feature value", "featureKey", k, "oldValue", old, "newValue", v)
			}
			featuresFromFiles[k] = v
		}

		// Merge labels from hooks and files
		for k, v := range labelsFromHooks {
			if old, ok := labelsFromFiles[k]; ok {
				slog.Info("overriding label value", "labelKey", k, "oldValue", old, "newValue", v)
			}
			labelsFromHooks[k] = v
		}
	}

	s.features.Attributes[LabelFeature] = nfdv1alpha1.NewAttributeFeatures(labelsFromFiles)
	s.features.Attributes[RawFeature] = nfdv1alpha1.NewAttributeFeatures(featuresFromFiles)

	slog.Debug("discovered features", "featureSource", s.Name(), "features", utils.DelayedDumper(s.features))

	return nil
}

// GetFeatures method of the FeatureSource Interface
func (s *localSource) GetFeatures() *nfdv1alpha1.Features {
	if s.features == nil {
		s.features = nfdv1alpha1.NewFeatures()
	}
	return s.features
}

func parseDirectives(line string, opts *parsingOpts) error {
	if !strings.HasPrefix(line, DirectivePrefix) {
		return nil
	}

	directive := line[len(DirectivePrefix):]
	split := strings.SplitN(directive, "=", 2)
	key := split[0]

	switch key {
	case DirectiveExpiryTime:
		if len(split) == 1 {
			return fmt.Errorf("invalid directive format in %q, should be '# +expiry-time=value'", line)
		}
		value := split[1]
		expiryDate, err := time.Parse(time.RFC3339, strings.TrimSpace(value))
		if err != nil {
			return fmt.Errorf("failed to parse expiry-date directive: %w", err)
		}
		opts.ExpiryTime = expiryDate
	case DirectiveNoFeature:
		opts.SkipFeature = true
	case DirectiveNoLabel:
		opts.SkipLabel = true
	default:
		return fmt.Errorf("unknown feature file directive %q", key)
	}

	return nil
}

func parseFeatureFile(lines [][]byte, fileName string) (map[string]string, map[string]string) {
	features := make(map[string]string)
	labels := make(map[string]string)

	now := time.Now()
	parsingOpts := &parsingOpts{
		ExpiryTime:  now,
		SkipLabel:   false,
		SkipFeature: false,
	}

	for _, l := range lines {
		line := strings.TrimSpace(string(l))
		if len(line) > 0 {
			if strings.HasPrefix(line, "#") {
				// Parse directives
				err := parseDirectives(line, parsingOpts)
				if err != nil {
					slog.Any(fmt.Sprintf("error while parsing directives fileName", fileName), err)
				}

				continue
			}

			// handle expiration
			if parsingOpts.ExpiryTime.Before(now) {
				continue
			}

			lineSplit := strings.SplitN(line, "=", 2)

			key := lineSplit[0]

			if !parsingOpts.SkipFeature {
				updateFeatures(features, lineSplit)
			} else {
				delete(features, key)
			}

			if !parsingOpts.SkipLabel {
				updateFeatures(labels, lineSplit)
			} else {
				delete(labels, key)
			}
			// SkipFeature and SkipLabel only take effect for one feature
			parsingOpts.SkipFeature = false
			parsingOpts.SkipLabel = false
		}
	}

	return features, labels
}

func updateFeatures(m map[string]string, lineSplit []string) {
	key := lineSplit[0]
	// Check if it's a boolean value
	if len(lineSplit) == 1 {
		m[key] = "true"

	} else {
		m[key] = lineSplit[1]
	}
}

// Run all hooks and get features
func getFeaturesFromHooks() (map[string]string, map[string]string, error) {

	features := make(map[string]string)
	labels := make(map[string]string)

	files, err := os.ReadDir(hookDir)
	if err != nil {
		if os.IsNotExist(err) {
			slog.Info("hook directory does not exist", "path", hookDir)
			return features, labels, nil
		}
		return features, labels, fmt.Errorf("unable to access %v: %w", hookDir, err)
	}
	if len(files) > 0 {
		slog.Info("hooks are DEPRECATED since v0.12.0 and support will be removed in a future release; use feature files instead")
	}

	for _, file := range files {
		fileName := file.Name()
		// ignore hidden feature file
		if strings.HasPrefix(fileName, ".") {
			continue
		}
		lines, err := runHook(fileName)
		if err != nil {
			slog.Any(fmt.Sprintf("failed to run hook fileName %s", fileName), err)
			continue
		}

		// Append features
		fileFeatures, fileLabels := parseFeatureFile(lines, fileName)
		slog.Info("hook executed", "fileName", fileName, "features", utils.DelayedDumper(fileFeatures), "labels", utils.DelayedDumper(fileLabels))
		for k, v := range fileFeatures {
			if old, ok := features[k]; ok {
				slog.Info("overriding feature value from another hook", "featureKey", k, "oldValue", old, "newValue", v, "fileName", fileName)
			}
			features[k] = v
		}

		for k, v := range fileLabels {
			if old, ok := labels[k]; ok {
				slog.Info("overriding label value from another hook", "labelKey", k, "oldValue", old, "newValue", v, "fileName", fileName)
			}
			labels[k] = v
		}
	}

	return features, labels, nil
}

// Run one hook
func runHook(file string) ([][]byte, error) {
	var lines [][]byte

	path := filepath.Join(hookDir, file)
	filestat, err := os.Stat(path)
	if err != nil {
		slog.Any(fmt.Sprintf("failed to get filestat, skipping hook path %s", path), err)
		return lines, err
	}

	if filestat.Mode().IsRegular() {
		cmd := exec.Command(path)
		var stdout bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		// Run hook
		err = cmd.Run()

		// Forward stderr to our logger
		errLines := bytes.Split(stderr.Bytes(), []byte("\n"))
		for i, line := range errLines {
			if i == len(errLines)-1 && len(line) == 0 {
				// Don't print the last empty string
				break
			}
			slog.Info(fmt.Sprintf("%s: %s", file, line))
		}

		// Do not return any lines if an error occurred
		if err != nil {
			return lines, err
		}
		lines = bytes.Split(stdout.Bytes(), []byte("\n"))
	}

	return lines, nil
}

// Read all files to get features
func getFeaturesFromFiles() (map[string]string, map[string]string, error) {
	features := make(map[string]string)
	labels := make(map[string]string)

	files, err := os.ReadDir(featureFilesDir)
	if err != nil {
		if os.IsNotExist(err) {
			slog.Info("features directory does not exist", "path", featureFilesDir)
			return features, labels, nil
		}
		return features, labels, fmt.Errorf("unable to access %v: %w", featureFilesDir, err)
	}

	for _, file := range files {
		fileName := file.Name()
		// ignore hidden feature file
		if strings.HasPrefix(fileName, ".") {
			continue
		}
		lines, err := getFileContent(fileName)
		if err != nil {
			slog.Any(fmt.Sprintf("failed to read file fileName %s", fileName), err)
			continue
		}

		// Append features
		fileFeatures, fileLabels := parseFeatureFile(lines, fileName)

		slog.Info("feature file read", "fileName", fileName, "features", utils.DelayedDumper(fileFeatures))
		for k, v := range fileFeatures {
			if old, ok := features[k]; ok {
				slog.Info("overriding label value from another feature file", "featureKey", k, "oldValue", old, "newValue", v, "fileName", fileName)
			}
			features[k] = v
		}

		for k, v := range fileLabels {
			if old, ok := labels[k]; ok {
				slog.Info("overriding label value from another feature file", "labelKey", k, "oldValue", old, "newValue", v, "fileName", fileName)
			}
			labels[k] = v
		}
	}

	return features, labels, nil
}

// Read one file
func getFileContent(fileName string) ([][]byte, error) {
	var lines [][]byte

	path := filepath.Join(featureFilesDir, fileName)
	filestat, err := os.Stat(path)
	if err != nil {
		slog.Any(fmt.Sprintf("failed to get filestat, skipping features file path %s", path), err)
		return lines, err
	}

	if filestat.Mode().IsRegular() {
		if filestat.Size() > MaxFeatureFileSize {
			return lines, fmt.Errorf("file size limit exceeded: %d bytes > %d bytes", filestat.Size(), MaxFeatureFileSize)
		}

		fileContent, err := os.ReadFile(path)

		// Do not return any lines if an error occurred
		if err != nil {
			return lines, err
		}
		lines = bytes.Split(fileContent, []byte("\n"))
	}

	return lines, nil
}

func init() {
	source.Register(&src)
}
