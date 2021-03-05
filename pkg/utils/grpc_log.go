/*
Copyright 2021 The Kubernetes Authors.

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

package utils

import (
	"google.golang.org/grpc/grpclog"
	"k8s.io/klog/v2"
)

// ConfigureGrpcKlog wraps grpc logging to use klog
func ConfigureGrpcKlog() {
	grpclog.SetLoggerV2(new(grpcLogger))
}

// grpcLogger implements the LoggerV2 interface from grpclog
type grpcLogger struct{}

func (g grpcLogger) Error(args ...interface{}) {
	klog.Error(args...)
}

func (g grpcLogger) Errorf(format string, args ...interface{}) {
	klog.Errorf(format, args...)
}

func (g grpcLogger) Errorln(args ...interface{}) {
	klog.Errorln(args...)
}

func (g grpcLogger) Fatal(args ...interface{}) {
	klog.Fatal(args...)
}

func (g grpcLogger) Fatalf(format string, args ...interface{}) {
	klog.Fatalf(format, args...)
}

func (g grpcLogger) Fatalln(args ...interface{}) {
	klog.Fatalln(args...)
}

func (g grpcLogger) Info(args ...interface{}) {
	klog.Info(args...)
}

func (g grpcLogger) Infof(format string, args ...interface{}) {
	klog.Infof(format, args...)
}

func (g grpcLogger) Infoln(args ...interface{}) {
	klog.Infoln(args...)
}

func (g grpcLogger) Warning(args ...interface{}) {
	klog.Warning(args...)
}

func (g grpcLogger) Warningf(format string, args ...interface{}) {
	klog.Warningf(format, args...)
}

func (g grpcLogger) Warningln(args ...interface{}) {
	klog.Warningln(args...)
}

func (g grpcLogger) V(l int) bool {
	return klog.V(klog.Level(l)).Enabled()
}
