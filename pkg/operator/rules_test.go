// Copyright 2022 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operator

import (
	"os"
	"strings"
	"testing"

	"github.com/go-kit/log"
	"k8s.io/apimachinery/pkg/util/intstr"

	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
)

func TestMakeRulesConfigMaps(t *testing.T) {
	t.Run("shouldAcceptRuleWithValidPartialResponseStrategyValue", shouldAcceptRuleWithValidPartialResponseStrategyValue)
	t.Run("shouldRejectRuleWithInvalidPartialResponseStrategyValue", shouldRejectRuleWithInvalidPartialResponseStrategyValue)
	t.Run("ShouldAcceptValidRule", shouldAcceptValidRule)
	t.Run("shouldRejectRuleWithInvalidLabels", shouldRejectRuleWithInvalidLabels)
	t.Run("shouldRejectRuleWithInvalidExpression", shouldRejectRuleWithInvalidExpression)
	t.Run("shouldResetRuleWithPartialResponseStrategySet", shouldResetRuleWithPartialResponseStrategySet)
	t.Run("validateFieldInAdmissionError", validateFieldInAdmissionError)
}

func shouldRejectRuleWithInvalidPartialResponseStrategyValue(t *testing.T) {
	rules := monitoringv1.PrometheusRuleSpec{Groups: []monitoringv1.RuleGroup{
		{
			Name:                    "group",
			PartialResponseStrategy: "invalid",
			Rules: []monitoringv1.Rule{
				{
					Alert: "alert",
					Expr:  intstr.FromString("vector(1)"),
				},
			},
		},
	}}
	_, err := generateRulesConfiguration(ThanosFormat, rules, log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout)))
	if err == nil {
		t.Fatalf("expected errors when parsing rule with invalid partial_response_strategy value")
	}
}

func shouldAcceptRuleWithValidPartialResponseStrategyValue(t *testing.T) {
	rules := monitoringv1.PrometheusRuleSpec{Groups: []monitoringv1.RuleGroup{
		{
			Name:                    "group",
			PartialResponseStrategy: "warn",
			Rules: []monitoringv1.Rule{
				{
					Alert: "alert",
					Expr:  intstr.FromString("vector(1)"),
				},
			},
		},
	}}
	content, _ := generateRulesConfiguration(ThanosFormat, rules, log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout)))
	if !strings.Contains(content, "partial_response_strategy: warn") {
		t.Fatalf("expected `partial_response_strategy` to be set in PrometheusRule as `warn`")

	}
}

func shouldAcceptValidRule(t *testing.T) {
	rules := monitoringv1.PrometheusRuleSpec{Groups: []monitoringv1.RuleGroup{
		{
			Name: "group",
			Rules: []monitoringv1.Rule{
				{
					Alert: "alert",
					Expr:  intstr.FromString("vector(1)"),
					Labels: map[string]string{
						"valid_label": "valid_value",
					},
				},
			},
		},
	}}
	_, err := generateRulesConfiguration(PrometheusFormat, rules, log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout)))
	if err != nil {
		t.Fatalf("expected no errors when parsing valid rule")
	}
}

func shouldRejectRuleWithInvalidLabels(t *testing.T) {
	rules := monitoringv1.PrometheusRuleSpec{Groups: []monitoringv1.RuleGroup{
		{
			Name: "group",
			Rules: []monitoringv1.Rule{
				{
					Alert: "alert",
					Expr:  intstr.FromString("vector(1)"),
					Labels: map[string]string{
						"invalid/label": "value",
					},
				},
			},
		},
	}}
	_, err := generateRulesConfiguration(PrometheusFormat, rules, log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout)))
	if err == nil {
		t.Fatalf("expected errors when parsing rule with invalid labels")
	}
}

func shouldRejectRuleWithInvalidExpression(t *testing.T) {
	rules := monitoringv1.PrometheusRuleSpec{Groups: []monitoringv1.RuleGroup{
		{
			Name: "group",
			Rules: []monitoringv1.Rule{
				{
					Alert: "alert",
					Expr:  intstr.FromString("invalidfn(1)"),
				},
			},
		},
	}}
	_, err := generateRulesConfiguration(PrometheusFormat, rules, log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout)))
	if err == nil {
		t.Fatalf("expected errors when parsing rule with invalid expression")
	}
}

func shouldResetRuleWithPartialResponseStrategySet(t *testing.T) {
	rules := monitoringv1.PrometheusRuleSpec{Groups: []monitoringv1.RuleGroup{
		{
			Name:                    "group",
			PartialResponseStrategy: "warn",
			Rules: []monitoringv1.Rule{
				{
					Alert: "alert",
					Expr:  intstr.FromString("vector(1)"),
				},
			},
		},
	}}
	content, _ := generateRulesConfiguration(PrometheusFormat, rules, log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout)))
	if strings.Contains(content, "partial_response_strategy") {
		t.Fatalf("expected `partial_response_strategy` removed from PrometheusRule")
	}
}

func validateFieldInAdmissionError(t *testing.T) {
	for _, tc := range []struct {
		name          string
		ruleSpec      monitoringv1.PrometheusRuleSpec
		expectedField string
	}{
		{
			name: "Invalid PartialResponseStrategy",
			ruleSpec: monitoringv1.PrometheusRuleSpec{Groups: []monitoringv1.RuleGroup{
				{
					Name:                    "group",
					PartialResponseStrategy: "invalid",
					Rules: []monitoringv1.Rule{
						{
							Alert: "alert",
							Expr:  intstr.FromString("vector(1)"),
						},
					},
				},
			}},
			expectedField: "groups[0].partial_response_strategy",
		},
		{
			name: "Invalid Rule",
			ruleSpec: monitoringv1.PrometheusRuleSpec{Groups: []monitoringv1.RuleGroup{
				{
					Name: "group",
					Rules: []monitoringv1.Rule{
						{
							Alert: "alert",
							Expr:  intstr.FromString("vector(1"),
						},
					},
				},
			}},
			expectedField: "groups[0].rules[0]",
		},
		{
			name: "Invalid Rule in second rule",
			ruleSpec: monitoringv1.PrometheusRuleSpec{Groups: []monitoringv1.RuleGroup{
				{
					Name: "group",
					Rules: []monitoringv1.Rule{
						{
							Alert: "alert",
							Expr:  intstr.FromString("vector(1)"),
						},
						{
							Alert: "alert",
							Expr:  intstr.FromString("vector(1"),
						},
					},
				},
			}},
			expectedField: "groups[0].rules[1]",
		},
		{
			name: "Invalid Rule in second group",
			ruleSpec: monitoringv1.PrometheusRuleSpec{Groups: []monitoringv1.RuleGroup{
				{
					Name: "group 1",
					Rules: []monitoringv1.Rule{
						{
							Alert: "alert",
							Expr:  intstr.FromString("vector(1)"),
						},
					},
				},
				{
					Name: "group 2",
					Rules: []monitoringv1.Rule{
						{
							Alert: "alert",
							Expr:  intstr.FromString("vector(1"),
						},
					},
				},
			}},
			expectedField: "groups[1].rules[0]",
		},
		{
			name: "Repeated group name",
			ruleSpec: monitoringv1.PrometheusRuleSpec{Groups: []monitoringv1.RuleGroup{
				{
					Name: "group",
					Rules: []monitoringv1.Rule{
						{
							Alert: "alert",
							Expr:  intstr.FromString("vector(1)"),
						},
					},
				},
				{
					Name: "group",
					Rules: []monitoringv1.Rule{
						{
							Alert: "alert",
							Expr:  intstr.FromString("vector(1)"),
						},
					},
				},
			}},
			expectedField: "groups",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			admissionErrors := ValidateRule(tc.ruleSpec)
			if len(admissionErrors) == 0 {
				t.Fatalf("expected errors when parsing invalid rule")
			}
			for _, admissionError := range admissionErrors {
				if tc.expectedField != admissionError.Field {
					t.Fatalf("field in admissionError doesn't match expected value: expected %s got %s", tc.expectedField, admissionError.Field)
				}
			}
		})
	}
}
