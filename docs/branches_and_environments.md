# Branches / Environments

| Environment | Branch  | URL |
| ----------- | ------- | --- |
| Staging     | staging | TBD |
| Production  | main    | TBD |

## Development

For development we use the `develop` branch. Whenever starting to develop a new feature, please create the feature branch from the `develop` branch.
The naming rule is as follows,

```
feature-xxxxx-xxxxx(add something that describes the new feature)
ie: `feature-filter-locations-by-restaurants`
```

Also don't forget to create an issue for the respective feature you are working on, and attach the branch name, so the whole team can see in which branch you are working.

---

## Staging

When you want to test and deploy some feature in the staging environment, please follow the steps below:

1. Create a new branch from `staging` branch (ie: `merge-feature-xxxxx-into-staging`)
2. Merge the feature branch you want to test and create a Pull Request to `staging` branch
3. At least you'll need one reviewer approve in order to merge the changes
4. After merging into `staging` branch it would be deployed automatically to the staging environment

---

## Production

When you want to deploy some feature into production environment, please follow the steps below:

1. Create a new branch from `main` branch (ie: `merge-feature-xxxxx-into-main`)
2. Merge the feature branch you want to deploy and create a Pull Request to `main` branch
3. You'll need all reviewers approve in order to merge the changes
4. After merging into `main` branch it would be deployed automatically to the production environment
