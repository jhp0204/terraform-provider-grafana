---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "grafana_user Resource - terraform-provider-grafana"
subcategory: ""
description: |-
  Official documentation https://grafana.com/docs/grafana/latest/administration/manage-users-and-permissions/manage-server-users/HTTP API https://grafana.com/docs/grafana/latest/http_api/user/
  This resource uses Grafana's admin APIs for creating and updating users which
  does not currently work with API Tokens. You must use basic auth.
---

# grafana_user (Resource)

* [Official documentation](https://grafana.com/docs/grafana/latest/administration/manage-users-and-permissions/manage-server-users/)
* [HTTP API](https://grafana.com/docs/grafana/latest/http_api/user/)

This resource uses Grafana's admin APIs for creating and updating users which
does not currently work with API Tokens. You must use basic auth.

## Example Usage

```terraform
resource "grafana_user" "staff" {
  email    = "staff.name@example.com"
  name     = "Staff Name"
  login    = "staff"
  password = "my-password"
  is_admin = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `email` (String) The email address of the Grafana user.
- `password` (String, Sensitive) The password for the Grafana user.

### Optional

- `is_admin` (Boolean) Whether to make user an admin. Defaults to `false`.
- `login` (String) The username for the Grafana user.
- `name` (String) The display name for the Grafana user.

### Read-Only

- `id` (String) The ID of this resource.
- `user_id` (Number) The numerical ID of the Grafana user.

## Import

Import is supported using the following syntax:

```shell
terraform import grafana_user.user_name {{user_id}}
```
