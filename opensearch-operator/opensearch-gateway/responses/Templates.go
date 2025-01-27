package responses

import "opensearch.opster.io/opensearch-gateway/requests"

type GetIndexTemplatesResponse struct {
	IndexTemplates []IndexTemplate `json:"index_templates"`
}

type IndexTemplate struct {
	Name          string                 `json:"name"`
	IndexTemplate requests.IndexTemplate `json:"index_template"`
}

type GetComponentTemplatesResponse struct {
	ComponentTemplates []ComponentTemplate `json:"component_templates"`
}

type ComponentTemplate struct {
	Name              string                     `json:"name"`
	ComponentTemplate requests.ComponentTemplate `json:"component_template"`
}
