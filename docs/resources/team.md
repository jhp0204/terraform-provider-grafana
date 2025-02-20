---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "grafana_team Resource - terraform-provider-grafana"
subcategory: ""
description: |-
  Official documentation https://grafana.com/docs/grafana/latest/administration/manage-users-and-permissions/manage-teams/HTTP API https://grafana.com/docs/grafana/latest/http_api/team/
---

# grafana_team (Resource)

* [Official documentation](https://grafana.com/docs/grafana/latest/administration/manage-users-and-permissions/manage-teams/)
* [HTTP API](https://grafana.com/docs/grafana/latest/http_api/team/)

## Example Usage

```terraform
resource "grafana_team" "test-team" {
  name  = "Test Team"
  email = "teamemail@example.com"
  members = [
    "viewer-01@example.com"
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The display name for the Grafana team created.

### Optional

- `email` (String) An email address for the team.
- `members` (Set of String) A set of email addresses corresponding to users who should be given membership
to the team. Note: users specified here must already exist in Grafana.

### Read-Only

- `id` (String) The ID of this resource.
- `team_id` (Number) The team id assigned to this team by Grafana.


