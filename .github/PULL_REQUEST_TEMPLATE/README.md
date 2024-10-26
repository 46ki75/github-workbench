## Default PR Template

The default PR template used is `.github/pull_request_template.md`.

## Using Multiple PR Templates

It's a bit inconvenient, but by specifying a template in the `.github/PULL_REQUEST_TEMPLATE/` directory with query parameters, you can use different PR templates for different purposes. However, you must specify the branch name, which adds to the inconvenience, and unfortunately, GitHub’s UI doesn’t provide a dropdown menu to select templates.

Here’s an example with query parameters specified: [Create PR (release -> main)](https://github.com/46ki75/github-workbench/compare/main...release?quick_pull=1&template=release.md)

For more on query parameters, see [here](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/using-query-parameters-to-create-a-pull-request).

## Official Documentation

For further details, refer to the [official documentation](https://docs.github.com/en/communities/using-templates-to-encourage-useful-issues-and-pull-requests/creating-a-pull-request-template-for-your-repository).
