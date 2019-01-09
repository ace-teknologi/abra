  Name: {{.Name}} {{if eq .IsCurrentIndicator "Y"}}(current){{end}}
  ABN: {{.ABN.IdentifierValue}}{{if .ABN.IdentifierStatus}} {{.ABN.IdentifierStatus}}{{end}}{{if eq .ABN.IsCurrentIndicator "Y"}} (Current){{end}}{{if .ABN.ReplacedIdentifierValue}} {{.ABN.ReplacedIdentifierValue}}{{end}}{{if eq .ABN.ReplacedFrom.IsZero false}} {{.ABN.ReplacedFrom.Format "2006-01-02"}}{{end}}
  Addresses: {{.MainBusinessPhysicalAddress.StateCode}} {{.MainBusinessPhysicalAddress.Postcode}}
  Score: {{.Score}}%
