/*
Copyright 2017 The Kubernetes Authors.

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

package cpu

import (
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/exp/slog"

	"github.com/converged-computing/nfd-source/pkg/utils/hostpath"
)

// Discover p-state related features such as turbo boost.
func detectPstate() (map[string]string, error) {
	// Check that sysfs is available
	sysfsBase := hostpath.SysfsDir.Path("devices/system/cpu")
	if _, err := os.Stat(sysfsBase); err != nil {
		return nil, err
	}
	pstateDir := filepath.Join(sysfsBase, "intel_pstate")
	if _, err := os.Stat(pstateDir); os.IsNotExist(err) {
		slog.Info("intel pstate driver not enabled")
		return nil, nil
	}

	// Get global pstate status
	data, err := os.ReadFile(filepath.Join(pstateDir, "status"))
	if err != nil {
		return nil, err
	}
	status := strings.TrimSpace(string(data))
	if status == "off" {
		// No need to check other pstate features
		slog.Info("intel_pstate driver is not in use")
		return nil, nil
	}
	features := map[string]string{"status": status}

	// Check turbo boost
	bytes, err := os.ReadFile(filepath.Join(pstateDir, "no_turbo"))
	if err != nil {
		slog.Any("can't detect whether turbo boost is enabled", err)
	} else {
		features["turbo"] = "false"
		if bytes[0] == byte('0') {
			features["turbo"] = "true"
		}
	}

	if status != "active" {
		// Don't check other features which depend on active state
		return features, nil
	}

	// Determine scaling governor that is being used
	cpufreqDir := filepath.Join(sysfsBase, "cpufreq")
	policies, err := os.ReadDir(cpufreqDir)
	if err != nil {
		slog.Any("failed to read cpufreq directory", err)
		return features, nil
	}

	scaling := ""
	for _, policy := range policies {
		// Ensure at least one cpu is using this policy
		cpus, err := os.ReadFile(filepath.Join(cpufreqDir, policy.Name(), "affected_cpus"))
		if err != nil {
			slog.Info("could not read cpufreq affected_cpus", "cpufreqPolicyName", policy.Name())
			continue
		}
		if strings.TrimSpace(string(cpus)) == "" {
			slog.Info("cpufreq policy has no associated cpus", "cpufreqPolicyName", policy.Name())
			continue
		}

		data, err := os.ReadFile(filepath.Join(cpufreqDir, policy.Name(), "scaling_governor"))
		if err != nil {
			slog.Info("could not read cpufreq scaling_governor", "cpufreqPolicyName", policy.Name())
			continue
		}
		policy_scaling := strings.TrimSpace(string(data))
		// Check that all of the policies have the same scaling governor, if not don't set feature
		if scaling != "" && scaling != policy_scaling {
			slog.Info("scaling_governor for cpufreq policy doesn't match prior policy", "cpufreqPolicyName", policy.Name())
			scaling = ""
			break
		}
		scaling = policy_scaling
	}

	if scaling != "" {
		features["scaling_governor"] = scaling
	}

	return features, nil
}
