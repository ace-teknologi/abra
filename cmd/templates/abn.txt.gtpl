  ===
  Name: {{.Name}}
  {{- if .ASICNumber}}
  ACN:  {{.ASICNumber}}
  {{- end}}
  ABNs:
  {{- range .ABNs}}
  - {{.IdentifierValue}}{{if .IdentifierStatus}} {{.IdentifierStatus}}{{end}}{{if eq .IsCurrentIndicator "Y"}} (Current){{end}}{{if .ReplacedIdentifierValue}} {{.ReplacedIdentifierValue}}{{end}}{{if eq .ReplacedFrom.IsZero false}} {{.ReplacedFrom.Format "2006-01-02"}}{{end}}
  {{- end}}
  Addresses:
  {{- range .PhysicalAddresses}}
  - {{.StateCode}} {{.Postcode}}
  {{- end}}
  Type: {{.EntityType.EntityTypeCode}} ({{.EntityType.EntityDescription}})
  Statuses:
  {{- range .EntityStatuses}}
  - {{.EntityStatusCode}}({{.EffectiveFrom.Format "2006-01-02"}} {{if eq .EffectiveTo.IsZero false}}to {{.EffectiveTo.Format "2006-01-02"}}{{else}}onwards{{end}})
  {{- end}}
  All Names:
  {{- range .HumanNames}}
  - {{.GivenName}} {{.OtherGivenName}} {{.FamilyName}}
  {{- end}}
  {{- range .BusinessNames}}
  - {{.OrganisationName}}
  {{- end}}
  {{- range .MainNames}}
  - {{.OrganisationName}}
  {{- end}}
  {{- range .MainTradingNames}}
  - {{.OrganisationName}}
  {{- end}}
  {{- range .OtherTradingNames}}
  - {{.OrganisationName}}
  {{- end}}
  ---
  ABR Links:
  {{- range .ABNs}}
  - https://abr.business.gov.au/ABN/View?abn={{.IdentifierValue}}{{if .IdentifierStatus}} {{.IdentifierStatus}}{{end}}{{if eq .IsCurrentIndicator "Y"}} (Current){{end}}
  {{- end}}
  ===
