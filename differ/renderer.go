/*
Copyright © 2021, 2022 Red Hat, Inc.

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

package differ

// Generated documentation is available at:
// https://pkg.go.dev/github.com/RedHatInsights/ccx-notification-service/differ
//
// Documentation in literate-programming-style is available at:
// https://redhatinsights.github.io/ccx-notification-service/packages/differ/renderer.html

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/RedHatInsights/ccx-notification-service/conf"
	"github.com/RedHatInsights/ccx-notification-service/types"
	httputils "github.com/RedHatInsights/insights-operator-utils/http"
	utypes "github.com/RedHatInsights/insights-results-types"
)

func renderReportsForCluster(
	config conf.DependenciesConfiguration,
	clusterName types.ClusterName,
	reports []types.ReportItem,
	ruleContent types.RulesMap) ([]types.RenderedReport, error) {

	req, err := createTemplateRendererRequest(ruleContent, reports, clusterName, config.TemplateRendererURL)
	if err != nil {
		log.Error().Err(err).Msg("Request to content template renderer could not be created")
		return nil, err
	}

	body, err := httputils.SendRequest(req, 10*time.Second)
	if err != nil {
		log.Error().Err(err).Msg("Request to content template renderer could not be processed")
		return nil, err
	}

	var receivedResult types.TemplateRendererOutput

	err = json.Unmarshal(body, &receivedResult)
	if err != nil {
		log.Error().Err(err).Msg("Error trying to decode template renderer output from received answer")
		return nil, err
	}

	return receivedResult.Reports[clusterName], nil
}

func getAllContentFromMap(ruleContent types.RulesMap) []utypes.RuleContent {
	contents := make([]utypes.RuleContent, len(ruleContent))

	i := 0
	for key := range ruleContent {
		contents[i] = ruleContent[key]
		i++
	}

	return contents
}

func createTemplateRendererRequest(
	ruleContent types.RulesMap,
	reports []types.ReportItem,
	clusterName types.ClusterName,
	rendererURL string) (*http.Request, error) {

	content := getAllContentFromMap(ruleContent)

	requestBody := types.TemplateRendererRequestBody{
		Content: content,
		ReportData: types.ReportData{
			Reports: map[types.ClusterName]types.Report{
				clusterName: {Reports: reports},
			},
			Clusters: []types.ClusterName{clusterName},
		},
	}
	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		log.Error().Err(err).Msg("Got error while creating json with content and report data")
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, rendererURL, bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Error().Err(err).Msg("Got error while setting up HTTP request for template renderer")
		return nil, err
	}
	return req, nil
}
