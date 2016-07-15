# Releasing

1. Verify the build by running `make build test docker`.
2. Update the version in `README.md`.
3. Update the `CHANGELOG.md` for the impending release.
4. Run `git commit -am "Prepare for release X.Y.Z."` (where X.Y.Z is the new version).
5. Run `git tag -a X.Y.Z -m "Version X.Y.Z"` (where X.Y.Z is the new version).
6. Run `git push && git push --tags`.
7. Run `make docker-push`.
8. Run `github-release release --user segmentio --repo segment-proxy --tag X.Y.Z --name "X.Y.Z"`
8. Run `github-release upload --user segmentio --repo segment-proxy --tag X.Y.Z --name "X.Y.Z" --file bin/segment-proxy`.
