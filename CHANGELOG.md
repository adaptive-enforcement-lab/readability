# Changelog

## [1.10.1](https://github.com/adaptive-enforcement-lab/readability/compare/v1.10.0...v1.10.1) (2025-12-14)


### Bug Fixes

* remove duplicate source archive signing from releases ([#146](https://github.com/adaptive-enforcement-lab/readability/issues/146)) ([67b8f11](https://github.com/adaptive-enforcement-lab/readability/commit/67b8f112cbb5ac7ccd0a11f941fc7e27dfcd05c4))

## [1.10.0](https://github.com/adaptive-enforcement-lab/readability/compare/v1.9.0...v1.10.0) (2025-12-14)


### Features

* add unified container SBOM with Trivy attestation ([#140](https://github.com/adaptive-enforcement-lab/readability/issues/140)) ([e7617b1](https://github.com/adaptive-enforcement-lab/readability/commit/e7617b162d2c2cde80a9c7e964598c5fe631661a))


### Maintenance

* **deps:** update actions/download-artifact digest to d3f86a1 ([#143](https://github.com/adaptive-enforcement-lab/readability/issues/143)) ([6f38f0d](https://github.com/adaptive-enforcement-lab/readability/commit/6f38f0d49e6b68cc4d79bef9b48547352d0ff041))
* **deps:** update docker/login-action digest to 5e57cd1 ([#137](https://github.com/adaptive-enforcement-lab/readability/issues/137)) ([8a00393](https://github.com/adaptive-enforcement-lab/readability/commit/8a003937b31114f2b29fb83f3cb4aae8f5e5425d))
* **deps:** update docker/metadata-action digest to c299e40 ([#138](https://github.com/adaptive-enforcement-lab/readability/issues/138)) ([4d210e6](https://github.com/adaptive-enforcement-lab/readability/commit/4d210e6b1cac55e6ea126dabceac2ebfb3f67259))
* **deps:** update docker/setup-buildx-action digest to e468171 ([#144](https://github.com/adaptive-enforcement-lab/readability/issues/144)) ([25cc3ba](https://github.com/adaptive-enforcement-lab/readability/commit/25cc3bafb640050e088a0c4137b0a2d8126e0a40))

## [1.9.0](https://github.com/adaptive-enforcement-lab/readability/compare/v1.8.1...v1.9.0) (2025-12-14)


### Features

* add container publishing to ghcr.io with Cosign signing ([#136](https://github.com/adaptive-enforcement-lab/readability/issues/136)) ([a2b78ff](https://github.com/adaptive-enforcement-lab/readability/commit/a2b78ff190b27ca609a0725ce9a12728e5865134))

## [1.8.1](https://github.com/adaptive-enforcement-lab/readability/compare/v1.8.0...v1.8.1) (2025-12-14)


### Bug Fixes

* trigger Scorecard after Release workflow completes ([#134](https://github.com/adaptive-enforcement-lab/readability/issues/134)) ([5450628](https://github.com/adaptive-enforcement-lab/readability/commit/5450628b523fab274dcc2acdec3eee69e2986da1))

## [1.8.0](https://github.com/adaptive-enforcement-lab/readability/compare/v1.7.1...v1.8.0) (2025-12-14)


### Features

* add Go fuzz tests for OpenSSF Scorecard compliance ([#132](https://github.com/adaptive-enforcement-lab/readability/issues/132)) ([dd4d038](https://github.com/adaptive-enforcement-lab/readability/commit/dd4d038d9750751846e1da2842cd8e15db4ab920))

## [1.7.1](https://github.com/adaptive-enforcement-lab/readability/compare/v1.7.0...v1.7.1) (2025-12-14)


### Code Refactoring

* simplify SLSA provenance workflow ([#130](https://github.com/adaptive-enforcement-lab/readability/issues/130)) ([0a03624](https://github.com/adaptive-enforcement-lab/readability/commit/0a03624d444e726b7105214f1d7dc083cac54d1f))


### Maintenance

* remove sign-releases workflow ([#129](https://github.com/adaptive-enforcement-lab/readability/issues/129)) ([e18457f](https://github.com/adaptive-enforcement-lab/readability/commit/e18457f7a29a304674f4d8661367c3d7663cdf6c))

## [1.7.0](https://github.com/adaptive-enforcement-lab/readability/compare/v1.6.2...v1.7.0) (2025-12-14)


### Features

* add SLSA provenance generation to releases ([#127](https://github.com/adaptive-enforcement-lab/readability/issues/127)) ([1f5c92d](https://github.com/adaptive-enforcement-lab/readability/commit/1f5c92d9af7de2f112b09bca5e84949ce77078ed))
* add workflow to sign existing releases ([#121](https://github.com/adaptive-enforcement-lab/readability/issues/121)) ([449c4a5](https://github.com/adaptive-enforcement-lab/readability/commit/449c4a5597abe8a557d9bcf06fee4878834f2a05))


### Bug Fixes

* move permissions to job level in sign-releases workflow ([#126](https://github.com/adaptive-enforcement-lab/readability/issues/126)) ([2565c22](https://github.com/adaptive-enforcement-lab/readability/commit/2565c22fe4a16340632238eb2274044079bccb0a))
* sign source code archives in releases ([#122](https://github.com/adaptive-enforcement-lab/readability/issues/122)) ([c86288e](https://github.com/adaptive-enforcement-lab/readability/commit/c86288eb2a91f7a10ecc1a736010a80f9f352d1b))


### Maintenance

* **deps:** update sigstore/cosign-installer action to v3.10.1 ([#123](https://github.com/adaptive-enforcement-lab/readability/issues/123)) ([bdba92e](https://github.com/adaptive-enforcement-lab/readability/commit/bdba92eaaaa8a780abd0c94a7a19735ef73052d4))
* **deps:** update sigstore/cosign-installer action to v4 ([#124](https://github.com/adaptive-enforcement-lab/readability/issues/124)) ([b264d48](https://github.com/adaptive-enforcement-lab/readability/commit/b264d4829f05726052bb9266568f35b5869a6788))

## [1.6.2](https://github.com/adaptive-enforcement-lab/readability/compare/v1.6.1...v1.6.2) (2025-12-14)


### Bug Fixes

* use .sig extension for cosign signatures ([#119](https://github.com/adaptive-enforcement-lab/readability/issues/119)) ([135e617](https://github.com/adaptive-enforcement-lab/readability/commit/135e6177d11847c01ca1ae61078be01c6b1dff0a))

## [1.6.1](https://github.com/adaptive-enforcement-lab/readability/compare/v1.6.0...v1.6.1) (2025-12-14)


### Bug Fixes

* update cosign signing for v3 API ([#117](https://github.com/adaptive-enforcement-lab/readability/issues/117)) ([ad12a83](https://github.com/adaptive-enforcement-lab/readability/commit/ad12a83406b0e9c23738063e715d657afe14bd12))

## [1.6.0](https://github.com/adaptive-enforcement-lab/readability/compare/v1.5.1...v1.6.0) (2025-12-13)


### Features

* add cosign signing for release artifacts ([#116](https://github.com/adaptive-enforcement-lab/readability/issues/116)) ([300bd9b](https://github.com/adaptive-enforcement-lab/readability/commit/300bd9b30d6dd5a2d6be9247a288ef5faeff7a38))


### Bug Fixes

* move workflow permissions to job level for least privilege ([#106](https://github.com/adaptive-enforcement-lab/readability/issues/106)) ([b168c1f](https://github.com/adaptive-enforcement-lab/readability/commit/b168c1ffeed27d4e34a0582f3d6ad32c64925115))
* use version tag for scorecard-action (cannot pin to SHA) ([#111](https://github.com/adaptive-enforcement-lab/readability/issues/111)) ([d07300c](https://github.com/adaptive-enforcement-lab/readability/commit/d07300cf0a2df11e379dec9d711c0675a2e1f599))
* use version tags for all actions in scorecard.yml ([#112](https://github.com/adaptive-enforcement-lab/readability/issues/112)) ([39ce338](https://github.com/adaptive-enforcement-lab/readability/commit/39ce338d46de276ba453e92cbfc1bddd51f750c4))


### Maintenance

* **deps:** update github/codeql-action action to v4 ([#104](https://github.com/adaptive-enforcement-lab/readability/issues/104)) ([8c0e6e3](https://github.com/adaptive-enforcement-lab/readability/commit/8c0e6e3864ce26cdfb3bc041c619a8049b3b505e))
* **deps:** update github/codeql-action digest to 1b168cd ([#109](https://github.com/adaptive-enforcement-lab/readability/issues/109)) ([7dec92d](https://github.com/adaptive-enforcement-lab/readability/commit/7dec92d951ef565391564f5ea6ef8b9c719ced6a))
* **deps:** update googleapis/release-please-action digest to 16a9c90 ([#110](https://github.com/adaptive-enforcement-lab/readability/issues/110)) ([1faf2e9](https://github.com/adaptive-enforcement-lab/readability/commit/1faf2e92bb72ccda6beb7171d1f03d9d63b2e33b))
* **deps:** update ossf/scorecard-action action to v2.4.3 ([#103](https://github.com/adaptive-enforcement-lab/readability/issues/103)) ([6d3834a](https://github.com/adaptive-enforcement-lab/readability/commit/6d3834a04bb8d4ad3a25e7789fbf79ec0191b0a0))
* improve OpenSSF Scorecard score ([#114](https://github.com/adaptive-enforcement-lab/readability/issues/114)) ([011a4d7](https://github.com/adaptive-enforcement-lab/readability/commit/011a4d73432583357c5fc58aec3436e9501f800e))
* pin all GitHub Actions to commit SHAs ([#108](https://github.com/adaptive-enforcement-lab/readability/issues/108)) ([b9840c7](https://github.com/adaptive-enforcement-lab/readability/commit/b9840c71f8adcc20995624f541859f4f0271cb08))

## [1.5.1](https://github.com/adaptive-enforcement-lab/readability/compare/1.5.0...v1.5.1) (2025-12-13)


### Bug Fixes

* enable blank issues for free-form issue creation ([#96](https://github.com/adaptive-enforcement-lab/readability/issues/96)) ([e5c50a2](https://github.com/adaptive-enforcement-lab/readability/commit/e5c50a2b02539c99416b4bcb8091f96e1adf1475))
* enable v prefix in release tags for Go module compliance ([#101](https://github.com/adaptive-enforcement-lab/readability/issues/101)) ([e3ab84e](https://github.com/adaptive-enforcement-lab/readability/commit/e3ab84ece59e2708fb847ec7c46a8d1fe46c57b9))

## [1.5.0](https://github.com/adaptive-enforcement-lab/readability/compare/1.4.0...1.5.0) (2025-12-13)


### Features

* increase coverage threshold to 95% with single source of truth ([#91](https://github.com/adaptive-enforcement-lab/readability/issues/91)) ([0a1266c](https://github.com/adaptive-enforcement-lab/readability/commit/0a1266c76acc1bfc60621811c5e6ebb90154200f))


### Bug Fixes

* pass CODECOV_TOKEN to reusable CI workflow ([#88](https://github.com/adaptive-enforcement-lab/readability/issues/88)) ([a98ba23](https://github.com/adaptive-enforcement-lab/readability/commit/a98ba2390f359e3f2bf74f1244bd21497ed4aef3))
* update component paths to use regex patterns ([#90](https://github.com/adaptive-enforcement-lab/readability/issues/90)) ([0931deb](https://github.com/adaptive-enforcement-lab/readability/commit/0931deb216eb31a09c4f2277319e53d6acdb775c))

## [1.4.0](https://github.com/adaptive-enforcement-lab/readability/compare/1.3.1...1.4.0) (2025-12-13)


### Features

* add Codecov configuration with components and test analytics ([#86](https://github.com/adaptive-enforcement-lab/readability/issues/86)) ([333c094](https://github.com/adaptive-enforcement-lab/readability/commit/333c09499608b0096e9dd590af1b5d4355f55a10))

## [1.3.1](https://github.com/adaptive-enforcement-lab/readability/compare/1.3.0...1.3.1) (2025-12-13)


### Bug Fixes

* improve Go Report Card compliance and pre-commit hooks ([#84](https://github.com/adaptive-enforcement-lab/readability/issues/84)) ([d42947d](https://github.com/adaptive-enforcement-lab/readability/commit/d42947d1ed580c4df356241e5c809170d2ba61ef))
* update CI badge to point to ci.yml ([#82](https://github.com/adaptive-enforcement-lab/readability/issues/82)) ([9a64c38](https://github.com/adaptive-enforcement-lab/readability/commit/9a64c3819fcafad4707e9e0bbf6047e2afc172a6))


### Code Refactoring

* reduce cyclomatic complexity and enforce strict gocyclo ([#85](https://github.com/adaptive-enforcement-lab/readability/issues/85)) ([bed6cae](https://github.com/adaptive-enforcement-lab/readability/commit/bed6caea4524a04b3de3fabc3932377acaa8f81d))

## [1.3.0](https://github.com/adaptive-enforcement-lab/readability/compare/1.2.2...1.3.0) (2025-12-13)


### Features

* add Trivy security scanning and SBOM generation ([#80](https://github.com/adaptive-enforcement-lab/readability/issues/80)) ([844464f](https://github.com/adaptive-enforcement-lab/readability/commit/844464f6bd793d51b7610c8ec355596aace6d119))


### Bug Fixes

* remove duplicate push trigger from ci.yml ([#78](https://github.com/adaptive-enforcement-lab/readability/issues/78)) ([63dd296](https://github.com/adaptive-enforcement-lab/readability/commit/63dd2963bf2cba27af0a68c55733365a6a9c96cc))


### Maintenance

* **deps:** update github artifact actions ([#81](https://github.com/adaptive-enforcement-lab/readability/issues/81)) ([4cde10f](https://github.com/adaptive-enforcement-lab/readability/commit/4cde10fcc28839113c04fd378e60e36dcecdd189))

## [1.2.2](https://github.com/adaptive-enforcement-lab/readability/compare/1.2.1...1.2.2) (2025-12-13)


### Code Refactoring

* unify CI/CD with reusable workflow pattern ([#76](https://github.com/adaptive-enforcement-lab/readability/issues/76)) ([df8511a](https://github.com/adaptive-enforcement-lab/readability/commit/df8511a1a028a9d3cae6277a88582ba02f466011))

## [1.2.1](https://github.com/adaptive-enforcement-lab/readability/compare/1.2.0...1.2.1) (2025-12-13)


### Bug Fixes

* configure Codecov with OIDC authentication ([#75](https://github.com/adaptive-enforcement-lab/readability/issues/75)) ([894311e](https://github.com/adaptive-enforcement-lab/readability/commit/894311e7cbfca3ad7bedcd776067e2575f0b40d9))


### Maintenance

* **deps:** update github artifact actions ([#71](https://github.com/adaptive-enforcement-lab/readability/issues/71)) ([2ae0a49](https://github.com/adaptive-enforcement-lab/readability/commit/2ae0a492c33972c75b21efefe7256b9d4ff66a1d))

## [1.2.0](https://github.com/adaptive-enforcement-lab/readability/compare/1.1.1...1.2.0) (2025-12-09)


### Features

* publish to GitHub Marketplace ([#70](https://github.com/adaptive-enforcement-lab/readability/issues/70)) ([4577d6d](https://github.com/adaptive-enforcement-lab/readability/commit/4577d6d7c84b65adc0476c41133fbd53b76f2bed))


### Bug Fixes

* rename action for GitHub Marketplace uniqueness ([f3df887](https://github.com/adaptive-enforcement-lab/readability/commit/f3df887f1aa10378ce2841647d3d9cb1fe8496d7))
* shorten action description for Marketplace limit ([e18fd8d](https://github.com/adaptive-enforcement-lab/readability/commit/e18fd8d0e38979b57b8ed3dcfd574ec65029cad2))


### Maintenance

* ignore .cache directory ([430480c](https://github.com/adaptive-enforcement-lab/readability/commit/430480c956a0fbd644976e6be6c58f5a8909c728))

## [1.1.1](https://github.com/adaptive-enforcement-lab/readability/compare/1.1.0...1.1.1) (2025-12-09)


### Bug Fixes

* add pillow and cairosvg for social cards in CI ([5cba44e](https://github.com/adaptive-enforcement-lab/readability/commit/5cba44e02abe1b68a5e4a412439f02de8ae9e37a))

## [1.1.0](https://github.com/adaptive-enforcement-lab/readability/compare/1.0.0...1.1.0) (2025-12-09)


### Features

* add social cards plugin and fix duplicate nav entry ([7935b61](https://github.com/adaptive-enforcement-lab/readability/commit/7935b617b0538168ec3d718701a49edcb8735c04))

## [1.0.0](https://github.com/adaptive-enforcement-lab/readability/compare/0.11.1...1.0.0) (2025-12-09)


### ⚠ BREAKING CHANGES

* Documentation structure reorganized. New introduction and use-cases pages added. All existing documentation rewritten for clarity and accessibility.

### Bug Fixes

* add admonitions to nav and fix anchor links ([#64](https://github.com/adaptive-enforcement-lab/readability/issues/64)) ([e4f8c8f](https://github.com/adaptive-enforcement-lab/readability/commit/e4f8c8fe2615f33b80016b3e194b4fc62b6bcfa0))


### Documentation

* comprehensive documentation overhaul for newcomer accessibility ([#65](https://github.com/adaptive-enforcement-lab/readability/issues/65)) ([9d7e446](https://github.com/adaptive-enforcement-lab/readability/commit/9d7e446054dd30cfeb4dbbaf3024f4ea45b52faf))

## [0.11.1](https://github.com/adaptive-enforcement-lab/readability/compare/0.11.0...0.11.1) (2025-12-09)


### Bug Fixes

* remove homepage override so logo links to docs site ([#62](https://github.com/adaptive-enforcement-lab/readability/issues/62)) ([772654e](https://github.com/adaptive-enforcement-lab/readability/commit/772654ee11a8da18df2af71aec770662e9f35b7c))

## [0.11.0](https://github.com/adaptive-enforcement-lab/readability/compare/0.10.1...0.11.0) (2025-12-07)


### Features

* add Issues column to markdown results table ([#60](https://github.com/adaptive-enforcement-lab/readability/issues/60)) ([692b522](https://github.com/adaptive-enforcement-lab/readability/commit/692b522350e4a37e63d4dc81559846814befde2b))

## [0.10.1](https://github.com/adaptive-enforcement-lab/readability/compare/0.10.0...0.10.1) (2025-12-07)


### Bug Fixes

* complete reading time ceiling division fixes ([#59](https://github.com/adaptive-enforcement-lab/readability/issues/59)) ([cfd4cdb](https://github.com/adaptive-enforcement-lab/readability/commit/cfd4cdbb84903596a3f1e29f6bf31cd216a1d0f9))
* use ceiling division for reading time calculation ([#57](https://github.com/adaptive-enforcement-lab/readability/issues/57)) ([2e8b644](https://github.com/adaptive-enforcement-lab/readability/commit/2e8b644262e95d519b6e557ddb5fedb78c97890a))

## [0.10.0](https://github.com/adaptive-enforcement-lab/readability/compare/0.9.2...0.10.0) (2025-12-07)


### Features

* add linter-style diagnostic output format ([#55](https://github.com/adaptive-enforcement-lab/readability/issues/55)) ([2e756e1](https://github.com/adaptive-enforcement-lab/readability/commit/2e756e139e6ddf4b98a3efd6a7310e6cdabdb85d))

## [0.9.2](https://github.com/adaptive-enforcement-lab/readability/compare/0.9.1...0.9.2) (2025-12-07)


### Bug Fixes

* **action:** resolve bash syntax errors and stdout duplication ([#51](https://github.com/adaptive-enforcement-lab/readability/issues/51)) ([8f26ed6](https://github.com/adaptive-enforcement-lab/readability/commit/8f26ed634971abc83f921b9f0ce9306925b12eff))

## [0.9.1](https://github.com/adaptive-enforcement-lab/readability/compare/0.9.0...0.9.1) (2025-12-06)


### Bug Fixes

* handle absolute paths in override path matching ([#49](https://github.com/adaptive-enforcement-lab/readability/issues/49)) ([105c5c0](https://github.com/adaptive-enforcement-lab/readability/commit/105c5c01f779b38ec61f329f87e22564fc09eadd))

## [0.9.0](https://github.com/adaptive-enforcement-lab/readability/compare/0.8.0...0.9.0) (2025-12-06)


### Features

* add MkDocs-style admonition detection and threshold check ([#47](https://github.com/adaptive-enforcement-lab/readability/issues/47)) ([9a41aac](https://github.com/adaptive-enforcement-lab/readability/commit/9a41aac9004438bc3ef3542d0573c33aa5a9ff33))

## [0.8.0](https://github.com/adaptive-enforcement-lab/readability/compare/0.7.2...0.8.0) (2025-12-06)


### Features

* add readability improvement hints on check failure ([2d5fc16](https://github.com/adaptive-enforcement-lab/readability/commit/2d5fc169cd12a6b66ed9bc8164b9026e81a55219))
* add warning to split files instead of removing content ([035ebed](https://github.com/adaptive-enforcement-lab/readability/commit/035ebed63245dd5e2085014a2df07b4aea1e2bb8))

## [0.7.2](https://github.com/adaptive-enforcement-lab/readability/compare/0.7.1...0.7.2) (2025-12-06)


### Maintenance

* add site/ to gitignore ([37bc3af](https://github.com/adaptive-enforcement-lab/readability/commit/37bc3afa4e3045370a17163a46c50a3d5b80ba00))
* **deps:** update actions/checkout action to v6 ([#40](https://github.com/adaptive-enforcement-lab/readability/issues/40)) ([edc40d1](https://github.com/adaptive-enforcement-lab/readability/commit/edc40d1cb49ac607c2b8b8098a6cf09bff617ee6))
* **deps:** update actions/setup-go action to v6 ([#41](https://github.com/adaptive-enforcement-lab/readability/issues/41)) ([9276f52](https://github.com/adaptive-enforcement-lab/readability/commit/9276f52e9e0ea05ef70f569eaad4c3993667d1fd))
* **deps:** update actions/setup-python action to v6 ([#42](https://github.com/adaptive-enforcement-lab/readability/issues/42)) ([cc545e4](https://github.com/adaptive-enforcement-lab/readability/commit/cc545e405a12b85a71a715e28b4871ac580c09b0))
* **deps:** update dependency go to 1.25 ([#36](https://github.com/adaptive-enforcement-lab/readability/issues/36)) ([9754456](https://github.com/adaptive-enforcement-lab/readability/commit/9754456098823a64532982cd8ed97c788a21b358))
* **deps:** update dependency python to 3.14 ([#37](https://github.com/adaptive-enforcement-lab/readability/issues/37)) ([6c8abbe](https://github.com/adaptive-enforcement-lab/readability/commit/6c8abbe768e219f0797b36e65b8daefe75a5eb9b))
* **deps:** update github artifact actions ([#43](https://github.com/adaptive-enforcement-lab/readability/issues/43)) ([f498a0e](https://github.com/adaptive-enforcement-lab/readability/commit/f498a0e3e4a82cc87187cfa6af0a43efae648f84))
* **deps:** update golangci/golangci-lint-action action to v9 ([#44](https://github.com/adaptive-enforcement-lab/readability/issues/44)) ([4a57751](https://github.com/adaptive-enforcement-lab/readability/commit/4a57751e0734850aff04e99462c1260d5fd61426))
* **deps:** update tj-actions/changed-files action to v47 ([#45](https://github.com/adaptive-enforcement-lab/readability/issues/45)) ([4333c5e](https://github.com/adaptive-enforcement-lab/readability/commit/4333c5eb13f5c6b39a416daed40af2f55c08c417))

## [0.7.1](https://github.com/adaptive-enforcement-lab/readability/compare/0.7.0...0.7.1) (2025-12-06)


### Maintenance

* add MIT license ([098e115](https://github.com/adaptive-enforcement-lab/readability/commit/098e115fa2ec30c2cfd24df09cb7d6decafeb757))

## [0.7.0](https://github.com/adaptive-enforcement-lab/readability/compare/0.6.0...0.7.0) (2025-12-06)


### Features

* add floating version tag aliases after release ([#31](https://github.com/adaptive-enforcement-lab/readability/issues/31)) ([733cf04](https://github.com/adaptive-enforcement-lab/readability/commit/733cf04a51602019fbb9d1a884473b8c64277caa))

## [0.6.0](https://github.com/adaptive-enforcement-lab/readability/compare/0.5.0...0.6.0) (2025-12-06)


### Features

* use pre-built binary and add pre-commit hook support ([#28](https://github.com/adaptive-enforcement-lab/readability/issues/28)) ([6e7bfc7](https://github.com/adaptive-enforcement-lab/readability/commit/6e7bfc7586f5e120c67f3138101d835826a24e75))

## [0.5.0](https://github.com/adaptive-enforcement-lab/readability/compare/0.4.0...0.5.0) (2025-12-06)


### Features

* enhance summary table with lines, reading time, and metric links ([#25](https://github.com/adaptive-enforcement-lab/readability/issues/25)) ([cf62405](https://github.com/adaptive-enforcement-lab/readability/commit/cf62405c31cafe4663df46c017c2fca52762eeaa))

## [0.4.0](https://github.com/adaptive-enforcement-lab/readability/compare/0.3.2...0.4.0) (2025-12-06)


### Features

* add automatic job summary generation ([#23](https://github.com/adaptive-enforcement-lab/readability/issues/23)) ([315d631](https://github.com/adaptive-enforcement-lab/readability/commit/315d6317f6bbd6c28ea4c0276b99e5faa76deb30)), closes [#22](https://github.com/adaptive-enforcement-lab/readability/issues/22)

## [0.3.2](https://github.com/adaptive-enforcement-lab/readability/compare/0.3.1...0.3.2) (2025-12-06)


### Bug Fixes

* add value mappings for composite action outputs ([#20](https://github.com/adaptive-enforcement-lab/readability/issues/20)) ([99cfb35](https://github.com/adaptive-enforcement-lab/readability/commit/99cfb35f4c469dd3d60d0f450dc98ab6477b1c21))

## [0.3.1](https://github.com/adaptive-enforcement-lab/readability/compare/v0.3.0...0.3.1) (2025-12-06)


### Maintenance

* remove v prefix from release tags ([#18](https://github.com/adaptive-enforcement-lab/readability/issues/18)) ([b431939](https://github.com/adaptive-enforcement-lab/readability/commit/b43193980346a87efa000977bc94523ac63cd9a4))
* rename CI to Build and run on PRs only ([#17](https://github.com/adaptive-enforcement-lab/readability/issues/17)) ([8f9c13e](https://github.com/adaptive-enforcement-lab/readability/commit/8f9c13e485336f56c24d320cd75ea3fed35532b5))

## [0.3.0](https://github.com/adaptive-enforcement-lab/readability/compare/v0.2.0...v0.3.0) (2025-12-06)


### Features

* add --version flag with ldflags injection ([#16](https://github.com/adaptive-enforcement-lab/readability/issues/16)) ([f08e561](https://github.com/adaptive-enforcement-lab/readability/commit/f08e561e92ead48c5eade4b1f3f755a6fab77c47))


### Bug Fixes

* action config auto-detection and outputs ([#14](https://github.com/adaptive-enforcement-lab/readability/issues/14)) ([f4edc89](https://github.com/adaptive-enforcement-lab/readability/commit/f4edc8968d7771341d1a362634075b807ac8c6bb))

## [0.2.0](https://github.com/adaptive-enforcement-lab/readability/compare/v0.1.1...v0.2.0) (2025-12-06)


### Features

* add release-please and MkDocs Material documentation ([#2](https://github.com/adaptive-enforcement-lab/readability/issues/2)) ([25d37ea](https://github.com/adaptive-enforcement-lab/readability/commit/25d37eadebc6da81c0433b20928e0c68e6053ae9))
