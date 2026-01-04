# Changelog

## [3.0.1](https://github.com/adaptive-enforcement-lab/readability/compare/v3.0.0...v3.0.1) (2026-01-04)


### Bug Fixes

* skip excluded files in diagnostic output ([#223](https://github.com/adaptive-enforcement-lab/readability/issues/223)) ([2a9e324](https://github.com/adaptive-enforcement-lab/readability/commit/2a9e324ee01bd989b78805369070c98eeac41308))

## [3.0.0](https://github.com/adaptive-enforcement-lab/readability/compare/v2.0.0...v3.0.0) (2026-01-04)


### ⚠ BREAKING CHANGES

* Removed hardcoded CHANGELOG.md and CONTRIBUTING.md skip logic. These files must now be explicitly excluded in .readability.yml or they will be analyzed and may fail threshold checks. Add the following to your config to preserve the old behavior:
* Documentation structure reorganized. New introduction and use-cases pages added. All existing documentation rewritten for clarity and accessibility.

### Features

* add --version flag with ldflags injection ([#16](https://github.com/adaptive-enforcement-lab/readability/issues/16)) ([f08e561](https://github.com/adaptive-enforcement-lab/readability/commit/f08e561e92ead48c5eade4b1f3f755a6fab77c47))
* add automatic job summary generation ([#23](https://github.com/adaptive-enforcement-lab/readability/issues/23)) ([315d631](https://github.com/adaptive-enforcement-lab/readability/commit/315d6317f6bbd6c28ea4c0276b99e5faa76deb30)), closes [#22](https://github.com/adaptive-enforcement-lab/readability/issues/22)
* add Codecov configuration with components and test analytics ([#86](https://github.com/adaptive-enforcement-lab/readability/issues/86)) ([333c094](https://github.com/adaptive-enforcement-lab/readability/commit/333c09499608b0096e9dd590af1b5d4355f55a10))
* add container publishing to ghcr.io with Cosign signing ([#136](https://github.com/adaptive-enforcement-lab/readability/issues/136)) ([a2b78ff](https://github.com/adaptive-enforcement-lab/readability/commit/a2b78ff190b27ca609a0725ce9a12728e5865134))
* add cosign signing for release artifacts ([#116](https://github.com/adaptive-enforcement-lab/readability/issues/116)) ([300bd9b](https://github.com/adaptive-enforcement-lab/readability/commit/300bd9b30d6dd5a2d6be9247a288ef5faeff7a38))
* add exclude field to path overrides ([#216](https://github.com/adaptive-enforcement-lab/readability/issues/216)) ([#218](https://github.com/adaptive-enforcement-lab/readability/issues/218)) ([e913d3b](https://github.com/adaptive-enforcement-lab/readability/commit/e913d3beae7d880a8f5be0147ffc5fca84b0730b))
* add floating version tag aliases after release ([#31](https://github.com/adaptive-enforcement-lab/readability/issues/31)) ([733cf04](https://github.com/adaptive-enforcement-lab/readability/commit/733cf04a51602019fbb9d1a884473b8c64277caa))
* add Go fuzz tests for OpenSSF Scorecard compliance ([#132](https://github.com/adaptive-enforcement-lab/readability/issues/132)) ([dd4d038](https://github.com/adaptive-enforcement-lab/readability/commit/dd4d038d9750751846e1da2842cd8e15db4ab920))
* add Issues column to markdown results table ([#60](https://github.com/adaptive-enforcement-lab/readability/issues/60)) ([692b522](https://github.com/adaptive-enforcement-lab/readability/commit/692b522350e4a37e63d4dc81559846814befde2b))
* add JSON Schema for .readability.yml configuration ([#180](https://github.com/adaptive-enforcement-lab/readability/issues/180)) ([f80543f](https://github.com/adaptive-enforcement-lab/readability/commit/f80543ff0fba7a6e77f524ef16d20d662625a026)), closes [#160](https://github.com/adaptive-enforcement-lab/readability/issues/160)
* add linter-style diagnostic output format ([#55](https://github.com/adaptive-enforcement-lab/readability/issues/55)) ([2e756e1](https://github.com/adaptive-enforcement-lab/readability/commit/2e756e139e6ddf4b98a3efd6a7310e6cdabdb85d))
* add MkDocs-style admonition detection and threshold check ([#47](https://github.com/adaptive-enforcement-lab/readability/issues/47)) ([9a41aac](https://github.com/adaptive-enforcement-lab/readability/commit/9a41aac9004438bc3ef3542d0573c33aa5a9ff33))
* add readability improvement hints on check failure ([2d5fc16](https://github.com/adaptive-enforcement-lab/readability/commit/2d5fc169cd12a6b66ed9bc8164b9026e81a55219))
* add release-please and MkDocs Material documentation ([#2](https://github.com/adaptive-enforcement-lab/readability/issues/2)) ([25d37ea](https://github.com/adaptive-enforcement-lab/readability/commit/25d37eadebc6da81c0433b20928e0c68e6053ae9))
* add runtime JSON Schema validation for configuration files ([#198](https://github.com/adaptive-enforcement-lab/readability/issues/198)) ([df643db](https://github.com/adaptive-enforcement-lab/readability/commit/df643dbe7d621f459392b983ff9c6565421a253d))
* add schema references to YAML configuration examples ([#185](https://github.com/adaptive-enforcement-lab/readability/issues/185)) ([6258b6a](https://github.com/adaptive-enforcement-lab/readability/commit/6258b6ac548175d090bb0ad96eed8f5f096ffc9b))
* add SLSA provenance generation to releases ([#127](https://github.com/adaptive-enforcement-lab/readability/issues/127)) ([1f5c92d](https://github.com/adaptive-enforcement-lab/readability/commit/1f5c92d9af7de2f112b09bca5e84949ce77078ed))
* add social cards plugin and fix duplicate nav entry ([7935b61](https://github.com/adaptive-enforcement-lab/readability/commit/7935b617b0538168ec3d718701a49edcb8735c04))
* add Trivy security scanning and SBOM generation ([#80](https://github.com/adaptive-enforcement-lab/readability/issues/80)) ([844464f](https://github.com/adaptive-enforcement-lab/readability/commit/844464f6bd793d51b7610c8ec355596aace6d119))
* add unified container SBOM with Trivy attestation ([#140](https://github.com/adaptive-enforcement-lab/readability/issues/140)) ([e7617b1](https://github.com/adaptive-enforcement-lab/readability/commit/e7617b162d2c2cde80a9c7e964598c5fe631661a))
* add warning to split files instead of removing content ([035ebed](https://github.com/adaptive-enforcement-lab/readability/commit/035ebed63245dd5e2085014a2df07b4aea1e2bb8))
* add workflow to sign existing releases ([#121](https://github.com/adaptive-enforcement-lab/readability/issues/121)) ([449c4a5](https://github.com/adaptive-enforcement-lab/readability/commit/449c4a5597abe8a557d9bcf06fee4878834f2a05))
* complete Components 5 & 6 - testing strategy and documentation ([#199](https://github.com/adaptive-enforcement-lab/readability/issues/199)) ([1074aa8](https://github.com/adaptive-enforcement-lab/readability/commit/1074aa8ce4bc34cbb2859d757be5ff4136415ae5)), closes [#160](https://github.com/adaptive-enforcement-lab/readability/issues/160)
* Detect AI slop via mid-sentence dash patterns ([#162](https://github.com/adaptive-enforcement-lab/readability/issues/162)) ([3dfc09f](https://github.com/adaptive-enforcement-lab/readability/commit/3dfc09f31a6d6592da331609a54bcbd5f659d81b))
* enhance summary table with lines, reading time, and metric links ([#25](https://github.com/adaptive-enforcement-lab/readability/issues/25)) ([cf62405](https://github.com/adaptive-enforcement-lab/readability/commit/cf62405c31cafe4663df46c017c2fca52762eeaa))
* increase coverage threshold to 95% with single source of truth ([#91](https://github.com/adaptive-enforcement-lab/readability/issues/91)) ([0a1266c](https://github.com/adaptive-enforcement-lab/readability/commit/0a1266c76acc1bfc60621811c5e6ebb90154200f))
* initial release of readability analyzer ([b55bca4](https://github.com/adaptive-enforcement-lab/readability/commit/b55bca4797f8f9648fa70f9ca4295fe19eee6cc5))
* publish JSON Schema to MkDocs documentation site ([#182](https://github.com/adaptive-enforcement-lab/readability/issues/182)) ([645f6bc](https://github.com/adaptive-enforcement-lab/readability/commit/645f6bcf736ce75e86103790c6983fd9483c35da))
* publish to GitHub Marketplace ([#70](https://github.com/adaptive-enforcement-lab/readability/issues/70)) ([4577d6d](https://github.com/adaptive-enforcement-lab/readability/commit/4577d6d7c84b65adc0476c41133fbd53b76f2bed))
* use pre-built binary and add pre-commit hook support ([#28](https://github.com/adaptive-enforcement-lab/readability/issues/28)) ([6e7bfc7](https://github.com/adaptive-enforcement-lab/readability/commit/6e7bfc7586f5e120c67f3138101d835826a24e75))


### Bug Fixes

* action config auto-detection and outputs ([#14](https://github.com/adaptive-enforcement-lab/readability/issues/14)) ([f4edc89](https://github.com/adaptive-enforcement-lab/readability/commit/f4edc8968d7771341d1a362634075b807ac8c6bb))
* **action:** resolve bash syntax errors and stdout duplication ([#51](https://github.com/adaptive-enforcement-lab/readability/issues/51)) ([8f26ed6](https://github.com/adaptive-enforcement-lab/readability/commit/8f26ed634971abc83f921b9f0ce9306925b12eff))
* add admonitions to nav and fix anchor links ([#64](https://github.com/adaptive-enforcement-lab/readability/issues/64)) ([e4f8c8f](https://github.com/adaptive-enforcement-lab/readability/commit/e4f8c8fe2615f33b80016b3e194b4fc62b6bcfa0))
* add pillow and cairosvg for social cards in CI ([5cba44e](https://github.com/adaptive-enforcement-lab/readability/commit/5cba44e02abe1b68a5e4a412439f02de8ae9e37a))
* add v prefix to mkdocs version URLs for consistency ([#166](https://github.com/adaptive-enforcement-lab/readability/issues/166)) ([0d40d20](https://github.com/adaptive-enforcement-lab/readability/commit/0d40d2041d0b6ea60dc3bf920672467d44f5fa97))
* add value mappings for composite action outputs ([#20](https://github.com/adaptive-enforcement-lab/readability/issues/20)) ([99cfb35](https://github.com/adaptive-enforcement-lab/readability/commit/99cfb35f4c469dd3d60d0f450dc98ab6477b1c21))
* complete reading time ceiling division fixes ([#59](https://github.com/adaptive-enforcement-lab/readability/issues/59)) ([cfd4cdb](https://github.com/adaptive-enforcement-lab/readability/commit/cfd4cdbb84903596a3f1e29f6bf31cd216a1d0f9))
* configure Codecov with OIDC authentication ([#75](https://github.com/adaptive-enforcement-lab/readability/issues/75)) ([894311e](https://github.com/adaptive-enforcement-lab/readability/commit/894311e7cbfca3ad7bedcd776067e2575f0b40d9))
* correct CLI flags and lint errors ([#1](https://github.com/adaptive-enforcement-lab/readability/issues/1)) ([b23bb82](https://github.com/adaptive-enforcement-lab/readability/commit/b23bb82cf8e2b2051ce7e1ede974539148a76ba7))
* **deps:** update module github.com/yuin/goldmark to v1.7.14 ([#217](https://github.com/adaptive-enforcement-lab/readability/issues/217)) ([4f784a5](https://github.com/adaptive-enforcement-lab/readability/commit/4f784a570b11cfe080205d56314f93060b94283a))
* enable blank issues for free-form issue creation ([#96](https://github.com/adaptive-enforcement-lab/readability/issues/96)) ([e5c50a2](https://github.com/adaptive-enforcement-lab/readability/commit/e5c50a2b02539c99416b4bcb8091f96e1adf1475))
* enable v prefix in release tags for Go module compliance ([#101](https://github.com/adaptive-enforcement-lab/readability/issues/101)) ([e3ab84e](https://github.com/adaptive-enforcement-lab/readability/commit/e3ab84ece59e2708fb847ec7c46a8d1fe46c57b9))
* Exclude frontmatter from prose extraction ([#164](https://github.com/adaptive-enforcement-lab/readability/issues/164)) ([afa3a36](https://github.com/adaptive-enforcement-lab/readability/commit/afa3a361c12e6b633f5ea82bde3b2716f5c05f1c))
* generate schema from Go structs with embedded validation ([#210](https://github.com/adaptive-enforcement-lab/readability/issues/210)) ([7e88510](https://github.com/adaptive-enforcement-lab/readability/commit/7e8851059dfc579b02aa413eaefad1456f5a95b4))
* handle absolute paths in override path matching ([#49](https://github.com/adaptive-enforcement-lab/readability/issues/49)) ([105c5c0](https://github.com/adaptive-enforcement-lab/readability/commit/105c5c01f779b38ec61f329f87e22564fc09eadd))
* improve Go Report Card compliance and pre-commit hooks ([#84](https://github.com/adaptive-enforcement-lab/readability/issues/84)) ([d42947d](https://github.com/adaptive-enforcement-lab/readability/commit/d42947d1ed580c4df356241e5c809170d2ba61ef))
* inline container publishing in release.yml for Scorecard detection ([#150](https://github.com/adaptive-enforcement-lab/readability/issues/150)) ([b676272](https://github.com/adaptive-enforcement-lab/readability/commit/b676272ea785b03ed16f3d1bfabe748fa652bc3e))
* make container binaries executable after artifact download ([#156](https://github.com/adaptive-enforcement-lab/readability/issues/156)) ([1bd5ccd](https://github.com/adaptive-enforcement-lab/readability/commit/1bd5ccd14a15d0d53551880cec0af8d5f7ab0d14))
* make release binaries executable after artifact download ([#158](https://github.com/adaptive-enforcement-lab/readability/issues/158)) ([02fc306](https://github.com/adaptive-enforcement-lab/readability/commit/02fc3062f9a76e4d7cdf60cf69470337d47d1df9))
* move permissions to job level in sign-releases workflow ([#126](https://github.com/adaptive-enforcement-lab/readability/issues/126)) ([2565c22](https://github.com/adaptive-enforcement-lab/readability/commit/2565c22fe4a16340632238eb2274044079bccb0a))
* move workflow permissions to job level for least privilege ([#106](https://github.com/adaptive-enforcement-lab/readability/issues/106)) ([b168c1f](https://github.com/adaptive-enforcement-lab/readability/commit/b168c1ffeed27d4e34a0582f3d6ad32c64925115))
* pass CODECOV_TOKEN to reusable CI workflow ([#88](https://github.com/adaptive-enforcement-lab/readability/issues/88)) ([a98ba23](https://github.com/adaptive-enforcement-lab/readability/commit/a98ba2390f359e3f2bf74f1244bd21497ed4aef3))
* remove duplicate push trigger from ci.yml ([#78](https://github.com/adaptive-enforcement-lab/readability/issues/78)) ([63dd296](https://github.com/adaptive-enforcement-lab/readability/commit/63dd2963bf2cba27af0a68c55733365a6a9c96cc))
* remove duplicate source archive signing from releases ([#146](https://github.com/adaptive-enforcement-lab/readability/issues/146)) ([67b8f11](https://github.com/adaptive-enforcement-lab/readability/commit/67b8f112cbb5ac7ccd0a11f941fc7e27dfcd05c4))
* remove homepage override so logo links to docs site ([#62](https://github.com/adaptive-enforcement-lab/readability/issues/62)) ([772654e](https://github.com/adaptive-enforcement-lab/readability/commit/772654ee11a8da18df2af71aec770662e9f35b7c))
* rename action for GitHub Marketplace uniqueness ([f3df887](https://github.com/adaptive-enforcement-lab/readability/commit/f3df887f1aa10378ce2841647d3d9cb1fe8496d7))
* shorten action description for Marketplace limit ([e18fd8d](https://github.com/adaptive-enforcement-lab/readability/commit/e18fd8d0e38979b57b8ed3dcfd574ec65029cad2))
* sign source code archives in releases ([#122](https://github.com/adaptive-enforcement-lab/readability/issues/122)) ([c86288e](https://github.com/adaptive-enforcement-lab/readability/commit/c86288eb2a91f7a10ecc1a736010a80f9f352d1b))
* test action builds from source instead of downloading release ([#221](https://github.com/adaptive-enforcement-lab/readability/issues/221)) ([25dc0b7](https://github.com/adaptive-enforcement-lab/readability/commit/25dc0b7c5ee79207371a968fdd302b5bebb5ab6c))
* trigger Scorecard after Release workflow completes ([#134](https://github.com/adaptive-enforcement-lab/readability/issues/134)) ([5450628](https://github.com/adaptive-enforcement-lab/readability/commit/5450628b523fab274dcc2acdec3eee69e2986da1))
* update CI badge to point to ci.yml ([#82](https://github.com/adaptive-enforcement-lab/readability/issues/82)) ([9a64c38](https://github.com/adaptive-enforcement-lab/readability/commit/9a64c3819fcafad4707e9e0bbf6047e2afc172a6))
* update component paths to use regex patterns ([#90](https://github.com/adaptive-enforcement-lab/readability/issues/90)) ([0931deb](https://github.com/adaptive-enforcement-lab/readability/commit/0931deb216eb31a09c4f2277319e53d6acdb775c))
* update cosign signing for v3 API ([#117](https://github.com/adaptive-enforcement-lab/readability/issues/117)) ([ad12a83](https://github.com/adaptive-enforcement-lab/readability/commit/ad12a83406b0e9c23738063e715d657afe14bd12))
* update schema URL to work with mike versioning ([#183](https://github.com/adaptive-enforcement-lab/readability/issues/183)) ([c14271c](https://github.com/adaptive-enforcement-lab/readability/commit/c14271cfa816eec0c0cf780c599461228830e6d3)), closes [#160](https://github.com/adaptive-enforcement-lab/readability/issues/160)
* use .sig extension for cosign signatures ([#119](https://github.com/adaptive-enforcement-lab/readability/issues/119)) ([135e617](https://github.com/adaptive-enforcement-lab/readability/commit/135e6177d11847c01ca1ae61078be01c6b1dff0a))
* use ceiling division for reading time calculation ([#57](https://github.com/adaptive-enforcement-lab/readability/issues/57)) ([2e8b644](https://github.com/adaptive-enforcement-lab/readability/commit/2e8b644262e95d519b6e557ddb5fedb78c97890a))
* use version tag for scorecard-action (cannot pin to SHA) ([#111](https://github.com/adaptive-enforcement-lab/readability/issues/111)) ([d07300c](https://github.com/adaptive-enforcement-lab/readability/commit/d07300cf0a2df11e379dec9d711c0675a2e1f599))
* use version tags for all actions in scorecard.yml ([#112](https://github.com/adaptive-enforcement-lab/readability/issues/112)) ([39ce338](https://github.com/adaptive-enforcement-lab/readability/commit/39ce338d46de276ba453e92cbfc1bddd51f750c4))


### Code Refactoring

* reduce cyclomatic complexity and enforce strict gocyclo ([#85](https://github.com/adaptive-enforcement-lab/readability/issues/85)) ([bed6cae](https://github.com/adaptive-enforcement-lab/readability/commit/bed6caea4524a04b3de3fabc3932377acaa8f81d))
* reorganize docs navigation - make Configuration a top-level section ([#203](https://github.com/adaptive-enforcement-lab/readability/issues/203)) ([b25af08](https://github.com/adaptive-enforcement-lab/readability/commit/b25af089b85bee62c1db7f5acb333ebc24366d8b))
* run release-please before CI in release workflow ([#219](https://github.com/adaptive-enforcement-lab/readability/issues/219)) ([7e64800](https://github.com/adaptive-enforcement-lab/readability/commit/7e64800019df79ad5700895a0c94181b1eb24ee2))
* simplify SLSA provenance workflow ([#130](https://github.com/adaptive-enforcement-lab/readability/issues/130)) ([0a03624](https://github.com/adaptive-enforcement-lab/readability/commit/0a03624d444e726b7105214f1d7dc083cac54d1f))
* unify CI/CD with reusable workflow pattern ([#76](https://github.com/adaptive-enforcement-lab/readability/issues/76)) ([df8511a](https://github.com/adaptive-enforcement-lab/readability/commit/df8511a1a028a9d3cae6277a88582ba02f466011))


### Documentation

* comprehensive documentation overhaul for newcomer accessibility ([#65](https://github.com/adaptive-enforcement-lab/readability/issues/65)) ([9d7e446](https://github.com/adaptive-enforcement-lab/readability/commit/9d7e446054dd30cfeb4dbbaf3024f4ea45b52faf))


### Maintenance

* add MIT license ([098e115](https://github.com/adaptive-enforcement-lab/readability/commit/098e115fa2ec30c2cfd24df09cb7d6decafeb757))
* add site/ to gitignore ([37bc3af](https://github.com/adaptive-enforcement-lab/readability/commit/37bc3afa4e3045370a17163a46c50a3d5b80ba00))
* **deps:** update actions/checkout action to v6 ([#40](https://github.com/adaptive-enforcement-lab/readability/issues/40)) ([edc40d1](https://github.com/adaptive-enforcement-lab/readability/commit/edc40d1cb49ac607c2b8b8098a6cf09bff617ee6))
* **deps:** update actions/download-artifact action to v7 ([#152](https://github.com/adaptive-enforcement-lab/readability/issues/152)) ([f43b7a6](https://github.com/adaptive-enforcement-lab/readability/commit/f43b7a66adb6b55019ca64350981a73d9bd8af37))
* **deps:** update actions/download-artifact digest to d3f86a1 ([#143](https://github.com/adaptive-enforcement-lab/readability/issues/143)) ([6f38f0d](https://github.com/adaptive-enforcement-lab/readability/commit/6f38f0d49e6b68cc4d79bef9b48547352d0ff041))
* **deps:** update actions/setup-go action to v6 ([#41](https://github.com/adaptive-enforcement-lab/readability/issues/41)) ([9276f52](https://github.com/adaptive-enforcement-lab/readability/commit/9276f52e9e0ea05ef70f569eaad4c3993667d1fd))
* **deps:** update actions/setup-python action to v6 ([#204](https://github.com/adaptive-enforcement-lab/readability/issues/204)) ([4965742](https://github.com/adaptive-enforcement-lab/readability/commit/496574216efe4066f63c9f429732d95e479cd3ef))
* **deps:** update actions/setup-python action to v6 ([#42](https://github.com/adaptive-enforcement-lab/readability/issues/42)) ([cc545e4](https://github.com/adaptive-enforcement-lab/readability/commit/cc545e405a12b85a71a715e28b4871ac580c09b0))
* **deps:** update actions/setup-python digest to a26af69 ([#200](https://github.com/adaptive-enforcement-lab/readability/issues/200)) ([8b77d6a](https://github.com/adaptive-enforcement-lab/readability/commit/8b77d6adf18f30e82222257152839ee6a66f073b))
* **deps:** update codecov/test-results-action digest to 0fa95f0 ([#172](https://github.com/adaptive-enforcement-lab/readability/issues/172)) ([e3cd7f2](https://github.com/adaptive-enforcement-lab/readability/commit/e3cd7f264ee66e94cbea97e780ca0439541dce01))
* **deps:** update dependency go to 1.25 ([#36](https://github.com/adaptive-enforcement-lab/readability/issues/36)) ([9754456](https://github.com/adaptive-enforcement-lab/readability/commit/9754456098823a64532982cd8ed97c788a21b358))
* **deps:** update dependency mkdocs-material to v9.7.1 ([#206](https://github.com/adaptive-enforcement-lab/readability/issues/206)) ([0f1aea5](https://github.com/adaptive-enforcement-lab/readability/commit/0f1aea5e8fda5724b6e0a920d1168e65cbae0cbc))
* **deps:** update dependency pillow to v12.1.0 ([#214](https://github.com/adaptive-enforcement-lab/readability/issues/214)) ([2957e52](https://github.com/adaptive-enforcement-lab/readability/commit/2957e52dd75fb50303ee86402f09c10bb367898c))
* **deps:** update dependency python to 3.14 ([#201](https://github.com/adaptive-enforcement-lab/readability/issues/201)) ([cd7eeda](https://github.com/adaptive-enforcement-lab/readability/commit/cd7eeda72c10ee1f9abc0f6166dd941e7e481818))
* **deps:** update dependency python to 3.14 ([#37](https://github.com/adaptive-enforcement-lab/readability/issues/37)) ([6c8abbe](https://github.com/adaptive-enforcement-lab/readability/commit/6c8abbe768e219f0797b36e65b8daefe75a5eb9b))
* **deps:** update docker/login-action digest to 5e57cd1 ([#137](https://github.com/adaptive-enforcement-lab/readability/issues/137)) ([8a00393](https://github.com/adaptive-enforcement-lab/readability/commit/8a003937b31114f2b29fb83f3cb4aae8f5e5425d))
* **deps:** update docker/metadata-action digest to c299e40 ([#138](https://github.com/adaptive-enforcement-lab/readability/issues/138)) ([4d210e6](https://github.com/adaptive-enforcement-lab/readability/commit/4d210e6b1cac55e6ea126dabceac2ebfb3f67259))
* **deps:** update docker/setup-buildx-action digest to 8d2750c ([#207](https://github.com/adaptive-enforcement-lab/readability/issues/207)) ([943e8b4](https://github.com/adaptive-enforcement-lab/readability/commit/943e8b440ad8ed22fb440a237cae3b79e247331e))
* **deps:** update docker/setup-buildx-action digest to e468171 ([#144](https://github.com/adaptive-enforcement-lab/readability/issues/144)) ([25cc3ba](https://github.com/adaptive-enforcement-lab/readability/commit/25cc3bafb640050e088a0c4137b0a2d8126e0a40))
* **deps:** update docker/setup-qemu-action digest to c7c5346 ([#151](https://github.com/adaptive-enforcement-lab/readability/issues/151)) ([f491d10](https://github.com/adaptive-enforcement-lab/readability/commit/f491d109567899355006cafe8b68ccdf0e5b7970))
* **deps:** update github artifact actions ([#43](https://github.com/adaptive-enforcement-lab/readability/issues/43)) ([f498a0e](https://github.com/adaptive-enforcement-lab/readability/commit/f498a0e3e4a82cc87187cfa6af0a43efae648f84))
* **deps:** update github artifact actions ([#71](https://github.com/adaptive-enforcement-lab/readability/issues/71)) ([2ae0a49](https://github.com/adaptive-enforcement-lab/readability/commit/2ae0a492c33972c75b21efefe7256b9d4ff66a1d))
* **deps:** update github artifact actions ([#81](https://github.com/adaptive-enforcement-lab/readability/issues/81)) ([4cde10f](https://github.com/adaptive-enforcement-lab/readability/commit/4cde10fcc28839113c04fd378e60e36dcecdd189))
* **deps:** update github/codeql-action action to v4 ([#104](https://github.com/adaptive-enforcement-lab/readability/issues/104)) ([8c0e6e3](https://github.com/adaptive-enforcement-lab/readability/commit/8c0e6e3864ce26cdfb3bc041c619a8049b3b505e))
* **deps:** update github/codeql-action digest to 1b168cd ([#109](https://github.com/adaptive-enforcement-lab/readability/issues/109)) ([7dec92d](https://github.com/adaptive-enforcement-lab/readability/commit/7dec92d951ef565391564f5ea6ef8b9c719ced6a))
* **deps:** update github/codeql-action digest to 5d4e8d1 ([#176](https://github.com/adaptive-enforcement-lab/readability/issues/176)) ([466842e](https://github.com/adaptive-enforcement-lab/readability/commit/466842e34959e98be0657446ca393c1b5c12691f))
* **deps:** update golangci/golangci-lint-action action to v9 ([#44](https://github.com/adaptive-enforcement-lab/readability/issues/44)) ([4a57751](https://github.com/adaptive-enforcement-lab/readability/commit/4a57751e0734850aff04e99462c1260d5fd61426))
* **deps:** update googleapis/release-please-action digest to 16a9c90 ([#110](https://github.com/adaptive-enforcement-lab/readability/issues/110)) ([1faf2e9](https://github.com/adaptive-enforcement-lab/readability/commit/1faf2e92bb72ccda6beb7171d1f03d9d63b2e33b))
* **deps:** update ossf/scorecard-action action to v2.4.3 ([#103](https://github.com/adaptive-enforcement-lab/readability/issues/103)) ([6d3834a](https://github.com/adaptive-enforcement-lab/readability/commit/6d3834a04bb8d4ad3a25e7789fbf79ec0191b0a0))
* **deps:** update sigstore/cosign-installer action to v3.10.1 ([#123](https://github.com/adaptive-enforcement-lab/readability/issues/123)) ([bdba92e](https://github.com/adaptive-enforcement-lab/readability/commit/bdba92eaaaa8a780abd0c94a7a19735ef73052d4))
* **deps:** update sigstore/cosign-installer action to v4 ([#124](https://github.com/adaptive-enforcement-lab/readability/issues/124)) ([b264d48](https://github.com/adaptive-enforcement-lab/readability/commit/b264d4829f05726052bb9266568f35b5869a6788))
* **deps:** update tj-actions/changed-files action to v47 ([#45](https://github.com/adaptive-enforcement-lab/readability/issues/45)) ([4333c5e](https://github.com/adaptive-enforcement-lab/readability/commit/4333c5eb13f5c6b39a416daed40af2f55c08c417))
* **deps:** update trufflesecurity/trufflehog digest to 8aea6cd ([#169](https://github.com/adaptive-enforcement-lab/readability/issues/169)) ([14653a6](https://github.com/adaptive-enforcement-lab/readability/commit/14653a6c80503a19e838c4aaf2b1383f28502307))
* **deps:** update trufflesecurity/trufflehog digest to 8c1219a ([#177](https://github.com/adaptive-enforcement-lab/readability/issues/177)) ([b731949](https://github.com/adaptive-enforcement-lab/readability/commit/b731949172072c82e236f5576934ca7b565674f0))
* **deps:** update trufflesecurity/trufflehog digest to a633174 ([#212](https://github.com/adaptive-enforcement-lab/readability/issues/212)) ([0d6633e](https://github.com/adaptive-enforcement-lab/readability/commit/0d6633ea9af7f9899b9dfccdad919a51bc94ed76))
* **deps:** update trufflesecurity/trufflehog digest to ef6e76c ([#208](https://github.com/adaptive-enforcement-lab/readability/issues/208)) ([bf6c56c](https://github.com/adaptive-enforcement-lab/readability/commit/bf6c56cb3d68bb00b6c4960535e2e26f486015d9))
* ignore .cache directory ([430480c](https://github.com/adaptive-enforcement-lab/readability/commit/430480c956a0fbd644976e6be6c58f5a8909c728))
* improve OpenSSF Scorecard score ([#114](https://github.com/adaptive-enforcement-lab/readability/issues/114)) ([011a4d7](https://github.com/adaptive-enforcement-lab/readability/commit/011a4d73432583357c5fc58aec3436e9501f800e))
* **main:** release 0.10.0 ([#56](https://github.com/adaptive-enforcement-lab/readability/issues/56)) ([1ec22c8](https://github.com/adaptive-enforcement-lab/readability/commit/1ec22c8b92af2ea45bb43d1fdcaeb54636c31094))
* **main:** release 0.10.1 ([#58](https://github.com/adaptive-enforcement-lab/readability/issues/58)) ([c2f938c](https://github.com/adaptive-enforcement-lab/readability/commit/c2f938c4f56f3acfb86cdd2e5caafad81bf3ba3d))
* **main:** release 0.11.0 ([#61](https://github.com/adaptive-enforcement-lab/readability/issues/61)) ([e8afe02](https://github.com/adaptive-enforcement-lab/readability/commit/e8afe028d241dd7e7f14c50cb9e2e5dc34bb8c53))
* **main:** release 0.11.1 ([#63](https://github.com/adaptive-enforcement-lab/readability/issues/63)) ([1f5d2c0](https://github.com/adaptive-enforcement-lab/readability/commit/1f5d2c0bbf12b5d81a0428ffa96aec3566b451f8))
* **main:** release 0.2.0 ([#3](https://github.com/adaptive-enforcement-lab/readability/issues/3)) ([54f247d](https://github.com/adaptive-enforcement-lab/readability/commit/54f247d69e7051d1b987c3d78cb5c227326b8b3e))
* **main:** release 0.3.0 ([#15](https://github.com/adaptive-enforcement-lab/readability/issues/15)) ([094dfa6](https://github.com/adaptive-enforcement-lab/readability/commit/094dfa647cadc896dba9bb442b1ef83f8c08f45e))
* **main:** release 0.3.1 ([#19](https://github.com/adaptive-enforcement-lab/readability/issues/19)) ([0f9f618](https://github.com/adaptive-enforcement-lab/readability/commit/0f9f6185634054fd3246eaa3b8b24c9e263bef69))
* **main:** release 0.3.2 ([#21](https://github.com/adaptive-enforcement-lab/readability/issues/21)) ([40b069c](https://github.com/adaptive-enforcement-lab/readability/commit/40b069c0c8bef42917c3ccc7966365847124baa2))
* **main:** release 0.4.0 ([#24](https://github.com/adaptive-enforcement-lab/readability/issues/24)) ([4ccf918](https://github.com/adaptive-enforcement-lab/readability/commit/4ccf91801bc04c807315f9ff558ad699d37736a7))
* **main:** release 0.5.0 ([#26](https://github.com/adaptive-enforcement-lab/readability/issues/26)) ([eef2669](https://github.com/adaptive-enforcement-lab/readability/commit/eef26692410ee2f5317bc54fad24c5de02add5a5))
* **main:** release 0.6.0 ([#29](https://github.com/adaptive-enforcement-lab/readability/issues/29)) ([8c41504](https://github.com/adaptive-enforcement-lab/readability/commit/8c415040d23dc3e5a13397e8cf0f0c734c2176c9))
* **main:** release 0.7.0 ([#32](https://github.com/adaptive-enforcement-lab/readability/issues/32)) ([c3b5a4c](https://github.com/adaptive-enforcement-lab/readability/commit/c3b5a4cca1430a620762830667f89698d1e83082))
* **main:** release 0.7.1 ([#33](https://github.com/adaptive-enforcement-lab/readability/issues/33)) ([e0c9193](https://github.com/adaptive-enforcement-lab/readability/commit/e0c9193ce2609acfb5a001ff52bb16d59f5f0296))
* **main:** release 0.7.2 ([#39](https://github.com/adaptive-enforcement-lab/readability/issues/39)) ([64ffe3a](https://github.com/adaptive-enforcement-lab/readability/commit/64ffe3a9bc5bae7200a81a1c053678954cd0b478))
* **main:** release 0.8.0 ([#46](https://github.com/adaptive-enforcement-lab/readability/issues/46)) ([92858c8](https://github.com/adaptive-enforcement-lab/readability/commit/92858c8049ac34990b37e31e387e1091723b170c))
* **main:** release 0.9.0 ([#48](https://github.com/adaptive-enforcement-lab/readability/issues/48)) ([e4d02a6](https://github.com/adaptive-enforcement-lab/readability/commit/e4d02a69d40ea673fd5efc0fb6e8c80ef8758860))
* **main:** release 0.9.1 ([#50](https://github.com/adaptive-enforcement-lab/readability/issues/50)) ([5ff04ab](https://github.com/adaptive-enforcement-lab/readability/commit/5ff04ab77e0043c5f398f8c797cd1043cc450266))
* **main:** release 0.9.2 ([#52](https://github.com/adaptive-enforcement-lab/readability/issues/52)) ([7578bcb](https://github.com/adaptive-enforcement-lab/readability/commit/7578bcbee92d2c4c6cd87fadf9d1af936bd2b14a))
* **main:** release 1.0.0 ([#66](https://github.com/adaptive-enforcement-lab/readability/issues/66)) ([a76a528](https://github.com/adaptive-enforcement-lab/readability/commit/a76a5282f322b4a98b0c5406731a6779c76e0aa3))
* **main:** release 1.1.0 ([#67](https://github.com/adaptive-enforcement-lab/readability/issues/67)) ([adebc50](https://github.com/adaptive-enforcement-lab/readability/commit/adebc507309b675d0d78bf2838dd13046a2a03a0))
* **main:** release 1.1.1 ([#68](https://github.com/adaptive-enforcement-lab/readability/issues/68)) ([4a82009](https://github.com/adaptive-enforcement-lab/readability/commit/4a82009e8c23ac810a2151a23fbdb8f6d85e509a))
* **main:** release 1.10.0 ([#141](https://github.com/adaptive-enforcement-lab/readability/issues/141)) ([08e620e](https://github.com/adaptive-enforcement-lab/readability/commit/08e620e0ceae989e7aa46cf508622b68a8cb0f18))
* **main:** release 1.10.1 ([#147](https://github.com/adaptive-enforcement-lab/readability/issues/147)) ([22cbc2c](https://github.com/adaptive-enforcement-lab/readability/commit/22cbc2c41694b88a57f3da1c2be5e89488b83a54))
* **main:** release 1.10.2 ([#153](https://github.com/adaptive-enforcement-lab/readability/issues/153)) ([95204a5](https://github.com/adaptive-enforcement-lab/readability/commit/95204a5cd3615ce95fae7453a6c0398ff0671a99))
* **main:** release 1.10.3 ([#155](https://github.com/adaptive-enforcement-lab/readability/issues/155)) ([7ab52a4](https://github.com/adaptive-enforcement-lab/readability/commit/7ab52a40a3b76cb409dd355ee11de31b65dd9d91))
* **main:** release 1.10.4 ([#157](https://github.com/adaptive-enforcement-lab/readability/issues/157)) ([e8eff00](https://github.com/adaptive-enforcement-lab/readability/commit/e8eff007ec7a27a5a5fdacd8daf00c189c0036a5))
* **main:** release 1.10.5 ([#159](https://github.com/adaptive-enforcement-lab/readability/issues/159)) ([3847ce3](https://github.com/adaptive-enforcement-lab/readability/commit/3847ce3d04e86d12c318aed414f484a8f3f307ad))
* **main:** release 1.11.0 ([#163](https://github.com/adaptive-enforcement-lab/readability/issues/163)) ([2f0d1a0](https://github.com/adaptive-enforcement-lab/readability/commit/2f0d1a01845897e3fb91329db8931a77f4abb2f2))
* **main:** release 1.11.1 ([#165](https://github.com/adaptive-enforcement-lab/readability/issues/165)) ([8d17033](https://github.com/adaptive-enforcement-lab/readability/commit/8d17033784809c0a763f2f6c5ebce665fb1d5d63))
* **main:** release 1.11.2 ([#167](https://github.com/adaptive-enforcement-lab/readability/issues/167)) ([1fcccff](https://github.com/adaptive-enforcement-lab/readability/commit/1fcccff39f52c340dcee3e7b2272da7e5a9d18f4))
* **main:** release 1.11.3 ([#173](https://github.com/adaptive-enforcement-lab/readability/issues/173)) ([a4f9344](https://github.com/adaptive-enforcement-lab/readability/commit/a4f9344a495df8010ea3b399d04daa18aa355a31))
* **main:** release 1.11.4 ([#178](https://github.com/adaptive-enforcement-lab/readability/issues/178)) ([f8352b8](https://github.com/adaptive-enforcement-lab/readability/commit/f8352b8f474833e7480841bb8a015a9d249d0c5e))
* **main:** release 1.12.0 ([#181](https://github.com/adaptive-enforcement-lab/readability/issues/181)) ([7a27eba](https://github.com/adaptive-enforcement-lab/readability/commit/7a27ebaa12c984f3cb9cfef26e1424a41d5c6e3c))
* **main:** release 1.12.1 ([#184](https://github.com/adaptive-enforcement-lab/readability/issues/184)) ([3487085](https://github.com/adaptive-enforcement-lab/readability/commit/34870854f592ca99616b46487ed63534e4fc2af9))
* **main:** release 1.13.0 ([#187](https://github.com/adaptive-enforcement-lab/readability/issues/187)) ([5d1653a](https://github.com/adaptive-enforcement-lab/readability/commit/5d1653ae972b1318fed0ff227537be7bcfec0b32))
* **main:** release 1.14.0 ([#202](https://github.com/adaptive-enforcement-lab/readability/issues/202)) ([de00c0e](https://github.com/adaptive-enforcement-lab/readability/commit/de00c0e1a573f54afbc767080f9fbee388f21e30))
* **main:** release 1.14.1 ([#205](https://github.com/adaptive-enforcement-lab/readability/issues/205)) ([bb36627](https://github.com/adaptive-enforcement-lab/readability/commit/bb366278f009b62834b09d9a9dd22f553f429b4d))
* **main:** release 1.14.2 ([#211](https://github.com/adaptive-enforcement-lab/readability/issues/211)) ([2c942fe](https://github.com/adaptive-enforcement-lab/readability/commit/2c942fe21cc0c6d612a7cf85483a8359f7182c6a))
* **main:** release 1.14.3 ([#213](https://github.com/adaptive-enforcement-lab/readability/issues/213)) ([07e37e9](https://github.com/adaptive-enforcement-lab/readability/commit/07e37e9cd183f0af8ace41bb62c91fbd159219f4))
* **main:** release 1.2.0 ([#69](https://github.com/adaptive-enforcement-lab/readability/issues/69)) ([1bc239f](https://github.com/adaptive-enforcement-lab/readability/commit/1bc239f94085c170779c11efd3a928524921bc5d))
* **main:** release 1.2.1 ([#72](https://github.com/adaptive-enforcement-lab/readability/issues/72)) ([39f6a90](https://github.com/adaptive-enforcement-lab/readability/commit/39f6a90cd0e817f6383f68ad2f6ed51dd7f3ebda))
* **main:** release 1.2.2 ([#77](https://github.com/adaptive-enforcement-lab/readability/issues/77)) ([823797a](https://github.com/adaptive-enforcement-lab/readability/commit/823797a60e5aea7f0fcdbb384305857ba6f8b992))
* **main:** release 1.3.0 ([#79](https://github.com/adaptive-enforcement-lab/readability/issues/79)) ([66fb203](https://github.com/adaptive-enforcement-lab/readability/commit/66fb20367697233dab3ecdb129b1d38f6c905f73))
* **main:** release 1.3.1 ([#83](https://github.com/adaptive-enforcement-lab/readability/issues/83)) ([150bae0](https://github.com/adaptive-enforcement-lab/readability/commit/150bae0f0413b53fe0af62b65dc3212ea5536fde))
* **main:** release 1.4.0 ([#87](https://github.com/adaptive-enforcement-lab/readability/issues/87)) ([da9f594](https://github.com/adaptive-enforcement-lab/readability/commit/da9f5949ea66b4defdf5d11090127e8fbb76fcb6))
* **main:** release 1.5.0 ([#89](https://github.com/adaptive-enforcement-lab/readability/issues/89)) ([530ed75](https://github.com/adaptive-enforcement-lab/readability/commit/530ed75b7f8a2dda8cfb9f2750eed2729a270cbf))
* **main:** release 1.5.1 ([#97](https://github.com/adaptive-enforcement-lab/readability/issues/97)) ([20b6020](https://github.com/adaptive-enforcement-lab/readability/commit/20b6020ba13ab57bd4e8008e971df5111207ca9f))
* **main:** release 1.6.0 ([#107](https://github.com/adaptive-enforcement-lab/readability/issues/107)) ([d0473c7](https://github.com/adaptive-enforcement-lab/readability/commit/d0473c7a422736e65f415e2f70cab22f41f4cee1))
* **main:** release 1.6.1 ([#118](https://github.com/adaptive-enforcement-lab/readability/issues/118)) ([6a1231c](https://github.com/adaptive-enforcement-lab/readability/commit/6a1231c7372fa459594241cbb4eb6e61c3340d70))
* **main:** release 1.6.2 ([#120](https://github.com/adaptive-enforcement-lab/readability/issues/120)) ([ebf29af](https://github.com/adaptive-enforcement-lab/readability/commit/ebf29af5dda23bc581ee0acc5c362b78c3284a06))
* **main:** release 1.7.0 ([#125](https://github.com/adaptive-enforcement-lab/readability/issues/125)) ([88ab47b](https://github.com/adaptive-enforcement-lab/readability/commit/88ab47bb2b2d0506bdb4dd548a4525544bd453c8))
* **main:** release 1.7.1 ([#131](https://github.com/adaptive-enforcement-lab/readability/issues/131)) ([15dab4a](https://github.com/adaptive-enforcement-lab/readability/commit/15dab4a45dd82c7c5eb28e2f89a83ac1794e97b9))
* **main:** release 1.8.0 ([#133](https://github.com/adaptive-enforcement-lab/readability/issues/133)) ([f450727](https://github.com/adaptive-enforcement-lab/readability/commit/f450727407a23a280351b476d8673665af07356c))
* **main:** release 1.8.1 ([#135](https://github.com/adaptive-enforcement-lab/readability/issues/135)) ([ff14e12](https://github.com/adaptive-enforcement-lab/readability/commit/ff14e126e47821f720ac2895b63461c56261620f))
* **main:** release 1.9.0 ([#139](https://github.com/adaptive-enforcement-lab/readability/issues/139)) ([cc4796e](https://github.com/adaptive-enforcement-lab/readability/commit/cc4796eea4af5bd79a9d38790e9ef6ca15ceec9a))
* **main:** release 2.0.0 ([#220](https://github.com/adaptive-enforcement-lab/readability/issues/220)) ([2f1e7de](https://github.com/adaptive-enforcement-lab/readability/commit/2f1e7ded1fcb54f287e68653433f144dfd3061a7))
* pin all GitHub Actions to commit SHAs ([#108](https://github.com/adaptive-enforcement-lab/readability/issues/108)) ([b9840c7](https://github.com/adaptive-enforcement-lab/readability/commit/b9840c71f8adcc20995624f541859f4f0271cb08))
* remove deprecated publish-container workflow ([#154](https://github.com/adaptive-enforcement-lab/readability/issues/154)) ([923ea88](https://github.com/adaptive-enforcement-lab/readability/commit/923ea8894c65fbca56229384a090b23dbdb63186))
* remove sign-releases workflow ([#129](https://github.com/adaptive-enforcement-lab/readability/issues/129)) ([e18457f](https://github.com/adaptive-enforcement-lab/readability/commit/e18457f7a29a304674f4d8661367c3d7663cdf6c))
* remove v prefix from release tags ([#18](https://github.com/adaptive-enforcement-lab/readability/issues/18)) ([b431939](https://github.com/adaptive-enforcement-lab/readability/commit/b43193980346a87efa000977bc94523ac63cd9a4))
* rename CI to Build and run on PRs only ([#17](https://github.com/adaptive-enforcement-lab/readability/issues/17)) ([8f9c13e](https://github.com/adaptive-enforcement-lab/readability/commit/8f9c13e485336f56c24d320cd75ea3fed35532b5))

## [2.0.0](https://github.com/adaptive-enforcement-lab/readability/compare/v1.14.3...v2.0.0) (2026-01-04)


### ⚠ BREAKING CHANGES

* Removed hardcoded CHANGELOG.md and CONTRIBUTING.md skip logic. These files must now be explicitly excluded in .readability.yml or they will be analyzed and may fail threshold checks. Add the following to your config to preserve the old behavior:

### Features

* add exclude field to path overrides ([#216](https://github.com/adaptive-enforcement-lab/readability/issues/216)) ([#218](https://github.com/adaptive-enforcement-lab/readability/issues/218)) ([e913d3b](https://github.com/adaptive-enforcement-lab/readability/commit/e913d3beae7d880a8f5be0147ffc5fca84b0730b))


### Code Refactoring

* run release-please before CI in release workflow ([#219](https://github.com/adaptive-enforcement-lab/readability/issues/219)) ([7e64800](https://github.com/adaptive-enforcement-lab/readability/commit/7e64800019df79ad5700895a0c94181b1eb24ee2))

## [1.14.3](https://github.com/adaptive-enforcement-lab/readability/compare/v1.14.2...v1.14.3) (2026-01-03)


### Maintenance

* **deps:** update dependency pillow to v12.1.0 ([#214](https://github.com/adaptive-enforcement-lab/readability/issues/214)) ([2957e52](https://github.com/adaptive-enforcement-lab/readability/commit/2957e52dd75fb50303ee86402f09c10bb367898c))
* **deps:** update trufflesecurity/trufflehog digest to a633174 ([#212](https://github.com/adaptive-enforcement-lab/readability/issues/212)) ([0d6633e](https://github.com/adaptive-enforcement-lab/readability/commit/0d6633ea9af7f9899b9dfccdad919a51bc94ed76))

## [1.14.2](https://github.com/adaptive-enforcement-lab/readability/compare/v1.14.1...v1.14.2) (2025-12-19)


### Bug Fixes

* generate schema from Go structs with embedded validation ([#210](https://github.com/adaptive-enforcement-lab/readability/issues/210)) ([7e88510](https://github.com/adaptive-enforcement-lab/readability/commit/7e8851059dfc579b02aa413eaefad1456f5a95b4))


### Maintenance

* **deps:** update dependency mkdocs-material to v9.7.1 ([#206](https://github.com/adaptive-enforcement-lab/readability/issues/206)) ([0f1aea5](https://github.com/adaptive-enforcement-lab/readability/commit/0f1aea5e8fda5724b6e0a920d1168e65cbae0cbc))
* **deps:** update docker/setup-buildx-action digest to 8d2750c ([#207](https://github.com/adaptive-enforcement-lab/readability/issues/207)) ([943e8b4](https://github.com/adaptive-enforcement-lab/readability/commit/943e8b440ad8ed22fb440a237cae3b79e247331e))
* **deps:** update trufflesecurity/trufflehog digest to ef6e76c ([#208](https://github.com/adaptive-enforcement-lab/readability/issues/208)) ([bf6c56c](https://github.com/adaptive-enforcement-lab/readability/commit/bf6c56cb3d68bb00b6c4960535e2e26f486015d9))

## [1.14.1](https://github.com/adaptive-enforcement-lab/readability/compare/v1.14.0...v1.14.1) (2025-12-18)


### Code Refactoring

* reorganize docs navigation - make Configuration a top-level section ([#203](https://github.com/adaptive-enforcement-lab/readability/issues/203)) ([b25af08](https://github.com/adaptive-enforcement-lab/readability/commit/b25af089b85bee62c1db7f5acb333ebc24366d8b))


### Maintenance

* **deps:** update actions/setup-python action to v6 ([#204](https://github.com/adaptive-enforcement-lab/readability/issues/204)) ([4965742](https://github.com/adaptive-enforcement-lab/readability/commit/496574216efe4066f63c9f429732d95e479cd3ef))

## [1.14.0](https://github.com/adaptive-enforcement-lab/readability/compare/v1.13.0...v1.14.0) (2025-12-18)


### Features

* complete Components 5 & 6 - testing strategy and documentation ([#199](https://github.com/adaptive-enforcement-lab/readability/issues/199)) ([1074aa8](https://github.com/adaptive-enforcement-lab/readability/commit/1074aa8ce4bc34cbb2859d757be5ff4136415ae5)), closes [#160](https://github.com/adaptive-enforcement-lab/readability/issues/160)


### Maintenance

* **deps:** update actions/setup-python digest to a26af69 ([#200](https://github.com/adaptive-enforcement-lab/readability/issues/200)) ([8b77d6a](https://github.com/adaptive-enforcement-lab/readability/commit/8b77d6adf18f30e82222257152839ee6a66f073b))
* **deps:** update dependency python to 3.14 ([#201](https://github.com/adaptive-enforcement-lab/readability/issues/201)) ([cd7eeda](https://github.com/adaptive-enforcement-lab/readability/commit/cd7eeda72c10ee1f9abc0f6166dd941e7e481818))

## [1.13.0](https://github.com/adaptive-enforcement-lab/readability/compare/v1.12.1...v1.13.0) (2025-12-18)


### Features

* add runtime JSON Schema validation for configuration files ([#198](https://github.com/adaptive-enforcement-lab/readability/issues/198)) ([df643db](https://github.com/adaptive-enforcement-lab/readability/commit/df643dbe7d621f459392b983ff9c6565421a253d))
* add schema references to YAML configuration examples ([#185](https://github.com/adaptive-enforcement-lab/readability/issues/185)) ([6258b6a](https://github.com/adaptive-enforcement-lab/readability/commit/6258b6ac548175d090bb0ad96eed8f5f096ffc9b))

## [1.12.1](https://github.com/adaptive-enforcement-lab/readability/compare/v1.12.0...v1.12.1) (2025-12-17)


### Bug Fixes

* update schema URL to work with mike versioning ([#183](https://github.com/adaptive-enforcement-lab/readability/issues/183)) ([c14271c](https://github.com/adaptive-enforcement-lab/readability/commit/c14271cfa816eec0c0cf780c599461228830e6d3)), closes [#160](https://github.com/adaptive-enforcement-lab/readability/issues/160)

## [1.12.0](https://github.com/adaptive-enforcement-lab/readability/compare/v1.11.4...v1.12.0) (2025-12-17)


### Features

* add JSON Schema for .readability.yml configuration ([#180](https://github.com/adaptive-enforcement-lab/readability/issues/180)) ([f80543f](https://github.com/adaptive-enforcement-lab/readability/commit/f80543ff0fba7a6e77f524ef16d20d662625a026)), closes [#160](https://github.com/adaptive-enforcement-lab/readability/issues/160)
* publish JSON Schema to MkDocs documentation site ([#182](https://github.com/adaptive-enforcement-lab/readability/issues/182)) ([645f6bc](https://github.com/adaptive-enforcement-lab/readability/commit/645f6bcf736ce75e86103790c6983fd9483c35da))

## [1.11.4](https://github.com/adaptive-enforcement-lab/readability/compare/v1.11.3...v1.11.4) (2025-12-17)


### Maintenance

* **deps:** update github/codeql-action digest to 5d4e8d1 ([#176](https://github.com/adaptive-enforcement-lab/readability/issues/176)) ([466842e](https://github.com/adaptive-enforcement-lab/readability/commit/466842e34959e98be0657446ca393c1b5c12691f))
* **deps:** update trufflesecurity/trufflehog digest to 8c1219a ([#177](https://github.com/adaptive-enforcement-lab/readability/issues/177)) ([b731949](https://github.com/adaptive-enforcement-lab/readability/commit/b731949172072c82e236f5576934ca7b565674f0))

## [1.11.3](https://github.com/adaptive-enforcement-lab/readability/compare/v1.11.2...v1.11.3) (2025-12-15)


### Maintenance

* **deps:** update codecov/test-results-action digest to 0fa95f0 ([#172](https://github.com/adaptive-enforcement-lab/readability/issues/172)) ([e3cd7f2](https://github.com/adaptive-enforcement-lab/readability/commit/e3cd7f264ee66e94cbea97e780ca0439541dce01))
* **deps:** update trufflesecurity/trufflehog digest to 8aea6cd ([#169](https://github.com/adaptive-enforcement-lab/readability/issues/169)) ([14653a6](https://github.com/adaptive-enforcement-lab/readability/commit/14653a6c80503a19e838c4aaf2b1383f28502307))

## [1.11.2](https://github.com/adaptive-enforcement-lab/readability/compare/v1.11.1...v1.11.2) (2025-12-14)


### Bug Fixes

* add v prefix to mkdocs version URLs for consistency ([#166](https://github.com/adaptive-enforcement-lab/readability/issues/166)) ([0d40d20](https://github.com/adaptive-enforcement-lab/readability/commit/0d40d2041d0b6ea60dc3bf920672467d44f5fa97))

## [1.11.1](https://github.com/adaptive-enforcement-lab/readability/compare/v1.11.0...v1.11.1) (2025-12-14)


### Bug Fixes

* Exclude frontmatter from prose extraction ([#164](https://github.com/adaptive-enforcement-lab/readability/issues/164)) ([afa3a36](https://github.com/adaptive-enforcement-lab/readability/commit/afa3a361c12e6b633f5ea82bde3b2716f5c05f1c))

## [1.11.0](https://github.com/adaptive-enforcement-lab/readability/compare/v1.10.5...v1.11.0) (2025-12-14)


### Features

* Detect AI slop via mid-sentence dash patterns ([#162](https://github.com/adaptive-enforcement-lab/readability/issues/162)) ([3dfc09f](https://github.com/adaptive-enforcement-lab/readability/commit/3dfc09f31a6d6592da331609a54bcbd5f659d81b))

## [1.10.5](https://github.com/adaptive-enforcement-lab/readability/compare/v1.10.4...v1.10.5) (2025-12-14)


### Bug Fixes

* make release binaries executable after artifact download ([#158](https://github.com/adaptive-enforcement-lab/readability/issues/158)) ([02fc306](https://github.com/adaptive-enforcement-lab/readability/commit/02fc3062f9a76e4d7cdf60cf69470337d47d1df9))

## [1.10.4](https://github.com/adaptive-enforcement-lab/readability/compare/v1.10.3...v1.10.4) (2025-12-14)


### Bug Fixes

* make container binaries executable after artifact download ([#156](https://github.com/adaptive-enforcement-lab/readability/issues/156)) ([1bd5ccd](https://github.com/adaptive-enforcement-lab/readability/commit/1bd5ccd14a15d0d53551880cec0af8d5f7ab0d14))

## [1.10.3](https://github.com/adaptive-enforcement-lab/readability/compare/v1.10.2...v1.10.3) (2025-12-14)


### Maintenance

* **deps:** update actions/download-artifact action to v7 ([#152](https://github.com/adaptive-enforcement-lab/readability/issues/152)) ([f43b7a6](https://github.com/adaptive-enforcement-lab/readability/commit/f43b7a66adb6b55019ca64350981a73d9bd8af37))
* **deps:** update docker/setup-qemu-action digest to c7c5346 ([#151](https://github.com/adaptive-enforcement-lab/readability/issues/151)) ([f491d10](https://github.com/adaptive-enforcement-lab/readability/commit/f491d109567899355006cafe8b68ccdf0e5b7970))
* remove deprecated publish-container workflow ([#154](https://github.com/adaptive-enforcement-lab/readability/issues/154)) ([923ea88](https://github.com/adaptive-enforcement-lab/readability/commit/923ea8894c65fbca56229384a090b23dbdb63186))

## [1.10.2](https://github.com/adaptive-enforcement-lab/readability/compare/v1.10.1...v1.10.2) (2025-12-14)


### Bug Fixes

* inline container publishing in release.yml for Scorecard detection ([#150](https://github.com/adaptive-enforcement-lab/readability/issues/150)) ([b676272](https://github.com/adaptive-enforcement-lab/readability/commit/b676272ea785b03ed16f3d1bfabe748fa652bc3e))

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
