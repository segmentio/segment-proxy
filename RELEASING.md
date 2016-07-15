# Releasing

1. Verify the build by running `make build test docker`.
2. Update the `CHANGELOG.md` for the impending release.
3. Run `git commit -am "Prepare for release X.Y.Z."` (where X.Y.Z is the new version).
4. Run `git tag -a X.Y.Z -m "Version X.Y.Z"` (where X.Y.Z is the new version).
5. Run `git push && git push --tags`.
6. Run `make docker-push`.
7. Create a release on [Github](https://github.com/segmentio/segment-proxy/releases).
8. Upload the files from the `bin` folder.
