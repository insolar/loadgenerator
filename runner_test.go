/*
 *    Copyright [2020] Sergey Kudasov
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package loadgen

import (
	"context"
	"net/http"
)

type TestAttack struct {
	WithRunner
	client *http.Client
}

func (a *TestAttack) Setup(hc RunnerConfig) error {
	a.client = NewLoggingHTTPClient(a.GetManager().SuiteConfig.DumpTransport, 60)
	return nil
}
func (a *TestAttack) Do(ctx context.Context) DoResult {
	if _, err := a.client.Get(a.GetManager().GeneratorConfig.Generator.Target); err != nil {
		return DoResult{
			Error:        err,
			RequestLabel: "test1",
		}
	}
	return DoResult{
		RequestLabel: "test1",
	}
}
func (a *TestAttack) Clone(r *Runner) Attack {
	return &TestAttack{WithRunner: WithRunner{R: r}}
}

func AttackerFromName(name string) Attack {
	switch name {
	case "test1":
		return WithMonitor(new(TestAttack))
	default:
		log.Fatalf("unknown attacker type: %s", name)
		return nil
	}
}

//
// func TestRunnerShutdown(t *testing.T) {
// 	ta := new(TestAttack)
// 	suiteCfg := SuiteConfig{
// 		DumpTransport:  false,
// 		GoroutinesDump: false,
// 		HttpTimeout:    20,
// 		Steps:          []Step{
// 			{
// 				Name: "step1",
// 				Handles: []RunnerConfig{
//
// 				},
// 			},
// 		},
// 	}
// 	runnerConfig := RunnerConfig{
// 		HandleName:      "test",
// 		RPS:             10,
// 		AttackTimeSec:   10,
// 		RampUpTimeSec:   10,
// 		RampUpStrategy:  "linear",
// 		MaxAttackers:    10,
// 		Verbose:         true,
// 		Metadata:        nil,
// 		DoTimeoutSec:    40,
// 		StoreData:       false,
// 		RecycleData:     false,
// 		ReadFromCsvName: "",
// 		WriteToCsvName:  "",
// 		HandleParams:    nil,
// 		IsValidationRun: false,
// 		StopIf:          nil,
// 		Validation:      Validation{},
// 		DebugSleep:      0,
// 	}
// 	lm := NewLoadManager(suiteCfg, )
// 	r := NewRunner("testRunner", lm, ta, nil, runnerConfig)
// }
