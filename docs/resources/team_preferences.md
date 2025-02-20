---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "grafana_team_preferences Resource - terraform-provider-grafana"
subcategory: ""
description: |-
  Official documentation https://grafana.com/docs/grafana/latest/administration/preferences/HTTP API https://grafana.com/docs/grafana/latest/http_api/team/
---

# grafana_team_preferences (Resource)

* [Official documentation](https://grafana.com/docs/grafana/latest/administration/preferences/)
* [HTTP API](https://grafana.com/docs/grafana/latest/http_api/team/)

## Example Usage

```terraform
resource "grafana_dashboard" "metrics" {
  config_json = file("grafana-dashboard.json")
}

resource "grafana_team" "team" {
  name = "Team Name"
}

resource "grafana_team_preferences" "team_preferences" {
  team_id           = grafana_team.team.id
  theme             = "dark"
  timezone          = "browser"
  home_dashboard_id = grafana_dashboard.metrics.dashboard_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `team_id` (Number) The numeric team ID.

### Optional

- `home_dashboard_id` (Number) The numeric ID of the dashboard to display when a team member logs in.
- `theme` (String) The theme for the specified team. Available themes are `light`, `dark`, or an empty string for the default theme.
- `timezone` (String) The timezone for the specified team. Available values are `utc`, `browser`, or an empty string for the default.

### Read-Only

- `id` (String) The ID of this resource.


