# gitops-backend

This is the backend that provides an API for fetching information about
pipelines defined in Git repositories.

## API

### Authentication

The request parameters `secretNS` and `secretName` are used to identify a secret
with a token in the `token` field, this token is used to authenticate requests
to the upstream Git Hosting service e.g. GitHub or GitLab.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: my-secet
  namespace: testing
type: Opaque
data:
  token: <base64 encoded token>
```

### Parsing a pipelines file.

```
GET /pipelines?url=https://github.com/my-org/my-repo.git&ref=main&secretNS=testing&secretName=my-secret
```

This will fetch and parse the pipelines.yaml from the repo
`https://github.com/my-org/my-repo.git`, specifically the `main` branch.

Response:

```json
{
  "applications": [
    {
      "name": "my-application",
      "repo_url": "https://github.com/my-org/source-repo.git",
      environments: [
        "development",
        "staging",
        "production"
      ]
    }
  ]
}
```
