# Changelog

## [Unreleased](https://github.com/argoproj/argo-cd/tree/HEAD)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.4.3...HEAD)

**Implemented enhancements:**

- Ignore extra objects dynamically created in same namespace [\#3498](https://github.com/argoproj/argo-cd/issues/3498)
- CLI: allow --server to point to an alternative base URL [\#3483](https://github.com/argoproj/argo-cd/issues/3483)
- Sorely missed features in application files [\#3469](https://github.com/argoproj/argo-cd/issues/3469)
- Upgrade bundled kustomize version [\#3458](https://github.com/argoproj/argo-cd/issues/3458)
- Ability to add local directory as source [\#3432](https://github.com/argoproj/argo-cd/issues/3432)
- Make it harder to brute force passwords [\#3423](https://github.com/argoproj/argo-cd/issues/3423)
- Support separate Kustomize version per application [\#3342](https://github.com/argoproj/argo-cd/issues/3342)
- Upgrade dex to v2.23.0 [\#3339](https://github.com/argoproj/argo-cd/issues/3339)
- Update to jsonnet v1.15.0 [\#3277](https://github.com/argoproj/argo-cd/issues/3277)
- Named JWT tokens [\#3213](https://github.com/argoproj/argo-cd/issues/3213)
- Add whitelisted namespaced resources support [\#2900](https://github.com/argoproj/argo-cd/issues/2900)
- Integrate OpenShift OAuth Server for Argo CD SSO [\#2574](https://github.com/argoproj/argo-cd/issues/2574)

**Fixed bugs:**

- ArgoCD-Repo-Server does not close https connections to git-server [\#3522](https://github.com/argoproj/argo-cd/issues/3522)
- Dex fails on startup due to a missing secret in the argo-cd-cli configuration [\#3519](https://github.com/argoproj/argo-cd/issues/3519)
- The `argocd sync` does not take into account IgnoreExtraneous annotation [\#3485](https://github.com/argoproj/argo-cd/issues/3485)
- ArgoCD runs crazy and hammers against the k8s API server [\#3481](https://github.com/argoproj/argo-cd/issues/3481)
- ArgoCD deleted all applications during the GitHub service degradation [\#3473](https://github.com/argoproj/argo-cd/issues/3473)
- argocd cli v1.0.0 not working [\#3471](https://github.com/argoproj/argo-cd/issues/3471)
- Status panel the text overflows. [\#3459](https://github.com/argoproj/argo-cd/issues/3459)
- Accounts list page crashes if users does not have accounts access [\#3452](https://github.com/argoproj/argo-cd/issues/3452)
- Default grafana dashboard is invalid [\#3419](https://github.com/argoproj/argo-cd/issues/3419)
- argocd repo add = Unimplemented desc = unknown method CreateRepository" [\#3417](https://github.com/argoproj/argo-cd/issues/3417)
- Evaluate attack vector of GHSA-qm7j-c969-7j4q on ArgoCD \(CVE-2020-5260\) [\#3416](https://github.com/argoproj/argo-cd/issues/3416)
- Inconsistent resource overrides key parsing [\#3406](https://github.com/argoproj/argo-cd/issues/3406)
- error: code = Unimplemented desc = unknown method ListRepositories [\#3402](https://github.com/argoproj/argo-cd/issues/3402)
- helm --set-file does not work [\#3365](https://github.com/argoproj/argo-cd/issues/3365)
- Cannot sync same resource to multiple clusters [\#3328](https://github.com/argoproj/argo-cd/issues/3328)
- App history call fails when using --basehref and Nginx rewrite annotation [\#3319](https://github.com/argoproj/argo-cd/issues/3319)
- syncPolicy validate being deleted upon failed sync [\#3208](https://github.com/argoproj/argo-cd/issues/3208)
- History and Rollback with Helm Charts [\#3082](https://github.com/argoproj/argo-cd/issues/3082)
- UI with multiple SSO provider render one button and name it with first provider found  [\#3009](https://github.com/argoproj/argo-cd/issues/3009)
- Login cookie clash on two ArgoCD instances sharing same domain [\#3005](https://github.com/argoproj/argo-cd/issues/3005)
- v1.3.x: RBAC applications get permission now shows error in UI [\#2805](https://github.com/argoproj/argo-cd/issues/2805)
- OIDC project role binding is not working [\#2723](https://github.com/argoproj/argo-cd/issues/2723)
- `no repository definition` in Plugins but declared in argocd-cm already [\#1797](https://github.com/argoproj/argo-cd/issues/1797)

**Closed issues:**

- how can create namespace [\#3490](https://github.com/argoproj/argo-cd/issues/3490)
- How to add Google Cloud Source Repo to ArgoCD hosted in GKE? [\#3379](https://github.com/argoproj/argo-cd/issues/3379)
- Application destination without specifying namespace [\#3285](https://github.com/argoproj/argo-cd/issues/3285)
- Enhancement proposal : Helm repo add to support --insecure-skip-verify [\#3113](https://github.com/argoproj/argo-cd/issues/3113)

**Merged pull requests:**

- fix: Disable keep-alive for HTTPS connection to Git \(\#3522\) [\#3531](https://github.com/argoproj/argo-cd/pull/3531) ([jannfis](https://github.com/jannfis))
- fix: revert "fix: Text overflow when the application status panel item was  too big \(\#3460\)" [\#3530](https://github.com/argoproj/argo-cd/pull/3530) ([alexmt](https://github.com/alexmt))
- docs: Adding OneLogin integration docs [\#3529](https://github.com/argoproj/argo-cd/pull/3529) ([christopappas](https://github.com/christopappas))
- docs: v1.5.3 changelog [\#3526](https://github.com/argoproj/argo-cd/pull/3526) ([alexmt](https://github.com/alexmt))
- docs: update document for --rootpath support for argoCD server and UI [\#3525](https://github.com/argoproj/argo-cd/pull/3525) ([mayzhang2000](https://github.com/mayzhang2000))
- Fix: Remove extra backtick\(`\) in JSX [\#3518](https://github.com/argoproj/argo-cd/pull/3518) ([jqlu](https://github.com/jqlu))
- feat: upgrade helm 3 to v3.2.0; user --insecure-verify-flag [\#3514](https://github.com/argoproj/argo-cd/pull/3514) ([alexmt](https://github.com/alexmt))
- docs: describe upgrading process and version breaking changes [\#3512](https://github.com/argoproj/argo-cd/pull/3512) ([alexmt](https://github.com/alexmt))
- fix: api incorrectly verifies if auto-sync is enabled and reject sync local request [\#3506](https://github.com/argoproj/argo-cd/pull/3506) ([alexmt](https://github.com/alexmt))
- fix: redis timeout should be more than client timeout [\#3505](https://github.com/argoproj/argo-cd/pull/3505) ([alexmt](https://github.com/alexmt))
- fix: to support --rootpath [\#3503](https://github.com/argoproj/argo-cd/pull/3503) ([mayzhang2000](https://github.com/mayzhang2000))
- fix: adding path to argoCD cookie [\#3501](https://github.com/argoproj/argo-cd/pull/3501) ([mayzhang2000](https://github.com/mayzhang2000))
- feat: add redis metrics to application controller and api server [\#3500](https://github.com/argoproj/argo-cd/pull/3500) ([alexmt](https://github.com/alexmt))
- docs: added MTN Group as a user of ArgoCD [\#3497](https://github.com/argoproj/argo-cd/pull/3497) ([beliot-atos](https://github.com/beliot-atos))
- fix: 'argocd sync' does not take into account IgnoreExtraneous annotation [\#3486](https://github.com/argoproj/argo-cd/pull/3486) ([alexmt](https://github.com/alexmt))
- fix: CLI renders flipped diff results [\#3480](https://github.com/argoproj/argo-cd/pull/3480) ([alexmt](https://github.com/alexmt))
- fix: Set root path [\#3475](https://github.com/argoproj/argo-cd/pull/3475) ([mayzhang2000](https://github.com/mayzhang2000))
- feat: upgrade kustomize version to 3.5.4 [\#3472](https://github.com/argoproj/argo-cd/pull/3472) ([alexmt](https://github.com/alexmt))
- docs: Update README.md for manifest installs [\#3470](https://github.com/argoproj/argo-cd/pull/3470) ([rumstead](https://github.com/rumstead))
- feat: downgrade dex to 2.22.0 and revert bug workaround [\#3468](https://github.com/argoproj/argo-cd/pull/3468) ([alexmt](https://github.com/alexmt))
- feat: limit the maximum number of concurrent login attempts [\#3467](https://github.com/argoproj/argo-cd/pull/3467) ([alexmt](https://github.com/alexmt))
- Fix: Text overflow when the application status panel item was too big. \#3459 [\#3460](https://github.com/argoproj/argo-cd/pull/3460) ([martinjaimem](https://github.com/martinjaimem))
- docs: update incorrect sync options example [\#3457](https://github.com/argoproj/argo-cd/pull/3457) ([alexmt](https://github.com/alexmt))
- fix: GetApplicationSyncWindows API should not validate project permissions [\#3456](https://github.com/argoproj/argo-cd/pull/3456) ([alexmt](https://github.com/alexmt))
- fix: javascript error on accounts list page [\#3453](https://github.com/argoproj/argo-cd/pull/3453) ([alexmt](https://github.com/alexmt))
- feat: upgrade dex to v2.23.0 [\#3448](https://github.com/argoproj/argo-cd/pull/3448) ([alexmt](https://github.com/alexmt))
- fix: argocd-util kubeconfig should use RawRestConfig to export config [\#3447](https://github.com/argoproj/argo-cd/pull/3447) ([alexmt](https://github.com/alexmt))
- chore: clean docker images before building Argo CD image to avoid no disk space left failure [\#3446](https://github.com/argoproj/argo-cd/pull/3446) ([alexmt](https://github.com/alexmt))
- chore: Disable lint in CI due to OOM fails [\#3444](https://github.com/argoproj/argo-cd/pull/3444) ([jannfis](https://github.com/jannfis))
- fix: Revert "feat: metrics, argocd\_app\_info adding syncpolicy info, argocd\_cluster\_info adding clustername" [\#3443](https://github.com/argoproj/argo-cd/pull/3443) ([jannfis](https://github.com/jannfis))
- fix: sort imports in knowntypes\_normalizer.go [\#3440](https://github.com/argoproj/argo-cd/pull/3440) ([alexmt](https://github.com/alexmt))
- Update OWNERS [\#3439](https://github.com/argoproj/argo-cd/pull/3439) ([edlee2121](https://github.com/edlee2121))
- fix: Grafana dashboard should dynamically load list of clusters [\#3435](https://github.com/argoproj/argo-cd/pull/3435) ([alexmt](https://github.com/alexmt))
- fix: use $datasource variable as a source for all dashboard panels [\#3434](https://github.com/argoproj/argo-cd/pull/3434) ([alexmt](https://github.com/alexmt))
- fix: support both \<group\>/\<kind\> as well as \<kind\> as a resource override key [\#3433](https://github.com/argoproj/argo-cd/pull/3433) ([alexmt](https://github.com/alexmt))
- fix: Updating to jsonnet v1.15.0 fix issue \#3277 [\#3431](https://github.com/argoproj/argo-cd/pull/3431) ([SouleBA](https://github.com/SouleBA))
- chore: Run lint in CircleCI due to shutdown of golangci.com [\#3427](https://github.com/argoproj/argo-cd/pull/3427) ([jannfis](https://github.com/jannfis))
- feat: support user specified account token ids [\#3425](https://github.com/argoproj/argo-cd/pull/3425) ([alexmt](https://github.com/alexmt))
- docs: Update documentation for CVE-2020-5260 [\#3421](https://github.com/argoproj/argo-cd/pull/3421) ([jannfis](https://github.com/jannfis))
- fix: helm repo add with flag --insecure-skip-server-verification. \#3113 [\#3420](https://github.com/argoproj/argo-cd/pull/3420) ([mayzhang2000](https://github.com/mayzhang2000))
- fix: Make CLI downwards compatible using old repository API \(\#3417 and \#3402\) [\#3418](https://github.com/argoproj/argo-cd/pull/3418) ([jannfis](https://github.com/jannfis))
- feat: support separate Kustomize version per application [\#3414](https://github.com/argoproj/argo-cd/pull/3414) ([alexmt](https://github.com/alexmt))
- fix: update min CLI version to 1.4.0 [\#3413](https://github.com/argoproj/argo-cd/pull/3413) ([alexmt](https://github.com/alexmt))
- chore: Fix a bunch of lint issues [\#3412](https://github.com/argoproj/argo-cd/pull/3412) ([jannfis](https://github.com/jannfis))
- feat: metrics, argocd\_app\_info adding syncpolicy info, argocd\_cluster\_info adding clustername [\#3411](https://github.com/argoproj/argo-cd/pull/3411) ([wecger](https://github.com/wecger))
- docs: Updated best practices to remove typos. [\#3410](https://github.com/argoproj/argo-cd/pull/3410) ([peteski22](https://github.com/peteski22))
- fix: app diff --local support for helm repo. \#3151 [\#3407](https://github.com/argoproj/argo-cd/pull/3407) ([mayzhang2000](https://github.com/mayzhang2000))
- fix: Add initial implementation for rate limiting failed logins [\#3404](https://github.com/argoproj/argo-cd/pull/3404) ([jannfis](https://github.com/jannfis))
- docs: Beautify & update security considerations doc [\#3400](https://github.com/argoproj/argo-cd/pull/3400) ([jannfis](https://github.com/jannfis))
- chore: Fix some Sonarcloud related quirks [\#3399](https://github.com/argoproj/argo-cd/pull/3399) ([jannfis](https://github.com/jannfis))
- feat: add settings troubleshooting commands to the 'argocd-util' binary [\#3398](https://github.com/argoproj/argo-cd/pull/3398) ([alexmt](https://github.com/alexmt))
- chore: Workaround for CircleCI bug [\#3397](https://github.com/argoproj/argo-cd/pull/3397) ([jannfis](https://github.com/jannfis))
- docs: Fix azure AD integration doc [\#3396](https://github.com/argoproj/argo-cd/pull/3396) ([toshyak](https://github.com/toshyak))
- docs: Add bug triage process proposal to docs [\#3394](https://github.com/argoproj/argo-cd/pull/3394) ([jannfis](https://github.com/jannfis))
- chore: Integrate sonarqube with our CI checks [\#3392](https://github.com/argoproj/argo-cd/pull/3392) ([jannfis](https://github.com/jannfis))
- fix: Display link between OLM ClusterServiceVersion and it's OperatorGroup. \(issue \#2999\) [\#3390](https://github.com/argoproj/argo-cd/pull/3390) ([dgoodwin](https://github.com/dgoodwin))
- feat: Let user to define meaningful unique JWT token name [\#3388](https://github.com/argoproj/argo-cd/pull/3388) ([rachelwang20](https://github.com/rachelwang20))
- fix:  \#3286 updated with the display message to say "app syncing" when syncing apps [\#3385](https://github.com/argoproj/argo-cd/pull/3385) ([ageekymonk](https://github.com/ageekymonk))
- docs: Update diffing\_known\_types.txt invalid link [\#3381](https://github.com/argoproj/argo-cd/pull/3381) ([djajcevic](https://github.com/djajcevic))
- docs: Add Dex and missing url field to FAQ [\#3380](https://github.com/argoproj/argo-cd/pull/3380) ([jannfis](https://github.com/jannfis))
- fix: Fix for Jsonnet when it is located in nested subdirectory. Closes \#3228 [\#3372](https://github.com/argoproj/argo-cd/pull/3372) ([mayzhang2000](https://github.com/mayzhang2000))
- chore: Code coverage offensive 05: util/clusterauth [\#3371](https://github.com/argoproj/argo-cd/pull/3371) ([jannfis](https://github.com/jannfis))
- fix: Update 4.5.3 redis-ha helm manifest [\#3370](https://github.com/argoproj/argo-cd/pull/3370) ([yutachaos](https://github.com/yutachaos))
- fix: return 401 error code if username does not exist [\#3369](https://github.com/argoproj/argo-cd/pull/3369) ([alexmt](https://github.com/alexmt))
- fix: Do not panic while running hooks with short revision \(\#3362\) [\#3368](https://github.com/argoproj/argo-cd/pull/3368) ([jannfis](https://github.com/jannfis))
- chore: Keep autogenerated message in upstream.yaml for HA manifests [\#3367](https://github.com/argoproj/argo-cd/pull/3367) ([jannfis](https://github.com/jannfis))
- docs: Add note about .git suffixes for GitLab repository URLs [\#3363](https://github.com/argoproj/argo-cd/pull/3363) ([jannfis](https://github.com/jannfis))
- docs: Clarify RBAC requirement for local users [\#3361](https://github.com/argoproj/argo-cd/pull/3361) ([jannfis](https://github.com/jannfis))
- docs: mention issue \#3358 in 1.5 changelog [\#3359](https://github.com/argoproj/argo-cd/pull/3359) ([alexmt](https://github.com/alexmt))
- feat: support normalizing CRD fields that use known built-in K8S types [\#3357](https://github.com/argoproj/argo-cd/pull/3357) ([alexmt](https://github.com/alexmt))
- fix: Increase HAProxy check interval to prevent intermittent failures [\#3356](https://github.com/argoproj/argo-cd/pull/3356) ([jannfis](https://github.com/jannfis))
- docs: Added detail on how to use the caData field with Okta SAML \#3352 [\#3351](https://github.com/argoproj/argo-cd/pull/3351) ([saranicole](https://github.com/saranicole))
- chore: Fix flaky unit test TestWatchCacheUpdated \(race condition\) [\#3350](https://github.com/argoproj/argo-cd/pull/3350) ([jannfis](https://github.com/jannfis))
- fix: Helm v3 CRD are not deployed [\#3345](https://github.com/argoproj/argo-cd/pull/3345) ([mayzhang2000](https://github.com/mayzhang2000))
- docs: document built-in user limitations and workaround [\#3341](https://github.com/argoproj/argo-cd/pull/3341) ([alexmt](https://github.com/alexmt))
- chore: Code coverage offensive 04: util/cert [\#3340](https://github.com/argoproj/argo-cd/pull/3340) ([jannfis](https://github.com/jannfis))
- fix: set MaxCallSendMsgSize and MaxCallRecvMsgSize to MaxGRPCMessageSize for the GRPC RepoServerClient \(fixes \#3117\) [\#3337](https://github.com/argoproj/argo-cd/pull/3337) ([Subreptivus](https://github.com/Subreptivus))
-  chore: Code coverage offensive 03: util/cache [\#3335](https://github.com/argoproj/argo-cd/pull/3335) ([jannfis](https://github.com/jannfis))
- feat: Whitelisted namespace in CLI [\#3333](https://github.com/argoproj/argo-cd/pull/3333) ([rachelwang20](https://github.com/rachelwang20))
- feat: Add v0.8 changes to Rollouts healthcheck [\#3331](https://github.com/argoproj/argo-cd/pull/3331) ([dthomson25](https://github.com/dthomson25))
- chore: fix invalid mount in 'make start' [\#3330](https://github.com/argoproj/argo-cd/pull/3330) ([alexmt](https://github.com/alexmt))
- chore: Code coverage offensive 02: util/argo [\#3329](https://github.com/argoproj/argo-cd/pull/3329) ([jannfis](https://github.com/jannfis))
- docs: improved Grafana dashboard [\#3327](https://github.com/argoproj/argo-cd/pull/3327) ([alexmt](https://github.com/alexmt))
- fix: Fixing could not find plugin issue when app sync and app diff [\#3326](https://github.com/argoproj/argo-cd/pull/3326) ([mayzhang2000](https://github.com/mayzhang2000))
- fix: argocd fails to connect clusters with IAM authentication configuration [\#3325](https://github.com/argoproj/argo-cd/pull/3325) ([alexmt](https://github.com/alexmt))
- docs: mention metrics changes and add legacy grafana dashboard [\#3324](https://github.com/argoproj/argo-cd/pull/3324) ([alexmt](https://github.com/alexmt))
- fix: update documentation for adding environment variable KUBE\_VERSION and KUBE\_API\_VERSION [\#3323](https://github.com/argoproj/argo-cd/pull/3323) ([mayzhang2000](https://github.com/mayzhang2000))
- feat: ArgoCD plugin: add environment variable KUBEVERSION and KUBE\_API\_VERS… [\#3318](https://github.com/argoproj/argo-cd/pull/3318) ([mayzhang2000](https://github.com/mayzhang2000))
- fix: avoid nil pointer dereference in badge handler [\#3316](https://github.com/argoproj/argo-cd/pull/3316) ([alexmt](https://github.com/alexmt))
- docs: \(and chore\) Add UI lint instructions and accompanying targets to Makefile [\#3315](https://github.com/argoproj/argo-cd/pull/3315) ([jannfis](https://github.com/jannfis))
- feat: Whitelisted namespace in UI [\#3314](https://github.com/argoproj/argo-cd/pull/3314) ([rachelwang20](https://github.com/rachelwang20))
- fix: pass APIVersions value to manifest generation request during app validation and during app manifests loading [\#3312](https://github.com/argoproj/argo-cd/pull/3312) ([alexmt](https://github.com/alexmt))
- fix: update help info about argcd account can-i [\#3310](https://github.com/argoproj/argo-cd/pull/3310) ([haoshuwei](https://github.com/haoshuwei))
- chore: Update argocd-test-tools to Go v1.14.1 [\#3306](https://github.com/argoproj/argo-cd/pull/3306) ([jannfis](https://github.com/jannfis))
- chore: Code coverage offensive \#1: util/dex [\#3305](https://github.com/argoproj/argo-cd/pull/3305) ([jannfis](https://github.com/jannfis))
- chore: Upgrade golang version from v1.14.0 to v1.14.1 [\#3304](https://github.com/argoproj/argo-cd/pull/3304) ([d-kuro](https://github.com/d-kuro))
- fix: Fix possible panic when generating Dex config from malformed YAML [\#3303](https://github.com/argoproj/argo-cd/pull/3303) ([jannfis](https://github.com/jannfis))
- chore: Run "dep check" in CircleCI pipeline to detect for changes in Gopkg.lock [\#3301](https://github.com/argoproj/argo-cd/pull/3301) ([jannfis](https://github.com/jannfis))
- fix: convert 'namesuffix', 'nameprefix' string flags to boolean flags in 'argocd app unset' [\#3300](https://github.com/argoproj/argo-cd/pull/3300) ([alexmt](https://github.com/alexmt))
- fix: use pagination while loading initial cluster state to avoid memory spikes [\#3299](https://github.com/argoproj/argo-cd/pull/3299) ([alexmt](https://github.com/alexmt))
- fix: SSO user unable to change local account password \(\#3297\) [\#3298](https://github.com/argoproj/argo-cd/pull/3298) ([alexmt](https://github.com/alexmt))
- fix: fix Cannot read property 'length' of undefined error [\#3296](https://github.com/argoproj/argo-cd/pull/3296) ([alexmt](https://github.com/alexmt))
- feat: Including namespace whiteliste resources support [\#3292](https://github.com/argoproj/argo-cd/pull/3292) ([rachelwang20](https://github.com/rachelwang20))
- fix: implement workaround for  helm/helm\#6870 bug [\#3290](https://github.com/argoproj/argo-cd/pull/3290) ([alexmt](https://github.com/alexmt))
- feat: Argocd App Unset Kustomize Override [\#3289](https://github.com/argoproj/argo-cd/pull/3289) ([mayzhang2000](https://github.com/mayzhang2000))
- docs: document Traefik 2 ingress requires disabling TLS [\#3284](https://github.com/argoproj/argo-cd/pull/3284) ([Addono](https://github.com/Addono))
- chore: Fix unparam errors from linter [\#3283](https://github.com/argoproj/argo-cd/pull/3283) ([jannfis](https://github.com/jannfis))
- fix: increase max connections count to support clusters with very large number of CRDs [\#3278](https://github.com/argoproj/argo-cd/pull/3278) ([alexmt](https://github.com/alexmt))
- docs: add slack link to issue template [\#3273](https://github.com/argoproj/argo-cd/pull/3273) ([alexmt](https://github.com/alexmt))
- docs: fix resource hooks docs layout \(\#3266\) [\#3272](https://github.com/argoproj/argo-cd/pull/3272) ([alexmt](https://github.com/alexmt))
- improvement: remove app name and project labels from reconcliation histogram to reduce cardinality [\#3271](https://github.com/argoproj/argo-cd/pull/3271) ([jessesuen](https://github.com/jessesuen))
- chore: update example dashboard to use updated metrics [\#3264](https://github.com/argoproj/argo-cd/pull/3264) ([alexmt](https://github.com/alexmt))
- Refine docs: How to explicitly select a tool [\#3261](https://github.com/argoproj/argo-cd/pull/3261) ([jomeier](https://github.com/jomeier))
- Add Healy to USERS.md [\#3250](https://github.com/argoproj/argo-cd/pull/3250) ([Therianthropie](https://github.com/Therianthropie))
- Add support for dex prometheus metrics [\#3249](https://github.com/argoproj/argo-cd/pull/3249) ([bclermont](https://github.com/bclermont))
- chore: Containerize complete build & test toolchain [\#3245](https://github.com/argoproj/argo-cd/pull/3245) ([jannfis](https://github.com/jannfis))

## [v1.4.3](https://github.com/argoproj/argo-cd/tree/v1.4.3) (2020-04-15)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.5.2...v1.4.3)

## [v1.5.2](https://github.com/argoproj/argo-cd/tree/v1.5.2) (2020-04-15)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.5.1...v1.5.2)

**Implemented enhancements:**

- support CRDs created by controllers [\#2873](https://github.com/argoproj/argo-cd/issues/2873)
- Improve diff support for resource requests/limits [\#1615](https://github.com/argoproj/argo-cd/issues/1615)

**Fixed bugs:**

- argocd-application-controller unexpectedly exit  [\#3395](https://github.com/argoproj/argo-cd/issues/3395)
- Helm deployments don't evaluate `.Capabilities.APIVersions.Has` correctly [\#3374](https://github.com/argoproj/argo-cd/issues/3374)
- ArgoCD v1.5.0 occured warning kustomize build [\#3348](https://github.com/argoproj/argo-cd/issues/3348)
- Syncing apps incorrectly states "app synced", but this is not true [\#3286](https://github.com/argoproj/argo-cd/issues/3286)
- jsonnet filename is passed without directory leading to odd import behavior with nested jsonnet directories [\#3228](https://github.com/argoproj/argo-cd/issues/3228)
- Helm repositories are not added for a diff against local manifests [\#3151](https://github.com/argoproj/argo-cd/issues/3151)
- ArgoCD config-management-plugin process substitution not possible [\#3007](https://github.com/argoproj/argo-cd/issues/3007)
- Argo-CD not detecting/linking generated objects [\#2999](https://github.com/argoproj/argo-cd/issues/2999)
- cert-manager crd are always unsynced [\#2239](https://github.com/argoproj/argo-cd/issues/2239)

**Closed issues:**

- Create SECURITY.md \(add security disclosure process\) [\#3384](https://github.com/argoproj/argo-cd/issues/3384)
- Error while trying to Login to argocd [\#3373](https://github.com/argoproj/argo-cd/issues/3373)

## [v1.5.1](https://github.com/argoproj/argo-cd/tree/v1.5.1) (2020-04-06)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/stable...v1.5.1)

**Implemented enhancements:**

- Section on caData in Okta SAML doc needs enhancement [\#3352](https://github.com/argoproj/argo-cd/issues/3352)
- Script build AMD64/ARM/ARM64 [\#3344](https://github.com/argoproj/argo-cd/issues/3344)
- Add security contacts to GitHub repo [\#3105](https://github.com/argoproj/argo-cd/issues/3105)
- Argocd autosync with cron schedule to enable nightly synchronizations [\#2804](https://github.com/argoproj/argo-cd/issues/2804)

**Fixed bugs:**

- Helm 3 issue with api version support [\#3364](https://github.com/argoproj/argo-cd/issues/3364)
- Panic on PostSync [\#3362](https://github.com/argoproj/argo-cd/issues/3362)
- HAProxy mode causes Redis intermittent failures [\#3358](https://github.com/argoproj/argo-cd/issues/3358)
- ArgoCD v1.5.0 Ignores kustomize field when creating application programmatically  [\#3355](https://github.com/argoproj/argo-cd/issues/3355)
- Helm v3 CRD are not deployed. [\#3336](https://github.com/argoproj/argo-cd/issues/3336)
- Config management plugin with name ... is not supported \(LOCAL SYNC\) [\#3293](https://github.com/argoproj/argo-cd/issues/3293)
- argocd-dex-server degrades to Init:CrashLoopBackOff [\#3070](https://github.com/argoproj/argo-cd/issues/3070)
- Setting TargetRevision to 'refs/tags/xxx' doesn't work [\#3067](https://github.com/argoproj/argo-cd/issues/3067)
- Argo CD Slack invite link not working [\#2357](https://github.com/argoproj/argo-cd/issues/2357)

## [stable](https://github.com/argoproj/argo-cd/tree/stable) (2020-04-02)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.5.0...stable)

## [v1.5.0](https://github.com/argoproj/argo-cd/tree/v1.5.0) (2020-04-02)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.5.0-rc3...v1.5.0)

**Implemented enhancements:**

- \[Feature request\] Pause and Resume [\#3317](https://github.com/argoproj/argo-cd/issues/3317)
- set KUBE\_VERSION and KUBE\_API\_VERSIONS env variables for custom plugins [\#2997](https://github.com/argoproj/argo-cd/issues/2997)

**Fixed bugs:**

- argocd-cli sso login basehref \[impossible\] [\#3334](https://github.com/argoproj/argo-cd/issues/3334)

## [v1.5.0-rc3](https://github.com/argoproj/argo-cd/tree/v1.5.0-rc3) (2020-03-30)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.5.0-rc2...v1.5.0-rc3)

**Implemented enhancements:**

- Question: does argo-cd allows to work in pull architecture? [\#3268](https://github.com/argoproj/argo-cd/issues/3268)

**Fixed bugs:**

- Creating app with plug as an option does not work unmarshal manifest: error [\#3311](https://github.com/argoproj/argo-cd/issues/3311)
- need update help examples about account can-i [\#3309](https://github.com/argoproj/argo-cd/issues/3309)
- is there a way to delete an existing cluster in argocd? [\#3302](https://github.com/argoproj/argo-cd/issues/3302)
- SSO user is unable to change local user password [\#3297](https://github.com/argoproj/argo-cd/issues/3297)
- Failed to unmarshal manifest: error unmarshaling JSON: json: cannot unmarshal string into Go value of type map\[string\]interface {} [\#3294](https://github.com/argoproj/argo-cd/issues/3294)
- Badges do not work properly in App-of-Apps deployments [\#3287](https://github.com/argoproj/argo-cd/issues/3287)
- Config management plugin with name 'helm-secrets' is not supported. [\#3267](https://github.com/argoproj/argo-cd/issues/3267)
- Prometheus Operator Custom Resource Definitions won't sync correctly [\#3240](https://github.com/argoproj/argo-cd/issues/3240)
- Unset a Kustomize Override [\#2546](https://github.com/argoproj/argo-cd/issues/2546)

**Closed issues:**

- 308 infinite redirect [\#3274](https://github.com/argoproj/argo-cd/issues/3274)

## [v1.5.0-rc2](https://github.com/argoproj/argo-cd/tree/v1.5.0-rc2) (2020-03-26)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.5.0-rc1...v1.5.0-rc2)

**Implemented enhancements:**

- Submit argo-cd to the CDF \(CD foundation\) [\#3265](https://github.com/argoproj/argo-cd/issues/3265)
- Git Submodules using Helm for multi environments  [\#3083](https://github.com/argoproj/argo-cd/issues/3083)
- Add target/current revision to status badge [\#2445](https://github.com/argoproj/argo-cd/issues/2445)

**Fixed bugs:**

- helm dependency build failed [\#3280](https://github.com/argoproj/argo-cd/issues/3280)
- Broken document layout on resource\_hook sample code. [\#3266](https://github.com/argoproj/argo-cd/issues/3266)
- impossible to add external HELM repos [\#3055](https://github.com/argoproj/argo-cd/issues/3055)
- Can't use ArgoCD CLI with custom certs? [\#2580](https://github.com/argoproj/argo-cd/issues/2580)

**Closed issues:**

- --insecure flag not working  [\#3270](https://github.com/argoproj/argo-cd/issues/3270)
- argocd clusterrole and binding  [\#3197](https://github.com/argoproj/argo-cd/issues/3197)
- Auto update helm chart to the latest version [\#3146](https://github.com/argoproj/argo-cd/issues/3146)

## [v1.5.0-rc1](https://github.com/argoproj/argo-cd/tree/v1.5.0-rc1) (2020-03-20)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.4.2...v1.5.0-rc1)

**Implemented enhancements:**

- Local users/account UI [\#3221](https://github.com/argoproj/argo-cd/issues/3221)
- Helm Version Tracking [\#3212](https://github.com/argoproj/argo-cd/issues/3212)
- Argo CD Service Account / Local Users [\#3185](https://github.com/argoproj/argo-cd/issues/3185)
- Add ability to skip cluster RBAC setup in CLI [\#3183](https://github.com/argoproj/argo-cd/issues/3183)
- App-of-apps how to set name of root app? [\#3167](https://github.com/argoproj/argo-cd/issues/3167)
- Allow for larger cookies [\#3126](https://github.com/argoproj/argo-cd/issues/3126)
- Enhancement Proposal \(UI\): provide docs to sync policy options present in create application panel.  [\#3098](https://github.com/argoproj/argo-cd/issues/3098)
- Add dest cluster and namespace in the Events [\#3084](https://github.com/argoproj/argo-cd/issues/3084)
- Disable admin login [\#3019](https://github.com/argoproj/argo-cd/issues/3019)
- Replace aws-iam-authenticator with aws cli to fix EKS service account roles [\#2996](https://github.com/argoproj/argo-cd/issues/2996)
- Replace `kubctl apply` with API invocation to avoid shell fork [\#2822](https://github.com/argoproj/argo-cd/issues/2822)
- Add webhook support for Bitbucket server with ssh [\#2773](https://github.com/argoproj/argo-cd/issues/2773)
- Add helm --set-file support [\#2751](https://github.com/argoproj/argo-cd/issues/2751)
- Upgrade to Go v1.13.x [\#2522](https://github.com/argoproj/argo-cd/issues/2522)
- Rollback should disable autosync [\#2441](https://github.com/argoproj/argo-cd/issues/2441)
- Support Helm v3 [\#2383](https://github.com/argoproj/argo-cd/issues/2383)
- Wildcard sync using CLI [\#2228](https://github.com/argoproj/argo-cd/issues/2228)
- Need Windows argocd cmd [\#2121](https://github.com/argoproj/argo-cd/issues/2121)
-  Helm Capabilities built-in object is not supported by repository server [\#1251](https://github.com/argoproj/argo-cd/issues/1251)
- Add ability to kubectl apply objects using `--validate=false` [\#1063](https://github.com/argoproj/argo-cd/issues/1063)
- feat: Supports the validate-false option at an app level. Closes \#1063 [\#2542](https://github.com/argoproj/argo-cd/pull/2542) ([alexec](https://github.com/alexec))

**Fixed bugs:**

- App reconciliation fails with panic: index out of range  [\#3232](https://github.com/argoproj/argo-cd/issues/3232)
- Helm 3 pull timeout for stable repo [\#3229](https://github.com/argoproj/argo-cd/issues/3229)
- Could not create new object of type X, Kind=Y, no Kind=Y is registered for version Z in scheme  \"k8s.io/client-go/kubernetes/scheme/register.go:67\" [\#3216](https://github.com/argoproj/argo-cd/issues/3216)
- Unable to connect to repository parse \<repo\_name\> first path segment in URL cannot contain colon [\#3210](https://github.com/argoproj/argo-cd/issues/3210)
- unable to use namespace isolation [\#3191](https://github.com/argoproj/argo-cd/issues/3191)
- Failed helm fetch call prints sensitive information in logs [\#3188](https://github.com/argoproj/argo-cd/issues/3188)
- Default password is not set to pod name [\#3187](https://github.com/argoproj/argo-cd/issues/3187)
- Required privileges that should not be required [\#3182](https://github.com/argoproj/argo-cd/issues/3182)
- Description of non-cascade command is incorrect [\#3173](https://github.com/argoproj/argo-cd/issues/3173)
- Error on Sync: PROTOCOL\_ERROR [\#3170](https://github.com/argoproj/argo-cd/issues/3170)
- Github SSO login not working with v1.4.2 [\#3169](https://github.com/argoproj/argo-cd/issues/3169)
- Dex + SSO login does not work if Argo CD is behind reverse proxy with a different base href [\#3163](https://github.com/argoproj/argo-cd/issues/3163)
- Long string in value for application info item causes UI to hang [\#3159](https://github.com/argoproj/argo-cd/issues/3159)
- cannot access helm repositories via proxy [\#3158](https://github.com/argoproj/argo-cd/issues/3158)
- github dex oauth2 not working: server response missing access\_token [\#3149](https://github.com/argoproj/argo-cd/issues/3149)
- ArgoCD UI: login redirect loop [\#3137](https://github.com/argoproj/argo-cd/issues/3137)
- Error Sync on large manifests: grpc: message larger than max [\#3117](https://github.com/argoproj/argo-cd/issues/3117)
- cluster add: x509: certificate signed by unknown authority [\#3116](https://github.com/argoproj/argo-cd/issues/3116)
- New release from Helm repo not immediately picked up on resync [\#3107](https://github.com/argoproj/argo-cd/issues/3107)
- No AutoSync is triggered for app with Helm source [\#3104](https://github.com/argoproj/argo-cd/issues/3104)
- Disaster Recovery export fails with permission denied loading kube config [\#3094](https://github.com/argoproj/argo-cd/issues/3094)
- Can't apply --label in CLI app create [\#3090](https://github.com/argoproj/argo-cd/issues/3090)
- Question about webhook [\#3089](https://github.com/argoproj/argo-cd/issues/3089)
- Can't create new App [\#3077](https://github.com/argoproj/argo-cd/issues/3077)
- ArgoCD can't find `Kustomization.yaml` \(capitalized\) [\#3075](https://github.com/argoproj/argo-cd/issues/3075)
- Connecting repositories using git@server-Style SSH URLs lead to undefined behaviour [\#3061](https://github.com/argoproj/argo-cd/issues/3061)
- The bug when using multiple configManagementPlugin [\#3058](https://github.com/argoproj/argo-cd/issues/3058)
- SSO failure on argocd-cli \(plain oidc, no dex\) [\#3054](https://github.com/argoproj/argo-cd/issues/3054)
- PVC created by volumeClaimTemplates always pruned [\#3050](https://github.com/argoproj/argo-cd/issues/3050)
- Labels not being deleted via UI [\#3046](https://github.com/argoproj/argo-cd/issues/3046)
- Unable to parse kubectl pre-release version strings [\#3034](https://github.com/argoproj/argo-cd/issues/3034)
- Helm "values: |" are ignored [\#3032](https://github.com/argoproj/argo-cd/issues/3032)
- K8s secrets for repository credential templates are not deleted when credential template is deleted [\#3028](https://github.com/argoproj/argo-cd/issues/3028)
- Opening in new tab bad key binding on Linux [\#3020](https://github.com/argoproj/argo-cd/issues/3020)
- v1.4 SSH credential template not working [\#3016](https://github.com/argoproj/argo-cd/issues/3016)
- Include resource group for Event's InvolvedObject.APIVersion [\#3011](https://github.com/argoproj/argo-cd/issues/3011)
- argocd 1.4: ui referring to /api/version using absolute path [\#3003](https://github.com/argoproj/argo-cd/issues/3003)
- Helm values.yaml are ignored [\#2967](https://github.com/argoproj/argo-cd/issues/2967)
- Cannot connect to Private Gitlab Repo [\#2890](https://github.com/argoproj/argo-cd/issues/2890)
- Cannot add https://source.cloud.google.com repo [\#2814](https://github.com/argoproj/argo-cd/issues/2814)
- argocd-server and argocd-repo-server do not parse SSH URLs the same ways [\#2576](https://github.com/argoproj/argo-cd/issues/2576)
- ArgoCD always shows login prompt when first accessed, even if current browser session has already authenticated with same OIDC identity provider [\#2535](https://github.com/argoproj/argo-cd/issues/2535)
- Intermittent unknown state. [\#2472](https://github.com/argoproj/argo-cd/issues/2472)
- Can't create Applications in GUi and CLI [\#2384](https://github.com/argoproj/argo-cd/issues/2384)
- JWT invalid =\> Password for superuser has changed since token issued  [\#2108](https://github.com/argoproj/argo-cd/issues/2108)
- Adding multiple OpenShift clusters [\#1761](https://github.com/argoproj/argo-cd/issues/1761)
- Server ignores initial TLS cert [\#995](https://github.com/argoproj/argo-cd/issues/995)
- fix: Hide bindPW in dex config [\#3025](https://github.com/argoproj/argo-cd/pull/3025) ([dnascimento](https://github.com/dnascimento))

**Closed issues:**

- DEX ldap connector returns blank groups [\#3112](https://github.com/argoproj/argo-cd/issues/3112)
- Problem to Acess Argo CD UI \(invalid x509 key pair tls.crt/tls.key\) [\#3078](https://github.com/argoproj/argo-cd/issues/3078)
- send logs to splunk [\#2934](https://github.com/argoproj/argo-cd/issues/2934)

**Merged pull requests:**

- fix: prevent syncing the same cluster cache multiple times due to race condition [\#3256](https://github.com/argoproj/argo-cd/pull/3256) ([alexmt](https://github.com/alexmt))
- fix: upgrade argoproj/pkg version [\#3255](https://github.com/argoproj/argo-cd/pull/3255) ([alexmt](https://github.com/alexmt))
- feat: add operation,sync\_status,health\_status labels to argocd\_app\_info metric [\#3254](https://github.com/argoproj/argo-cd/pull/3254) ([jessesuen](https://github.com/jessesuen))
- fix: stop trigging auto-syncing if app only need to prune resources and automated prune is disabled [\#3253](https://github.com/argoproj/argo-cd/pull/3253) ([alexmt](https://github.com/alexmt))
- feat: store the name of operation initator in the app CRD [\#3252](https://github.com/argoproj/argo-cd/pull/3252) ([alexmt](https://github.com/alexmt))
- fix: UI crashes when trying to enable auto-sync [\#3251](https://github.com/argoproj/argo-cd/pull/3251) ([alexmt](https://github.com/alexmt))
- fix: applications.List API should sort apps by name [\#3248](https://github.com/argoproj/argo-cd/pull/3248) ([alexmt](https://github.com/alexmt))
- feat: Introduce sync-option SkipDryRunOnMissingResource=true \(\#2873\) [\#3247](https://github.com/argoproj/argo-cd/pull/3247) ([a-hat](https://github.com/a-hat))
- perf: reduce unnecessary  K8s calls for CRDs during reconciliation [\#3246](https://github.com/argoproj/argo-cd/pull/3246) ([jessesuen](https://github.com/jessesuen))
- docs: add v1.5 change log [\#3244](https://github.com/argoproj/argo-cd/pull/3244) ([alexmt](https://github.com/alexmt))
- feat: use --api-versions during helm chart rendering [\#3243](https://github.com/argoproj/argo-cd/pull/3243) ([alexmt](https://github.com/alexmt))
- docs: Add missing parentheses to Webhook docs [\#3239](https://github.com/argoproj/argo-cd/pull/3239) ([HeroCC](https://github.com/HeroCC))
- fix: Test for nil to prevent nil pointer dereference in state.go [\#3237](https://github.com/argoproj/argo-cd/pull/3237) ([jannfis](https://github.com/jannfis))
- fix: fix broken URL regex expression [\#3236](https://github.com/argoproj/argo-cd/pull/3236) ([alexmt](https://github.com/alexmt))
- refactor: use http forwarders from argoproj/pkg repository [\#3235](https://github.com/argoproj/argo-cd/pull/3235) ([alexmt](https://github.com/alexmt))
- fix: 'requires pruning' is not rendered on app details page for resources that should be pruned [\#3234](https://github.com/argoproj/argo-cd/pull/3234) ([alexmt](https://github.com/alexmt))
- fix: app reconciliation fails with panic: index out of [\#3233](https://github.com/argoproj/argo-cd/pull/3233) ([alexmt](https://github.com/alexmt))
- fix: upgrade argoproj/pkg version to fix leaked sensitive informationin logs [\#3230](https://github.com/argoproj/argo-cd/pull/3230) ([alexmt](https://github.com/alexmt))
- fix: really skip app requeing for some types of resources [\#3225](https://github.com/argoproj/argo-cd/pull/3225) ([bclermont](https://github.com/bclermont))
- docs: add docs on mapping different scopes for microsoft [\#3224](https://github.com/argoproj/argo-cd/pull/3224) ([bergur88](https://github.com/bergur88))
- feat: improve api-server and controller performance [\#3222](https://github.com/argoproj/argo-cd/pull/3222) ([jessesuen](https://github.com/jessesuen))
- improvement: Surface failure reasons for Rollouts/AnalysisRuns [\#3219](https://github.com/argoproj/argo-cd/pull/3219) ([khhirani](https://github.com/khhirani))
- feat: implement Service Account / Local Users [\#3215](https://github.com/argoproj/argo-cd/pull/3215) ([alexmt](https://github.com/alexmt))
- chore: Update testify to v1.5.1 [\#3209](https://github.com/argoproj/argo-cd/pull/3209) ([jannfis](https://github.com/jannfis))
- docs: Fixed documentation fault in user-guide/projects/ [\#3207](https://github.com/argoproj/argo-cd/pull/3207) ([CBytelabs](https://github.com/CBytelabs))
- docs: Create USERS.md [\#3206](https://github.com/argoproj/argo-cd/pull/3206) ([saradhis](https://github.com/saradhis))
- feat\(ui\): add docs to sync policy options present in create application panel \(Close \#3098\) [\#3203](https://github.com/argoproj/argo-cd/pull/3203) ([Elgarni](https://github.com/Elgarni))
- fix: codegen task is broken locally [\#3200](https://github.com/argoproj/argo-cd/pull/3200) ([alexmt](https://github.com/alexmt))
- docs\(README\): Add Prudential to the customer list in README [\#3199](https://github.com/argoproj/argo-cd/pull/3199) ([hayorov](https://github.com/hayorov))
- fix: error message on login page if login is disabled [\#3198](https://github.com/argoproj/argo-cd/pull/3198) ([alexmt](https://github.com/alexmt))
- docs: Add comment for helm valueFiles path reference [\#3196](https://github.com/argoproj/argo-cd/pull/3196) ([sysadmind](https://github.com/sysadmind))
- fix: stop caching helm index [\#3193](https://github.com/argoproj/argo-cd/pull/3193) ([alexmt](https://github.com/alexmt))
- docs: document resource.inclusions setting [\#3190](https://github.com/argoproj/argo-cd/pull/3190) ([alexmt](https://github.com/alexmt))
- fix: set MaxCallSendMsgSize to MaxGRPCMessageSize for the GRPC caller \(fixes \#3117\) [\#3189](https://github.com/argoproj/argo-cd/pull/3189) ([matthyx](https://github.com/matthyx))
- feat: add "service-account" flag to "cluster add" command \(\#3183\) [\#3184](https://github.com/argoproj/argo-cd/pull/3184) ([alexandrfox](https://github.com/alexandrfox))
- feat: Disable Admin Login \(fixes \#3019\) [\#3179](https://github.com/argoproj/argo-cd/pull/3179) ([OmerKahani](https://github.com/OmerKahani))
- feat: support helm3 \(\#2383\) [\#3178](https://github.com/argoproj/argo-cd/pull/3178) ([alexmt](https://github.com/alexmt))
- docs: Fixed description of non-cascade command in App Deletion \(\#3173\) [\#3174](https://github.com/argoproj/argo-cd/pull/3174) ([m3y](https://github.com/m3y))
- docs: correct typo in diffing section on HPAs [\#3172](https://github.com/argoproj/argo-cd/pull/3172) ([aboyett](https://github.com/aboyett))
- fix: dex proxy should forward request to dex preserving the basehref [\#3165](https://github.com/argoproj/argo-cd/pull/3165) ([alexmt](https://github.com/alexmt))
- fix: set default login redirect to baseHRef \(\#3163\) [\#3164](https://github.com/argoproj/argo-cd/pull/3164) ([gpaul](https://github.com/gpaul))
- chore: Update Go version from v1.13.8 to v1.14.0 [\#3162](https://github.com/argoproj/argo-cd/pull/3162) ([d-kuro](https://github.com/d-kuro))
- fix: Unhang UI on long app info items by using more sane URL match pattern \(\#3159\) [\#3161](https://github.com/argoproj/argo-cd/pull/3161) ([jannfis](https://github.com/jannfis))
- improvement: examples/dashboard.json: removing filling, add increase to grpc metrics [\#3154](https://github.com/argoproj/argo-cd/pull/3154) ([wecger](https://github.com/wecger))
- fix: upgrade redis-ha chart and enable haproxy [\#3147](https://github.com/argoproj/argo-cd/pull/3147) ([shelby-moore](https://github.com/shelby-moore))
- chore\(argocd-ui\): fix broken url for argo logo in README.md and resize it [\#3145](https://github.com/argoproj/argo-cd/pull/3145) ([nurinamu](https://github.com/nurinamu))
- improvement: example/dashboard.json: added labels to legends [\#3140](https://github.com/argoproj/argo-cd/pull/3140) ([wecger](https://github.com/wecger))
- fix: don't double-prepend basehref to redirect URLs \(fixes \#3137\) [\#3138](https://github.com/argoproj/argo-cd/pull/3138) ([gpaul](https://github.com/gpaul))
- make dex server deployment init container resilient to restarts \(fixes \#3070\) [\#3136](https://github.com/argoproj/argo-cd/pull/3136) ([shelby-moore](https://github.com/shelby-moore))
- feat: Add auth0 docs [\#3133](https://github.com/argoproj/argo-cd/pull/3133) ([davidkarlsen](https://github.com/davidkarlsen))
- chore: Update Go version from v1.13.7 to v1.13.8 [\#3129](https://github.com/argoproj/argo-cd/pull/3129) ([d-kuro](https://github.com/d-kuro))
- feat: Add revision to status badge \(\#2445\) [\#3128](https://github.com/argoproj/argo-cd/pull/3128) ([milesarmstrong](https://github.com/milesarmstrong))
- docs: Improve docs for Hooks [\#3115](https://github.com/argoproj/argo-cd/pull/3115) ([agrimmer](https://github.com/agrimmer))
- docs: add missing CHANGELOG entry for 1.2.5 [\#3109](https://github.com/argoproj/argo-cd/pull/3109) ([rowleyaj](https://github.com/rowleyaj))
- \[docs\] add Bank-Vaults to Secret Management [\#3106](https://github.com/argoproj/argo-cd/pull/3106) ([bonifaido](https://github.com/bonifaido))
- Adding ThousandEyes to user-list in README.md [\#3103](https://github.com/argoproj/argo-cd/pull/3103) ([astein-te](https://github.com/astein-te))
- docs: Fix typo in custom\_tools.md [\#3100](https://github.com/argoproj/argo-cd/pull/3100) ([soupdiver](https://github.com/soupdiver))
- fix: argocd-util backup produced truncated backups [\#3096](https://github.com/argoproj/argo-cd/pull/3096) ([jessesuen](https://github.com/jessesuen))
- fix: install exact version of mkdocs and mkdocs\_material during docs site building [\#3095](https://github.com/argoproj/argo-cd/pull/3095) ([alexmt](https://github.com/alexmt))
- feat: add dest cluster and namespace in the Events [\#3093](https://github.com/argoproj/argo-cd/pull/3093) ([alexmt](https://github.com/alexmt))
- fix: ui referring to /api/version using absolute path [\#3092](https://github.com/argoproj/argo-cd/pull/3092) ([alexmt](https://github.com/alexmt))
- fix: reduct secret values of manifests stored in git [\#3088](https://github.com/argoproj/argo-cd/pull/3088) ([alexmt](https://github.com/alexmt))
- docs: json pointer in docs [\#3087](https://github.com/argoproj/argo-cd/pull/3087) ([duboisf](https://github.com/duboisf))
- fix: labels not being deleted via UI [\#3081](https://github.com/argoproj/argo-cd/pull/3081) ([alexmt](https://github.com/alexmt))
- fix: when user tries to rollback an app with auto-sync enabled UI should ask for user confirmation once [\#3080](https://github.com/argoproj/argo-cd/pull/3080) ([alexmt](https://github.com/alexmt))
- fix: correct a typo in kustomize.md [\#3079](https://github.com/argoproj/argo-cd/pull/3079) ([Elgarni](https://github.com/Elgarni))
- fix: Allow multiple hostnames per SSH known hosts entry and also allow IPv6 \(\#2814\) [\#3074](https://github.com/argoproj/argo-cd/pull/3074) ([jannfis](https://github.com/jannfis))
- chore: Use 'latest' as default tag for Docker image builds [\#3072](https://github.com/argoproj/argo-cd/pull/3072) ([jannfis](https://github.com/jannfis))
- chore: Update CHANGELOG.md up to v1.4.2 [\#3069](https://github.com/argoproj/argo-cd/pull/3069) ([jannfis](https://github.com/jannfis))
- fix: Correct usage text for repo add command regarding insecure repos [\#3068](https://github.com/argoproj/argo-cd/pull/3068) ([jannfis](https://github.com/jannfis))
- docs: Add MOO Print to list of users [\#3065](https://github.com/argoproj/argo-cd/pull/3065) ([milesarmstrong](https://github.com/milesarmstrong))
- fix: Ensure SSH private key is written out with a final newline character \(\#2890\) [\#3064](https://github.com/argoproj/argo-cd/pull/3064) ([jannfis](https://github.com/jannfis))
- fix: HTTP|HTTPS|NO\_PROXY env variable reading \#3055 [\#3063](https://github.com/argoproj/argo-cd/pull/3063) ([zeph](https://github.com/zeph))
- fix: Handle SSH URLs in 'git@server:org/repo' notation correctly [\#3062](https://github.com/argoproj/argo-cd/pull/3062) ([jannfis](https://github.com/jannfis))
- fix: Even if define multiple configManagementPlugin, you can only get one plugin \(\#3058\) [\#3059](https://github.com/argoproj/argo-cd/pull/3059) ([d-kuro](https://github.com/d-kuro))
- fix\(UI\): SSO condition when several SSO connectors has been configured [\#3057](https://github.com/argoproj/argo-cd/pull/3057) ([rayanebel](https://github.com/rayanebel))
- chore: Upgrade Go version from v1.13.6 to v1.13.7 [\#3056](https://github.com/argoproj/argo-cd/pull/3056) ([d-kuro](https://github.com/d-kuro))
- chore: fix codegen-local [\#3053](https://github.com/argoproj/argo-cd/pull/3053) ([alexmt](https://github.com/alexmt))
- chore: publish image on every push to master [\#3051](https://github.com/argoproj/argo-cd/pull/3051) ([alexmt](https://github.com/alexmt))
- docs: Clarify syntax for diff customization should be json-patch [\#3044](https://github.com/argoproj/argo-cd/pull/3044) ([aarongorka](https://github.com/aarongorka))
- docs: removing extra dot in docs [\#3043](https://github.com/argoproj/argo-cd/pull/3043) ([pablokbs](https://github.com/pablokbs))
- feat: Enable to open application details in new tab using ctrl-click \(\#3020\) [\#3038](https://github.com/argoproj/argo-cd/pull/3038) ([jannfis](https://github.com/jannfis))
- docs: Fix getting started 404, add useful links [\#3037](https://github.com/argoproj/argo-cd/pull/3037) ([GlenTiki](https://github.com/GlenTiki))
- feat: check ssh and https url in bitbucketserver webhook \#2773 [\#3036](https://github.com/argoproj/argo-cd/pull/3036) ([eSamS](https://github.com/eSamS))
- fix: Allow dash character in output of kubectl version \(\#3034\) [\#3035](https://github.com/argoproj/argo-cd/pull/3035) ([jannfis](https://github.com/jannfis))
- fix: Include resource group for events InvolvedObject.APIVersion \#3011 [\#3031](https://github.com/argoproj/argo-cd/pull/3031) ([d-kuro](https://github.com/d-kuro))
- docs: Clarify that non-standard SSH ports have to use ssh:// URLs [\#3030](https://github.com/argoproj/argo-cd/pull/3030) ([jannfis](https://github.com/jannfis))
- fix: Better handling of K8s secrets for repository credentials & templates \(\#3016 \#3028\) [\#3029](https://github.com/argoproj/argo-cd/pull/3029) ([jannfis](https://github.com/jannfis))
- fix: correctly replace cache in namespace isolation mode [\#3023](https://github.com/argoproj/argo-cd/pull/3023) ([alexmt](https://github.com/alexmt))
- fix: Jsonnet TLA parameters of same type are overwritten [\#3022](https://github.com/argoproj/argo-cd/pull/3022) ([dbeal-wiser](https://github.com/dbeal-wiser))
- docs: Update disaster\_recovery.md to fix verbiage around argocd-util. [\#3021](https://github.com/argoproj/argo-cd/pull/3021) ([rumstead](https://github.com/rumstead))
- fix: run dep ensure [\#3018](https://github.com/argoproj/argo-cd/pull/3018) ([alexmt](https://github.com/alexmt))
- fix: Provide correct example for repocreds add command  \(\#3016\) [\#3017](https://github.com/argoproj/argo-cd/pull/3017) ([jannfis](https://github.com/jannfis))
- feat: Add argocd cmd for Windows \#2121 [\#3015](https://github.com/argoproj/argo-cd/pull/3015) ([masa213f](https://github.com/masa213f))
- chore: add jannfis to approvers list [\#3014](https://github.com/argoproj/argo-cd/pull/3014) ([alexmt](https://github.com/alexmt))
- fix:  impossible to config RBAC if group name includes ',' [\#3013](https://github.com/argoproj/argo-cd/pull/3013) ([alexmt](https://github.com/alexmt))
- fix: Replace aws-iam-authenticator to support IRSA [\#3010](https://github.com/argoproj/argo-cd/pull/3010) ([jdmulloy](https://github.com/jdmulloy))
- chore: Upgrade Go version from v1.12.6 to v1.13.6 \#2929 [\#2994](https://github.com/argoproj/argo-cd/pull/2994) ([d-kuro](https://github.com/d-kuro))
- docs: update v1.4 contributors list [\#2993](https://github.com/argoproj/argo-cd/pull/2993) ([alexmt](https://github.com/alexmt))
- fix: Fixes username matching for SSH repositories for webhooks \(fixes \#2988\) [\#2990](https://github.com/argoproj/argo-cd/pull/2990) ([machgo](https://github.com/machgo))
- feat: upgrade dex to v2.21.0 [\#2985](https://github.com/argoproj/argo-cd/pull/2985) ([alexmt](https://github.com/alexmt))
- fix: sync apps panel fails with 'No App Selected' message if name contains '.' [\#2983](https://github.com/argoproj/argo-cd/pull/2983) ([alexmt](https://github.com/alexmt))
- docs: add notifications.md with recommandation about notifications [\#2979](https://github.com/argoproj/argo-cd/pull/2979) ([alexmt](https://github.com/alexmt))
- fix: Rendering CRD acronym [\#2978](https://github.com/argoproj/argo-cd/pull/2978) ([alexmt](https://github.com/alexmt))
- fix: remove 'total' suffix from gauge prom metric [\#2976](https://github.com/argoproj/argo-cd/pull/2976) ([alexmt](https://github.com/alexmt))
- fix: fix nil pointer dereference in CreateRepositoryCredentials method [\#2975](https://github.com/argoproj/argo-cd/pull/2975) ([alexmt](https://github.com/alexmt))
- docs: Fix a broken link to Helm Hooks [\#2970](https://github.com/argoproj/argo-cd/pull/2970) ([pbrit](https://github.com/pbrit))
- docs: v1.3 and v1.4 changelog [\#2952](https://github.com/argoproj/argo-cd/pull/2952) ([alexmt](https://github.com/alexmt))
- feat: Support setting Helm file parameters using CLI/UI [\#2752](https://github.com/argoproj/argo-cd/pull/2752) ([tomcruise81](https://github.com/tomcruise81))
- Rollback disables auto sync issue \#2441 [\#2591](https://github.com/argoproj/argo-cd/pull/2591) ([adamjohnson01](https://github.com/adamjohnson01))

## [v1.4.2](https://github.com/argoproj/argo-cd/tree/v1.4.2) (2020-01-24)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.4.1...v1.4.2)

**Fixed bugs:**

- Update Dockerfile to build with more recent version of Go [\#2929](https://github.com/argoproj/argo-cd/issues/2929)

## [v1.4.1](https://github.com/argoproj/argo-cd/tree/v1.4.1) (2020-01-22)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.4.0...v1.4.1)

**Implemented enhancements:**

- call argocd-server from client - navigating auth [\#2982](https://github.com/argoproj/argo-cd/issues/2982)

**Fixed bugs:**

- It is impossible to config RBAC if group name includes ',' [\#3012](https://github.com/argoproj/argo-cd/issues/3012)
- ArgoCD config-management-plugin process substitution [\#3006](https://github.com/argoproj/argo-cd/issues/3006)

## [v1.4.0](https://github.com/argoproj/argo-cd/tree/v1.4.0) (2020-01-18)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.4.0-rc1...v1.4.0)

**Implemented enhancements:**

- Community Meetings [\#2984](https://github.com/argoproj/argo-cd/issues/2984)

**Fixed bugs:**

- fetch repo timeout [\#2991](https://github.com/argoproj/argo-cd/issues/2991)
- SSH repo URL with a user different from `git` is not matched correctly when resolving a webhook [\#2988](https://github.com/argoproj/argo-cd/issues/2988)
- Invalid icon for CRD resources in list view [\#2977](https://github.com/argoproj/argo-cd/issues/2977)
- panic adding credentials template [\#2974](https://github.com/argoproj/argo-cd/issues/2974)
- ArgoCD isn't templating helm loops properly [\#2973](https://github.com/argoproj/argo-cd/issues/2973)
- Some build environment variables are not passed in a custom plugin [\#2971](https://github.com/argoproj/argo-cd/issues/2971)
- Sync Apps "No App Selected" when name container "." [\#2962](https://github.com/argoproj/argo-cd/issues/2962)
- Dex version 2.14.0 cause login without respecting gitlab groups policy [\#2958](https://github.com/argoproj/argo-cd/issues/2958)

## [v1.4.0-rc1](https://github.com/argoproj/argo-cd/tree/v1.4.0-rc1) (2020-01-13)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.3.6...v1.4.0-rc1)

**Implemented enhancements:**

- Support helm 3 [\#2864](https://github.com/argoproj/argo-cd/issues/2864)
- Remove `latest` from UI [\#2757](https://github.com/argoproj/argo-cd/issues/2757)
- Add support for CRD icons and tidy up app view [\#2625](https://github.com/argoproj/argo-cd/issues/2625)
- Expand hook status to query resource health [\#2515](https://github.com/argoproj/argo-cd/issues/2515)
- Surface ControllerRevision revisions in UI [\#2306](https://github.com/argoproj/argo-cd/issues/2306)
- ArgoCD Namespace Isolation [\#2176](https://github.com/argoproj/argo-cd/issues/2176)
- feat: Displays controllerrevsion's revision in the UI. Closes \#2306 [\#2702](https://github.com/argoproj/argo-cd/pull/2702) ([alexec](https://github.com/alexec))
- Improves the "Sync Apps" panel UI. Closes \#2698 [\#2699](https://github.com/argoproj/argo-cd/pull/2699) ([alexec](https://github.com/alexec))
- Updates UI icons. Closes \#2625 and \#2757 [\#2653](https://github.com/argoproj/argo-cd/pull/2653) ([alexec](https://github.com/alexec))
- UI fixes for "Sync Apps" panel. [\#2604](https://github.com/argoproj/argo-cd/pull/2604) ([alexec](https://github.com/alexec))

**Fixed bugs:**

- Timeouts / Possible memory issues? [\#2957](https://github.com/argoproj/argo-cd/issues/2957)
- Autosync without pruning seems to cause CPU spikes [\#2941](https://github.com/argoproj/argo-cd/issues/2941)
- Failed to apply normalization: jsonpatch remove operation does not apply: doc is missing path [\#2933](https://github.com/argoproj/argo-cd/issues/2933)
- The 'repo add' command incorrectly requires upsert flag [\#2925](https://github.com/argoproj/argo-cd/issues/2925)
- Unable to ignore differences in metadata annotations [\#2918](https://github.com/argoproj/argo-cd/issues/2918)
- Application health status is Running while no resources are in Running state [\#2908](https://github.com/argoproj/argo-cd/issues/2908)
- Dex client secret not masked in logs for Microsoft [\#2904](https://github.com/argoproj/argo-cd/issues/2904)
- \[UI\] - commit message shown as \(m.message} on mouse hover [\#2896](https://github.com/argoproj/argo-cd/issues/2896)
- Trying to move from one application's manifests to another with no downtime on resources [\#2885](https://github.com/argoproj/argo-cd/issues/2885)
- \[cli\] Cannot create helm app from cmdline due to no way to specify chart-version [\#2884](https://github.com/argoproj/argo-cd/issues/2884)
- \[CLI\] misleading docs for `--helm-set` [\#2883](https://github.com/argoproj/argo-cd/issues/2883)
- RevisionMetadata too expensive to be called against all commits in history [\#2878](https://github.com/argoproj/argo-cd/issues/2878)
- IgnoreDifferences doesn't work on docker registry secrets [\#2875](https://github.com/argoproj/argo-cd/issues/2875)
- Webhook does not work when ssh URL have a custom port [\#2866](https://github.com/argoproj/argo-cd/issues/2866)
- Resource Diff view crashes [\#2863](https://github.com/argoproj/argo-cd/issues/2863)
- Git Submodule support on official documentation site, with no mention of it being v1.4 [\#2854](https://github.com/argoproj/argo-cd/issues/2854)
- Application label/environment variables editing UI is hard to use [\#2853](https://github.com/argoproj/argo-cd/issues/2853)
- argocd app edit fails when EDITOR has arguments [\#2833](https://github.com/argoproj/argo-cd/issues/2833)
- Application history can't be appended and new changes are never applied [\#2790](https://github.com/argoproj/argo-cd/issues/2790)
- How to disable tls for the purpose of own tls termination on tls proxy ? [\#2753](https://github.com/argoproj/argo-cd/issues/2753)
- Test `TestAutoSyncSelfHealEnabled` is failing [\#2739](https://github.com/argoproj/argo-cd/issues/2739)
- application-controller doesn't deal with rm/add same cluster gracefully \(x509 unknown\) [\#2389](https://github.com/argoproj/argo-cd/issues/2389)
- git contention leads applications into Unknown state [\#2351](https://github.com/argoproj/argo-cd/issues/2351)
- Controller panics due to nil pointer error [\#2290](https://github.com/argoproj/argo-cd/issues/2290)
- Fix bug where manifests are not cached. Fixes \#2770 [\#2771](https://github.com/argoproj/argo-cd/pull/2771) ([alexec](https://github.com/alexec))
- Fixes bug whereby retry does not work for CLI. Fixes \#2767 [\#2768](https://github.com/argoproj/argo-cd/pull/2768) ([alexec](https://github.com/alexec))
- Ensures that Helm charts are correctly resolved before sync. Fixes \#2758 [\#2760](https://github.com/argoproj/argo-cd/pull/2760) ([alexec](https://github.com/alexec))
- Make BeforeHookCreation the default. Fixes \#2754 [\#2759](https://github.com/argoproj/argo-cd/pull/2759) ([alexec](https://github.com/alexec))
- Allow dot in project policy. Closes \#2724 [\#2755](https://github.com/argoproj/argo-cd/pull/2755) ([alexec](https://github.com/alexec))
- Make directory enforcer more lenient and add flag [\#2716](https://github.com/argoproj/argo-cd/pull/2716) ([simster7](https://github.com/simster7))
- Adds support for /api/v1/account\* via HTTP. Fixes \#2664 [\#2701](https://github.com/argoproj/argo-cd/pull/2701) ([alexec](https://github.com/alexec))
- Revert "Use Kustomize 3 to generate manifetsts. Closes \#2487 \(\#2510\)" [\#2696](https://github.com/argoproj/argo-cd/pull/2696) ([alexec](https://github.com/alexec))
- Bump min client cache version [\#2619](https://github.com/argoproj/argo-cd/pull/2619) ([alexmt](https://github.com/alexmt))

**Closed issues:**

- `argocd cluster get` throws error [\#2860](https://github.com/argoproj/argo-cd/issues/2860)

**Merged pull requests:**

- Sort imports in controller/sync\_hooks.go and run 'dep ensure' [\#2972](https://github.com/argoproj/argo-cd/pull/2972) ([alexmt](https://github.com/alexmt))
- chore: remove unnecessary gh-pages filter from circle config [\#2966](https://github.com/argoproj/argo-cd/pull/2966) ([alexmt](https://github.com/alexmt))
- chore: add github action that publish docs changes [\#2965](https://github.com/argoproj/argo-cd/pull/2965) ([alexmt](https://github.com/alexmt))
- fix: rename cluster prometheus metrics according to the naming convention [\#2964](https://github.com/argoproj/argo-cd/pull/2964) ([alexmt](https://github.com/alexmt))
- fix: app create --upsert should take finalizers into account [\#2963](https://github.com/argoproj/argo-cd/pull/2963) ([alexmt](https://github.com/alexmt))
- fix: application upsert should retry on concurrent modification error [\#2959](https://github.com/argoproj/argo-cd/pull/2959) ([alexmt](https://github.com/alexmt))
- fix: sort conditions to avoid reconciliation loop [\#2955](https://github.com/argoproj/argo-cd/pull/2955) ([alexmt](https://github.com/alexmt))
- docs: Add MLB to users list! [\#2951](https://github.com/argoproj/argo-cd/pull/2951) ([mgoodness](https://github.com/mgoodness))
- fix: self-heal should retry syncing an application after specified delay [\#2950](https://github.com/argoproj/argo-cd/pull/2950) ([alexmt](https://github.com/alexmt))
- feat: Add support for ssh-with-port repo url \(\#2866\) [\#2948](https://github.com/argoproj/argo-cd/pull/2948) ([whs](https://github.com/whs))
- fix: limit number of parallel kubectl apply [\#2944](https://github.com/argoproj/argo-cd/pull/2944) ([alexmt](https://github.com/alexmt))
- fix: argocd diff should use NormalizedLiveState [\#2943](https://github.com/argoproj/argo-cd/pull/2943) ([alexmt](https://github.com/alexmt))
- fix: diff local ignore kustomize build options [\#2942](https://github.com/argoproj/argo-cd/pull/2942) ([alexmt](https://github.com/alexmt))
- fix: property retry clusters secret watch [\#2940](https://github.com/argoproj/argo-cd/pull/2940) ([alexmt](https://github.com/alexmt))
- fix: update argocd-util import was not working properly [\#2939](https://github.com/argoproj/argo-cd/pull/2939) ([jessesuen](https://github.com/jessesuen))
- fix: use resource health for hook status evaluation [\#2938](https://github.com/argoproj/argo-cd/pull/2938) ([alexmt](https://github.com/alexmt))
- fix: stop logging dex config secrets \#\(2904\) [\#2937](https://github.com/argoproj/argo-cd/pull/2937) ([alexmt](https://github.com/alexmt))
- fix: remove unnecessary scroll from filter panel [\#2935](https://github.com/argoproj/argo-cd/pull/2935) ([alexmt](https://github.com/alexmt))
- fix: status badge should display 'Not Found' if application does not exist [\#2927](https://github.com/argoproj/argo-cd/pull/2927) ([alexmt](https://github.com/alexmt))
- fix: the 'repo add' command incorrectly requires upsert flag [\#2926](https://github.com/argoproj/argo-cd/pull/2926) ([alexmt](https://github.com/alexmt))
- fix: collect cluster metrics in background; remove obsolete metrics [\#2923](https://github.com/argoproj/argo-cd/pull/2923) ([alexmt](https://github.com/alexmt))
- docs: correct wordings [\#2916](https://github.com/argoproj/argo-cd/pull/2916) ([SHLo](https://github.com/SHLo))
- fix: cmd line docs [\#2912](https://github.com/argoproj/argo-cd/pull/2912) ([davidkarlsen](https://github.com/davidkarlsen))
- chore: a few lint fixes [\#2911](https://github.com/argoproj/argo-cd/pull/2911) ([davidkarlsen](https://github.com/davidkarlsen))
- fix: prevent user from seeing/deleting resources not permitted in project [\#2910](https://github.com/argoproj/argo-cd/pull/2910) ([alexmt](https://github.com/alexmt))
- fix: add missing networking.k8s.io/Ingress health check \(\#2908\) [\#2909](https://github.com/argoproj/argo-cd/pull/2909) ([alexmt](https://github.com/alexmt))
- Add documentation for Azure AD SSO configuration [\#2905](https://github.com/argoproj/argo-cd/pull/2905) ([jrgirvan](https://github.com/jrgirvan))
- Add Peloton Interactive as user to Readme [\#2901](https://github.com/argoproj/argo-cd/pull/2901) ([Lixxia](https://github.com/Lixxia))
- fix: commit is rendered as undefinied is app revision is not specified explicitly [\#2899](https://github.com/argoproj/argo-cd/pull/2899) ([alexmt](https://github.com/alexmt))
- fix: removes redundant mutex usage in controller cache and adds cluster cache metrics [\#2898](https://github.com/argoproj/argo-cd/pull/2898) ([alexmt](https://github.com/alexmt))
- refactor: Export autocomplete to argo-ui [\#2897](https://github.com/argoproj/argo-cd/pull/2897) ([simster7](https://github.com/simster7))
- fix: Improve cli-docs. Fixes issue \#2884 [\#2895](https://github.com/argoproj/argo-cd/pull/2895) ([davidkarlsen](https://github.com/davidkarlsen))
- fix: expand/collapse annotations on application details page [\#2894](https://github.com/argoproj/argo-cd/pull/2894) ([alexmt](https://github.com/alexmt))
- fix: add missing git checkout in GetRevisionMetadata method [\#2893](https://github.com/argoproj/argo-cd/pull/2893) ([alexmt](https://github.com/alexmt))
- fix: UI should re-trigger SSO login if SSO JWT token expires [\#2891](https://github.com/argoproj/argo-cd/pull/2891) ([alexmt](https://github.com/alexmt))
- fix: UI should cache version info to avoid having loading indicator on every navigation [\#2889](https://github.com/argoproj/argo-cd/pull/2889) ([alexmt](https://github.com/alexmt))
- Fix links in documentation  [\#2888](https://github.com/argoproj/argo-cd/pull/2888) ([OmerKahani](https://github.com/OmerKahani))
- refactor: upgrade argo ui [\#2887](https://github.com/argoproj/argo-cd/pull/2887) ([alexmt](https://github.com/alexmt))
- fix: prevent loading commit metadata is it is missing in sync results [\#2886](https://github.com/argoproj/argo-cd/pull/2886) ([alexmt](https://github.com/alexmt))
- Add tooltip to version info [\#2882](https://github.com/argoproj/argo-cd/pull/2882) ([alexmt](https://github.com/alexmt))
- fix: stop loading history metadata for the whole deployment history \(\#2878\) [\#2880](https://github.com/argoproj/argo-cd/pull/2880) ([alexmt](https://github.com/alexmt))
- fix: clarify cluster cli command arguments to avoid confusion [\#2879](https://github.com/argoproj/argo-cd/pull/2879) ([alexmt](https://github.com/alexmt))
- fix: git contention leads applications into Unknown state \(\#2351\) [\#2877](https://github.com/argoproj/argo-cd/pull/2877) ([alexmt](https://github.com/alexmt))
- fix: Fix flaky TestAutoSyncSelfHealEnabled test [\#2876](https://github.com/argoproj/argo-cd/pull/2876) ([alexmt](https://github.com/alexmt))
- Add Syncier as user of Argo CD [\#2872](https://github.com/argoproj/argo-cd/pull/2872) ([torstenwalter](https://github.com/torstenwalter))
- docs: fix typo in webhook doc [\#2870](https://github.com/argoproj/argo-cd/pull/2870) ([qingbo](https://github.com/qingbo))
- fix: stop using jsondiffpatch on clientside to render resource difference  [\#2869](https://github.com/argoproj/argo-cd/pull/2869) ([alexmt](https://github.com/alexmt))
- Update README.md [\#2868](https://github.com/argoproj/argo-cd/pull/2868) ([vdinesh2461990](https://github.com/vdinesh2461990))
- Added our company \(ARZ\) to the official users [\#2867](https://github.com/argoproj/argo-cd/pull/2867) ([jkleinlercher](https://github.com/jkleinlercher))
- feat: Use kubectl apply library instead of forking binary [\#2861](https://github.com/argoproj/argo-cd/pull/2861) ([jessesuen](https://github.com/jessesuen))
- Grafana Dashboard: Template Out Prometheus Data Source [\#2859](https://github.com/argoproj/argo-cd/pull/2859) ([abhishekjiitr](https://github.com/abhishekjiitr))
- docs: add cybozu to list of users [\#2858](https://github.com/argoproj/argo-cd/pull/2858) ([aoi1](https://github.com/aoi1))
- Fix TestAutoSyncSelfHealEnabled test [\#2857](https://github.com/argoproj/argo-cd/pull/2857) ([alexmt](https://github.com/alexmt))
- Issue \#2853 - Improve application env variables/labels editing [\#2856](https://github.com/argoproj/argo-cd/pull/2856) ([alexmt](https://github.com/alexmt))
- docs: Issue: \#2854: Add version information for git submodule support [\#2855](https://github.com/argoproj/argo-cd/pull/2855) ([hongymagic](https://github.com/hongymagic))
- Issue \#2683 - Make sure app update don't fail due to concurrent modification [\#2852](https://github.com/argoproj/argo-cd/pull/2852) ([alexmt](https://github.com/alexmt))
- who uses: biobox analytics [\#2851](https://github.com/argoproj/argo-cd/pull/2851) ([bx-julian](https://github.com/bx-julian))
- Issue \#2833: use editor arguments in InteractiveEditor [\#2850](https://github.com/argoproj/argo-cd/pull/2850) ([cabrinha](https://github.com/cabrinha))
- Issue \#2848 - Application Deployment history panel shows incorrect info for recent releases [\#2849](https://github.com/argoproj/argo-cd/pull/2849) ([alexmt](https://github.com/alexmt))
- \[doc\] Add KSOPS to Secret Management List [\#2846](https://github.com/argoproj/argo-cd/pull/2846) ([devstein](https://github.com/devstein))
- Updates releasing guide. [\#2844](https://github.com/argoproj/argo-cd/pull/2844) ([alexec](https://github.com/alexec))
- Update rollout healthcheck to use pause conditons [\#2842](https://github.com/argoproj/argo-cd/pull/2842) ([dthomson25](https://github.com/dthomson25))
- doc: Add Viaduct to "Who Uses" List [\#2841](https://github.com/argoproj/argo-cd/pull/2841) ([devstein](https://github.com/devstein))
- feat: namespace isolation \(\#2176\) [\#2839](https://github.com/argoproj/argo-cd/pull/2839) ([alexmt](https://github.com/alexmt))
- doc: Add Pipefy to list of current companies using Argocd [\#2838](https://github.com/argoproj/argo-cd/pull/2838) ([yrosaguiar](https://github.com/yrosaguiar))
- doc: add celonis to list of users [\#2836](https://github.com/argoproj/argo-cd/pull/2836) ([muenchdo](https://github.com/muenchdo))
- doc: add Walkbase to list of users [\#2835](https://github.com/argoproj/argo-cd/pull/2835) ([LinAnt](https://github.com/LinAnt))
- doc: add elium to list of users [\#2834](https://github.com/argoproj/argo-cd/pull/2834) ([julienp](https://github.com/julienp))
- doc: add hipages to list of users [\#2832](https://github.com/argoproj/argo-cd/pull/2832) ([estahn](https://github.com/estahn))
- Add Max Kelsen to Who Uses Argo CD [\#2831](https://github.com/argoproj/argo-cd/pull/2831) ([ofaz](https://github.com/ofaz))
- Add Baloise as Argo CD user [\#2830](https://github.com/argoproj/argo-cd/pull/2830) ([niiku](https://github.com/niiku))
- Add Adevinta in README.md [\#2829](https://github.com/argoproj/argo-cd/pull/2829) ([jaimejorge](https://github.com/jaimejorge))
- Update k8s dependency to v1.16 [\#2828](https://github.com/argoproj/argo-cd/pull/2828) ([jessesuen](https://github.com/jessesuen))
- Added EDF Renewables to list of organizations [\#2823](https://github.com/argoproj/argo-cd/pull/2823) ([Kampe](https://github.com/Kampe))
- Add support for hidden directories with directory enforcer [\#2821](https://github.com/argoproj/argo-cd/pull/2821) ([simster7](https://github.com/simster7))
- Make ConvertToVersion maybe 1090% faster on average [\#2820](https://github.com/argoproj/argo-cd/pull/2820) ([alexec](https://github.com/alexec))
- Fixes logging of tracing option in CLI [\#2819](https://github.com/argoproj/argo-cd/pull/2819) ([alexec](https://github.com/alexec))
- Adds revision history limit. Closes \#2790 [\#2818](https://github.com/argoproj/argo-cd/pull/2818) ([alexec](https://github.com/alexec))
- Fixes error in docs that prevents publishing them [\#2817](https://github.com/argoproj/argo-cd/pull/2817) ([alexec](https://github.com/alexec))
- Issue \#2721 Optimize helm repo querying [\#2816](https://github.com/argoproj/argo-cd/pull/2816) ([alexmt](https://github.com/alexmt))
- SAML/Azure integration [\#2815](https://github.com/argoproj/argo-cd/pull/2815) ([zeph](https://github.com/zeph))
- Fix typo in docs [\#2813](https://github.com/argoproj/argo-cd/pull/2813) ([posquit0](https://github.com/posquit0))
- Adds tracing to key external invocations. [\#2811](https://github.com/argoproj/argo-cd/pull/2811) ([alexec](https://github.com/alexec))
- Issue \#2721 - cache parsed repositories, repo credentials to avoid unnecessary yaml parsing [\#2809](https://github.com/argoproj/argo-cd/pull/2809) ([alexmt](https://github.com/alexmt))
- More Detailed argocd CLI Install Instructions for Mac and Linux [\#2807](https://github.com/argoproj/argo-cd/pull/2807) ([danielhelfand](https://github.com/danielhelfand))
- Updates Language and UI Photos in Getting Started Tutorial [\#2801](https://github.com/argoproj/argo-cd/pull/2801) ([danielhelfand](https://github.com/danielhelfand))
- Fix directory traversal edge case and enhance tests. Fixes \#2796 [\#2797](https://github.com/argoproj/argo-cd/pull/2797) ([simster7](https://github.com/simster7))
- Fix incorrectly formatted link [\#2788](https://github.com/argoproj/argo-cd/pull/2788) ([alexmt](https://github.com/alexmt))
- Usage of Argo-cd at Saloodo! GmbH [\#2786](https://github.com/argoproj/argo-cd/pull/2786) ([nitinpatil1992](https://github.com/nitinpatil1992))
- fixed docs with proper keys for secret data fix \#2776 [\#2777](https://github.com/argoproj/argo-cd/pull/2777) ([nouseforaname](https://github.com/nouseforaname))
- docs fix label for ServiceMonitor CRD for API server metrics. Fixes: \#2774 [\#2775](https://github.com/argoproj/argo-cd/pull/2775) ([greenstatic](https://github.com/greenstatic))
- Update README.md to add KubeCon workshop [\#2766](https://github.com/argoproj/argo-cd/pull/2766) ([alexec](https://github.com/alexec))
- Removes log warning regarding indexer and may improve perf. Closes \#1345 [\#2761](https://github.com/argoproj/argo-cd/pull/2761) ([alexec](https://github.com/alexec))
- In docs, corrected version of Argo CD that build env is available in. See \#2749 [\#2750](https://github.com/argoproj/argo-cd/pull/2750) ([alexec](https://github.com/alexec))
- Allow you to sync local Helm apps.  Fixes \#2741 [\#2747](https://github.com/argoproj/argo-cd/pull/2747) ([alexec](https://github.com/alexec))
- Add Helm values file block example [\#2745](https://github.com/argoproj/argo-cd/pull/2745) ([ibiqlik](https://github.com/ibiqlik))
- Fix a bug with cluster add when token secret is not first in list. [\#2744](https://github.com/argoproj/argo-cd/pull/2744) ([dgoodwin](https://github.com/dgoodwin))
- Fix 'Open application' link when using basehref [\#2729](https://github.com/argoproj/argo-cd/pull/2729) ([cbanek](https://github.com/cbanek))
- Shows chart name in apps tiles and apps table pages. Closes \#2726 [\#2728](https://github.com/argoproj/argo-cd/pull/2728) ([alexec](https://github.com/alexec))
- Update CONTRIBUTING.md to add `goreman` [\#2727](https://github.com/argoproj/argo-cd/pull/2727) ([alexec](https://github.com/alexec))
- modify pre-requisites kustomize version \#2717 [\#2725](https://github.com/argoproj/argo-cd/pull/2725) ([oke-py](https://github.com/oke-py))
- Clarify the need for namespace during export [\#2722](https://github.com/argoproj/argo-cd/pull/2722) ([dmaciel](https://github.com/dmaciel))
- Fixed KustomizeBuildOptions not getting synced [\#2714](https://github.com/argoproj/argo-cd/pull/2714) ([simster7](https://github.com/simster7))
- Set X-Frame-Options on serving static assets \(\#2706\) [\#2711](https://github.com/argoproj/argo-cd/pull/2711) ([jannfis](https://github.com/jannfis))
- Remove references to argocd-ui image during manifests generation [\#2710](https://github.com/argoproj/argo-cd/pull/2710) ([olivierlemasle](https://github.com/olivierlemasle))
- Disables TestAutoSyncSelfHealEnabled. [\#2703](https://github.com/argoproj/argo-cd/pull/2703) ([alexec](https://github.com/alexec))
- Restore 'argocd app run action' backward compatibility [\#2700](https://github.com/argoproj/argo-cd/pull/2700) ([alexmt](https://github.com/alexmt))
- Fixes the failing CI master build [\#2697](https://github.com/argoproj/argo-cd/pull/2697) ([alexec](https://github.com/alexec))
- Only delete resources during app delete cascade if permitted to \(fixes \#2693\) [\#2695](https://github.com/argoproj/argo-cd/pull/2695) ([jannfis](https://github.com/jannfis))
- Issue \#2691 - Remove annoying toolbar flickering [\#2692](https://github.com/argoproj/argo-cd/pull/2692) ([alexmt](https://github.com/alexmt))
- Adds support for the custom health checks for the new cert-manager API. Closes \#2667 [\#2689](https://github.com/argoproj/argo-cd/pull/2689) ([binoue](https://github.com/binoue))
- Simplify using Argo CD without users/SSO/UI \(agent mode\) [\#2688](https://github.com/argoproj/argo-cd/pull/2688) ([alexmt](https://github.com/alexmt))
- Rename deprecated deadline option to timeout [\#2686](https://github.com/argoproj/argo-cd/pull/2686) ([oke-py](https://github.com/oke-py))
- modified make target \#2680 [\#2685](https://github.com/argoproj/argo-cd/pull/2685) ([oke-py](https://github.com/oke-py))
- Issue \#2673 - Application controller  flag is broken [\#2674](https://github.com/argoproj/argo-cd/pull/2674) ([alexmt](https://github.com/alexmt))
- Issue \#2670 - API server does not allow creating role with action 'action/\*' [\#2671](https://github.com/argoproj/argo-cd/pull/2671) ([alexmt](https://github.com/alexmt))
- Issue \#2668 - Delete a specified context [\#2669](https://github.com/argoproj/argo-cd/pull/2669) ([masa213f](https://github.com/masa213f))
- Returns a clearer error on invalid Helm version. Closes \#2665 and \#2736 [\#2666](https://github.com/argoproj/argo-cd/pull/2666) ([alexec](https://github.com/alexec))
- Issue \#2662 - Don't parse kustomize version outout [\#2663](https://github.com/argoproj/argo-cd/pull/2663) ([alexmt](https://github.com/alexmt))
- Adds support for testing .tsx files. Closes \#2610 [\#2661](https://github.com/argoproj/argo-cd/pull/2661) ([alexec](https://github.com/alexec))
- Issue \#2659 - Fix 1.3 login regressions [\#2660](https://github.com/argoproj/argo-cd/pull/2660) ([alexmt](https://github.com/alexmt))
- Issue \#2559 - Add gauge Prometheus metric which represents the number of pending manifest requests. [\#2658](https://github.com/argoproj/argo-cd/pull/2658) ([alexmt](https://github.com/alexmt))
- Issue \#2655 - Application list page is not updated automatically [\#2656](https://github.com/argoproj/argo-cd/pull/2656) ([alexmt](https://github.com/alexmt))
- Issue \#2645 - /api/version should not fail if unable to load tool version [\#2654](https://github.com/argoproj/argo-cd/pull/2654) ([alexmt](https://github.com/alexmt))
- Adds note to docs that build enviroment is available in v1.3. See \#2651 [\#2652](https://github.com/argoproj/argo-cd/pull/2652) ([alexec](https://github.com/alexec))
- Modify docs for ingress ssl passthrough and SSO [\#2649](https://github.com/argoproj/argo-cd/pull/2649) ([alantang888](https://github.com/alantang888))
- Revert "Fix Typo: Filter Label To Get Server Pod Name \(\#2597\)" [\#2637](https://github.com/argoproj/argo-cd/pull/2637) ([yujunz](https://github.com/yujunz))
- Issue \#2635 - Custom actions are disabled in Argo CD UI [\#2636](https://github.com/argoproj/argo-cd/pull/2636) ([alexmt](https://github.com/alexmt))
- Issue \#2633 - Application list page incorrectly filter apps by label selector [\#2634](https://github.com/argoproj/argo-cd/pull/2634) ([alexmt](https://github.com/alexmt))
- Issue \#2592 - Remove transitive dependency on packr [\#2631](https://github.com/argoproj/argo-cd/pull/2631) ([alexmt](https://github.com/alexmt))
- Update README.md [\#2630](https://github.com/argoproj/argo-cd/pull/2630) ([saradhis](https://github.com/saradhis))
- Certmanager docs [\#2629](https://github.com/argoproj/argo-cd/pull/2629) ([genebean](https://github.com/genebean))
- Assume git as default repository type \(fixes \#2622\) [\#2628](https://github.com/argoproj/argo-cd/pull/2628) ([jannfis](https://github.com/jannfis))
- Issue \#2626 - Repo server executes unnecessary ls-remotes [\#2627](https://github.com/argoproj/argo-cd/pull/2627) ([alexmt](https://github.com/alexmt))
- Issue \#2620 - Cluster list page fails if any cluster is not reachable [\#2621](https://github.com/argoproj/argo-cd/pull/2621) ([alexmt](https://github.com/alexmt))
- Issue \#2616 - argocd app diff prints only first difference [\#2617](https://github.com/argoproj/argo-cd/pull/2617) ([alexmt](https://github.com/alexmt))
- argocd-util should allow editing project policies in bulk  [\#2615](https://github.com/argoproj/argo-cd/pull/2615) ([alexmt](https://github.com/alexmt))
- Allow '@'-character in SSH usernames when connecting a repository [\#2612](https://github.com/argoproj/argo-cd/pull/2612) ([jannfis](https://github.com/jannfis))
- Upgrade kustomize to 3.2.1 \#2609 [\#2607](https://github.com/argoproj/argo-cd/pull/2607) ([hongymagic](https://github.com/hongymagic))
- Adds support for Helm charts to be a semver range. Closes \#2552 [\#2606](https://github.com/argoproj/argo-cd/pull/2606) ([alexec](https://github.com/alexec))
- Executes app label filtering on client side [\#2605](https://github.com/argoproj/argo-cd/pull/2605) ([alexmt](https://github.com/alexmt))
- Adds a status icon for the op to the UI. Closes \#2596 [\#2601](https://github.com/argoproj/argo-cd/pull/2601) ([alexec](https://github.com/alexec))
- correct the spelling of hashicorp [\#2599](https://github.com/argoproj/argo-cd/pull/2599) ([dcherman](https://github.com/dcherman))
- Update secret-management doc [\#2598](https://github.com/argoproj/argo-cd/pull/2598) ([niqdev](https://github.com/niqdev))
- Fix Typo: Filter Label To Get Server Pod Name [\#2597](https://github.com/argoproj/argo-cd/pull/2597) ([abhishekjiitr](https://github.com/abhishekjiitr))
- Fix Typo in Docs [\#2595](https://github.com/argoproj/argo-cd/pull/2595) ([abhishekjiitr](https://github.com/abhishekjiitr))
- Upgrade Helm to v2.15.2. Closes \#2587 [\#2590](https://github.com/argoproj/argo-cd/pull/2590) ([alexec](https://github.com/alexec))
- Fix date in CHANGELOG.md [\#2584](https://github.com/argoproj/argo-cd/pull/2584) ([mnaser](https://github.com/mnaser))
- docs\(app cmd\): updating --local text to include helm and kustomize [\#2582](https://github.com/argoproj/argo-cd/pull/2582) ([igaskin](https://github.com/igaskin))
- Issue \#2339 - Don't update 'status.reconciledAt' unless compared with latest git version [\#2581](https://github.com/argoproj/argo-cd/pull/2581) ([alexmt](https://github.com/alexmt))
- Add AnalysisRun and Experiment HealthCheck [\#2579](https://github.com/argoproj/argo-cd/pull/2579) ([dthomson25](https://github.com/dthomson25))
- Sets app status to unknown if there is an error. Closes \#2577 [\#2578](https://github.com/argoproj/argo-cd/pull/2578) ([alexec](https://github.com/alexec))
- Adds timeout to Helm commands. [\#2570](https://github.com/argoproj/argo-cd/pull/2570) ([alexec](https://github.com/alexec))
- Allows Helm charts to have arbitrary file names. Fixes \#2549 [\#2569](https://github.com/argoproj/argo-cd/pull/2569) ([alexec](https://github.com/alexec))
- Fixes panic when creating repo. Closes \#2567 [\#2568](https://github.com/argoproj/argo-cd/pull/2568) ([alexec](https://github.com/alexec))
- Display app conditions timestamp in CLI and UI. \(\#737\) [\#2565](https://github.com/argoproj/argo-cd/pull/2565) ([dgoodwin](https://github.com/dgoodwin))
- Fixes blocking SSO login via CLI \(\#2563\) [\#2564](https://github.com/argoproj/argo-cd/pull/2564) ([masa213f](https://github.com/masa213f))
- Reduces total build time from 20m to 14m [\#2560](https://github.com/argoproj/argo-cd/pull/2560) ([alexec](https://github.com/alexec))
- Adds the option to output in YAML and JSON to several CLI commands. Closes \#2534 and closes \#2534 [\#2551](https://github.com/argoproj/argo-cd/pull/2551) ([jannfis](https://github.com/jannfis))
- Adds doc on secrets, updates changelog, updates Github templates [\#2550](https://github.com/argoproj/argo-cd/pull/2550) ([alexec](https://github.com/alexec))
- Unknown child app should not affect app health [\#2544](https://github.com/argoproj/argo-cd/pull/2544) ([alexmt](https://github.com/alexmt))
- Issue \#2339 - Controller should compare with latest git revision if app has changed [\#2543](https://github.com/argoproj/argo-cd/pull/2543) ([alexmt](https://github.com/alexmt))
- Tidies up naming of variables in code.  [\#2541](https://github.com/argoproj/argo-cd/pull/2541) ([alexec](https://github.com/alexec))
- Redact secrets in dex logs [\#2538](https://github.com/argoproj/argo-cd/pull/2538) ([simster7](https://github.com/simster7))
- Added audit events for actions [\#2529](https://github.com/argoproj/argo-cd/pull/2529) ([simster7](https://github.com/simster7))
- Add CARFAX to list of Argo CD users [\#2527](https://github.com/argoproj/argo-cd/pull/2527) ([theeltahir](https://github.com/theeltahir))
- Fix `argocd app sync app-name --resource` not waiting for operation to finish [\#2526](https://github.com/argoproj/argo-cd/pull/2526) ([simster7](https://github.com/simster7))
- Allows Helm parameters that contains arrays or maps.  [\#2525](https://github.com/argoproj/argo-cd/pull/2525) ([alexec](https://github.com/alexec))
- Enable prettier on UI source code [\#2524](https://github.com/argoproj/argo-cd/pull/2524) ([alexec](https://github.com/alexec))
- Adds `argocd auth can-i` command. Close \#2255 [\#2521](https://github.com/argoproj/argo-cd/pull/2521) ([alexec](https://github.com/alexec))
- Upgrade casbin/casbin to 1.9.1 [\#2517](https://github.com/argoproj/argo-cd/pull/2517) ([leominov](https://github.com/leominov))
- Update README.md added Universidad Mesoamericana to organization list [\#2514](https://github.com/argoproj/argo-cd/pull/2514) ([mgaroz](https://github.com/mgaroz))
- Update codecov to only complain at 2% drop [\#2513](https://github.com/argoproj/argo-cd/pull/2513) ([alexec](https://github.com/alexec))
- Use the same tools for `make image` to `make dev-tools-image`. [\#2511](https://github.com/argoproj/argo-cd/pull/2511) ([alexec](https://github.com/alexec))
- Use Kustomize 3 to generate manifests.  [\#2510](https://github.com/argoproj/argo-cd/pull/2510) ([alexec](https://github.com/alexec))
- Documentation update: Fix typo: grcp -\> grpc in FAQ [\#2509](https://github.com/argoproj/argo-cd/pull/2509) ([jannfis](https://github.com/jannfis))
- Update CHANGELOG.md to list v1.3.0-rc1 [\#2504](https://github.com/argoproj/argo-cd/pull/2504) ([alexec](https://github.com/alexec))
- Added SSO and RBAC [\#2503](https://github.com/argoproj/argo-cd/pull/2503) ([stgarf](https://github.com/stgarf))
- Shows version in UI.  [\#2502](https://github.com/argoproj/argo-cd/pull/2502) ([alexec](https://github.com/alexec))
- Fixes a bug where app kind was not show in UI YAML editor [\#2501](https://github.com/argoproj/argo-cd/pull/2501) ([alexec](https://github.com/alexec))
- Set cookie policy to SameSite=lax and httpOnly \(partial mitigation for \#2496\) [\#2498](https://github.com/argoproj/argo-cd/pull/2498) ([jannfis](https://github.com/jannfis))
- add git submodule support [\#2495](https://github.com/argoproj/argo-cd/pull/2495) ([adamjohnson01](https://github.com/adamjohnson01))
- add namesuffix for kustomize applications [\#2473](https://github.com/argoproj/argo-cd/pull/2473) ([jeffmhastings](https://github.com/jeffmhastings))
- Adds the ability to work with groups of apps using labels [\#2463](https://github.com/argoproj/argo-cd/pull/2463) ([alexec](https://github.com/alexec))
- Add Time to ApplicationCondition. [\#2417](https://github.com/argoproj/argo-cd/pull/2417) ([dgoodwin](https://github.com/dgoodwin))
- Adds support for ARGO\_CD\_\[TARGET\_REVISION|REVISION\] and pass to Custom Tool/Helm/Jsonnet [\#2415](https://github.com/argoproj/argo-cd/pull/2415) ([alexec](https://github.com/alexec))
- Makes cache timeouts configurable [\#2412](https://github.com/argoproj/argo-cd/pull/2412) ([alexec](https://github.com/alexec))
- Use vars for service name reference in commands [\#2408](https://github.com/argoproj/argo-cd/pull/2408) ([imranismail](https://github.com/imranismail))
- Add repository credential management API and CLI \(addresses \#2136\) [\#2207](https://github.com/argoproj/argo-cd/pull/2207) ([jannfis](https://github.com/jannfis))

## [v1.3.6](https://github.com/argoproj/argo-cd/tree/v1.3.6) (2019-12-10)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.3.5...v1.3.6)

**Fixed bugs:**

- Application Deployment history panel shows incorrect info for recent releases [\#2848](https://github.com/argoproj/argo-cd/issues/2848)
- 1.3.0-rc4 ui editing helm parameters are reset during user input [\#2683](https://github.com/argoproj/argo-cd/issues/2683)

## [v1.3.5](https://github.com/argoproj/argo-cd/tree/v1.3.5) (2019-12-09)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.3.4...v1.3.5)

**Implemented enhancements:**

- More Detailed Installation Instructions for CLI [\#2800](https://github.com/argoproj/argo-cd/issues/2800)
- Provide a progress meter for the Sync App panel [\#2698](https://github.com/argoproj/argo-cd/issues/2698)

**Fixed bugs:**

- UI Error - History & Rollback button - v1.3.0-rc5 [\#2694](https://github.com/argoproj/argo-cd/issues/2694)
- Project destination restrictions not respected during cascade deleting an application [\#2693](https://github.com/argoproj/argo-cd/issues/2693)

## [v1.3.4](https://github.com/argoproj/argo-cd/tree/v1.3.4) (2019-12-05)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.3.3...v1.3.4)

## [v1.3.3](https://github.com/argoproj/argo-cd/tree/v1.3.3) (2019-12-05)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.3.2...v1.3.3)

**Implemented enhancements:**

- argocd-util should allow editing project policies in bulk [\#2614](https://github.com/argoproj/argo-cd/issues/2614)

**Fixed bugs:**

- High CPU utilisation \(5 cores\) and spammy logs [\#2721](https://github.com/argoproj/argo-cd/issues/2721)

## [v1.3.2](https://github.com/argoproj/argo-cd/tree/v1.3.2) (2019-12-03)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.3.1...v1.3.2)

**Implemented enhancements:**

- Support for the custom health checks for the new cert-manager API [\#2667](https://github.com/argoproj/argo-cd/issues/2667)

**Fixed bugs:**

- Error: index out of range when Chart.yaml is at root of git repo [\#2796](https://github.com/argoproj/argo-cd/issues/2796)
- Login UI shows "Login" when already logged in [\#2781](https://github.com/argoproj/argo-cd/issues/2781)
- "cluster add" can break due to assumption about first Secret type in a ServiceAccounts list [\#2743](https://github.com/argoproj/argo-cd/issues/2743)
- Unable to apply cluster-install manifest using branch `release-1.3` but `v1.3.0-rc4` works [\#2690](https://github.com/argoproj/argo-cd/issues/2690)

## [v1.3.1](https://github.com/argoproj/argo-cd/tree/v1.3.1) (2019-12-02)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.3.0...v1.3.1)

**Implemented enhancements:**

- Is there a way to create an environment on the fly? [\#2740](https://github.com/argoproj/argo-cd/issues/2740)
- How to get bearerToken for connection to external k8s Cluster [\#2735](https://github.com/argoproj/argo-cd/issues/2735)
- Helm value files outside of chart path [\#2715](https://github.com/argoproj/argo-cd/issues/2715)
- Support Jsonnet Jpaths  [\#2679](https://github.com/argoproj/argo-cd/issues/2679)
- Documentation about best practices [\#2648](https://github.com/argoproj/argo-cd/issues/2648)
- Color or shape code for  K8s object types [\#2583](https://github.com/argoproj/argo-cd/issues/2583)
- How does arcocd support patching of existing resources [\#2437](https://github.com/argoproj/argo-cd/issues/2437)
- Omitting repoURL for Applications in the same repository [\#2171](https://github.com/argoproj/argo-cd/issues/2171)
- Client retries between internal services [\#959](https://github.com/argoproj/argo-cd/issues/959)

**Fixed bugs:**

- Secret substitution in OIDC config may not work. [\#2779](https://github.com/argoproj/argo-cd/issues/2779)
- Webhook Documentation [\#2776](https://github.com/argoproj/argo-cd/issues/2776)
- Always cache miss for manifests [\#2770](https://github.com/argoproj/argo-cd/issues/2770)
- Fix bug whereby retry does not work for CLI [\#2767](https://github.com/argoproj/argo-cd/issues/2767)
- Issues with github authentication with Argo-ha setup [\#2763](https://github.com/argoproj/argo-cd/issues/2763)
- Regression: Can not supply value files as URLs [\#2762](https://github.com/argoproj/argo-cd/issues/2762)
- We might not always resolve revision for Helm when we need to [\#2758](https://github.com/argoproj/argo-cd/issues/2758)
- BeforeHookCreation should be the default hook [\#2754](https://github.com/argoproj/argo-cd/issues/2754)
- Build environment variables not working in helm [\#2749](https://github.com/argoproj/argo-cd/issues/2749)
- argocd local sync cannot parse kubernetes version [\#2741](https://github.com/argoproj/argo-cd/issues/2741)
- Target Revision truncated [\#2736](https://github.com/argoproj/argo-cd/issues/2736)
- App list does not show chart for Helm app [\#2726](https://github.com/argoproj/argo-cd/issues/2726)
- Can't use `DNS-1123` compliant app name when creating project role [\#2724](https://github.com/argoproj/argo-cd/issues/2724)
- Contribution pre-requisites requires invalid version of kustomize [\#2717](https://github.com/argoproj/argo-cd/issues/2717)
- argocd app create fails with status.sync.comparedTo.source.path in body is required [\#2712](https://github.com/argoproj/argo-cd/issues/2712)
- Wrong installation commands in GitHub release notes v1.3.0 [\#2709](https://github.com/argoproj/argo-cd/issues/2709)
- Cert-manager:  [\#2684](https://github.com/argoproj/argo-cd/issues/2684)
- 1.3.0-rc4 Unable to provide helm valuesFiles from a git repo? [\#2682](https://github.com/argoproj/argo-cd/issues/2682)
- No error on `argocd app create` in  CLI if `--revision` is omitted [\#2665](https://github.com/argoproj/argo-cd/issues/2665)
- update account password from API resulted 404 [\#2664](https://github.com/argoproj/argo-cd/issues/2664)
- Helm cannot use build environment variables [\#2651](https://github.com/argoproj/argo-cd/issues/2651)
- Not all values from custom values files are applied [\#2562](https://github.com/argoproj/argo-cd/issues/2562)
- Argo CD fails to sync a specific custom resource \(Calico GlobalNetworkPolicy\) [\#2505](https://github.com/argoproj/argo-cd/issues/2505)
- Argocd v1.2.2 not starting! [\#2358](https://github.com/argoproj/argo-cd/issues/2358)
- AutoSync doesn't work anymore [\#2339](https://github.com/argoproj/argo-cd/issues/2339)
- Inconsistent sync result from UI and CLI [\#2321](https://github.com/argoproj/argo-cd/issues/2321)
- "invalid cookie, longer than max length 4093" [\#2165](https://github.com/argoproj/argo-cd/issues/2165)

**Closed issues:**

- Docs metrics wrong selector for API Server Metrics [\#2774](https://github.com/argoproj/argo-cd/issues/2774)
- How to configure ignoreDifferences for replica count? [\#2748](https://github.com/argoproj/argo-cd/issues/2748)
- We're rather busy with KubeCon at the moment [\#2704](https://github.com/argoproj/argo-cd/issues/2704)
- argocd-application-controller: can not retrieve list of objects using index : Index with name namespace does not exist [\#1345](https://github.com/argoproj/argo-cd/issues/1345)

## [v1.3.0](https://github.com/argoproj/argo-cd/tree/v1.3.0) (2019-11-13)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.3.0-rc5...v1.3.0)

**Implemented enhancements:**

- From scratch / distroless containers [\#2676](https://github.com/argoproj/argo-cd/issues/2676)
- Update tooling to use Kustomize v3 [\#2487](https://github.com/argoproj/argo-cd/issues/2487)

**Fixed bugs:**

- Annoying toolbar flickering [\#2691](https://github.com/argoproj/argo-cd/issues/2691)
- Running `make lint` fails [\#2687](https://github.com/argoproj/argo-cd/issues/2687)
- CLI - Can not delete a specified context [\#2668](https://github.com/argoproj/argo-cd/issues/2668)
- namespace not respected when purge-syncing a project via app.kubernetes.io/instance [\#2650](https://github.com/argoproj/argo-cd/issues/2650)

## [v1.3.0-rc5](https://github.com/argoproj/argo-cd/tree/v1.3.0-rc5) (2019-11-11)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.3.0-rc4...v1.3.0-rc5)

**Implemented enhancements:**

- Display wave and image in UI [\#2643](https://github.com/argoproj/argo-cd/issues/2643)
- Tech Debt: Configure testing of `\*.tsx` files [\#2610](https://github.com/argoproj/argo-cd/issues/2610)
- Upgrade kustomize from v3.1.0 to v3.2.1 [\#2609](https://github.com/argoproj/argo-cd/issues/2609)
- Make `SettingsManager ` mockable [\#2608](https://github.com/argoproj/argo-cd/issues/2608)
- Argo CD in the DigitalOcean Marketplace [\#2585](https://github.com/argoproj/argo-cd/issues/2585)
- Make App diff configMapGenerator aware [\#2397](https://github.com/argoproj/argo-cd/issues/2397)
- Application conditions should have timestamp [\#737](https://github.com/argoproj/argo-cd/issues/737)

**Fixed bugs:**

- 1.3.0-rc4 ui oauth flow only completes after second attempt [\#2681](https://github.com/argoproj/argo-cd/issues/2681)
- No rule to make target `dev-builder-image' [\#2680](https://github.com/argoproj/argo-cd/issues/2680)
- Application controller `kubectl-parallelism-limit` flag is broken [\#2673](https://github.com/argoproj/argo-cd/issues/2673)
- API server does not allow creating role with action `action/\*` [\#2670](https://github.com/argoproj/argo-cd/issues/2670)
- Regression: Cannot return Kustomize version for 3.1.0 [\#2662](https://github.com/argoproj/argo-cd/issues/2662)
- Login regression issues [\#2659](https://github.com/argoproj/argo-cd/issues/2659)
- Regression: Cannot tab out of path for new app [\#2657](https://github.com/argoproj/argo-cd/issues/2657)
- Application list page is not updated automatically anymore [\#2655](https://github.com/argoproj/argo-cd/issues/2655)
- Unable to add AWS CodeCommit repository from argocd cli [\#2646](https://github.com/argoproj/argo-cd/issues/2646)
- Failure of `argocd version` in the self-building container image [\#2645](https://github.com/argoproj/argo-cd/issues/2645)
- Incorrect Helm version appears in Docker image [\#2641](https://github.com/argoproj/argo-cd/issues/2641)
- Old ReplicaSets are not deleted when sync completes [\#2639](https://github.com/argoproj/argo-cd/issues/2639)
- Custom actions are disabled in Argo CD UI [\#2635](https://github.com/argoproj/argo-cd/issues/2635)
- Application list page incorrectly filter apps by label selector [\#2633](https://github.com/argoproj/argo-cd/issues/2633)
- 1.3.0-rc3 ui crash when changing project [\#2632](https://github.com/argoproj/argo-cd/issues/2632)
- Applications randomly becoming "OutOfSync" [\#2613](https://github.com/argoproj/argo-cd/issues/2613)
- Remove transitive dependency on packr [\#2592](https://github.com/argoproj/argo-cd/issues/2592)
- diff customization for cert-manager doesn't work [\#2586](https://github.com/argoproj/argo-cd/issues/2586)
- Variable substitution not working as expected [\#2539](https://github.com/argoproj/argo-cd/issues/2539)

## [v1.3.0-rc4](https://github.com/argoproj/argo-cd/tree/v1.3.0-rc4) (2019-11-04)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.3.0-rc3...v1.3.0-rc4)

**Implemented enhancements:**

- Use `helm template --is-upgrade` [\#2602](https://github.com/argoproj/argo-cd/issues/2602)
- Update Helm to latest version [\#2594](https://github.com/argoproj/argo-cd/issues/2594)
- Clean-up and generify `health.go` [\#2588](https://github.com/argoproj/argo-cd/issues/2588)
- Upgrade Helm  [\#2587](https://github.com/argoproj/argo-cd/issues/2587)
- WiP: Generalized Plugins [\#2571](https://github.com/argoproj/argo-cd/issues/2571)
- Helm Chart update available [\#2552](https://github.com/argoproj/argo-cd/issues/2552)
- Provide consistent output formatting options in the CLI [\#2534](https://github.com/argoproj/argo-cd/issues/2534)
- We need rollback beyond 10 deployment [\#2285](https://github.com/argoproj/argo-cd/issues/2285)
- Can I? [\#2255](https://github.com/argoproj/argo-cd/issues/2255)
- Better support for json and yaml outputs for cli [\#1986](https://github.com/argoproj/argo-cd/issues/1986)

**Fixed bugs:**

- Repo server executes unnecessary ls-remotes [\#2626](https://github.com/argoproj/argo-cd/issues/2626)
- Repository type should be mandatory for repo add command in CLI [\#2622](https://github.com/argoproj/argo-cd/issues/2622)
- Cluster list page fails if any cluster is not reachable [\#2620](https://github.com/argoproj/argo-cd/issues/2620)
- argocd app diff prints only first difference [\#2616](https://github.com/argoproj/argo-cd/issues/2616)
- Unable to add a private Google Source Repository with SSH [\#2611](https://github.com/argoproj/argo-cd/issues/2611)
- invalid resource: status.sync.comparedTo.source.path: Required value [\#2603](https://github.com/argoproj/argo-cd/issues/2603)
- UI application tiles should reflect corresponding app health status when presync hooks are running [\#2596](https://github.com/argoproj/argo-cd/issues/2596)
- sso login via cli always failing \(on my machine\) [\#2018](https://github.com/argoproj/argo-cd/issues/2018)

## [v1.3.0-rc3](https://github.com/argoproj/argo-cd/tree/v1.3.0-rc3) (2019-10-29)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.2.5...v1.3.0-rc3)

**Fixed bugs:**

- UI incorrectly mark resources as `Required Pruning` [\#2577](https://github.com/argoproj/argo-cd/issues/2577)
- argocd DeadlineExceeded when adding a new app [\#2553](https://github.com/argoproj/argo-cd/issues/2553)

## [v1.2.5](https://github.com/argoproj/argo-cd/tree/v1.2.5) (2019-10-29)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.2.4...v1.2.5)

**Implemented enhancements:**

- Trigger argocd pipeline on every new image push [\#2532](https://github.com/argoproj/argo-cd/issues/2532)
- Resource actions should show in audit trail as events [\#2180](https://github.com/argoproj/argo-cd/issues/2180)

**Fixed bugs:**

- Attempting to create a repo with password but not username panics [\#2567](https://github.com/argoproj/argo-cd/issues/2567)
- SSO login via CLI is blocked by the opened browser [\#2563](https://github.com/argoproj/argo-cd/issues/2563)
- Timestamp in Helm package file name causes error in Application with Helm source [\#2549](https://github.com/argoproj/argo-cd/issues/2549)
- argocd-dex-server sends to logs configuration with secrets [\#2536](https://github.com/argoproj/argo-cd/issues/2536)

**Closed issues:**

- Question: How can I control Applications defined in other namespaces? [\#2557](https://github.com/argoproj/argo-cd/issues/2557)

## [v1.2.4](https://github.com/argoproj/argo-cd/tree/v1.2.4) (2019-10-23)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.3.0-rc2...v1.2.4)

**Fixed bugs:**

- PostSync Hook not triggered when syncing from UI [\#2547](https://github.com/argoproj/argo-cd/issues/2547)

**Closed issues:**

- Helm Application source values [\#2531](https://github.com/argoproj/argo-cd/issues/2531)

## [v1.3.0-rc2](https://github.com/argoproj/argo-cd/tree/v1.3.0-rc2) (2019-10-23)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.3.0-rc1...v1.3.0-rc2)

**Implemented enhancements:**

- Opinionated formatting of typescript [\#2523](https://github.com/argoproj/argo-cd/issues/2523)
- Update root `Dockerfile` to use the `hack/install.sh` [\#2488](https://github.com/argoproj/argo-cd/issues/2488)
- Print version in interface [\#2483](https://github.com/argoproj/argo-cd/issues/2483)
- Consider using kustomize vars for manifests [\#2302](https://github.com/argoproj/argo-cd/issues/2302)
- Name suffix option [\#2301](https://github.com/argoproj/argo-cd/issues/2301)
- Support git submodules [\#2209](https://github.com/argoproj/argo-cd/issues/2209)
- Add repo credentials management to CLI and UI [\#2136](https://github.com/argoproj/argo-cd/issues/2136)
- App-of-apps: sync child apps [\#1932](https://github.com/argoproj/argo-cd/issues/1932)
- Improved docs for security set-up [\#1866](https://github.com/argoproj/argo-cd/issues/1866)
- `argocd app actions list APPNAME` should work without specify flags [\#1588](https://github.com/argoproj/argo-cd/issues/1588)
- Application label filtering [\#312](https://github.com/argoproj/argo-cd/issues/312)

**Fixed bugs:**

- RBAC not supporting team name with space in it [\#2540](https://github.com/argoproj/argo-cd/issues/2540)
- "argocd app sync --resource" prints stale status [\#2519](https://github.com/argoproj/argo-cd/issues/2519)
- Maintenance window meaning is confusing [\#2398](https://github.com/argoproj/argo-cd/issues/2398)
- Unknown error when setting params with argocd app set on helm app [\#2046](https://github.com/argoproj/argo-cd/issues/2046)
- Support YAML editor on helm parameter [\#1305](https://github.com/argoproj/argo-cd/issues/1305)

## [v1.3.0-rc1](https://github.com/argoproj/argo-cd/tree/v1.3.0-rc1) (2019-10-16)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.2.3...v1.3.0-rc1)

**Implemented enhancements:**

- Upgrade golangci-lint to lower memory version 1.20.0 [\#2433](https://github.com/argoproj/argo-cd/issues/2433)
- Allow non-interactive editing of application details [\#2403](https://github.com/argoproj/argo-cd/issues/2403)
- RBAC policy for restart [\#2401](https://github.com/argoproj/argo-cd/issues/2401)
- argocd list command should have filter options like by project [\#2396](https://github.com/argoproj/argo-cd/issues/2396)
- Is there a way to configure polling frequency? [\#2369](https://github.com/argoproj/argo-cd/issues/2369)
- Surface tooling versions in `/api/version` [\#2311](https://github.com/argoproj/argo-cd/issues/2311)
- Helm: UI improvements [\#2261](https://github.com/argoproj/argo-cd/issues/2261)
- Config management plugin environment variable UI/CLI support [\#2203](https://github.com/argoproj/argo-cd/issues/2203)
- User Profile Page [\#2031](https://github.com/argoproj/argo-cd/issues/2031)
- Support values.yaml outside of chart folder [\#2020](https://github.com/argoproj/argo-cd/issues/2020)
- Resource actions need their own dynamic RBAC action/\* verbs [\#2002](https://github.com/argoproj/argo-cd/issues/2002)
- Add a maintenance window for Applications with automated syncing [\#1995](https://github.com/argoproj/argo-cd/issues/1995)
- \[UI\] Add application labels to Applications list and Applications details page [\#1099](https://github.com/argoproj/argo-cd/issues/1099)

**Fixed bugs:**

- ArgoCD failing to install CRDs in master from Helm Charts [\#2497](https://github.com/argoproj/argo-cd/issues/2497)
- `golangci-lint --version` breaking Makefile [\#2491](https://github.com/argoproj/argo-cd/issues/2491)
- Impossible to edit chart name using App details page [\#2484](https://github.com/argoproj/argo-cd/issues/2484)
- Helm Hook is executed twice if annotated with both pre-install and pre-upgrade annotations [\#2480](https://github.com/argoproj/argo-cd/issues/2480)
- UI don't allow to create window with `\* \* \* \* \*` schedule [\#2475](https://github.com/argoproj/argo-cd/issues/2475)
- Missing `Kind` when editing new application as YAML  [\#2471](https://github.com/argoproj/argo-cd/issues/2471)
- Render application label through kustomize instead of replacing after [\#2466](https://github.com/argoproj/argo-cd/issues/2466)
- Command `app actions list \<appName\>` fails if any app resource is missing [\#2461](https://github.com/argoproj/argo-cd/issues/2461)
- JSON/YAML output of `argocd app actions list \<appName\>` is incorrect [\#2460](https://github.com/argoproj/argo-cd/issues/2460)
- Logic that checks sync windows state in the cli is incorrect [\#2455](https://github.com/argoproj/argo-cd/issues/2455)
- Application controller sometimes accidentally removes duplicated/excluded resource warning condition [\#2453](https://github.com/argoproj/argo-cd/issues/2453)
- Custom resource actions cannot be executed from the UI [\#2448](https://github.com/argoproj/argo-cd/issues/2448)
- Support values.yaml outside of chart folder allows directory traversal attack [\#2447](https://github.com/argoproj/argo-cd/issues/2447)
- proj windows commands not consistent with other commands [\#2443](https://github.com/argoproj/argo-cd/issues/2443)
- compare-options annotation don't appear to work v1.2.0 [\#2427](https://github.com/argoproj/argo-cd/issues/2427)
- Too many vulnerabilities in Docker image [\#2425](https://github.com/argoproj/argo-cd/issues/2425)
- UI bug when targetRevision is ommited [\#2407](https://github.com/argoproj/argo-cd/issues/2407)
- "Unable to create K8s REST config: invalid configuration" error when running using make start [\#2379](https://github.com/argoproj/argo-cd/issues/2379)
- Creating an application from Helm repository should select "Helm" as source type [\#2378](https://github.com/argoproj/argo-cd/issues/2378)
- The 'helm.repositories' settings is dropped without migration path [\#2316](https://github.com/argoproj/argo-cd/issues/2316)
- Client Side error coming when adding .git in REPO url [\#2286](https://github.com/argoproj/argo-cd/issues/2286)
- Argo CD should not be tied to a workflow version [\#2270](https://github.com/argoproj/argo-cd/issues/2270)
- user with project scoped full permission cannot set parameter from argocd web ui [\#2225](https://github.com/argoproj/argo-cd/issues/2225)
- `Exec format error` when running CLI tool on WSL with Ubuntu 18.04 [\#2221](https://github.com/argoproj/argo-cd/issues/2221)
- Manual sync does not trigger Presync hooks [\#2185](https://github.com/argoproj/argo-cd/issues/2185)
- Error message "Unable to load data: key is missing" is confusing [\#1944](https://github.com/argoproj/argo-cd/issues/1944)
- Fixes bug that prevents you creating repos via UI/CLI.  [\#2308](https://github.com/argoproj/argo-cd/pull/2308) ([alexec](https://github.com/alexec))

**Security fixes:**

- Adds checks around valid paths for apps [\#2133](https://github.com/argoproj/argo-cd/pull/2133) ([alexec](https://github.com/alexec))

**Closed issues:**

- Error when deploying argocd on Openshift [\#2423](https://github.com/argoproj/argo-cd/issues/2423)
- e2e tests timing out after 10m [\#2272](https://github.com/argoproj/argo-cd/issues/2272)
- Other install methods?  [\#2226](https://github.com/argoproj/argo-cd/issues/2226)
- Performing automated syncs on daily basis [\#1971](https://github.com/argoproj/argo-cd/issues/1971)

**Merged pull requests:**

- Fix bug preventing Helm CRDs from being installed [\#2500](https://github.com/argoproj/argo-cd/pull/2500) ([alexec](https://github.com/alexec))
- Work-around golang cilint error [\#2499](https://github.com/argoproj/argo-cd/pull/2499) ([alexec](https://github.com/alexec))
- Update CONTRIBUTING.md [\#2494](https://github.com/argoproj/argo-cd/pull/2494) ([jannfis](https://github.com/jannfis))
- Fix UI crash on application list page [\#2490](https://github.com/argoproj/argo-cd/pull/2490) ([alexmt](https://github.com/alexmt))
- Speedup codegen on mac [\#2489](https://github.com/argoproj/argo-cd/pull/2489) ([alexmt](https://github.com/alexmt))
- Final optimisations [\#2486](https://github.com/argoproj/argo-cd/pull/2486) ([alexec](https://github.com/alexec))
- Issue \#2484 - Impossible to edit chart name using App details page [\#2485](https://github.com/argoproj/argo-cd/pull/2485) ([alexmt](https://github.com/alexmt))
- Optimize codegen [\#2482](https://github.com/argoproj/argo-cd/pull/2482) ([alexec](https://github.com/alexec))
- Issue \#2480 - Helm Hook is executed twice if annotated with both pre-install and pre-upgrade annotations [\#2481](https://github.com/argoproj/argo-cd/pull/2481) ([alexmt](https://github.com/alexmt))
- Optimize linting [\#2479](https://github.com/argoproj/argo-cd/pull/2479) ([alexec](https://github.com/alexec))
- Update issue and PR templates [\#2478](https://github.com/argoproj/argo-cd/pull/2478) ([alexec](https://github.com/alexec))
- Issue \#2185 - Manual sync don't trigger hooks [\#2477](https://github.com/argoproj/argo-cd/pull/2477) ([alexmt](https://github.com/alexmt))
- Issue \#2475 - UI don't allow to create window with '\* \* \* \* \*' schedule [\#2476](https://github.com/argoproj/argo-cd/pull/2476) ([alexmt](https://github.com/alexmt))
- Optimizes e2e tests [\#2474](https://github.com/argoproj/argo-cd/pull/2474) ([alexec](https://github.com/alexec))
- Change "available" to "disabled" in actions, make them available by default [\#2470](https://github.com/argoproj/argo-cd/pull/2470) ([simster7](https://github.com/simster7))
- Allow collapse/expand helm values text [\#2469](https://github.com/argoproj/argo-cd/pull/2469) ([alexmt](https://github.com/alexmt))
- add support for --additional-headers cli flag [\#2467](https://github.com/argoproj/argo-cd/pull/2467) ([oboukili](https://github.com/oboukili))
- Add Latest Version Badge [\#2465](https://github.com/argoproj/argo-cd/pull/2465) ([codenio](https://github.com/codenio))
- Issue \#1944 - Gracefully handle missing cached app state [\#2464](https://github.com/argoproj/argo-cd/pull/2464) ([alexmt](https://github.com/alexmt))
- Hide sync windows if there are none [\#2459](https://github.com/argoproj/argo-cd/pull/2459) ([alexec](https://github.com/alexec))
- Issue \#2321 - Hook deletion should not fail if error message is not found [\#2458](https://github.com/argoproj/argo-cd/pull/2458) ([alexmt](https://github.com/alexmt))
- Fix logic when checking sync window status using the cli [\#2456](https://github.com/argoproj/argo-cd/pull/2456) ([adamjohnson01](https://github.com/adamjohnson01))
- Issue \#2453 - Application controller sometimes accidentally removes duplicated/excluded resource warning condition [\#2454](https://github.com/argoproj/argo-cd/pull/2454) ([alexmt](https://github.com/alexmt))
- Fix possible path traversal attack when supporting Helm `values.yaml` [\#2452](https://github.com/argoproj/argo-cd/pull/2452) ([simster7](https://github.com/simster7))
- Issue \#2110 - Disable menu item for non available resource actions on App Details page [\#2450](https://github.com/argoproj/argo-cd/pull/2450) ([alexmt](https://github.com/alexmt))
- Issue \#2448 - Custom resource actions cannot be executed from the UI [\#2449](https://github.com/argoproj/argo-cd/pull/2449) ([alexmt](https://github.com/alexmt))
- Fix formatting issues in docs [\#2446](https://github.com/argoproj/argo-cd/pull/2446) ([alexec](https://github.com/alexec))
- make proj windows commands consistent [\#2444](https://github.com/argoproj/argo-cd/pull/2444) ([adamjohnson01](https://github.com/adamjohnson01))
- Fix flaky TestExcludedResource [\#2440](https://github.com/argoproj/argo-cd/pull/2440) ([alexmt](https://github.com/alexmt))
- Error with new  `actions run` suggestion [\#2434](https://github.com/argoproj/argo-cd/pull/2434) ([simster7](https://github.com/simster7))
- Fix JS error on application creation page if no plugins configured [\#2432](https://github.com/argoproj/argo-cd/pull/2432) ([alexmt](https://github.com/alexmt))
- Update base image for Docker to Debian buster \(addresses \#2425\) [\#2431](https://github.com/argoproj/argo-cd/pull/2431) ([jannfis](https://github.com/jannfis))
- Add application labels to Applications list and Applications details page [\#2430](https://github.com/argoproj/argo-cd/pull/2430) ([alexmt](https://github.com/alexmt))
- Detach ArgoCD from specific workflow API [\#2428](https://github.com/argoproj/argo-cd/pull/2428) ([simster7](https://github.com/simster7))
- Fixes flakey test [\#2426](https://github.com/argoproj/argo-cd/pull/2426) ([alexec](https://github.com/alexec))
- Issue \#2396 argocd list command should have filter options like by pr… [\#2421](https://github.com/argoproj/argo-cd/pull/2421) ([AalokAhluwalia](https://github.com/AalokAhluwalia))
- App status panel shows metadata of current revision in git instead ofmost recent reconciled revision [\#2419](https://github.com/argoproj/argo-cd/pull/2419) ([alexmt](https://github.com/alexmt))
- Convert maintenance windows to sync windows [\#2416](https://github.com/argoproj/argo-cd/pull/2416) ([adamjohnson01](https://github.com/adamjohnson01))
- Add a hook example for sending a Slack message [\#2414](https://github.com/argoproj/argo-cd/pull/2414) ([alexec](https://github.com/alexec))
- Stop unnecessary re-loading clusters on every app list page re-render [\#2411](https://github.com/argoproj/argo-cd/pull/2411) ([alexmt](https://github.com/alexmt))
- Issue \#2407 - Improve Helm/Git app version rendering [\#2410](https://github.com/argoproj/argo-cd/pull/2410) ([alexmt](https://github.com/alexmt))
- Issue \#2378 - Creating an application from Helm repository should select Helm as source type [\#2409](https://github.com/argoproj/argo-cd/pull/2409) ([alexmt](https://github.com/alexmt))
- Adds support for plugin params.  [\#2406](https://github.com/argoproj/argo-cd/pull/2406) ([alexec](https://github.com/alexec))
- Issue \#2316 - support deprecated 'helm.repositories' config [\#2405](https://github.com/argoproj/argo-cd/pull/2405) ([alexmt](https://github.com/alexmt))
- Update FAQ for resource limits case [\#2404](https://github.com/argoproj/argo-cd/pull/2404) ([alexec](https://github.com/alexec))
- Updates FAQ for progressing sts scenario [\#2402](https://github.com/argoproj/argo-cd/pull/2402) ([alexec](https://github.com/alexec))
- Polish maintenance windows [\#2400](https://github.com/argoproj/argo-cd/pull/2400) ([alexec](https://github.com/alexec))
- Add release 1.2.1~1.2.3 changelog [\#2395](https://github.com/argoproj/argo-cd/pull/2395) ([alexmt](https://github.com/alexmt))
- Update faq.md [\#2394](https://github.com/argoproj/argo-cd/pull/2394) ([alexec](https://github.com/alexec))
- Fix broken helm test [\#2393](https://github.com/argoproj/argo-cd/pull/2393) ([alexmt](https://github.com/alexmt))
- Adds Traefik v2 documentation to ingress options [\#2392](https://github.com/argoproj/argo-cd/pull/2392) ([andrew-waters](https://github.com/andrew-waters))
- Add missing externalURL for networking.k8s.io Ingress type [\#2390](https://github.com/argoproj/argo-cd/pull/2390) ([reegnz](https://github.com/reegnz))
- Add dest-server and dest-namespace field to reconciliation logs [\#2388](https://github.com/argoproj/argo-cd/pull/2388) ([alexmt](https://github.com/alexmt))
- Stop loggin /repository.RepositoryService/ValidateAccess parameters [\#2387](https://github.com/argoproj/argo-cd/pull/2387) ([alexmt](https://github.com/alexmt))
- Use configured certificate to access helm repository [\#2385](https://github.com/argoproj/argo-cd/pull/2385) ([alexmt](https://github.com/alexmt))
- Add project level maintenance windows for applications [\#2380](https://github.com/argoproj/argo-cd/pull/2380) ([adamjohnson01](https://github.com/adamjohnson01))
- Refactor Helm client and unit test repo server [\#2377](https://github.com/argoproj/argo-cd/pull/2377) ([alexmt](https://github.com/alexmt))
- Add custom action example to argocd-cm.yaml [\#2375](https://github.com/argoproj/argo-cd/pull/2375) ([dthomson25](https://github.com/dthomson25))
- Fix API version for Deployment resources in e2e tests to app/v1 [\#2372](https://github.com/argoproj/argo-cd/pull/2372) ([jannfis](https://github.com/jannfis))
- Update Helm docs [\#2368](https://github.com/argoproj/argo-cd/pull/2368) ([alexec](https://github.com/alexec))
- Update api-docs.md [\#2365](https://github.com/argoproj/argo-cd/pull/2365) ([alexec](https://github.com/alexec))
- Issue \#2261 - Refactor Helm first class support [\#2364](https://github.com/argoproj/argo-cd/pull/2364) ([alexmt](https://github.com/alexmt))
- Make argo-cd docker images openshift friendly [\#2362](https://github.com/argoproj/argo-cd/pull/2362) ([duboisf](https://github.com/duboisf))
- Try out community icons.  [\#2349](https://github.com/argoproj/argo-cd/pull/2349) ([alexec](https://github.com/alexec))
- Issue \#2339 - Make sure controller uses latest git version if app reconciliation result expired [\#2346](https://github.com/argoproj/argo-cd/pull/2346) ([alexmt](https://github.com/alexmt))
- Fixes display of path in UI [\#2345](https://github.com/argoproj/argo-cd/pull/2345) ([alexec](https://github.com/alexec))
- Adds support for Github Enterprise URLs [\#2344](https://github.com/argoproj/argo-cd/pull/2344) ([alexec](https://github.com/alexec))
- Codegen [\#2343](https://github.com/argoproj/argo-cd/pull/2343) ([alexmt](https://github.com/alexmt))
- Don't fix imports in auto-generated files [\#2342](https://github.com/argoproj/argo-cd/pull/2342) ([alexmt](https://github.com/alexmt))
- docs: improve sso oidc documentation regarding client secret [\#2341](https://github.com/argoproj/argo-cd/pull/2341) ([greenstatic](https://github.com/greenstatic))
- chore\(dashboard.json\): adding argocd project as variable to grafana d… [\#2336](https://github.com/argoproj/argo-cd/pull/2336) ([igaskin](https://github.com/igaskin))
- Make `group` optional for `ignoreDifferences` setting [\#2335](https://github.com/argoproj/argo-cd/pull/2335) ([yujunz](https://github.com/yujunz))
- Fix Helm parameters with comma [\#2334](https://github.com/argoproj/argo-cd/pull/2334) ([olivierlemasle](https://github.com/olivierlemasle))
- Detypo architecture doc [\#2332](https://github.com/argoproj/argo-cd/pull/2332) ([squaremo](https://github.com/squaremo))
- Fix missing envs when updating application of content management plugin [\#2331](https://github.com/argoproj/argo-cd/pull/2331) ([yujunz](https://github.com/yujunz))
- Added Kustomize, Helm, and Kubectl to `argocd version` [\#2329](https://github.com/argoproj/argo-cd/pull/2329) ([simster7](https://github.com/simster7))
- Add cache-control HTTP header to badge response [\#2328](https://github.com/argoproj/argo-cd/pull/2328) ([greenstatic](https://github.com/greenstatic))
- util/localconfig: prefer HOME env var over os/user [\#2326](https://github.com/argoproj/argo-cd/pull/2326) ([gpaul](https://github.com/gpaul))
- Who uses ArgoCD: Add Lytt to the list [\#2324](https://github.com/argoproj/argo-cd/pull/2324) ([luizfonseca](https://github.com/luizfonseca))
- Fix docker image for dev [\#2319](https://github.com/argoproj/argo-cd/pull/2319) ([yujunz](https://github.com/yujunz))
- Add helm.repositories back to documentation [\#2314](https://github.com/argoproj/argo-cd/pull/2314) ([alexmt](https://github.com/alexmt))
- Fix broken links [\#2313](https://github.com/argoproj/argo-cd/pull/2313) ([jpreese](https://github.com/jpreese))
- Document flags/env variables useful for performance tuning [\#2312](https://github.com/argoproj/argo-cd/pull/2312) ([alexmt](https://github.com/alexmt))
- Auto-detect Helm repos + support Helm basic auth + fix bugs [\#2309](https://github.com/argoproj/argo-cd/pull/2309) ([alexec](https://github.com/alexec))
- Fixes bug in `argocd repo list` and tidy up UI [\#2307](https://github.com/argoproj/argo-cd/pull/2307) ([alexec](https://github.com/alexec))
- Clean-up the kube-version from Helm so that we can support GKE.  [\#2304](https://github.com/argoproj/argo-cd/pull/2304) ([alexec](https://github.com/alexec))
- Add Deployment action for `kubectl rollout restart` [\#2300](https://github.com/argoproj/argo-cd/pull/2300) ([jessesuen](https://github.com/jessesuen))
- Add --self-heal flag to argocd cli [\#2296](https://github.com/argoproj/argo-cd/pull/2296) ([mura-s](https://github.com/mura-s))
- Re-enable caching when listing apps.  [\#2295](https://github.com/argoproj/argo-cd/pull/2295) ([alexec](https://github.com/alexec))
- Fix TestAutoSyncSelfHealEnabled test flakiness [\#2293](https://github.com/argoproj/argo-cd/pull/2293) ([alexmt](https://github.com/alexmt))
- Issue \#2290 - Fix nil pointer dereference in application controller [\#2291](https://github.com/argoproj/argo-cd/pull/2291) ([alexmt](https://github.com/alexmt))
- Update OWNERS [\#2283](https://github.com/argoproj/argo-cd/pull/2283) ([alexec](https://github.com/alexec))
- Fix TestAutoSyncSelfHealEnabled test flakiness [\#2282](https://github.com/argoproj/argo-cd/pull/2282) ([alexmt](https://github.com/alexmt))
- Issue \#2245 - Intermittent "git ls-remote" request failures should not fail app reconciliation [\#2281](https://github.com/argoproj/argo-cd/pull/2281) ([alexmt](https://github.com/alexmt))
- Adding information to make local execution more accessible [\#2279](https://github.com/argoproj/argo-cd/pull/2279) ([augabet](https://github.com/augabet))
- Fix building error when following CONTRIBUTING.md [\#2278](https://github.com/argoproj/argo-cd/pull/2278) ([yujunz](https://github.com/yujunz))
- Support --kube-version.  [\#2276](https://github.com/argoproj/argo-cd/pull/2276) ([alexec](https://github.com/alexec))
- More helm [\#2274](https://github.com/argoproj/argo-cd/pull/2274) ([alexec](https://github.com/alexec))
- Increase e2e timeout to 15m.  [\#2273](https://github.com/argoproj/argo-cd/pull/2273) ([alexec](https://github.com/alexec))
- Fixes issue diffing secrets [\#2271](https://github.com/argoproj/argo-cd/pull/2271) ([alexec](https://github.com/alexec))
- Change Helm repo URLs to argoproj/argo-cd/master [\#2266](https://github.com/argoproj/argo-cd/pull/2266) ([twz123](https://github.com/twz123))
- Fix typo [\#2265](https://github.com/argoproj/argo-cd/pull/2265) ([ssbtn](https://github.com/ssbtn))
- Issue \#2022 - Support limiting number of concurrent kubectl fork/execs [\#2264](https://github.com/argoproj/argo-cd/pull/2264) ([alexmt](https://github.com/alexmt))
- API clients may use the HTTP Authorization header for auth.  [\#2262](https://github.com/argoproj/argo-cd/pull/2262) ([alexec](https://github.com/alexec))
- codegen [\#2254](https://github.com/argoproj/argo-cd/pull/2254) ([alexec](https://github.com/alexec))
- Add v1.2 Changelog [\#2252](https://github.com/argoproj/argo-cd/pull/2252) ([alexmt](https://github.com/alexmt))
- Fix degraded proxy support for http\(s\) git repository \(\#2243\) [\#2249](https://github.com/argoproj/argo-cd/pull/2249) ([mitsutaka](https://github.com/mitsutaka))
- Improve build stability [\#2247](https://github.com/argoproj/argo-cd/pull/2247) ([alexec](https://github.com/alexec))
- codegen [\#2244](https://github.com/argoproj/argo-cd/pull/2244) ([alexec](https://github.com/alexec))
- Update bug\_report.md [\#2242](https://github.com/argoproj/argo-cd/pull/2242) ([alexec](https://github.com/alexec))
- Fix typo. [\#2240](https://github.com/argoproj/argo-cd/pull/2240) ([ssbtn](https://github.com/ssbtn))
- docs/user-guide/projects.md: fix example policy [\#2233](https://github.com/argoproj/argo-cd/pull/2233) ([gpaul](https://github.com/gpaul))
- Grammar fixes. [\#2232](https://github.com/argoproj/argo-cd/pull/2232) ([cjyar](https://github.com/cjyar))
- Alter wording in Ingress docs to be more natural [\#2230](https://github.com/argoproj/argo-cd/pull/2230) ([k3ypad](https://github.com/k3ypad))
- Fix/grafana datasources [\#2229](https://github.com/argoproj/argo-cd/pull/2229) ([andrew-waters](https://github.com/andrew-waters))
- Issue \#2198 - Print empty string instead of Unknown in 'argocd app sync' output [\#2223](https://github.com/argoproj/argo-cd/pull/2223) ([alexmt](https://github.com/alexmt))
- Fix JS crash in EditablePanel component [\#2222](https://github.com/argoproj/argo-cd/pull/2222) ([alexmt](https://github.com/alexmt))
- Fix for displaying hooks in app diff view. Fixes \#2215 [\#2218](https://github.com/argoproj/argo-cd/pull/2218) ([alexec](https://github.com/alexec))
- If there is only one wave and no pre/post hooks, we should be synced.… [\#2217](https://github.com/argoproj/argo-cd/pull/2217) ([alexec](https://github.com/alexec))
- Issue \#2212 - Correctly handle trailing slash in configured URL while creating redirect URL  [\#2214](https://github.com/argoproj/argo-cd/pull/2214) ([alexmt](https://github.com/alexmt))
- Fix TestImmutableChange for running locally in microk8s [\#2213](https://github.com/argoproj/argo-cd/pull/2213) ([jannfis](https://github.com/jannfis))
- Fix flaky TestOrphanedResource test [\#2210](https://github.com/argoproj/argo-cd/pull/2210) ([alexmt](https://github.com/alexmt))
- Add path to externalURLs [\#2208](https://github.com/argoproj/argo-cd/pull/2208) ([rprin](https://github.com/rprin))
- Gzip JWTs and Adds New User Info Page [\#2204](https://github.com/argoproj/argo-cd/pull/2204) ([alexec](https://github.com/alexec))
- Create projects from manifests [\#2202](https://github.com/argoproj/argo-cd/pull/2202) ([Rayyis](https://github.com/Rayyis))
- Deals with race condition when deleting resource. Closes \#2141 [\#2200](https://github.com/argoproj/argo-cd/pull/2200) ([alexec](https://github.com/alexec))
- Issue \#1059 - Use ApplicationParameters panel on ApplicationCreatePanel [\#2197](https://github.com/argoproj/argo-cd/pull/2197) ([alexmt](https://github.com/alexmt))
- Remove duplicated DoNotIgnoreErrors method [\#2196](https://github.com/argoproj/argo-cd/pull/2196) ([alexmt](https://github.com/alexmt))
- Codegen [\#2195](https://github.com/argoproj/argo-cd/pull/2195) ([alexec](https://github.com/alexec))
- Fix broken links in Operator Manual declarative setup [\#2194](https://github.com/argoproj/argo-cd/pull/2194) ([jannfis](https://github.com/jannfis))
- Issue \#2192 - SyncError app condition disappears during app reconciliation [\#2193](https://github.com/argoproj/argo-cd/pull/2193) ([alexmt](https://github.com/alexmt))
- FAQ: Simplify admin password snippet a bit [\#2190](https://github.com/argoproj/argo-cd/pull/2190) ([twz123](https://github.com/twz123))
- Add missing labels to argocd-cm yaml in kustomize.md [\#2189](https://github.com/argoproj/argo-cd/pull/2189) ([alexmt](https://github.com/alexmt))
- Issue \#1167 - Document orphaned resources, update proj CLI [\#2188](https://github.com/argoproj/argo-cd/pull/2188) ([alexmt](https://github.com/alexmt))
- Add missing labels to configmap/secret in documentation [\#2187](https://github.com/argoproj/argo-cd/pull/2187) ([alexmt](https://github.com/alexmt))
- Issue \#1167 - Allow enabling/disabling orphaned resources using UI [\#2186](https://github.com/argoproj/argo-cd/pull/2186) ([alexmt](https://github.com/alexmt))
- Issue \#1167 - Excluded known orphaned resources exceptions [\#2178](https://github.com/argoproj/argo-cd/pull/2178) ([alexmt](https://github.com/alexmt))
- Issue \#2174 - Fix git repo url parsing on application list view [\#2175](https://github.com/argoproj/argo-cd/pull/2175) ([alexmt](https://github.com/alexmt))
- Temporary disable Git LFS test to unblock release [\#2172](https://github.com/argoproj/argo-cd/pull/2172) ([alexmt](https://github.com/alexmt))
- Issue \#2146 - Fix nil pointer dereference error during app reconciliation [\#2170](https://github.com/argoproj/argo-cd/pull/2170) ([alexmt](https://github.com/alexmt))
- Issue \#1167 - Controller should remove orphaned resources warning if app has no orphaned resources [\#2169](https://github.com/argoproj/argo-cd/pull/2169) ([alexmt](https://github.com/alexmt))
- Issue \#2114 - Fix history api fallback implementation to support app names with dots [\#2168](https://github.com/argoproj/argo-cd/pull/2168) ([alexmt](https://github.com/alexmt))
- Determine the manifest version from the VERSION file when on release … [\#2166](https://github.com/argoproj/argo-cd/pull/2166) ([alexec](https://github.com/alexec))
- Better detection for authorization\_code OIDC response type [\#2164](https://github.com/argoproj/argo-cd/pull/2164) ([twz123](https://github.com/twz123))
- Fixed routing issue for periods [\#2162](https://github.com/argoproj/argo-cd/pull/2162) ([simster7](https://github.com/simster7))
- Create "argocd" ns in `make start` [\#2161](https://github.com/argoproj/argo-cd/pull/2161) ([simster7](https://github.com/simster7))
- Added more health filters in UI [\#2160](https://github.com/argoproj/argo-cd/pull/2160) ([simster7](https://github.com/simster7))
- Updates app-of-apps docs [\#2159](https://github.com/argoproj/argo-cd/pull/2159) ([alexec](https://github.com/alexec))
- Update broken link [\#2158](https://github.com/argoproj/argo-cd/pull/2158) ([rytswd](https://github.com/rytswd))
- Adds test for updating immutable field, adds UI button to allow force from UI. See \#2150 [\#2155](https://github.com/argoproj/argo-cd/pull/2155) ([alexec](https://github.com/alexec))
- Added 'SyncFail' to possible HookTypes in UI [\#2153](https://github.com/argoproj/argo-cd/pull/2153) ([simster7](https://github.com/simster7))
- Fixes some code issues related to Kustomize build options. See \#2146 [\#2151](https://github.com/argoproj/argo-cd/pull/2151) ([alexec](https://github.com/alexec))
- Indicate that `SyncFail` hooks are on v1.2+ [\#2149](https://github.com/argoproj/argo-cd/pull/2149) ([simster7](https://github.com/simster7))
- Update helm.md [\#2145](https://github.com/argoproj/argo-cd/pull/2145) ([alexec](https://github.com/alexec))
- Improves BeforeHookCreation. Closes \#2141 [\#2142](https://github.com/argoproj/argo-cd/pull/2142) ([alexec](https://github.com/alexec))
- Ignore generated code coverage [\#2135](https://github.com/argoproj/argo-cd/pull/2135) ([alexmt](https://github.com/alexmt))
- Enhances cookie warning with actual length to help users fix their co… [\#2134](https://github.com/argoproj/argo-cd/pull/2134) ([alexec](https://github.com/alexec))
- Minor CLI bug fixes [\#2132](https://github.com/argoproj/argo-cd/pull/2132) ([alexec](https://github.com/alexec))
- Update faq.md [\#2131](https://github.com/argoproj/argo-cd/pull/2131) ([alexec](https://github.com/alexec))
- Update rbac.md [\#2130](https://github.com/argoproj/argo-cd/pull/2130) ([alexec](https://github.com/alexec))
- Issue \#2060 - Enpoint incorrectly considered top level managed resource [\#2129](https://github.com/argoproj/argo-cd/pull/2129) ([alexmt](https://github.com/alexmt))
- Fixed truncation of group in UI. Closes \#2006 [\#2128](https://github.com/argoproj/argo-cd/pull/2128) ([alexec](https://github.com/alexec))
- Updates hook delete policy docs [\#2127](https://github.com/argoproj/argo-cd/pull/2127) ([alexec](https://github.com/alexec))
- Fixes flaky e2e tests. Closes \#2086 [\#2126](https://github.com/argoproj/argo-cd/pull/2126) ([alexec](https://github.com/alexec))
- Adds a floating action button with help and chat links to every page.… [\#2125](https://github.com/argoproj/argo-cd/pull/2125) ([alexec](https://github.com/alexec))
- Add FuturePLC to List of companies using ArgoCD [\#2122](https://github.com/argoproj/argo-cd/pull/2122) ([warmfusion](https://github.com/warmfusion))
- Allow adding certs for hostnames ending on a dot \(fixes \#2116\) [\#2120](https://github.com/argoproj/argo-cd/pull/2120) ([jannfis](https://github.com/jannfis))
- Redact secrets using "+" rather than "\*" as this is base 64 compatiba… [\#2119](https://github.com/argoproj/argo-cd/pull/2119) ([alexec](https://github.com/alexec))
- Escape square brackets in pattern matching hostnames \(fixes \#2099\) [\#2113](https://github.com/argoproj/argo-cd/pull/2113) ([jannfis](https://github.com/jannfis))
- Correct some broken links in yaml [\#2111](https://github.com/argoproj/argo-cd/pull/2111) ([rytswd](https://github.com/rytswd))
- RBAC Support for Actions [\#2110](https://github.com/argoproj/argo-cd/pull/2110) ([simster7](https://github.com/simster7))
- Adds support for setting Helm string parameters via CLI. Closes \#2078 [\#2109](https://github.com/argoproj/argo-cd/pull/2109) ([alexec](https://github.com/alexec))
- Issue \#1167 - Implement orphan resources support  [\#2103](https://github.com/argoproj/argo-cd/pull/2103) ([alexmt](https://github.com/alexmt))
- Fix and enhance end-to-end testing for SSH repositories [\#2101](https://github.com/argoproj/argo-cd/pull/2101) ([jannfis](https://github.com/jannfis))
- Helm hooks. Closes \#355 [\#2069](https://github.com/argoproj/argo-cd/pull/2069) ([alexec](https://github.com/alexec))
- Adds support for a literal YAML block of Helm values. Closes \#1930  [\#2057](https://github.com/argoproj/argo-cd/pull/2057) ([alexec](https://github.com/alexec))
- Adds support for hook-delete-policy: BeforeHookCreation. Closes \#2036 [\#2048](https://github.com/argoproj/argo-cd/pull/2048) ([alexec](https://github.com/alexec))
- support OIDC claims request [\#1957](https://github.com/argoproj/argo-cd/pull/1957) ([sboschman](https://github.com/sboschman))
- Allow list actions to return yaml or json [\#1805](https://github.com/argoproj/argo-cd/pull/1805) ([dthomson25](https://github.com/dthomson25))

## [v1.2.3](https://github.com/argoproj/argo-cd/tree/v1.2.3) (2019-10-01)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.2.2...v1.2.3)

**Implemented enhancements:**

- Add cluster information into Splunk [\#2354](https://github.com/argoproj/argo-cd/issues/2354)
- Update Helm docs [\#2315](https://github.com/argoproj/argo-cd/issues/2315)
- Make `group` optional for `ignoreDifferences` config [\#2298](https://github.com/argoproj/argo-cd/issues/2298)
- Use community icons for resources [\#2277](https://github.com/argoproj/argo-cd/issues/2277)

**Fixed bugs:**

- The parameters of ValidateAccess GRPC method should not be logged  [\#2386](https://github.com/argoproj/argo-cd/issues/2386)
- Cannot add new application from Git - Application.argoproj.io "guestbook" is invalid [\#2374](https://github.com/argoproj/argo-cd/issues/2374)
- End-to-End tests not working with Kubernetes v1.16 [\#2371](https://github.com/argoproj/argo-cd/issues/2371)
- OpenShift route always OutOfSync [\#2370](https://github.com/argoproj/argo-cd/issues/2370)
- failed to update sync policy after deletion of repo [\#2363](https://github.com/argoproj/argo-cd/issues/2363)
- Helm values not set on child apps [\#2359](https://github.com/argoproj/argo-cd/issues/2359)
- Using CURL to get Status of projects/applications - no session information [\#2356](https://github.com/argoproj/argo-cd/issues/2356)
- accept ingress: null in OpenShift routes [\#2353](https://github.com/argoproj/argo-cd/issues/2353)
- adding repositories from argued cli failed [\#2348](https://github.com/argoproj/argo-cd/issues/2348)
- Child app not installed when using same private repo for parent and child, in app-of-apps pattern [\#2333](https://github.com/argoproj/argo-cd/issues/2333)

## [v1.2.2](https://github.com/argoproj/argo-cd/tree/v1.2.2) (2019-09-24)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.2.1...v1.2.2)

**Implemented enhancements:**

- Helm repo name should optional [\#2325](https://github.com/argoproj/argo-cd/issues/2325)
- Helm: auto-detect URLs [\#2260](https://github.com/argoproj/argo-cd/issues/2260)
- Resource action equivalent to `kubectl rollout restart` [\#2177](https://github.com/argoproj/argo-cd/issues/2177)

**Fixed bugs:**

- Failed edit application with plugin type requiring environment [\#2330](https://github.com/argoproj/argo-cd/issues/2330)
- ignoreDifferences not effective for Secret [\#2322](https://github.com/argoproj/argo-cd/issues/2322)
- Warning message on application view [\#2320](https://github.com/argoproj/argo-cd/issues/2320)
- Badge response does not contain cache control header [\#2317](https://github.com/argoproj/argo-cd/issues/2317)
- Must always supply credentials when adding repo despite common secret being present in repository.credentials [\#2305](https://github.com/argoproj/argo-cd/issues/2305)
- The Helm --kube-version support does not work on GKE: [\#2303](https://github.com/argoproj/argo-cd/issues/2303)
- Argocd wait failed on synced application [\#2297](https://github.com/argoproj/argo-cd/issues/2297)
- ListApps does not utilize cache [\#2287](https://github.com/argoproj/argo-cd/issues/2287)
- Multiple parallel app syncs causes OOM [\#2022](https://github.com/argoproj/argo-cd/issues/2022)
- GUI not parsing map values properly [\#1877](https://github.com/argoproj/argo-cd/issues/1877)

**Merged pull requests:**

- util/localconfig: prefer HOME env var over os/user \(\#2326\) [\#2350](https://github.com/argoproj/argo-cd/pull/2350) ([alexmt](https://github.com/alexmt))

## [v1.2.1](https://github.com/argoproj/argo-cd/tree/v1.2.1) (2019-09-12)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.2.0...v1.2.1)

**Implemented enhancements:**

- Scoping to namespace for remote clusters [\#2289](https://github.com/argoproj/argo-cd/issues/2289)
- Support `helm template --kube-version ` [\#2275](https://github.com/argoproj/argo-cd/issues/2275)
- Helm: remove `tlsClientCaData`  [\#2258](https://github.com/argoproj/argo-cd/issues/2258)
- helmfile support [\#2143](https://github.com/argoproj/argo-cd/issues/2143)
- HTTP API should allow JWT to be passed via Authorization header [\#1642](https://github.com/argoproj/argo-cd/issues/1642)
- Helm repository as first class Argo CD Application source [\#1145](https://github.com/argoproj/argo-cd/issues/1145)

**Fixed bugs:**

- Path autocompletion is broken on application create panel [\#2269](https://github.com/argoproj/argo-cd/issues/2269)
- Result of ListApps operation for Git repo is cached incorrectly [\#2263](https://github.com/argoproj/argo-cd/issues/2263)
- SNI support is not implemented and creates a redirect loop with the nginx ingress controller [\#2246](https://github.com/argoproj/argo-cd/issues/2246)
- Intermittent "git ls-remote" request failures should not fail app reconciliation [\#2245](https://github.com/argoproj/argo-cd/issues/2245)
- v1.2.0-rc2 does not retrieve http\(s\) based git repository behind the proxy [\#2243](https://github.com/argoproj/argo-cd/issues/2243)
- Brew installs macOS binaries on Linux [\#2241](https://github.com/argoproj/argo-cd/issues/2241)
- CONTRIBUTING documentation outdated [\#2231](https://github.com/argoproj/argo-cd/issues/2231)
- 1.2.0-rc2 Warning during secret diffing [\#2206](https://github.com/argoproj/argo-cd/issues/2206)

**Closed issues:**

- Poll users [\#2123](https://github.com/argoproj/argo-cd/issues/2123)
- Tell people about v1.2 \(when GA\) [\#2117](https://github.com/argoproj/argo-cd/issues/2117)

## [v1.2.0](https://github.com/argoproj/argo-cd/tree/v1.2.0) (2019-09-04)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.2.0-rc2...v1.2.0)

**Implemented enhancements:**

- Additional Helm values from configmaps or secrets [\#2183](https://github.com/argoproj/argo-cd/issues/2183)
- Flag to auto-sync even on sync failure [\#2156](https://github.com/argoproj/argo-cd/issues/2156)
- Request OIDC groups claim if groups scope is not supported [\#1956](https://github.com/argoproj/argo-cd/issues/1956)
- Deflake TestDeletingAppStuckInSync [\#1882](https://github.com/argoproj/argo-cd/issues/1882)
- Ability to create & upsert projects from spec [\#1852](https://github.com/argoproj/argo-cd/issues/1852)
- Add a layer of overrides [\#1784](https://github.com/argoproj/argo-cd/issues/1784)
- Allow adding repo with public key from the GUI [\#1638](https://github.com/argoproj/argo-cd/issues/1638)
- Ability to generate a warn/alert when a namespace deviates from the expected state [\#1167](https://github.com/argoproj/argo-cd/issues/1167)
- \[UI\] Enhance app creation page with Helm parameters overrides [\#1059](https://github.com/argoproj/argo-cd/issues/1059)

**Fixed bugs:**

- After port - forwarding execution, login ends up with tcp6 stream error [\#2220](https://github.com/argoproj/argo-cd/issues/2220)
- An app with a single resource and Sync hook remains progressing [\#2216](https://github.com/argoproj/argo-cd/issues/2216)
- Application summary diff page shows hooks [\#2215](https://github.com/argoproj/argo-cd/issues/2215)
- SSO redirect url is incorrect if configured Argo CD URL has trailing slash [\#2212](https://github.com/argoproj/argo-cd/issues/2212)
- Command `argocd account update-password` is broken [\#2205](https://github.com/argoproj/argo-cd/issues/2205)
- argocd app wait\sync prints 'Unknown' for resources without health [\#2198](https://github.com/argoproj/argo-cd/issues/2198)
- SyncError app condition disappears during app reconciliation [\#2192](https://github.com/argoproj/argo-cd/issues/2192)
- sso login with cli doesn't work [\#2179](https://github.com/argoproj/argo-cd/issues/2179)
- ArgoCD tries to delete the contents of argocd-cm.yaml with kustomize bases [\#2173](https://github.com/argoproj/argo-cd/issues/2173)
- Docs should state that BeforeHookCreation is only in v1.3 [\#2148](https://github.com/argoproj/argo-cd/issues/2148)
- Fix `TestHookDeleteBeforeCreation` [\#2141](https://github.com/argoproj/argo-cd/issues/2141)
- Secret key can not be omitted [\#1936](https://github.com/argoproj/argo-cd/issues/1936)

## [v1.2.0-rc2](https://github.com/argoproj/argo-cd/tree/v1.2.0-rc2) (2019-08-21)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.2.0-rc1...v1.2.0-rc2)

**Implemented enhancements:**

- Add configurable help link to every page [\#2124](https://github.com/argoproj/argo-cd/issues/2124)
- Add all health options to UI filter [\#2118](https://github.com/argoproj/argo-cd/issues/2118)
- Support setting Helm string parameters using CLI/UI [\#2078](https://github.com/argoproj/argo-cd/issues/2078)
- Support `helm.sh/hook` [\#2039](https://github.com/argoproj/argo-cd/issues/2039)
- Support `argocd.argoproj.io/hook-delete-policy: BeforeHookCreation` [\#2036](https://github.com/argoproj/argo-cd/issues/2036)
- Support `"helm.sh/hook-delete-policy": before-hook-creation` [\#2035](https://github.com/argoproj/argo-cd/issues/2035)
- Support for in-line block from helm chart values [\#1930](https://github.com/argoproj/argo-cd/issues/1930)
- Allow `argocd app wait` to continue if a specified number of replicas satisfy [\#1924](https://github.com/argoproj/argo-cd/issues/1924)
- Syncwaves at the application level [\#1888](https://github.com/argoproj/argo-cd/issues/1888)
- Anonymous readOnly access to WebUI [\#1620](https://github.com/argoproj/argo-cd/issues/1620)
- Map helm lifecycle hooks to ArgoCD pre/post/sync hooks [\#355](https://github.com/argoproj/argo-cd/issues/355)

**Fixed bugs:**

- v1.2.0-rc1 Applications List View doesn't work [\#2174](https://github.com/argoproj/argo-cd/issues/2174)
- 1.2.0-rc1 - Authentication Required error in Repo Server [\#2152](https://github.com/argoproj/argo-cd/issues/2152)
- Replacing services failure [\#2150](https://github.com/argoproj/argo-cd/issues/2150)
- SyncFail Resource Hooks not working [\#2147](https://github.com/argoproj/argo-cd/issues/2147)
- v1.2.0-rc1 nil pointer dereference when syncing [\#2146](https://github.com/argoproj/argo-cd/issues/2146)
- https://hub.docker.com/r/argoproj/argocd-ui/tags has no v1.2.0-rc1 Tag [\#2137](https://github.com/argoproj/argo-cd/issues/2137)
- Adding certs for hostnames ending with a dot \(.\) is not possible [\#2116](https://github.com/argoproj/argo-cd/issues/2116)
- Application 404s on names with periods [\#2114](https://github.com/argoproj/argo-cd/issues/2114)
- I can't add to repository with the ArgoCD [\#2112](https://github.com/argoproj/argo-cd/issues/2112)
- SSH known hosts entry cannot be deleted if contains shell pattern in name [\#2099](https://github.com/argoproj/argo-cd/issues/2099)
- Applications in unknown state for too many minutes, timeouts on repo server [\#2070](https://github.com/argoproj/argo-cd/issues/2070)
- Endpoint is no longer shown as a child of services [\#2060](https://github.com/argoproj/argo-cd/issues/2060)
- OIDC group bindings are truncated [\#2006](https://github.com/argoproj/argo-cd/issues/2006)
- Warning during secret diffing [\#1923](https://github.com/argoproj/argo-cd/issues/1923)

**Closed issues:**

- Fix `TestAutoSyncSelfHealEnabled` [\#2086](https://github.com/argoproj/argo-cd/issues/2086)
- StatefulSet's volumeClaimTemplates is displayed as PersistentVolumeClaim which status is OutOfSync [\#1729](https://github.com/argoproj/argo-cd/issues/1729)
- Support for weighted resource hooks: `helm.sh/hook-weight` [\#1396](https://github.com/argoproj/argo-cd/issues/1396)

## [v1.2.0-rc1](https://github.com/argoproj/argo-cd/tree/v1.2.0-rc1) (2019-08-06)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.1.2...v1.2.0-rc1)

**Implemented enhancements:**

- Fix `TestAppOfApps` [\#2089](https://github.com/argoproj/argo-cd/issues/2089)
- Upgrade to Kustomize v3.1 [\#2068](https://github.com/argoproj/argo-cd/issues/2068)
- Members of a group bound to a role should be able to generate tokens for that role [\#1977](https://github.com/argoproj/argo-cd/issues/1977)
- Expose kustomize image overrides in CLI [\#1843](https://github.com/argoproj/argo-cd/issues/1843)
- Document best practices for preview apps [\#1811](https://github.com/argoproj/argo-cd/issues/1811)
- System level `kustomize build` options [\#1789](https://github.com/argoproj/argo-cd/issues/1789)
- Kustomize v2.1.0 support [\#1783](https://github.com/argoproj/argo-cd/issues/1783)
- Remove support for kustomize 1 entirely [\#1573](https://github.com/argoproj/argo-cd/issues/1573)

**Fixed bugs:**

- argocd-server panic: interface conversion: runtime.Object is \*v1.Status, not \*v1alpha1.Application [\#2105](https://github.com/argoproj/argo-cd/issues/2105)
- Yoda test issues [\#2095](https://github.com/argoproj/argo-cd/issues/2095)
- Resource manifest view tab is empty [\#2081](https://github.com/argoproj/argo-cd/issues/2081)
- argocd-util missing repository TLS client and insecure export details [\#2075](https://github.com/argoproj/argo-cd/issues/2075)
- Missing UI to enable selfHealing in app syncing policy [\#2074](https://github.com/argoproj/argo-cd/issues/2074)
- App resources table view looks odd [\#2073](https://github.com/argoproj/argo-cd/issues/2073)
- Command `argocd app wait` prints incorrect sync status while app is in out-of-sync state [\#2049](https://github.com/argoproj/argo-cd/issues/2049)
- Problem syncing app sync failed without message [\#1887](https://github.com/argoproj/argo-cd/issues/1887)

**Closed issues:**

- Fix `TestCustomToolWithGitCreds` [\#2090](https://github.com/argoproj/argo-cd/issues/2090)
- Fix `TestArgoCDWaitEnsureAppIsNotCrashing` [\#2088](https://github.com/argoproj/argo-cd/issues/2088)
- Fix `TestDeletingAppStuckInSync` [\#2087](https://github.com/argoproj/argo-cd/issues/2087)

**Merged pull requests:**

- Run dockerized manifests generation during release [\#2107](https://github.com/argoproj/argo-cd/pull/2107) ([alexmt](https://github.com/alexmt))
- Do not panic if the type is not api.Status \(an error scenario\). Close… [\#2106](https://github.com/argoproj/argo-cd/pull/2106) ([alexec](https://github.com/alexec))
- Issue \#2073 - Fix status column width on resources list table [\#2104](https://github.com/argoproj/argo-cd/pull/2104) ([alexmt](https://github.com/alexmt))
- Update private repository documentation to reflect latest changes [\#2102](https://github.com/argoproj/argo-cd/pull/2102) ([jannfis](https://github.com/jannfis))
- Fix flaky TestDeletingAppStuckInSync test [\#2100](https://github.com/argoproj/argo-cd/pull/2100) ([alexmt](https://github.com/alexmt))
- Export cert settings and TLS client cert secrets in argocd-util \(closes \#2075\) [\#2098](https://github.com/argoproj/argo-cd/pull/2098) ([jannfis](https://github.com/jannfis))
- Correctly set required env vars for running E2E tests locally \(fixes \#2090\) [\#2097](https://github.com/argoproj/argo-cd/pull/2097) ([jannfis](https://github.com/jannfis))
- Support Jsonnet TLA code and string args [\#2096](https://github.com/argoproj/argo-cd/pull/2096) ([alexmt](https://github.com/alexmt))
- Fixes TestDeletingAppStuckInSync. Closes \#2087 [\#2093](https://github.com/argoproj/argo-cd/pull/2093) ([alexec](https://github.com/alexec))
- Fixes TestArgoCDWaitEnsureAppIsNotCrashing. Closes \#2088 [\#2092](https://github.com/argoproj/argo-cd/pull/2092) ([alexec](https://github.com/alexec))
- Skip TestAppOfApps. Closes \#2089 [\#2091](https://github.com/argoproj/argo-cd/pull/2091) ([alexec](https://github.com/alexec))
- Update releasing docs [\#2084](https://github.com/argoproj/argo-cd/pull/2084) ([alexmt](https://github.com/alexmt))
- Update VERSION file [\#2083](https://github.com/argoproj/argo-cd/pull/2083) ([alexmt](https://github.com/alexmt))
- Issue \#2081 - Fix empty manifest tab error [\#2082](https://github.com/argoproj/argo-cd/pull/2082) ([alexmt](https://github.com/alexmt))
- Issue \#2060 - Make sure endpoint is shown as a child of service [\#2080](https://github.com/argoproj/argo-cd/pull/2080) ([alexmt](https://github.com/alexmt))
- Update metrics doc [\#2079](https://github.com/argoproj/argo-cd/pull/2079) ([jvoltaww](https://github.com/jvoltaww))
- Remove support for Kustomize 1. Closes \#1573 [\#2077](https://github.com/argoproj/argo-cd/pull/2077) ([alexec](https://github.com/alexec))
- Issue \#2074 - Add UI to manage sync options [\#2076](https://github.com/argoproj/argo-cd/pull/2076) ([alexmt](https://github.com/alexmt))
- Upgrade kustomize to v3.1.0 [\#2072](https://github.com/argoproj/argo-cd/pull/2072) ([alexmt](https://github.com/alexmt))
- Change field names to camelCase in RepositoryCertificate of CertificateService API [\#2071](https://github.com/argoproj/argo-cd/pull/2071) ([jannfis](https://github.com/jannfis))
- Improve wrapped text rendering on application list and application sync status pages [\#2067](https://github.com/argoproj/argo-cd/pull/2067) ([alexmt](https://github.com/alexmt))
- Allow users to create tokens for projects where they have any role. C… [\#2066](https://github.com/argoproj/argo-cd/pull/2066) ([alexec](https://github.com/alexec))
- Make sure applications filter is expanded on xlarge and xxlarge screens [\#2065](https://github.com/argoproj/argo-cd/pull/2065) ([alexmt](https://github.com/alexmt))
- Bugfix: Ensure we have a valid hostname when adding certificates [\#2064](https://github.com/argoproj/argo-cd/pull/2064) ([jannfis](https://github.com/jannfis))
- Moves talks into one place [\#2063](https://github.com/argoproj/argo-cd/pull/2063) ([alexec](https://github.com/alexec))
- Update user documentation on private repositories feature [\#2062](https://github.com/argoproj/argo-cd/pull/2062) ([jannfis](https://github.com/jannfis))
- Make sure applications page has 4 tiles on very large screen [\#2059](https://github.com/argoproj/argo-cd/pull/2059) ([alexmt](https://github.com/alexmt))
- Bump Circle CI cache key versions. [\#2058](https://github.com/argoproj/argo-cd/pull/2058) ([alexec](https://github.com/alexec))
- Adds support for `argocd app set` for Kustomize. Closes \#1843 [\#2055](https://github.com/argoproj/argo-cd/pull/2055) ([alexec](https://github.com/alexec))
- Install go-junit-report dependency in CI [\#2054](https://github.com/argoproj/argo-cd/pull/2054) ([alexmt](https://github.com/alexmt))
- Improve refresh button on applications tiles view [\#2053](https://github.com/argoproj/argo-cd/pull/2053) ([alexmt](https://github.com/alexmt))
- Build ArgoCD UI into argocd Docker image [\#2052](https://github.com/argoproj/argo-cd/pull/2052) ([jannfis](https://github.com/jannfis))
- Add v1.x changelog [\#2051](https://github.com/argoproj/argo-cd/pull/2051) ([alexmt](https://github.com/alexmt))
- Issue \#2049 - 'argocd app wait' should print correct sync status [\#2050](https://github.com/argoproj/argo-cd/pull/2050) ([alexmt](https://github.com/alexmt))
- Move repo certificate info generation to server side [\#2044](https://github.com/argoproj/argo-cd/pull/2044) ([jannfis](https://github.com/jannfis))
- Add first-class support for connecting repositories via UI [\#2043](https://github.com/argoproj/argo-cd/pull/2043) ([jannfis](https://github.com/jannfis))
- Delete obsolete comment [\#2042](https://github.com/argoproj/argo-cd/pull/2042) ([alexmt](https://github.com/alexmt))
- Build argocd-repo-server and argocd-util with packr instead of go [\#2041](https://github.com/argoproj/argo-cd/pull/2041) ([jannfis](https://github.com/jannfis))
- Allow codegen to run as non-root user in the Docker container [\#2032](https://github.com/argoproj/argo-cd/pull/2032) ([jannfis](https://github.com/jannfis))
- Update OWNERS [\#2030](https://github.com/argoproj/argo-cd/pull/2030) ([edlee2121](https://github.com/edlee2121))
- Shows correct revision metadata for. Closes \#2028 [\#2029](https://github.com/argoproj/argo-cd/pull/2029) ([alexec](https://github.com/alexec))
- Adds a refresh button to list and tile view. Closes \#1606 [\#2027](https://github.com/argoproj/argo-cd/pull/2027) ([alexec](https://github.com/alexec))
- Word-wraps app info in the table and list views. Closes \#2004 [\#2026](https://github.com/argoproj/argo-cd/pull/2026) ([alexec](https://github.com/alexec))
- Removes width limitation and adjusts column widths for resource list … [\#2025](https://github.com/argoproj/argo-cd/pull/2025) ([alexec](https://github.com/alexec))
- Run yarn build on CI. [\#2024](https://github.com/argoproj/argo-cd/pull/2024) ([alexec](https://github.com/alexec))
- Issue \#1780 - Project source/destination removal should consider wildcards [\#2023](https://github.com/argoproj/argo-cd/pull/2023) ([alexmt](https://github.com/alexmt))
- Allow Apply Only on sync [\#2015](https://github.com/argoproj/argo-cd/pull/2015) ([leominov](https://github.com/leominov))
- Fixed date and version number of 1.1 release in changelog [\#2014](https://github.com/argoproj/argo-cd/pull/2014) ([crunchtime-ali](https://github.com/crunchtime-ali))
- Adds CLI support for adding and removing groups from project roles. C… [\#2013](https://github.com/argoproj/argo-cd/pull/2013) ([alexec](https://github.com/alexec))
- Update yarn.lock [\#2012](https://github.com/argoproj/argo-cd/pull/2012) ([alexec](https://github.com/alexec))
- Word wraps error message. Closes \#1942 [\#2011](https://github.com/argoproj/argo-cd/pull/2011) ([alexec](https://github.com/alexec))
- Issue \#2000 - Repo whitelisting in UI does not support wildcards [\#2010](https://github.com/argoproj/argo-cd/pull/2010) ([alexmt](https://github.com/alexmt))
- Issue \#2007 - UI should remember most recent selected tab on resource info panel [\#2009](https://github.com/argoproj/argo-cd/pull/2009) ([alexmt](https://github.com/alexmt))
- Regenerate lock file [\#2008](https://github.com/argoproj/argo-cd/pull/2008) ([alexmt](https://github.com/alexmt))
- Fix invalid RBAC configuration example \(add missing quotes\) [\#2003](https://github.com/argoproj/argo-cd/pull/2003) ([alexmt](https://github.com/alexmt))
- Issue \#1989 - Fix creating app resources graph using network connections on app details page [\#2001](https://github.com/argoproj/argo-cd/pull/2001) ([alexmt](https://github.com/alexmt))
- Issue \#1940 - Wait for CRD creation during sync process [\#1999](https://github.com/argoproj/argo-cd/pull/1999) ([alexmt](https://github.com/alexmt))
- Adds link to the project from the app summary page. Closes \#1911 [\#1998](https://github.com/argoproj/argo-cd/pull/1998) ([alexec](https://github.com/alexec))
- Update config-management-plugins.md. Closes \#1950 [\#1997](https://github.com/argoproj/argo-cd/pull/1997) ([alexec](https://github.com/alexec))
- Documents the instance label. Closes \#1482 [\#1996](https://github.com/argoproj/argo-cd/pull/1996) ([alexec](https://github.com/alexec))
- Add 1.1 changelog [\#1994](https://github.com/argoproj/argo-cd/pull/1994) ([alexmt](https://github.com/alexmt))
- Added a button to select out of sync items in the sync panel [\#1993](https://github.com/argoproj/argo-cd/pull/1993) ([Rayyis](https://github.com/Rayyis))
- Fix typo. [\#1991](https://github.com/argoproj/argo-cd/pull/1991) ([d-kuro](https://github.com/d-kuro))
- Issue \#1736 - Auto-sync should support self-healing option [\#1990](https://github.com/argoproj/argo-cd/pull/1990) ([alexmt](https://github.com/alexmt))
- Add UI functionality for certificate management [\#1987](https://github.com/argoproj/argo-cd/pull/1987) ([jannfis](https://github.com/jannfis))
- Issue \#1984 - Support 'override' action in UI/API [\#1985](https://github.com/argoproj/argo-cd/pull/1985) ([alexmt](https://github.com/alexmt))
- Issue \#1982 - Fix argocd app wait message [\#1983](https://github.com/argoproj/argo-cd/pull/1983) ([alexmt](https://github.com/alexmt))
- Add ability to specify insecure and LFS modes when adding repos via UI [\#1979](https://github.com/argoproj/argo-cd/pull/1979) ([jannfis](https://github.com/jannfis))
- Mention git-lfs as a build pre-requisite \(fixes \#1974\) [\#1978](https://github.com/argoproj/argo-cd/pull/1978) ([jannfis](https://github.com/jannfis))
- Update sample grafana dashboard with new metrics [\#1976](https://github.com/argoproj/argo-cd/pull/1976) ([alexmt](https://github.com/alexmt))
- Only run Git LFS tests on CI. [\#1975](https://github.com/argoproj/argo-cd/pull/1975) ([alexec](https://github.com/alexec))
- Do not ignore Argo hooks when there is a Helm hook. Closes \#1952 [\#1973](https://github.com/argoproj/argo-cd/pull/1973) ([alexec](https://github.com/alexec))
- Move kustomization files back to kustomize 2.0.3 [\#1972](https://github.com/argoproj/argo-cd/pull/1972) ([alexmt](https://github.com/alexmt))
- Bump jquery from 3.3.1 to 3.4.1 in /ui [\#1970](https://github.com/argoproj/argo-cd/pull/1970) ([dependabot[bot]](https://github.com/apps/dependabot))
- Bump lodash-es from 4.17.10 to 4.17.15 in /ui [\#1969](https://github.com/argoproj/argo-cd/pull/1969) ([dependabot[bot]](https://github.com/apps/dependabot))
- Bump fstream from 1.0.11 to 1.0.12 in /ui [\#1968](https://github.com/argoproj/argo-cd/pull/1968) ([dependabot[bot]](https://github.com/apps/dependabot))
- Bump lodash.mergewith from 4.6.1 to 4.6.2 in /ui [\#1967](https://github.com/argoproj/argo-cd/pull/1967) ([dependabot[bot]](https://github.com/apps/dependabot))
- Bump js-yaml from 3.12.0 to 3.13.1 in /ui [\#1966](https://github.com/argoproj/argo-cd/pull/1966) ([dependabot[bot]](https://github.com/apps/dependabot))
- Update CONTRIBUTING.md [\#1964](https://github.com/argoproj/argo-cd/pull/1964) ([stgarf](https://github.com/stgarf))
- Check that TLS is enabled when registering DEX Handlers [\#1963](https://github.com/argoproj/argo-cd/pull/1963) ([dmizelle](https://github.com/dmizelle))
- Issue \#1935 - `argocd app sync` hangs when cluster is not configured  [\#1962](https://github.com/argoproj/argo-cd/pull/1962) ([alexmt](https://github.com/alexmt))
- Add support for connecting repositories using TLS client certs \(fixes \#1945\) [\#1960](https://github.com/argoproj/argo-cd/pull/1960) ([jannfis](https://github.com/jannfis))
- Fix argocd app sync/get cli [\#1959](https://github.com/argoproj/argo-cd/pull/1959) ([alexmt](https://github.com/alexmt))
- Revision metadata api fails if specified revision is ambiguous [\#1958](https://github.com/argoproj/argo-cd/pull/1958) ([alexmt](https://github.com/alexmt))
- Do not allow app-of-app child app's Missing status to affect parent [\#1954](https://github.com/argoproj/argo-cd/pull/1954) ([jessesuen](https://github.com/jessesuen))
- Remove unnecessary details from sync errors [\#1951](https://github.com/argoproj/argo-cd/pull/1951) ([alexmt](https://github.com/alexmt))
- Change git prometheus counter name [\#1949](https://github.com/argoproj/argo-cd/pull/1949) ([alexmt](https://github.com/alexmt))
- Bump Kustomize v2.1.0 to v3.0.2 [\#1948](https://github.com/argoproj/argo-cd/pull/1948) ([dmizelle](https://github.com/dmizelle))
- Add Mambu - The SaaS Banking Engine to the list of users [\#1947](https://github.com/argoproj/argo-cd/pull/1947) ([lcostea](https://github.com/lcostea))
- E2E test should fail in action fails unless otherwise configured [\#1946](https://github.com/argoproj/argo-cd/pull/1946) ([alexmt](https://github.com/alexmt))
- Update README.md [\#1943](https://github.com/argoproj/argo-cd/pull/1943) ([saradhis](https://github.com/saradhis))
- Add support for Git LFS enabled repositories \(fixes \#1853\) [\#1941](https://github.com/argoproj/argo-cd/pull/1941) ([jannfis](https://github.com/jannfis))
- Add link to IBM blog about Argo CD [\#1939](https://github.com/argoproj/argo-cd/pull/1939) ([alexmt](https://github.com/alexmt))
- Update README.md [\#1937](https://github.com/argoproj/argo-cd/pull/1937) ([saradhis](https://github.com/saradhis))
- Issues \#1513 - Make sure insecure flag works for remote Kustomize bases [\#1934](https://github.com/argoproj/argo-cd/pull/1934) ([alexmt](https://github.com/alexmt))
- Update user and operator documentation for change introduced in \#1807 [\#1933](https://github.com/argoproj/argo-cd/pull/1933) ([jannfis](https://github.com/jannfis))
- Adress some change requests remaining from \#1807 [\#1931](https://github.com/argoproj/argo-cd/pull/1931) ([jannfis](https://github.com/jannfis))
- Issue \#1919 - Eliminate unnecessary git interactions for top-level resource changes [\#1929](https://github.com/argoproj/argo-cd/pull/1929) ([alexmt](https://github.com/alexmt))
- Make sure refresh icon is not hidden under status panel [\#1928](https://github.com/argoproj/argo-cd/pull/1928) ([alexmt](https://github.com/alexmt))
- Disable codecov/patch [\#1926](https://github.com/argoproj/argo-cd/pull/1926) ([alexmt](https://github.com/alexmt))
- Regenerate Gopkg.lock [\#1925](https://github.com/argoproj/argo-cd/pull/1925) ([alexmt](https://github.com/alexmt))
- Issue \#1841 Make the health check for APIService a built in [\#1921](https://github.com/argoproj/argo-cd/pull/1921) ([masa213f](https://github.com/masa213f))
- Issue \#897 - Secret data not redacted in last-applied-configuration [\#1920](https://github.com/argoproj/argo-cd/pull/1920) ([alexmt](https://github.com/alexmt))
- Issue \#1917 - App details page incorrect uses 'requires pruning' icon for out-of-sync resources [\#1918](https://github.com/argoproj/argo-cd/pull/1918) ([alexmt](https://github.com/alexmt))
- Issue \#1915 - App deployment history should load metadata of only sel… [\#1916](https://github.com/argoproj/argo-cd/pull/1916) ([alexmt](https://github.com/alexmt))
- Issue \#1912 - Add Prometheus metrics for git repo interactions [\#1914](https://github.com/argoproj/argo-cd/pull/1914) ([alexmt](https://github.com/alexmt))
- Introduces compact diff view. Closes \#1831 [\#1913](https://github.com/argoproj/argo-cd/pull/1913) ([alexec](https://github.com/alexec))
- Issue \#1909 - App controller should log additional information during app syncing [\#1910](https://github.com/argoproj/argo-cd/pull/1910) ([alexmt](https://github.com/alexmt))
- Update config-management-plugins.md [\#1908](https://github.com/argoproj/argo-cd/pull/1908) ([alexec](https://github.com/alexec))
- Added SyncFail hooks to docs [\#1907](https://github.com/argoproj/argo-cd/pull/1907) ([simster7](https://github.com/simster7))
- Upgrade argo ui version to pull dropdown fix [\#1906](https://github.com/argoproj/argo-cd/pull/1906) ([alexmt](https://github.com/alexmt))
- Attempts to make CI builds more reliable by reducing lint memory usage. [\#1905](https://github.com/argoproj/argo-cd/pull/1905) ([alexec](https://github.com/alexec))
- Improved the usability of an app-of-apps by adding a link to the chil… [\#1900](https://github.com/argoproj/argo-cd/pull/1900) ([alexec](https://github.com/alexec))
- Upgrade argo ui version to pull dropdown fix [\#1899](https://github.com/argoproj/argo-cd/pull/1899) ([alexmt](https://github.com/alexmt))
- Attempts to fix flaky TestCustomToolWithEnv [\#1893](https://github.com/argoproj/argo-cd/pull/1893) ([alexec](https://github.com/alexec))
- Fixes bug where repo was not displayed in UI. Closes \#1883 [\#1892](https://github.com/argoproj/argo-cd/pull/1892) ([alexec](https://github.com/alexec))
- Log more error information. See \#1887 [\#1891](https://github.com/argoproj/argo-cd/pull/1891) ([alexec](https://github.com/alexec))
- Allow Helm parameters to force ambiguously-typed values to be strings [\#1889](https://github.com/argoproj/argo-cd/pull/1889) ([jutley](https://github.com/jutley))
- Marks TestDeletingAppStuckInSync as flaky [\#1880](https://github.com/argoproj/argo-cd/pull/1880) ([alexec](https://github.com/alexec))
- Fixes garbage in e2e logging on CI [\#1879](https://github.com/argoproj/argo-cd/pull/1879) ([alexec](https://github.com/alexec))
- Attempts to fix flaky TestCustomToolWithEnv [\#1876](https://github.com/argoproj/argo-cd/pull/1876) ([alexec](https://github.com/alexec))
- Issue \#1874 - validate app spec before verifying app permissions [\#1875](https://github.com/argoproj/argo-cd/pull/1875) ([alexmt](https://github.com/alexmt))
- Enables unparam linter and fixes linting issues [\#1872](https://github.com/argoproj/argo-cd/pull/1872) ([alexec](https://github.com/alexec))
- Redacts Helm username and password. Closes \#1868 [\#1871](https://github.com/argoproj/argo-cd/pull/1871) ([alexec](https://github.com/alexec))
- Issue \#1867 - Fix JS error on project role edit panel [\#1869](https://github.com/argoproj/argo-cd/pull/1869) ([alexmt](https://github.com/alexmt))
- Adds support for Helm 1st-class. Closes \#1145 [\#1865](https://github.com/argoproj/argo-cd/pull/1865) ([alexec](https://github.com/alexec))
- Issue \#1620 - Support anonymous argocd access [\#1864](https://github.com/argoproj/argo-cd/pull/1864) ([alexmt](https://github.com/alexmt))
- Fix JS UI crash if user is not authenticated [\#1863](https://github.com/argoproj/argo-cd/pull/1863) ([alexmt](https://github.com/alexmt))
- Issue \#1621 - Proper handling of an excluded resource in an application [\#1862](https://github.com/argoproj/argo-cd/pull/1862) ([alexmt](https://github.com/alexmt))
- Adds support for environment variables to custom plugins [\#1860](https://github.com/argoproj/argo-cd/pull/1860) ([alexec](https://github.com/alexec))
- Issue 1858 - Support 'application/merge-patch+json' in 'argocd app patch' [\#1859](https://github.com/argoproj/argo-cd/pull/1859) ([alexmt](https://github.com/alexmt))
- Issue \#1855 - Fix jumping app status panel [\#1856](https://github.com/argoproj/argo-cd/pull/1856) ([alexmt](https://github.com/alexmt))
- Issue \#1159 - Different icon for resources which require pruning [\#1854](https://github.com/argoproj/argo-cd/pull/1854) ([alexmt](https://github.com/alexmt))
- Issue \#1614 - Stop reloading logs of stopped container [\#1850](https://github.com/argoproj/argo-cd/pull/1850) ([alexmt](https://github.com/alexmt))
- Attempt to fix flaky tests [\#1849](https://github.com/argoproj/argo-cd/pull/1849) ([alexec](https://github.com/alexec))
- Issue \#738 - Allow configuring google analytics tracking [\#1848](https://github.com/argoproj/argo-cd/pull/1848) ([alexmt](https://github.com/alexmt))
- Upgrade argo-ui version to fix dropdown position calculation [\#1847](https://github.com/argoproj/argo-cd/pull/1847) ([alexmt](https://github.com/alexmt))
- Improve status badge feature [\#1844](https://github.com/argoproj/argo-cd/pull/1844) ([alexmt](https://github.com/alexmt))
- Removes logging that appears when using the  CLI [\#1842](https://github.com/argoproj/argo-cd/pull/1842) ([alexec](https://github.com/alexec))
- Badgedocs [\#1840](https://github.com/argoproj/argo-cd/pull/1840) ([naynasiddharth](https://github.com/naynasiddharth))
- Fix docs typos
 [\#1838](https://github.com/argoproj/argo-cd/pull/1838) ([evrardjp](https://github.com/evrardjp))
- Enable monaco editor web service to improve editor performance [\#1837](https://github.com/argoproj/argo-cd/pull/1837) ([alexmt](https://github.com/alexmt))
- Fix validation of datetime fields [\#1836](https://github.com/argoproj/argo-cd/pull/1836) ([alexmt](https://github.com/alexmt))
- Explicitly specify version of every dev tool [\#1835](https://github.com/argoproj/argo-cd/pull/1835) ([alexmt](https://github.com/alexmt))
- Updates CI cache go-v16 [\#1834](https://github.com/argoproj/argo-cd/pull/1834) ([alexec](https://github.com/alexec))
- Issue \#1820 - Make sure api server to repo server grpc calls have timeout [\#1832](https://github.com/argoproj/argo-cd/pull/1832) ([alexmt](https://github.com/alexmt))
- Increase codecov threshold to 1% [\#1830](https://github.com/argoproj/argo-cd/pull/1830) ([alexmt](https://github.com/alexmt))
- Issue \#1676 - Move remove Repositories, RepositoryCredentials, HelmCredentials settings from ArgoCDSettings structure [\#1829](https://github.com/argoproj/argo-cd/pull/1829) ([alexmt](https://github.com/alexmt))
- Running application actions should require `override` privileges not `get` [\#1828](https://github.com/argoproj/argo-cd/pull/1828) ([jessesuen](https://github.com/jessesuen))
- Update declarative-setup.md [\#1825](https://github.com/argoproj/argo-cd/pull/1825) ([alexec](https://github.com/alexec))
- Adds a timeout to all external commands. Closes \#1821 [\#1823](https://github.com/argoproj/argo-cd/pull/1823) ([alexec](https://github.com/alexec))
- Fixes TestUserAgent [\#1822](https://github.com/argoproj/argo-cd/pull/1822) ([alexec](https://github.com/alexec))
- App status badge [\#1818](https://github.com/argoproj/argo-cd/pull/1818) ([naynasiddharth](https://github.com/naynasiddharth))
- Added Kustomize build options to settings/argocd-cm [\#1817](https://github.com/argoproj/argo-cd/pull/1817) ([simster7](https://github.com/simster7))
- Adds cluster name to UI. Closes \#1353 [\#1814](https://github.com/argoproj/argo-cd/pull/1814) ([alexec](https://github.com/alexec))
- Parameterize Argo UI base image [\#1813](https://github.com/argoproj/argo-cd/pull/1813) ([jessesuen](https://github.com/jessesuen))
- App status badge [\#1812](https://github.com/argoproj/argo-cd/pull/1812) ([naynasiddharth](https://github.com/naynasiddharth))
- Fix formatting on helm docs [\#1810](https://github.com/argoproj/argo-cd/pull/1810) ([alexec](https://github.com/alexec))
- Implement Bitbucket Server and Gogs webhook providers [\#1808](https://github.com/argoproj/argo-cd/pull/1808) ([johnmarcou](https://github.com/johnmarcou))
- Simplify server certificate and known hosts management [\#1807](https://github.com/argoproj/argo-cd/pull/1807) ([jannfis](https://github.com/jannfis))
- Update k8s libraries to v1.14 [\#1806](https://github.com/argoproj/argo-cd/pull/1806) ([jessesuen](https://github.com/jessesuen))
- bash autocompletion for argocd [\#1803](https://github.com/argoproj/argo-cd/pull/1803) ([edwinpjacques](https://github.com/edwinpjacques))
- Display the revision of resources in the labels. Closes \#1367 [\#1802](https://github.com/argoproj/argo-cd/pull/1802) ([alexec](https://github.com/alexec))
- Fixes TestKustomize2AppSource. Closes \#1800 [\#1801](https://github.com/argoproj/argo-cd/pull/1801) ([alexec](https://github.com/alexec))
- Issue \#1676 - Move remove AppInstanceLabelKey, ConfigManagementPlugins, ResourceOverrides, ResourceExclusions, ResourceInclusions settings from ArgoCDSettings structure [\#1799](https://github.com/argoproj/argo-cd/pull/1799) ([alexmt](https://github.com/alexmt))
- Fixes a bug where cluster objs could leave app is running op state. C… [\#1796](https://github.com/argoproj/argo-cd/pull/1796) ([alexec](https://github.com/alexec))
- Added SyncFail hooks [\#1795](https://github.com/argoproj/argo-cd/pull/1795) ([simster7](https://github.com/simster7))
- Move remarshaling to happen only during comparison, instead of manifest generation [\#1788](https://github.com/argoproj/argo-cd/pull/1788) ([jessesuen](https://github.com/jessesuen))
- Add health check to the controller deployment [\#1785](https://github.com/argoproj/argo-cd/pull/1785) ([alexmt](https://github.com/alexmt))
- Make status fields as optional fields [\#1779](https://github.com/argoproj/argo-cd/pull/1779) ([alexmt](https://github.com/alexmt))
- Updates docs for resetting password. Closes \#1775 [\#1777](https://github.com/argoproj/argo-cd/pull/1777) ([alexec](https://github.com/alexec))
- Use correct healthcheck for Rollout with empty steps list [\#1776](https://github.com/argoproj/argo-cd/pull/1776) ([dthomson25](https://github.com/dthomson25))
-  Generate CRD schema using github.com/kubernetes-sigs/controller-tools [\#1773](https://github.com/argoproj/argo-cd/pull/1773) ([alexmt](https://github.com/alexmt))
- Adds a check for whether or not the application path exists and return a clear error message [\#1772](https://github.com/argoproj/argo-cd/pull/1772) ([alexec](https://github.com/alexec))
- Added local sync to docs [\#1771](https://github.com/argoproj/argo-cd/pull/1771) ([simster7](https://github.com/simster7))
- Sync status button should be hidden if there is no sync operation [\#1770](https://github.com/argoproj/argo-cd/pull/1770) ([alexmt](https://github.com/alexmt))
- Support top level arguments for jsonnet on app create cli [\#1769](https://github.com/argoproj/argo-cd/pull/1769) ([lcostea](https://github.com/lcostea))
- Improve sync result messages. Closes \#1486 [\#1768](https://github.com/argoproj/argo-cd/pull/1768) ([alexec](https://github.com/alexec))
- Displays targetRevision in app dashboards. Closes \#1239 [\#1767](https://github.com/argoproj/argo-cd/pull/1767) ([alexec](https://github.com/alexec))
- Pin k8s.io/kube-openapi to a stable version [\#1765](https://github.com/argoproj/argo-cd/pull/1765) ([jessesuen](https://github.com/jessesuen))
- UI should allow editing repo URL [\#1763](https://github.com/argoproj/argo-cd/pull/1763) ([alexmt](https://github.com/alexmt))
- Adds more commit data. Closes \#1219 [\#1762](https://github.com/argoproj/argo-cd/pull/1762) ([alexec](https://github.com/alexec))
- Attempts to make CI builds more robust by increasing the expect timeout [\#1759](https://github.com/argoproj/argo-cd/pull/1759) ([alexec](https://github.com/alexec))
- Adds "namespace: argocd" to examples [\#1758](https://github.com/argoproj/argo-cd/pull/1758) ([alexec](https://github.com/alexec))
- Update Readme.md adds Saildrone org [\#1757](https://github.com/argoproj/argo-cd/pull/1757) ([jdeprin](https://github.com/jdeprin))
- Add example for overriding Helm values and release name [\#1756](https://github.com/argoproj/argo-cd/pull/1756) ([AndiDog](https://github.com/AndiDog))
- Add ANSTO - Australian Synchrotron to list of users [\#1755](https://github.com/argoproj/argo-cd/pull/1755) ([johnmarcou](https://github.com/johnmarcou))
- Increase stale bot inactivity period [\#1754](https://github.com/argoproj/argo-cd/pull/1754) ([alexmt](https://github.com/alexmt))
- Make E2E tests more configurable locally [\#1753](https://github.com/argoproj/argo-cd/pull/1753) ([jannfis](https://github.com/jannfis))
- Allow setting timeout for golangci-lint in build process from external [\#1752](https://github.com/argoproj/argo-cd/pull/1752) ([jannfis](https://github.com/jannfis))
- Adds Validate=false that disables validation when applying resources.… [\#1750](https://github.com/argoproj/argo-cd/pull/1750) ([alexec](https://github.com/alexec))
- Explicitly specify user root during argocd image build [\#1749](https://github.com/argoproj/argo-cd/pull/1749) ([alexmt](https://github.com/alexmt))
- Fix bug in release [\#1748](https://github.com/argoproj/argo-cd/pull/1748) ([alexec](https://github.com/alexec))
- codegen [\#1747](https://github.com/argoproj/argo-cd/pull/1747) ([alexec](https://github.com/alexec))
- Server side rotation of cluster bearer tokens [\#1744](https://github.com/argoproj/argo-cd/pull/1744) ([jessesuen](https://github.com/jessesuen))
- Make the `--grpc-web` flag more obvious [\#1703](https://github.com/argoproj/argo-cd/pull/1703) ([stgarf](https://github.com/stgarf))
- Added local path syncing [\#1578](https://github.com/argoproj/argo-cd/pull/1578) ([simster7](https://github.com/simster7))
- Issue \#1570 - Application controller is unable to delete self-referenced app [\#1574](https://github.com/argoproj/argo-cd/pull/1574) ([alexmt](https://github.com/alexmt))

## [v1.1.2](https://github.com/argoproj/argo-cd/tree/v1.1.2) (2019-07-30)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.1.1...v1.1.2)

**Implemented enhancements:**

- Support `helm.sh/hook-weight` [\#2037](https://github.com/argoproj/argo-cd/issues/2037)
- UI should remember most recent selected tab on resource info panel [\#2007](https://github.com/argoproj/argo-cd/issues/2007)
- Support Client HTTPS certifcates for private git repositories [\#1945](https://github.com/argoproj/argo-cd/issues/1945)
- Link to project from app details [\#1911](https://github.com/argoproj/argo-cd/issues/1911)
- sync only all outofsync items option [\#1902](https://github.com/argoproj/argo-cd/issues/1902)
- `argocd proj role add-group PROJECT ROLE-NAME GROUP-NAME` [\#1851](https://github.com/argoproj/argo-cd/issues/1851)
- \[Feature\] Allow adding private repositories with Custom CA via UI [\#1670](https://github.com/argoproj/argo-cd/issues/1670)
- \[UI\] Refresh button from table and card view [\#1606](https://github.com/argoproj/argo-cd/issues/1606)
- Investigate alternative notion of managed resources than `app.kubernetes.io/instance` label  [\#1482](https://github.com/argoproj/argo-cd/issues/1482)
- \[UI\] Support dry run and hook vs. apply strategy during sync [\#798](https://github.com/argoproj/argo-cd/issues/798)

**Fixed bugs:**

- OIDC Authentication failing after upgrading to 1.1.1 [\#2047](https://github.com/argoproj/argo-cd/issues/2047)
- Error initializing argo-dex-server and argo-repo-server deployments [\#2045](https://github.com/argoproj/argo-cd/issues/2045)
- Revision meta-data contains stale information for OutOfSync [\#2028](https://github.com/argoproj/argo-cd/issues/2028)
- Unable to login from CLI via ingress route [\#2019](https://github.com/argoproj/argo-cd/issues/2019)
- SSO Configuration using argocd-secret not working [\#2017](https://github.com/argoproj/argo-cd/issues/2017)
- Truncated names in application resource list view [\#2005](https://github.com/argoproj/argo-cd/issues/2005)
- Application names are truncated from table/card views [\#2004](https://github.com/argoproj/argo-cd/issues/2004)
- Repo whitelisting in UI does not support wildcards [\#2000](https://github.com/argoproj/argo-cd/issues/2000)
- Regression in network view [\#1989](https://github.com/argoproj/argo-cd/issues/1989)
- It is impossible to use `override` action in role policy [\#1984](https://github.com/argoproj/argo-cd/issues/1984)
- Argocd RBAC configmap scopes field not working properly [\#1981](https://github.com/argoproj/argo-cd/issues/1981)
- Argo CD hooks are ignored when helm hooks present [\#1952](https://github.com/argoproj/argo-cd/issues/1952)
- Application env vars missing in configManagementPlugins [\#1950](https://github.com/argoproj/argo-cd/issues/1950)
- Long sync error messages are truncated in UI [\#1942](https://github.com/argoproj/argo-cd/issues/1942)
- Argo CD sync might fail to create if app contains CRD and instance of the same CRD [\#1940](https://github.com/argoproj/argo-cd/issues/1940)
- Project destination removal does not consider wildcards [\#1780](https://github.com/argoproj/argo-cd/issues/1780)
- auto-sync apply only once for one commit [\#1736](https://github.com/argoproj/argo-cd/issues/1736)

## [v1.1.1](https://github.com/argoproj/argo-cd/tree/v1.1.1) (2019-07-24)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.1.0...v1.1.1)

**Implemented enhancements:**

- Update contributing guide to support Git LFS [\#1974](https://github.com/argoproj/argo-cd/issues/1974)
- Argo CD UI: Sync Only Missing and/or OutOfSync [\#1961](https://github.com/argoproj/argo-cd/issues/1961)
- Helm parameters should support Helm's `--set-string` flag [\#1846](https://github.com/argoproj/argo-cd/issues/1846)

**Fixed bugs:**

- `argocd app wait` shows incorrect status for non-hook resources [\#1982](https://github.com/argoproj/argo-cd/issues/1982)
- Connection to external k8s Cluster failes with "Unable to connect to cluster: specifying a root certificates file with the insecure flag is not allowed" [\#1965](https://github.com/argoproj/argo-cd/issues/1965)
- Simplify whitelisting self signed cert of git server [\#1514](https://github.com/argoproj/argo-cd/issues/1514)

## [v1.1.0](https://github.com/argoproj/argo-cd/tree/v1.1.0) (2019-07-22)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.1.0-rc8...v1.1.0)

## [v1.1.0-rc8](https://github.com/argoproj/argo-cd/tree/v1.1.0-rc8) (2019-07-19)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.1.0-rc7...v1.1.0-rc8)

**Implemented enhancements:**

- Validation for HPA and Deployment/replicas [\#1603](https://github.com/argoproj/argo-cd/issues/1603)

**Fixed bugs:**

- API Server communicates with the Git repository directly [\#1955](https://github.com/argoproj/argo-cd/issues/1955)
- `argocd app sync` hangs when cluster is not configured [\#1935](https://github.com/argoproj/argo-cd/issues/1935)
- ArgoCD Does Not Support LFS [\#1853](https://github.com/argoproj/argo-cd/issues/1853)
- CertManager becomes `SyncError` with ArgoCD v1.1.0-rc1 [\#1826](https://github.com/argoproj/argo-cd/issues/1826)

## [v1.1.0-rc7](https://github.com/argoproj/argo-cd/tree/v1.1.0-rc7) (2019-07-17)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.1.0-rc6...v1.1.0-rc7)

**Fixed bugs:**

- error message not showing in the UI [\#1897](https://github.com/argoproj/argo-cd/issues/1897)

## [v1.1.0-rc6](https://github.com/argoproj/argo-cd/tree/v1.1.0-rc6) (2019-07-16)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.1.0-rc5...v1.1.0-rc6)

**Implemented enhancements:**

- Eliminate unnecessary git interactions for top-level resource changes [\#1919](https://github.com/argoproj/argo-cd/issues/1919)
- Prometheus metrics for git repo interactions [\#1912](https://github.com/argoproj/argo-cd/issues/1912)
- App controller should log additional information during app syncing [\#1909](https://github.com/argoproj/argo-cd/issues/1909)
- Kustomize plugins, transformers and generators [\#1904](https://github.com/argoproj/argo-cd/issues/1904)
- Improve application in UI [\#1898](https://github.com/argoproj/argo-cd/issues/1898)
- Make the health check for APIService a built in  [\#1841](https://github.com/argoproj/argo-cd/issues/1841)
- Make it easier to understand what the application diffs are [\#1831](https://github.com/argoproj/argo-cd/issues/1831)
- Introduce `--insecure-skip-tls-verify` option for self-signed HTTP git URLs [\#1513](https://github.com/argoproj/argo-cd/issues/1513)

**Fixed bugs:**

- Cannot login with new install per instructions in Getting Started docs [\#1938](https://github.com/argoproj/argo-cd/issues/1938)
- unhandled exception returns exit code 0 [\#1922](https://github.com/argoproj/argo-cd/issues/1922)
- App details page incorrect uses 'requires pruning' icon for out-of-sync resources [\#1917](https://github.com/argoproj/argo-cd/issues/1917)
- Application details sends dozen on API queries to load revision metadata [\#1915](https://github.com/argoproj/argo-cd/issues/1915)
- Large number of git pulls from the same repo [\#1363](https://github.com/argoproj/argo-cd/issues/1363)
- Secret data not redacted in last-applied-configuration [\#897](https://github.com/argoproj/argo-cd/issues/897)

## [v1.1.0-rc5](https://github.com/argoproj/argo-cd/tree/v1.1.0-rc5) (2019-07-09)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.1.0-rc4...v1.1.0-rc5)

**Fixed bugs:**

- Intermittent status code 429 when using AWS CodeCommit repositories [\#1895](https://github.com/argoproj/argo-cd/issues/1895)
- Repo will not appear in UI if it is not a known provider [\#1883](https://github.com/argoproj/argo-cd/issues/1883)

**Closed issues:**

- gitops secrets [\#1364](https://github.com/argoproj/argo-cd/issues/1364)

## [v1.1.0-rc4](https://github.com/argoproj/argo-cd/tree/v1.1.0-rc4) (2019-07-03)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.1.0-rc3...v1.1.0-rc4)

**Implemented enhancements:**

- Support `application/merge-patch+json` patch type in `argocd app patch` command [\#1858](https://github.com/argoproj/argo-cd/issues/1858)
- Better handling of "excluded" resource in application manifests [\#1621](https://github.com/argoproj/argo-cd/issues/1621)
- Use cluster name rather than cluster URL in the UI [\#1353](https://github.com/argoproj/argo-cd/issues/1353)
- allow for parameters in configuration plugins [\#1323](https://github.com/argoproj/argo-cd/issues/1323)
- \[UI\] Non-Obvious OutOfSync resource due to pruning [\#1159](https://github.com/argoproj/argo-cd/issues/1159)
- Add health check for API Services [\#1017](https://github.com/argoproj/argo-cd/issues/1017)
- Make a SyncFailure hook available [\#815](https://github.com/argoproj/argo-cd/issues/815)
- Google analytics integration [\#738](https://github.com/argoproj/argo-cd/issues/738)

**Fixed bugs:**

- `argocd app create` should infer app destination for ksonnet app [\#1874](https://github.com/argoproj/argo-cd/issues/1874)
- argoexec does not redact username and password for Helm [\#1868](https://github.com/argoproj/argo-cd/issues/1868)
- \[UI\]Role edit panel is broken [\#1867](https://github.com/argoproj/argo-cd/issues/1867)
- resource customization  still out of sync [\#1857](https://github.com/argoproj/argo-cd/issues/1857)
- Application status panel height jumping [\#1855](https://github.com/argoproj/argo-cd/issues/1855)
- Application type dropdown is broken on Application create panel [\#1845](https://github.com/argoproj/argo-cd/issues/1845)
- Argo CD openapi validation spec is incomplete [\#1687](https://github.com/argoproj/argo-cd/issues/1687)
- \[UI\] Argo UI repeatedly spits out logs of init container [\#1614](https://github.com/argoproj/argo-cd/issues/1614)

**Closed issues:**

- Logging into ArgoCD using kubectl [\#1760](https://github.com/argoproj/argo-cd/issues/1760)

## [v1.1.0-rc3](https://github.com/argoproj/argo-cd/tree/v1.1.0-rc3) (2019-06-28)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.1.0-rc2...v1.1.0-rc3)

**Implemented enhancements:**

- Github Enterprise SSH Key [\#1824](https://github.com/argoproj/argo-cd/issues/1824)
- bash autocompletion [\#1798](https://github.com/argoproj/argo-cd/issues/1798)
- Badges for build status [\#1435](https://github.com/argoproj/argo-cd/issues/1435)
- \[UI\] Highlight the current replicaset of a deployment  [\#1367](https://github.com/argoproj/argo-cd/issues/1367)
- Support for BitBucket Server webhook events [\#1269](https://github.com/argoproj/argo-cd/issues/1269)

**Fixed bugs:**

- CLI - Can list apps but can't sync [\#1839](https://github.com/argoproj/argo-cd/issues/1839)
- Running app actions requires only read privileges [\#1827](https://github.com/argoproj/argo-cd/issues/1827)
- Forked tool processes should timeout [\#1821](https://github.com/argoproj/argo-cd/issues/1821)
- Fix API server to repo server timeout [\#1820](https://github.com/argoproj/argo-cd/issues/1820)
- Diffing Customization is not applied as expected [\#1815](https://github.com/argoproj/argo-cd/issues/1815)
- Argo CD drops new K8s v1.13+ fields during deploy/comparison [\#1781](https://github.com/argoproj/argo-cd/issues/1781)
- Unable to create application from private SSH repository: key is unknown [\#1743](https://github.com/argoproj/argo-cd/issues/1743)
- Argo CD don't handle well k8s objects which exceeds 1mb [\#1685](https://github.com/argoproj/argo-cd/issues/1685)
- Argo CD reformats own ConfigMap which make it is difficult to manage Argo CD declaratively [\#1676](https://github.com/argoproj/argo-cd/issues/1676)
- Context deadline exceeded when having more than 30 apps [\#1641](https://github.com/argoproj/argo-cd/issues/1641)
- Argo-cd synchronisation hangs when using PreSync hooks [\#1594](https://github.com/argoproj/argo-cd/issues/1594)

## [v1.1.0-rc2](https://github.com/argoproj/argo-cd/tree/v1.1.0-rc2) (2019-06-21)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.1.0-rc1...v1.1.0-rc2)

**Implemented enhancements:**

- Fix TestKustomize2AppSource [\#1800](https://github.com/argoproj/argo-cd/issues/1800)
- Improve messaging around non existent app paths [\#1778](https://github.com/argoproj/argo-cd/issues/1778)
- Improve documentation on resetting password [\#1775](https://github.com/argoproj/argo-cd/issues/1775)
- Support openshift oc client when applying resources \(i.e. routes\) [\#1734](https://github.com/argoproj/argo-cd/issues/1734)
- Add support for adding Jsonnet TLA arguments to ArgoCD CLI [\#1626](https://github.com/argoproj/argo-cd/issues/1626)
- Option to change namespace for "argocd-manager" ServiceAccount [\#1595](https://github.com/argoproj/argo-cd/issues/1595)
- Change model.conf to support email instead sub field [\#1444](https://github.com/argoproj/argo-cd/issues/1444)
- Improve rollback usability [\#1355](https://github.com/argoproj/argo-cd/issues/1355)
- display targetVersion of Application in Applications dashboard [\#1239](https://github.com/argoproj/argo-cd/issues/1239)
- Store more commit metadata in application revision history [\#1219](https://github.com/argoproj/argo-cd/issues/1219)
- Ability to rotate the bearer token used to manage external clusters [\#1084](https://github.com/argoproj/argo-cd/issues/1084)
- How to sync local changes with argocd [\#839](https://github.com/argoproj/argo-cd/issues/839)
- \[UI\] UI should convert UTC to local browser time [\#819](https://github.com/argoproj/argo-cd/issues/819)

**Fixed bugs:**

- Cannot delete my app created with an older version of ArgoCD [\#1790](https://github.com/argoproj/argo-cd/issues/1790)
- application controller needs a liveness probe [\#1782](https://github.com/argoproj/argo-cd/issues/1782)
- Data race in app controller [\#1774](https://github.com/argoproj/argo-cd/issues/1774)
- App Controller fails to save application due to spec validation error [\#1764](https://github.com/argoproj/argo-cd/issues/1764)
- Helm repositories do not support hot-reload [\#1751](https://github.com/argoproj/argo-cd/issues/1751)
- rpc error: code = Unknown desc -  when trying to create new application in ArgoCD. [\#1719](https://github.com/argoproj/argo-cd/issues/1719)
- \[bug\] no UI error when fail during kustomize build [\#1651](https://github.com/argoproj/argo-cd/issues/1651)
- Application CRD validation errors hard to diagnose from UI/CLI [\#1624](https://github.com/argoproj/argo-cd/issues/1624)
- \[usability\] error, but marked green [\#1587](https://github.com/argoproj/argo-cd/issues/1587)
- argocd-server crash constantly [\#1569](https://github.com/argoproj/argo-cd/issues/1569)
- I can't remove Application with error [\#1568](https://github.com/argoproj/argo-cd/issues/1568)
- hard to view the complete error message on SyncFailed [\#1486](https://github.com/argoproj/argo-cd/issues/1486)
- Support hot-reload for resource.exclusions [\#1199](https://github.com/argoproj/argo-cd/issues/1199)
- Unclear error messages in Sync result and app conditions [\#757](https://github.com/argoproj/argo-cd/issues/757)

**Closed issues:**

- Sync Qs [\#1461](https://github.com/argoproj/argo-cd/issues/1461)
- \[UI\] color indication for pending, unhealthy, unready pods [\#1056](https://github.com/argoproj/argo-cd/issues/1056)
- Update admin password fails on v0.10.6, kube v1.12.x [\#785](https://github.com/argoproj/argo-cd/issues/785)

## [v1.1.0-rc1](https://github.com/argoproj/argo-cd/tree/v1.1.0-rc1) (2019-06-14)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.0.2...v1.1.0-rc1)

**Implemented enhancements:**

- Added ability to sync specific labels from the command line [\#1501](https://github.com/argoproj/argo-cd/pull/1501) ([simster7](https://github.com/simster7))

**Merged pull requests:**

- Cluster registration inadvertently carried forward client-cert auth credentials [\#1742](https://github.com/argoproj/argo-cd/pull/1742) ([jessesuen](https://github.com/jessesuen))
- Support parameterizing argocd base image [\#1741](https://github.com/argoproj/argo-cd/pull/1741) ([alexmt](https://github.com/alexmt))
- Added `--async` flag to `argocd app sync` [\#1738](https://github.com/argoproj/argo-cd/pull/1738) ([simster7](https://github.com/simster7))
- Add Optoro to list of users [\#1737](https://github.com/argoproj/argo-cd/pull/1737) ([spencergilbert](https://github.com/spencergilbert))
- Adding Volvo Cars as officially using ArgoCD [\#1735](https://github.com/argoproj/argo-cd/pull/1735) ([iremmats](https://github.com/iremmats))
- Adds support for SSH keys with Kustomize remote bases WIP [\#1733](https://github.com/argoproj/argo-cd/pull/1733) ([alexec](https://github.com/alexec))
- Setting Helm params with commas now updates local spec [\#1732](https://github.com/argoproj/argo-cd/pull/1732) ([simster7](https://github.com/simster7))
- No longer waits for healthy before completing sync op. Closes \#1715 [\#1727](https://github.com/argoproj/argo-cd/pull/1727) ([alexec](https://github.com/alexec))
- Issue \#1375 - Error view instead of blank page in UI [\#1726](https://github.com/argoproj/argo-cd/pull/1726) ([AdityaGupta1](https://github.com/AdityaGupta1))
- Add ui/node\_modules to docker ignore [\#1725](https://github.com/argoproj/argo-cd/pull/1725) ([alexmt](https://github.com/alexmt))
- Issue \#1533 - Add e2e tests for self-referenced app edge case [\#1724](https://github.com/argoproj/argo-cd/pull/1724) ([alexmt](https://github.com/alexmt))
- Fix key generation loop when running server on insecure mode [\#1723](https://github.com/argoproj/argo-cd/pull/1723) ([simster7](https://github.com/simster7))
- Issue \#1693 - Project Editor: Whitelisted Cluster Resources doesn't s… [\#1722](https://github.com/argoproj/argo-cd/pull/1722) ([AdityaGupta1](https://github.com/AdityaGupta1))
- Order users alphabetically [\#1721](https://github.com/argoproj/argo-cd/pull/1721) ([alexec](https://github.com/alexec))
- Fixes non-escaped comma bug on Helm command arguments [\#1720](https://github.com/argoproj/argo-cd/pull/1720) ([simster7](https://github.com/simster7))
- Issue \#1711 - Upgrade argo ui version to get dropdown fix [\#1717](https://github.com/argoproj/argo-cd/pull/1717) ([alexmt](https://github.com/alexmt))
- Forward git credentials to config management plugins. Closes \#1628 [\#1716](https://github.com/argoproj/argo-cd/pull/1716) ([alexec](https://github.com/alexec))
- Issue \#1677 - Allow users to define app specific urls to expose in the UI [\#1714](https://github.com/argoproj/argo-cd/pull/1714) ([AdityaGupta1](https://github.com/AdityaGupta1))
- Issue \#1701 - UI will crash when create application without destination namespace [\#1713](https://github.com/argoproj/argo-cd/pull/1713) ([alexmt](https://github.com/alexmt))
- Adding Telsa to list of users [\#1712](https://github.com/argoproj/argo-cd/pull/1712) ([igaskin](https://github.com/igaskin))
- Adds documentation around repo connections [\#1709](https://github.com/argoproj/argo-cd/pull/1709) ([alexec](https://github.com/alexec))
- Restore reposerver in Procfile [\#1708](https://github.com/argoproj/argo-cd/pull/1708) ([alexmt](https://github.com/alexmt))
- Updates generated code [\#1707](https://github.com/argoproj/argo-cd/pull/1707) ([alexec](https://github.com/alexec))
- Add Mirantis as an official user [\#1702](https://github.com/argoproj/argo-cd/pull/1702) ([pbrit](https://github.com/pbrit))
- Handle nil obj when processing custom actions [\#1700](https://github.com/argoproj/argo-cd/pull/1700) ([dthomson25](https://github.com/dthomson25))
- Account for missing fields in Rollout HealthStatus [\#1699](https://github.com/argoproj/argo-cd/pull/1699) ([dthomson25](https://github.com/dthomson25))
- Name e2e apps after the test they run for, rather than random ID. [\#1698](https://github.com/argoproj/argo-cd/pull/1698) ([alexec](https://github.com/alexec))
- Move generated api code to pkg package [\#1696](https://github.com/argoproj/argo-cd/pull/1696) ([alexmt](https://github.com/alexmt))
- Bump base version to 1.0.1 for cluster-install [\#1695](https://github.com/argoproj/argo-cd/pull/1695) ([narg95](https://github.com/narg95))
- Adds custom port repo note [\#1694](https://github.com/argoproj/argo-cd/pull/1694) ([vaisov](https://github.com/vaisov))
- Improve Circle CI builds [\#1691](https://github.com/argoproj/argo-cd/pull/1691) ([alexec](https://github.com/alexec))
- Documents HA/DR [\#1690](https://github.com/argoproj/argo-cd/pull/1690) ([alexec](https://github.com/alexec))
- Tidy up \#1684 [\#1689](https://github.com/argoproj/argo-cd/pull/1689) ([alexec](https://github.com/alexec))
- Merge UI [\#1688](https://github.com/argoproj/argo-cd/pull/1688) ([alexec](https://github.com/alexec))
- add tZERO to organizations list [\#1686](https://github.com/argoproj/argo-cd/pull/1686) ([stevesea](https://github.com/stevesea))
- Add support for adding Kustomize commonLabels in Applications [\#1684](https://github.com/argoproj/argo-cd/pull/1684) ([twz123](https://github.com/twz123))
- Support to override helm release name [\#1682](https://github.com/argoproj/argo-cd/pull/1682) ([lcostea](https://github.com/lcostea))
- Update SUPPORT.md [\#1681](https://github.com/argoproj/argo-cd/pull/1681) ([alexec](https://github.com/alexec))
- Adds Prune=false and IgnoreExtraneous options [\#1680](https://github.com/argoproj/argo-cd/pull/1680) ([alexec](https://github.com/alexec))
- Added Codility to ArgoCD users [\#1679](https://github.com/argoproj/argo-cd/pull/1679) ([mjasion](https://github.com/mjasion))
- codegen [\#1674](https://github.com/argoproj/argo-cd/pull/1674) ([alexec](https://github.com/alexec))
- Issue \#1668 - Replicasets ordering is not stable on app tree view [\#1669](https://github.com/argoproj/argo-cd/pull/1669) ([alexmt](https://github.com/alexmt))
- Improve tests [\#1667](https://github.com/argoproj/argo-cd/pull/1667) ([alexmt](https://github.com/alexmt))
- Issue \#1665 - Stuck processor on App Controller after deleting application with incomplete operation [\#1666](https://github.com/argoproj/argo-cd/pull/1666) ([alexmt](https://github.com/alexmt))
- Adds docs about app deletion [\#1664](https://github.com/argoproj/argo-cd/pull/1664) ([alexec](https://github.com/alexec))
- Add ability to specify system namespace during cluster add operation [\#1661](https://github.com/argoproj/argo-cd/pull/1661) ([jannfis](https://github.com/jannfis))
- Terminates op before delete [\#1658](https://github.com/argoproj/argo-cd/pull/1658) ([alexec](https://github.com/alexec))
- Update releasing.md [\#1657](https://github.com/argoproj/argo-cd/pull/1657) ([alexec](https://github.com/alexec))
- Updated templates [\#1654](https://github.com/argoproj/argo-cd/pull/1654) ([alexec](https://github.com/alexec))
- 🎉 KintoHub using ArgoCD [\#1650](https://github.com/argoproj/argo-cd/pull/1650) ([bappr](https://github.com/bappr))
- Make listener and metrics ports configurable [\#1647](https://github.com/argoproj/argo-cd/pull/1647) ([jannfis](https://github.com/jannfis))
- Add END. to adopters in README.md [\#1643](https://github.com/argoproj/argo-cd/pull/1643) ([samcgardner](https://github.com/samcgardner))
- Build ArgoCD on CircleCI [\#1635](https://github.com/argoproj/argo-cd/pull/1635) ([alexec](https://github.com/alexec))
- Sync wave [\#1634](https://github.com/argoproj/argo-cd/pull/1634) ([alexec](https://github.com/alexec))
- Public git creds [\#1633](https://github.com/argoproj/argo-cd/pull/1633) ([alexec](https://github.com/alexec))
- Codegen [\#1632](https://github.com/argoproj/argo-cd/pull/1632) ([alexec](https://github.com/alexec))
- Make build options in Makefile settable from environment [\#1619](https://github.com/argoproj/argo-cd/pull/1619) ([jannfis](https://github.com/jannfis))
- Update v1.0.0 change log [\#1618](https://github.com/argoproj/argo-cd/pull/1618) ([alexmt](https://github.com/alexmt))
- Fixes e2e tests. Closes \#1616. [\#1617](https://github.com/argoproj/argo-cd/pull/1617) ([alexec](https://github.com/alexec))
- Supply resourceVersion to watch request to prevent reading of stale cache [\#1612](https://github.com/argoproj/argo-cd/pull/1612) ([jessesuen](https://github.com/jessesuen))
- Issue \#1533 - Prevent reconciliation loop for self-managed apps [\#1608](https://github.com/argoproj/argo-cd/pull/1608) ([alexmt](https://github.com/alexmt))
- Updates codegen [\#1601](https://github.com/argoproj/argo-cd/pull/1601) ([alexec](https://github.com/alexec))
- E2e test infra [\#1600](https://github.com/argoproj/argo-cd/pull/1600) ([alexec](https://github.com/alexec))
- Issue \#1586 - Ignore patch errors during diffing normalization [\#1599](https://github.com/argoproj/argo-cd/pull/1599) ([alexmt](https://github.com/alexmt))
- Updates issue template and Makefile [\#1598](https://github.com/argoproj/argo-cd/pull/1598) ([alexec](https://github.com/alexec))
- Issue \#1596 - SSH URLs support is partially broken [\#1597](https://github.com/argoproj/argo-cd/pull/1597) ([alexmt](https://github.com/alexmt))
- Issue \#1471 - Support configuring requested OIDC provider scopes and enforced RBAC scopes [\#1585](https://github.com/argoproj/argo-cd/pull/1585) ([alexmt](https://github.com/alexmt))
- Issue \#1552 - Improve rendering app image information [\#1584](https://github.com/argoproj/argo-cd/pull/1584) ([alexmt](https://github.com/alexmt))
- Added logout ability \(`argocd logout`\) [\#1582](https://github.com/argoproj/argo-cd/pull/1582) ([simster7](https://github.com/simster7))
- Issue \#1579 - Impossible to sync to HEAD from UI if auto-sync is enabled [\#1580](https://github.com/argoproj/argo-cd/pull/1580) ([alexmt](https://github.com/alexmt))
- add commonbond to users of argocd [\#1577](https://github.com/argoproj/argo-cd/pull/1577) ([jpresky](https://github.com/jpresky))
- Fix ingress browsable url formatting if port is not string [\#1576](https://github.com/argoproj/argo-cd/pull/1576) ([alexmt](https://github.com/alexmt))
- Issue \#1352 - Dedupe live resourced by UID instead of group/kind/namespace/name [\#1575](https://github.com/argoproj/argo-cd/pull/1575) ([alexmt](https://github.com/alexmt))
- Documents Kustomize. Closes \#1566 [\#1572](https://github.com/argoproj/argo-cd/pull/1572) ([alexec](https://github.com/alexec))
- Add GMETRI to organizations using ArgoCD [\#1564](https://github.com/argoproj/argo-cd/pull/1564) ([sahil87](https://github.com/sahil87))
- Issue \#1546 - Add liveness probe to repo server/api servers [\#1560](https://github.com/argoproj/argo-cd/pull/1560) ([alexmt](https://github.com/alexmt))
- Issue \#1540 - Fix kustomize manifest generation crash if manifest has image without version [\#1559](https://github.com/argoproj/argo-cd/pull/1559) ([alexmt](https://github.com/alexmt))
- Issue \#1557 - Controller incorrectly report health state of self managed application [\#1558](https://github.com/argoproj/argo-cd/pull/1558) ([alexmt](https://github.com/alexmt))
- Fix hardcoded 'git' user in `util/git.NewClient` [\#1556](https://github.com/argoproj/argo-cd/pull/1556) ([pbrit](https://github.com/pbrit))
- Improve Rollout health.lua [\#1554](https://github.com/argoproj/argo-cd/pull/1554) ([dthomson25](https://github.com/dthomson25))
- Fix invalid URL for ingress without hostname [\#1553](https://github.com/argoproj/argo-cd/pull/1553) ([alexmt](https://github.com/alexmt))
- Updates manifests. Closes \#1520 [\#1549](https://github.com/argoproj/argo-cd/pull/1549) ([alexec](https://github.com/alexec))
- Issue \#1533 - Prevent reconciliation loop for self-managed apps [\#1547](https://github.com/argoproj/argo-cd/pull/1547) ([alexmt](https://github.com/alexmt))
- Rollout health checks/actions should support v0.2 and v0.2+ versions [\#1543](https://github.com/argoproj/argo-cd/pull/1543) ([alexmt](https://github.com/alexmt))
- Fixes bug in normalizer [\#1542](https://github.com/argoproj/argo-cd/pull/1542) ([alexec](https://github.com/alexec))
- Add kustomize [\#1541](https://github.com/argoproj/argo-cd/pull/1541) ([OmerKahani](https://github.com/OmerKahani))
- fix typo in best practices [\#1538](https://github.com/argoproj/argo-cd/pull/1538) ([tom-256](https://github.com/tom-256))
- Adds missing section to docs [\#1537](https://github.com/argoproj/argo-cd/pull/1537) ([alexec](https://github.com/alexec))
- Issue 1476 - Avoid validating repository in application controller [\#1535](https://github.com/argoproj/argo-cd/pull/1535) ([alexmt](https://github.com/alexmt))
- Update CONTRIBUTING.md [\#1534](https://github.com/argoproj/argo-cd/pull/1534) ([alexec](https://github.com/alexec))
- Documents cluster bootstrapping. Close \#1481 [\#1530](https://github.com/argoproj/argo-cd/pull/1530) ([alexec](https://github.com/alexec))
- Fix flaky TestGetIngressInfo unit test [\#1529](https://github.com/argoproj/argo-cd/pull/1529) ([alexmt](https://github.com/alexmt))
- Issue \#1476 - Add repo server grpc call timeout [\#1528](https://github.com/argoproj/argo-cd/pull/1528) ([alexmt](https://github.com/alexmt))
- Issue \#1414 - Load target resource using K8S if conversion fails [\#1527](https://github.com/argoproj/argo-cd/pull/1527) ([alexmt](https://github.com/alexmt))
- Fix e2e [\#1526](https://github.com/argoproj/argo-cd/pull/1526) ([alexec](https://github.com/alexec))
- Ingress resource might get invalid ExternalURL \(\#1522\) [\#1523](https://github.com/argoproj/argo-cd/pull/1523) ([alexmt](https://github.com/alexmt))
- codegen [\#1521](https://github.com/argoproj/argo-cd/pull/1521) ([alexec](https://github.com/alexec))
- Add Network View description to changelog [\#1519](https://github.com/argoproj/argo-cd/pull/1519) ([alexmt](https://github.com/alexmt))
- Updated CHANGELOG.md [\#1518](https://github.com/argoproj/argo-cd/pull/1518) ([alexec](https://github.com/alexec))
- Whitelisting of resources [\#1509](https://github.com/argoproj/argo-cd/pull/1509) ([simster7](https://github.com/simster7))
- Adds support for configuring repo creds at a domain/org level. Closes… [\#1496](https://github.com/argoproj/argo-cd/pull/1496) ([alexec](https://github.com/alexec))

## [v1.0.2](https://github.com/argoproj/argo-cd/tree/v1.0.2) (2019-06-14)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.0.1...v1.0.2)

**Implemented enhancements:**

- Remove Server::CreateFromKubeConfig\(\) dead code [\#1739](https://github.com/argoproj/argo-cd/issues/1739)
- Document how to add Gitlab RBAC policies [\#1731](https://github.com/argoproj/argo-cd/issues/1731)
- argocd app sync APPNAME --async [\#1728](https://github.com/argoproj/argo-cd/issues/1728)
- Updates contribution guide to state versions of tools and add PR template [\#1706](https://github.com/argoproj/argo-cd/issues/1706)
- Project Editor: Whitelisted Cluster Resources doesn't strip whitespace [\#1693](https://github.com/argoproj/argo-cd/issues/1693)
- Allow users to define app specific urls to expose in the UI [\#1677](https://github.com/argoproj/argo-cd/issues/1677)
- Provide ability to ignore extra application objects \(e.g. kustomize nameSuffixHash\) [\#1629](https://github.com/argoproj/argo-cd/issues/1629)
- Forward git credentials to config management plugins [\#1628](https://github.com/argoproj/argo-cd/issues/1628)
- Improve Kustomize 2 parameters UI [\#1609](https://github.com/argoproj/argo-cd/issues/1609)
- kustomize access to private repo [\#1567](https://github.com/argoproj/argo-cd/issues/1567)
- Sync waves [\#1544](https://github.com/argoproj/argo-cd/issues/1544)
- diff for k8s secrets using CLI [\#1442](https://github.com/argoproj/argo-cd/issues/1442)
- Set kustomize image parameter through CLI [\#1400](https://github.com/argoproj/argo-cd/issues/1400)
- Error view instead of blank page in UI [\#1375](https://github.com/argoproj/argo-cd/issues/1375)
- Allow resources to be excluded from sync via annotation [\#1373](https://github.com/argoproj/argo-cd/issues/1373)
- Github integration? [\#1091](https://github.com/argoproj/argo-cd/issues/1091)
- Application controller should expose more cluster status [\#1011](https://github.com/argoproj/argo-cd/issues/1011)
- Using SSH keys to authenticate kustomize bases from git [\#827](https://github.com/argoproj/argo-cd/issues/827)

**Fixed bugs:**

- Sync operation unnecessary waits for a healthy state of all resources [\#1715](https://github.com/argoproj/argo-cd/issues/1715)
- \[ui small bug\] menu position outside block [\#1711](https://github.com/argoproj/argo-cd/issues/1711)
- \[ui bug\] healthy or not [\#1710](https://github.com/argoproj/argo-cd/issues/1710)
- Deployment is still progressing since Ingress resource is not considered as synced [\#1704](https://github.com/argoproj/argo-cd/issues/1704)
- UI will crash when create application without destination namespace [\#1701](https://github.com/argoproj/argo-cd/issues/1701)
- ArgoCD synchronization failed due to internal error [\#1697](https://github.com/argoproj/argo-cd/issues/1697)
- Persistent sync error: Unexpected error &http.httpError{err:"context deadline exceeded \(Client.Timeout exceeded while reading body\)" [\#1683](https://github.com/argoproj/argo-cd/issues/1683)
- Unable to use helm CLI against helm releases that have been deployed by ArgoCD [\#1672](https://github.com/argoproj/argo-cd/issues/1672)
- cant create new app [\#1671](https://github.com/argoproj/argo-cd/issues/1671)
- Replicasets ordering is not stable on app tree view [\#1668](https://github.com/argoproj/argo-cd/issues/1668)
- Stuck processor on App Controller after deleting application with incomplete operation [\#1665](https://github.com/argoproj/argo-cd/issues/1665)
- Role edit page fails with JS error [\#1662](https://github.com/argoproj/argo-cd/issues/1662)
- failed parsing on parameters with comma [\#1660](https://github.com/argoproj/argo-cd/issues/1660)
- \[bug\] not enough hardware in cluster, no UI error, issue after add hardware to cluster [\#1652](https://github.com/argoproj/argo-cd/issues/1652)
- Unable to use ssh gitlab connection for on prem gitlab [\#1322](https://github.com/argoproj/argo-cd/issues/1322)

**Closed issues:**

- Is it possible to run argocd on ARM ?  [\#1565](https://github.com/argoproj/argo-cd/issues/1565)
- Merge UI into main code base [\#1421](https://github.com/argoproj/argo-cd/issues/1421)
- Support EKS clusters with IAM creds [\#1304](https://github.com/argoproj/argo-cd/issues/1304)
- `argocd logout` [\#1210](https://github.com/argoproj/argo-cd/issues/1210)
- Make it possible to set release name different from argcd app name.  [\#1066](https://github.com/argoproj/argo-cd/issues/1066)
- Backup state [\#1048](https://github.com/argoproj/argo-cd/issues/1048)
- Migrate argocd to another k8s cluster [\#993](https://github.com/argoproj/argo-cd/issues/993)

## [v1.0.1](https://github.com/argoproj/argo-cd/tree/v1.0.1) (2019-05-28)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.0.0...v1.0.1)

**Implemented enhancements:**

- Ignore an entire Kind [\#1649](https://github.com/argoproj/argo-cd/issues/1649)
- CLI argocd app sync APPNAME -l foo=bar [\#1241](https://github.com/argoproj/argo-cd/issues/1241)

**Fixed bugs:**

- OIDC with RBAC not functioning [\#1645](https://github.com/argoproj/argo-cd/issues/1645)
- Cannot delete app that is already in operation [\#1644](https://github.com/argoproj/argo-cd/issues/1644)
- Possible goroutine leak in application controller [\#1627](https://github.com/argoproj/argo-cd/issues/1627)
- Global repo credentials doesn't work unless repo is listed [\#1625](https://github.com/argoproj/argo-cd/issues/1625)

**Closed issues:**

- Promote via GUI [\#1593](https://github.com/argoproj/argo-cd/issues/1593)
- Investigate what users might want for 1st class notifications [\#1480](https://github.com/argoproj/argo-cd/issues/1480)

## [v1.0.0](https://github.com/argoproj/argo-cd/tree/v1.0.0) (2019-05-16)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.0.0-rc3...v1.0.0)

**Implemented enhancements:**

- Fix skipped e2e tests [\#1616](https://github.com/argoproj/argo-cd/issues/1616)
- Support for OpenShift [\#1224](https://github.com/argoproj/argo-cd/issues/1224)

**Fixed bugs:**

- Argo doesn't recongize diffs from kubectl edit [\#1613](https://github.com/argoproj/argo-cd/issues/1613)
- Got error while deploy on listed cluster which connection status successful [\#1607](https://github.com/argoproj/argo-cd/issues/1607)
- Need to handle stale kubernetes watch cache [\#1605](https://github.com/argoproj/argo-cd/issues/1605)
- Could not deploy cert manager stable chart with argocd [\#1591](https://github.com/argoproj/argo-cd/issues/1591)
- UI crash in network view [\#1589](https://github.com/argoproj/argo-cd/issues/1589)
- The 'argocd app wait' might print incorrect finall app status [\#1562](https://github.com/argoproj/argo-cd/issues/1562)
- CLI not working on \(Alpine\) Linux [\#1561](https://github.com/argoproj/argo-cd/issues/1561)
- OpenID Connect using AWS Cognito cannot map Groups [\#1471](https://github.com/argoproj/argo-cd/issues/1471)

**Closed issues:**

- Dedupe live resourced by UID instead of group/kind/namespace/name [\#1352](https://github.com/argoproj/argo-cd/issues/1352)

## [v1.0.0-rc3](https://github.com/argoproj/argo-cd/tree/v1.0.0-rc3) (2019-05-09)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.0.0-rc2...v1.0.0-rc3)

**Implemented enhancements:**

- Ignore existing resources for diffing [\#1472](https://github.com/argoproj/argo-cd/issues/1472)

**Fixed bugs:**

- SSH URLs support is partially broken  [\#1596](https://github.com/argoproj/argo-cd/issues/1596)
- UI crashes when viewing app with no target state [\#1592](https://github.com/argoproj/argo-cd/issues/1592)
- Bad external URL link https://%!s\(\<nil\>\) [\#1590](https://github.com/argoproj/argo-cd/issues/1590)
- Can't ignore a non-existent pointer anymore [\#1586](https://github.com/argoproj/argo-cd/issues/1586)
- UI doesn't load for complex setup [\#1583](https://github.com/argoproj/argo-cd/issues/1583)
- Argo CD server should not generate SSH certificate if started in insecure mode [\#1581](https://github.com/argoproj/argo-cd/issues/1581)
- Impossible to sync to HEAD from UI if auto-sync is enabled [\#1579](https://github.com/argoproj/argo-cd/issues/1579)
- Cannot upgrade from v0.11 to v0.12 [\#1571](https://github.com/argoproj/argo-cd/issues/1571)
- Application controller is unable to delete self-referenced app [\#1570](https://github.com/argoproj/argo-cd/issues/1570)
- \[usability\] sync debugging [\#1566](https://github.com/argoproj/argo-cd/issues/1566)
- Network view crashes if any filter is set [\#1563](https://github.com/argoproj/argo-cd/issues/1563)
- Controller incorrectly report health state of self managed application [\#1557](https://github.com/argoproj/argo-cd/issues/1557)
- Adding a private repository with non-git user fails [\#1555](https://github.com/argoproj/argo-cd/issues/1555)
- Container Image URL name does not fit into tooltip [\#1552](https://github.com/argoproj/argo-cd/issues/1552)
- boolean Helm parameters in Application CRD make UI/CLI unusable [\#1551](https://github.com/argoproj/argo-cd/issues/1551)
- UI don't allow to see details of ClusterRole/ClusterRoleBinding with `:` character in name [\#1550](https://github.com/argoproj/argo-cd/issues/1550)
- The argocd-server, argocd-repo-server need liveness probe [\#1546](https://github.com/argoproj/argo-cd/issues/1546)
- StatefulSet with initContainers gives syntax error in 0.12.1 and up [\#1540](https://github.com/argoproj/argo-cd/issues/1540)
- When using kustomize, namePrefix should not be required [\#1520](https://github.com/argoproj/argo-cd/issues/1520)
- `argocd app sync APPNAME --revision ...` doesn't seem to work [\#1512](https://github.com/argoproj/argo-cd/issues/1512)
- Issue while connecting repo on kubernetes running on minikube + ec2 instance [\#1510](https://github.com/argoproj/argo-cd/issues/1510)

## [v1.0.0-rc2](https://github.com/argoproj/argo-cd/tree/v1.0.0-rc2) (2019-04-30)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.12.3...v1.0.0-rc2)

**Fixed bugs:**

- Authentication required when creating an application with GitLab deploy tokens [\#1491](https://github.com/argoproj/argo-cd/issues/1491)

## [v0.12.3](https://github.com/argoproj/argo-cd/tree/v0.12.3) (2019-04-30)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v1.0.0-rc1...v0.12.3)

**Implemented enhancements:**

- Implement whitelist option analogous to `resource.exclusions` [\#1490](https://github.com/argoproj/argo-cd/issues/1490)
- Documenting setting up private repos [\#1484](https://github.com/argoproj/argo-cd/issues/1484)
- Document the patterns for cluster bootstrapping [\#1481](https://github.com/argoproj/argo-cd/issues/1481)
- Sync failure hook [\#1478](https://github.com/argoproj/argo-cd/issues/1478)
- Ability to configure git credentials at a domain level [\#1332](https://github.com/argoproj/argo-cd/issues/1332)

**Fixed bugs:**

- Ingress resource might get invalid ExternalURL [\#1522](https://github.com/argoproj/argo-cd/issues/1522)
- How to actually deploy anything at all? [\#1516](https://github.com/argoproj/argo-cd/issues/1516)
- how to add private gitlab repo? [\#1493](https://github.com/argoproj/argo-cd/issues/1493)
- ArgoCd becomes unresponsive in Failed to watch {...} Start watch {...} loop [\#1476](https://github.com/argoproj/argo-cd/issues/1476)
- ComparisonError while try to deploy k8s resource kind: APIService [\#1414](https://github.com/argoproj/argo-cd/issues/1414)
- Synchronization failed when using non existing application namespace on cluster resources [\#1045](https://github.com/argoproj/argo-cd/issues/1045)

## [v1.0.0-rc1](https://github.com/argoproj/argo-cd/tree/v1.0.0-rc1) (2019-04-24)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.12.2...v1.0.0-rc1)

**Implemented enhancements:**

- Support for customizable resource actions as Lua scripts [\#86](https://github.com/argoproj/argo-cd/issues/86)

**Fixed bugs:**

- Selective sync is broken in UI [\#1507](https://github.com/argoproj/argo-cd/issues/1507)
- UI fails to load custom actions is resource is not deployed [\#1502](https://github.com/argoproj/argo-cd/issues/1502)
- External app links are not accurate [\#1499](https://github.com/argoproj/argo-cd/issues/1499)
- Unable to create an app from private repository [\#1411](https://github.com/argoproj/argo-cd/issues/1411)

**Closed issues:**

- No instructions for private repositry [\#1342](https://github.com/argoproj/argo-cd/issues/1342)

**Merged pull requests:**

- Update min client version and cache version to 1.0.0 [\#1517](https://github.com/argoproj/argo-cd/pull/1517) ([alexmt](https://github.com/alexmt))
- Issue \#1411 - Document private repository configuration [\#1515](https://github.com/argoproj/argo-cd/pull/1515) ([alexmt](https://github.com/alexmt))
- Issue \#1499 - Use ingress host information to populate application external URL [\#1511](https://github.com/argoproj/argo-cd/pull/1511) ([alexmt](https://github.com/alexmt))
- v0.12.2 Change log [\#1508](https://github.com/argoproj/argo-cd/pull/1508) ([alexmt](https://github.com/alexmt))
- Allow empty. Close \#1504 [\#1506](https://github.com/argoproj/argo-cd/pull/1506) ([alexec](https://github.com/alexec))
- Update CHANGELOG.md [\#1500](https://github.com/argoproj/argo-cd/pull/1500) ([alexec](https://github.com/alexec))
- Add riskified to organizations using ArgoCD [\#1497](https://github.com/argoproj/argo-cd/pull/1497) ([OmerKahani](https://github.com/OmerKahani))
- Issue \#86 - Custom actions bug fixing [\#1494](https://github.com/argoproj/argo-cd/pull/1494) ([alexmt](https://github.com/alexmt))
- Fix flaky e2e test. Again [\#1489](https://github.com/argoproj/argo-cd/pull/1489) ([alexmt](https://github.com/alexmt))
- Change loggin level in util function to Debug [\#1488](https://github.com/argoproj/argo-cd/pull/1488) ([alexmt](https://github.com/alexmt))
-  Issue \#1476 - Fix racing condition in controller cache [\#1485](https://github.com/argoproj/argo-cd/pull/1485) ([alexmt](https://github.com/alexmt))
- Adds a faster way to run e2e locally [\#1475](https://github.com/argoproj/argo-cd/pull/1475) ([alexec](https://github.com/alexec))
- Fix flaky e2e test [\#1474](https://github.com/argoproj/argo-cd/pull/1474) ([alexmt](https://github.com/alexmt))
- Change version to 1.0.0 [\#1473](https://github.com/argoproj/argo-cd/pull/1473) ([alexmt](https://github.com/alexmt))
- Updates CHANGELOG for v1.0.0 [\#1469](https://github.com/argoproj/argo-cd/pull/1469) ([alexec](https://github.com/alexec))
- Adds label to Github issue templates [\#1468](https://github.com/argoproj/argo-cd/pull/1468) ([alexec](https://github.com/alexec))
- Document steps to troubleshot cluster configuration [\#1467](https://github.com/argoproj/argo-cd/pull/1467) ([alexmt](https://github.com/alexmt))
- Add e2e tests for app with secrets [\#1466](https://github.com/argoproj/argo-cd/pull/1466) ([alexmt](https://github.com/alexmt))
- grammar change. added an 'if' [\#1465](https://github.com/argoproj/argo-cd/pull/1465) ([ryanfernandes09](https://github.com/ryanfernandes09))
- Issue \#1326 - Rollback UI is not showing correct ksonnet parameters in preview [\#1464](https://github.com/argoproj/argo-cd/pull/1464) ([alexmt](https://github.com/alexmt))
- more-information-needed [\#1463](https://github.com/argoproj/argo-cd/pull/1463) ([alexec](https://github.com/alexec))
- Fix e2e test flakyness [\#1462](https://github.com/argoproj/argo-cd/pull/1462) ([alexmt](https://github.com/alexmt))
- don't compare secrets, since argo-cd doesn't have access to their data [\#1459](https://github.com/argoproj/argo-cd/pull/1459) ([marcb1](https://github.com/marcb1))
- Fixes goroutine leak. Closes \#1381 [\#1457](https://github.com/argoproj/argo-cd/pull/1457) ([alexec](https://github.com/alexec))
- Add link to e2e testing on contributing guide [\#1456](https://github.com/argoproj/argo-cd/pull/1456) ([dthomson25](https://github.com/dthomson25))
- Generate random name for grpc proxy unix socket file instead of time stamp [\#1455](https://github.com/argoproj/argo-cd/pull/1455) ([alexmt](https://github.com/alexmt))
- Regenerate manifests [\#1454](https://github.com/argoproj/argo-cd/pull/1454) ([alexmt](https://github.com/alexmt))
- Added `--resource` flag to `argocd app wait` [\#1453](https://github.com/argoproj/argo-cd/pull/1453) ([simster7](https://github.com/simster7))
- Revert "Redis mastergroup name should be resolvable and argocd-redis-ha is" [\#1452](https://github.com/argoproj/argo-cd/pull/1452) ([alexmt](https://github.com/alexmt))
- Redis mastergroup name should be resolvable and argocd-redis-ha is [\#1450](https://github.com/argoproj/argo-cd/pull/1450) ([KarstenSiemer](https://github.com/KarstenSiemer))
- Issue \#1446 - Delete helm temp directories [\#1449](https://github.com/argoproj/argo-cd/pull/1449) ([alexmt](https://github.com/alexmt))
- Document how to use helm without internet access [\#1448](https://github.com/argoproj/argo-cd/pull/1448) ([alexmt](https://github.com/alexmt))
- Fix github reference to use mainline instead of fork [\#1445](https://github.com/argoproj/argo-cd/pull/1445) ([backjo](https://github.com/backjo))
- Issue \#1389 - Fix null pointer exception in secret normalization function [\#1443](https://github.com/argoproj/argo-cd/pull/1443) ([alexmt](https://github.com/alexmt))
- Docs [\#1441](https://github.com/argoproj/argo-cd/pull/1441) ([alexec](https://github.com/alexec))
- Introduces new RBAC permissions that are required for changing cluste… [\#1440](https://github.com/argoproj/argo-cd/pull/1440) ([alexec](https://github.com/alexec))
- Fix invalid ignoreDifferences config example [\#1437](https://github.com/argoproj/argo-cd/pull/1437) ([alexmt](https://github.com/alexmt))
- Overlay selector of argocd-redis-ha service [\#1436](https://github.com/argoproj/argo-cd/pull/1436) ([KarstenSiemer](https://github.com/KarstenSiemer))
- Shows the health of the application. Closes \#1433 [\#1434](https://github.com/argoproj/argo-cd/pull/1434) ([alexec](https://github.com/alexec))
- Issue \#1425 - Argo CD should not delete CRDs [\#1428](https://github.com/argoproj/argo-cd/pull/1428) ([alexmt](https://github.com/alexmt))
- Add CyberAgent and OpenSaaS Studio to organizations using ArgoCD [\#1427](https://github.com/argoproj/argo-cd/pull/1427) ([nghialv](https://github.com/nghialv))
- Displays resources that are being deleted as "Progressing". Closes \#1410 [\#1426](https://github.com/argoproj/argo-cd/pull/1426) ([alexec](https://github.com/alexec))
- Add Ticketmaster to "Who uses" section of README [\#1424](https://github.com/argoproj/argo-cd/pull/1424) ([mgoodness](https://github.com/mgoodness))
- Add v0.12.1 release notes [\#1423](https://github.com/argoproj/argo-cd/pull/1423) ([alexmt](https://github.com/alexmt))
- Perform health assessments on all resource nodes in the tree. Closes \#1382 [\#1422](https://github.com/argoproj/argo-cd/pull/1422) ([alexec](https://github.com/alexec))
- Enables Probot stale and no-respones plugins. Closes \#1418 [\#1419](https://github.com/argoproj/argo-cd/pull/1419) ([alexec](https://github.com/alexec))
- Run 'go fmt' for application.go and server.go [\#1417](https://github.com/argoproj/argo-cd/pull/1417) ([alexmt](https://github.com/alexmt))
- Add patch audit [\#1416](https://github.com/argoproj/argo-cd/pull/1416) ([dthomson25](https://github.com/dthomson25))
- Add networking test app [\#1409](https://github.com/argoproj/argo-cd/pull/1409) ([alexmt](https://github.com/alexmt))
- Nils health if the resource does not provide it. Closes \#1383 [\#1408](https://github.com/argoproj/argo-cd/pull/1408) ([alexec](https://github.com/alexec))
- Issue \#1406 - Don't try deleting application resource if it already have deletion timestamp set [\#1407](https://github.com/argoproj/argo-cd/pull/1407) ([alexmt](https://github.com/alexmt))
- Issue \#1404 - App controller unnecessary set namespace to cluster level resources [\#1405](https://github.com/argoproj/argo-cd/pull/1405) ([alexmt](https://github.com/alexmt))
- Fixes doc bugs. Closes \#1395 [\#1403](https://github.com/argoproj/argo-cd/pull/1403) ([alexec](https://github.com/alexec))
- Add KompiTech GmbH to organizations using Argo CD [\#1402](https://github.com/argoproj/argo-cd/pull/1402) ([pjediny](https://github.com/pjediny))
- issue \#1202: docs\(help examples\): adding template and first examples … [\#1398](https://github.com/argoproj/argo-cd/pull/1398) ([igaskin](https://github.com/igaskin))
- Mkdocs2 [\#1393](https://github.com/argoproj/argo-cd/pull/1393) ([alexec](https://github.com/alexec))
- Allow wait to return on health or suspended [\#1392](https://github.com/argoproj/argo-cd/pull/1392) ([dthomson25](https://github.com/dthomson25))
- Create docs website [\#1387](https://github.com/argoproj/argo-cd/pull/1387) ([alexec](https://github.com/alexec))
- Add Yieldlab to organzations using Argo CD [\#1385](https://github.com/argoproj/argo-cd/pull/1385) ([brushmate](https://github.com/brushmate))
- Fix project.yaml link in README.md [\#1384](https://github.com/argoproj/argo-cd/pull/1384) ([mjasion](https://github.com/mjasion))
- Issue \#1366 - Fix null pointer dereference error in 'argocd app wait' [\#1380](https://github.com/argoproj/argo-cd/pull/1380) ([alexmt](https://github.com/alexmt))
- Issue \#1374 - Add k8s objects circular dependency protection to getApp method [\#1379](https://github.com/argoproj/argo-cd/pull/1379) ([alexmt](https://github.com/alexmt))
- Removes componentParameterOverrides. Closes \#1372 [\#1378](https://github.com/argoproj/argo-cd/pull/1378) ([alexec](https://github.com/alexec))
- Magically increase the code coverage!! [\#1370](https://github.com/argoproj/argo-cd/pull/1370) ([dthomson25](https://github.com/dthomson25))
- Initial Custom Actions Implementation [\#1369](https://github.com/argoproj/argo-cd/pull/1369) ([dthomson25](https://github.com/dthomson25))
- Pod health [\#1365](https://github.com/argoproj/argo-cd/pull/1365) ([alexec](https://github.com/alexec))
- Add 'Who uses Argo CD?' section [\#1361](https://github.com/argoproj/argo-cd/pull/1361) ([alexmt](https://github.com/alexmt))
- Issue \#1012 - kubectl v1.13 fails to convert extensions/NetworkPolicy [\#1360](https://github.com/argoproj/argo-cd/pull/1360) ([alexmt](https://github.com/alexmt))
- Images [\#1351](https://github.com/argoproj/argo-cd/pull/1351) ([alexec](https://github.com/alexec))
- Fix Failing Linter [\#1350](https://github.com/argoproj/argo-cd/pull/1350) ([dthomson25](https://github.com/dthomson25))
- Add mapping to new canonical Ingress API group [\#1348](https://github.com/argoproj/argo-cd/pull/1348) ([twz123](https://github.com/twz123))
- Issue \#908 - Surface Service/Ingress external IPs, hostname to application [\#1347](https://github.com/argoproj/argo-cd/pull/1347) ([alexmt](https://github.com/alexmt))
- Copy-paste error: clusterResourceWhitelist -\> namespaceResourceBlacklist [\#1343](https://github.com/argoproj/argo-cd/pull/1343) ([coderanger](https://github.com/coderanger))
- gotestsum [\#1341](https://github.com/argoproj/argo-cd/pull/1341) ([alexec](https://github.com/alexec))
- Issue \#733 - 'argocd app wait' should fail sooner if app transitioned to Degraded state [\#1339](https://github.com/argoproj/argo-cd/pull/1339) ([alexmt](https://github.com/alexmt))
- Don't run lint after running codegen [\#1338](https://github.com/argoproj/argo-cd/pull/1338) ([alexmt](https://github.com/alexmt))
- Issue \#1294 - CLI diff should take into account resource customizations [\#1337](https://github.com/argoproj/argo-cd/pull/1337) ([alexmt](https://github.com/alexmt))
- Issue \#1218 - Allow using any name for secrets which store cluster credentials [\#1336](https://github.com/argoproj/argo-cd/pull/1336) ([alexmt](https://github.com/alexmt))
- Declarative setup doc update [\#1334](https://github.com/argoproj/argo-cd/pull/1334) ([daanemanz](https://github.com/daanemanz))
- Issue \#357 - Expose application nodes networking information [\#1333](https://github.com/argoproj/argo-cd/pull/1333) ([alexmt](https://github.com/alexmt))
- Adds "make build" target, and running lint,build,test [\#1331](https://github.com/argoproj/argo-cd/pull/1331) ([alexec](https://github.com/alexec))
- Lint [\#1330](https://github.com/argoproj/argo-cd/pull/1330) ([alexec](https://github.com/alexec))
- Update argocd-util import/export to support proper backup and restore [\#1328](https://github.com/argoproj/argo-cd/pull/1328) ([jessesuen](https://github.com/jessesuen))
- Fixs deps [\#1325](https://github.com/argoproj/argo-cd/pull/1325) ([alexec](https://github.com/alexec))
- Adds support for `kustomize edit set image`. Closes \#1275 [\#1324](https://github.com/argoproj/argo-cd/pull/1324) ([alexec](https://github.com/alexec))
- Use paused field in rollout health check [\#1321](https://github.com/argoproj/argo-cd/pull/1321) ([alexmt](https://github.com/alexmt))
- Issue \#1319 - Fix invalid group filtering in 'patch-resource' command [\#1320](https://github.com/argoproj/argo-cd/pull/1320) ([alexmt](https://github.com/alexmt))
- Issue \#1135 - Run e2e tests in throw-away kubernetes cluster [\#1318](https://github.com/argoproj/argo-cd/pull/1318) ([alexmt](https://github.com/alexmt))
- Update CHANGELOG.md for v0.12 release [\#1317](https://github.com/argoproj/argo-cd/pull/1317) ([jessesuen](https://github.com/jessesuen))
- Force color logging locally, start-up UI with services [\#1316](https://github.com/argoproj/argo-cd/pull/1316) ([alexec](https://github.com/alexec))
- Use Recreate deployment strategy for controller [\#1315](https://github.com/argoproj/argo-cd/pull/1315) ([jessesuen](https://github.com/jessesuen))
- Fix goroutine leak in RetryUntilSucceed [\#1314](https://github.com/argoproj/argo-cd/pull/1314) ([jessesuen](https://github.com/jessesuen))
- Copy-paste error: clusterResourceWhitelist -\> namespaceResourceBlacklist [\#1312](https://github.com/argoproj/argo-cd/pull/1312) ([twz123](https://github.com/twz123))
- Issue \#1308 - argo diff --local fails if live object does not exist [\#1309](https://github.com/argoproj/argo-cd/pull/1309) ([alexmt](https://github.com/alexmt))
- Support a separate OAuth2 CLI clientID different from server [\#1307](https://github.com/argoproj/argo-cd/pull/1307) ([jessesuen](https://github.com/jessesuen))
- helm honor os env [\#1306](https://github.com/argoproj/argo-cd/pull/1306) ([1337andre](https://github.com/1337andre))
- Unavailable cache should not prevent reconciling/syncing application [\#1303](https://github.com/argoproj/argo-cd/pull/1303) ([alexmt](https://github.com/alexmt))
- Update redis-ha chart to resolve redis failover issues [\#1301](https://github.com/argoproj/argo-cd/pull/1301) ([jessesuen](https://github.com/jessesuen))
- Fix sample dashboard link in metrics doc [\#1299](https://github.com/argoproj/argo-cd/pull/1299) ([alexmt](https://github.com/alexmt))
- Controller don't stop running watches on cluster resync [\#1298](https://github.com/argoproj/argo-cd/pull/1298) ([alexmt](https://github.com/alexmt))
- Update dashboard to have controller/repo-server stats. Collapsible rows [\#1295](https://github.com/argoproj/argo-cd/pull/1295) ([jessesuen](https://github.com/jessesuen))
- Issue \#1290 - Fix concurrent read/write error in state cache [\#1293](https://github.com/argoproj/argo-cd/pull/1293) ([alexmt](https://github.com/alexmt))
- Fix a goroutine leak in api-server application.PodLogs and application.Watch [\#1292](https://github.com/argoproj/argo-cd/pull/1292) ([jessesuen](https://github.com/jessesuen))
- Issue \#1287 - Fix local diff of non-namespaced resources. Also handleduplicates in local diff [\#1289](https://github.com/argoproj/argo-cd/pull/1289) ([alexmt](https://github.com/alexmt))
- only print to stdout, if there is a diff [\#1288](https://github.com/argoproj/argo-cd/pull/1288) ([marcb1](https://github.com/marcb1))
- Issue \#1258 - Disable CGO\_ENABLED for server/controller binaries [\#1286](https://github.com/argoproj/argo-cd/pull/1286) ([alexmt](https://github.com/alexmt))
- Fix documentation on diffing customization [\#1285](https://github.com/argoproj/argo-cd/pull/1285) ([yann-soubeyrand](https://github.com/yann-soubeyrand))
- Issue \#1070 - Handle duplicated resource definitions [\#1284](https://github.com/argoproj/argo-cd/pull/1284) ([alexmt](https://github.com/alexmt))
- Add golang prometheus metrics to controller and repo-server [\#1281](https://github.com/argoproj/argo-cd/pull/1281) ([jessesuen](https://github.com/jessesuen))
- Fix isssue where `argocd app set -p` required repo privileges. [\#1280](https://github.com/argoproj/argo-cd/pull/1280) ([jessesuen](https://github.com/jessesuen))
- MAGA: Make ArgoCD Golang Again! [\#1279](https://github.com/argoproj/argo-cd/pull/1279) ([jessesuen](https://github.com/jessesuen))
- Sample Grafana dashboard [\#1277](https://github.com/argoproj/argo-cd/pull/1277) ([hartman17](https://github.com/hartman17))
- Git cloning via SSH was not verifying host public key [\#1276](https://github.com/argoproj/argo-cd/pull/1276) ([jessesuen](https://github.com/jessesuen))
- Rename Application observedAt to reconciledAt and use observedAt to notify about partial app refresh [\#1270](https://github.com/argoproj/argo-cd/pull/1270) ([alexmt](https://github.com/alexmt))
- Bug fix: set 'Version' field while saving application resources tree [\#1268](https://github.com/argoproj/argo-cd/pull/1268) ([alexmt](https://github.com/alexmt))
- Avoid doing full reconciliation unless application 'managed' resource has changed [\#1267](https://github.com/argoproj/argo-cd/pull/1267) ([alexmt](https://github.com/alexmt))
- Support kustomize apps with remote bases in private repos in the same host [\#1264](https://github.com/argoproj/argo-cd/pull/1264) ([jessesuen](https://github.com/jessesuen))
- Add note about Kustomize1 [\#1263](https://github.com/argoproj/argo-cd/pull/1263) ([dthomson25](https://github.com/dthomson25))
- Enable debug logging for local development [\#1260](https://github.com/argoproj/argo-cd/pull/1260) ([alexec](https://github.com/alexec))
- Tweak lint [\#1259](https://github.com/argoproj/argo-cd/pull/1259) ([alexec](https://github.com/alexec))
- Fix link location [\#1257](https://github.com/argoproj/argo-cd/pull/1257) ([OmerKahani](https://github.com/OmerKahani))
- Add OpenAPI validation in CRD schema [\#1256](https://github.com/argoproj/argo-cd/pull/1256) ([jessesuen](https://github.com/jessesuen))
- Issue \#1252 - Application controller incorrectly build application objects tree [\#1253](https://github.com/argoproj/argo-cd/pull/1253) ([alexmt](https://github.com/alexmt))
- Issue \#1247 - Fix CRD creation/deletion handling [\#1249](https://github.com/argoproj/argo-cd/pull/1249) ([alexmt](https://github.com/alexmt))
- Replace git fetch implementation with git CLI \(from go-git\) [\#1244](https://github.com/argoproj/argo-cd/pull/1244) ([jessesuen](https://github.com/jessesuen))
- Fix nil pointer dereference in CompareAppState \(\#1234\) [\#1240](https://github.com/argoproj/argo-cd/pull/1240) ([alexmt](https://github.com/alexmt))
- Fix nil pointer dereference in CompareAppState [\#1234](https://github.com/argoproj/argo-cd/pull/1234) ([jessesuen](https://github.com/jessesuen))
- Issue \#1231 - Deprecated resource kinds from 'extensions' groups are not reconciled correctly [\#1232](https://github.com/argoproj/argo-cd/pull/1232) ([alexmt](https://github.com/alexmt))
- Issue \#1229 - App creation failed for public repository [\#1230](https://github.com/argoproj/argo-cd/pull/1230) ([alexmt](https://github.com/alexmt))
- Update link to config management plugins in custom\_tools.md [\#1228](https://github.com/argoproj/argo-cd/pull/1228) ([alexmt](https://github.com/alexmt))
- Update documentation for v0.12.0 [\#1227](https://github.com/argoproj/argo-cd/pull/1227) ([jessesuen](https://github.com/jessesuen))
- Migrates from gometalinter to golangci-lint. Closes \#1225 [\#1226](https://github.com/argoproj/argo-cd/pull/1226) ([alexec](https://github.com/alexec))
- Support talking to Dex using local cluster address instead of public address [\#1211](https://github.com/argoproj/argo-cd/pull/1211) ([jessesuen](https://github.com/jessesuen))

## [v0.12.2](https://github.com/argoproj/argo-cd/tree/v0.12.2) (2019-04-22)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.12.1...v0.12.2)

**Implemented enhancements:**

- Display number of errors on resource tab  [\#1477](https://github.com/argoproj/argo-cd/issues/1477)
- Display health of applications that are part of other apps [\#1433](https://github.com/argoproj/argo-cd/issues/1433)
- Whitelisting repos/clusters in projects should consider repo/cluster permissions [\#1432](https://github.com/argoproj/argo-cd/issues/1432)
- Resources with deletion timestamp set should be considered progressing [\#1410](https://github.com/argoproj/argo-cd/issues/1410)
- Core Concepts Documentation [\#1394](https://github.com/argoproj/argo-cd/issues/1394)
- Resources without health checks should not have health status [\#1383](https://github.com/argoproj/argo-cd/issues/1383)
- Perform health assessments on all resource nodes in the tree [\#1382](https://github.com/argoproj/argo-cd/issues/1382)
- `argocd app wait` should have `--resource` flag like sync [\#1206](https://github.com/argoproj/argo-cd/issues/1206)
- \[UI\] Service objects should show connection to pods which it is selecting [\#357](https://github.com/argoproj/argo-cd/issues/357)

**Fixed bugs:**

- Resource node details is crashing if live resource is missing [\#1505](https://github.com/argoproj/argo-cd/issues/1505)
- kustomization fields are all mandatory [\#1504](https://github.com/argoproj/argo-cd/issues/1504)
- Events tab title is not right if resources has no errors [\#1503](https://github.com/argoproj/argo-cd/issues/1503)
- Resources on sync panel and not sorted [\#1498](https://github.com/argoproj/argo-cd/issues/1498)
- how to add private gitlab repo? [\#1492](https://github.com/argoproj/argo-cd/issues/1492)
- Problem with MutatingAdmissionWebhookConfiguration object used by Istio [\#1487](https://github.com/argoproj/argo-cd/issues/1487)
- "bind: address already in use" after switching to gRPC-Web [\#1451](https://github.com/argoproj/argo-cd/issues/1451)
- 'argocd-repo-server' don't cleanup helm temp files [\#1446](https://github.com/argoproj/argo-cd/issues/1446)
- Auth expiry redirects don't respect basehref [\#1430](https://github.com/argoproj/argo-cd/issues/1430)
- UI is unable to load cluster level resource manifest [\#1429](https://github.com/argoproj/argo-cd/issues/1429)
- Argo CD deletes CRD if it is labeled with `app.kubernetes.io/instance` label [\#1425](https://github.com/argoproj/argo-cd/issues/1425)
- Annoying warning while using  `--grpc-web` flag [\#1420](https://github.com/argoproj/argo-cd/issues/1420)
- ComparisonError when using Umbrella charts [\#1412](https://github.com/argoproj/argo-cd/issues/1412)
- Application deletion might get stuck  [\#1406](https://github.com/argoproj/argo-cd/issues/1406)
- argocd cli segfault if ignoreDifferences is set [\#1389](https://github.com/argoproj/argo-cd/issues/1389)
- UI / API Errors Truncated, Time Out [\#1386](https://github.com/argoproj/argo-cd/issues/1386)
- API server goroutine leak [\#1381](https://github.com/argoproj/argo-cd/issues/1381)
- \[UI\] See details of applications fails with "r.nodes is undefined" [\#1371](https://github.com/argoproj/argo-cd/issues/1371)
- Rollback UI is not showing correct ksonnet parameters in preview [\#1326](https://github.com/argoproj/argo-cd/issues/1326)

**Closed issues:**

- Enable Probot [\#1418](https://github.com/argoproj/argo-cd/issues/1418)
- Unable to login against ingress: FATA\[0007\] rpc error: code = Internal desc = transport: received the unexpected content-type "text/plain; charset=utf-8 [\#1415](https://github.com/argoproj/argo-cd/issues/1415)
- Argo-cd & Argo-events [\#1401](https://github.com/argoproj/argo-cd/issues/1401)
- Out of sync resource deployment after syncing [\#1094](https://github.com/argoproj/argo-cd/issues/1094)
- Enable auth in e2e test [\#507](https://github.com/argoproj/argo-cd/issues/507)

## [v0.12.1](https://github.com/argoproj/argo-cd/tree/v0.12.1) (2019-04-09)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.12.0...v0.12.1)

**Implemented enhancements:**

- Create new documentation website [\#1390](https://github.com/argoproj/argo-cd/issues/1390)
- Surface pod status to tree view [\#1358](https://github.com/argoproj/argo-cd/issues/1358)
- \[UI\] default view should resource view instead of diff view [\#1354](https://github.com/argoproj/argo-cd/issues/1354)
- Surface Service/Ingress external IPs, hostname to application [\#908](https://github.com/argoproj/argo-cd/issues/908)

**Fixed bugs:**

- Patch APIs are not audited [\#1397](https://github.com/argoproj/argo-cd/issues/1397)
- Documentation feedback [\#1395](https://github.com/argoproj/argo-cd/issues/1395)
- agrocd-application-controller fails with `runtime: goroutine stack exceeds 1000000000-byte limit` [\#1374](https://github.com/argoproj/argo-cd/issues/1374)
- \[UI\] applications view blows up when user does not have `get projects` permissions [\#1368](https://github.com/argoproj/argo-cd/issues/1368)
- SIGSEGV in `argocd app sync` waitOnApplicationStatus [\#1366](https://github.com/argoproj/argo-cd/issues/1366)
- Dropdown menu should not have sync item for unmanaged resources [\#1357](https://github.com/argoproj/argo-cd/issues/1357)
- Cannot create app from local file [\#1327](https://github.com/argoproj/argo-cd/issues/1327)
- Local diff seems to ignore diffing customization [\#1294](https://github.com/argoproj/argo-cd/issues/1294)
- \[UI\] Improve input style [\#1217](https://github.com/argoproj/argo-cd/issues/1217)
- Unable to create app from private repo: x509: certificate signed by unknown authority  [\#1171](https://github.com/argoproj/argo-cd/issues/1171)
- kubectl v1.13 fails to convert extensions/NetworkPolicy [\#1012](https://github.com/argoproj/argo-cd/issues/1012)

**Closed issues:**

- Application Details page don't render cluster resources health/sync state [\#1404](https://github.com/argoproj/argo-cd/issues/1404)
- Remove deprecated componentParameterOverrides field [\#1372](https://github.com/argoproj/argo-cd/issues/1372)
- Diff wrong after sync, one line stays red [\#1362](https://github.com/argoproj/argo-cd/issues/1362)
- \[UI\] parameter preview is not considering revision [\#1356](https://github.com/argoproj/argo-cd/issues/1356)
- Not possible to add URI prefix to SSO link [\#1349](https://github.com/argoproj/argo-cd/issues/1349)
- Switching from RC to stable branch breaks application controller deployment [\#1346](https://github.com/argoproj/argo-cd/issues/1346)
- The `--group` flag is not working in `patch-resource` cli command [\#1319](https://github.com/argoproj/argo-cd/issues/1319)
- Does ClusterRole `argocd-application-controller` need such priviledges? [\#1311](https://github.com/argoproj/argo-cd/issues/1311)
- \[UI\] application table view needs to be sorted [\#1310](https://github.com/argoproj/argo-cd/issues/1310)
- NamespaceBlacklist typo in the docs [\#1302](https://github.com/argoproj/argo-cd/issues/1302)
- Controllers started failing syncs `net/http: request canceled \(Client.Timeout exceeded while awaiting headers\)` [\#1291](https://github.com/argoproj/argo-cd/issues/1291)
- Remove "sync" button from UI [\#1278](https://github.com/argoproj/argo-cd/issues/1278)
- Support overriding image name/tag in for Kustomize 2 apps [\#1275](https://github.com/argoproj/argo-cd/issues/1275)
- UI Enhancement Proposals Quick Wins [\#1274](https://github.com/argoproj/argo-cd/issues/1274)
- Custom Actions [\#1272](https://github.com/argoproj/argo-cd/issues/1272)
- PersistentVolumeClaims created by a CRD [\#1250](https://github.com/argoproj/argo-cd/issues/1250)
- Add protection to controller from malformed app/project specs [\#1245](https://github.com/argoproj/argo-cd/issues/1245)
- Improve e2e testing infrastructure [\#1135](https://github.com/argoproj/argo-cd/issues/1135)
- Get git-sha of application git repository [\#740](https://github.com/argoproj/argo-cd/issues/740)
- `wait` must throw an error if application deployment failed [\#733](https://github.com/argoproj/argo-cd/issues/733)
- SSO login with multiple argocd-server replicas might fail with "unknown app state" [\#709](https://github.com/argoproj/argo-cd/issues/709)
- Tests should be run with race-condition detector [\#505](https://github.com/argoproj/argo-cd/issues/505)
- Static assets are being browser cached between upgrades [\#489](https://github.com/argoproj/argo-cd/issues/489)
- Argo-cd perimeter/integration [\#92](https://github.com/argoproj/argo-cd/issues/92)

## [v0.12.0](https://github.com/argoproj/argo-cd/tree/v0.12.0) (2019-03-22)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.12.0-rc6...v0.12.0)

**Closed issues:**

- goroutine leak correlated to sync activity [\#1313](https://github.com/argoproj/argo-cd/issues/1313)
- Getting a crash when running argocd diff [\#1308](https://github.com/argoproj/argo-cd/issues/1308)
- Web login via dex works, cli login returns "Invalid client credentials. [\#1266](https://github.com/argoproj/argo-cd/issues/1266)
- Need to set HELM\_PLUGIN for helm plugin support [\#1153](https://github.com/argoproj/argo-cd/issues/1153)

## [v0.12.0-rc6](https://github.com/argoproj/argo-cd/tree/v0.12.0-rc6) (2019-03-20)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.12.0-rc5...v0.12.0-rc6)

**Closed issues:**

- HA: READONLY You can't write against a read only replica [\#1300](https://github.com/argoproj/argo-cd/issues/1300)
- Weird diff on Helm Charts [\#1297](https://github.com/argoproj/argo-cd/issues/1297)
- Serviceaccount can't list CRDs [\#1265](https://github.com/argoproj/argo-cd/issues/1265)
- Crashloop with v0.12-rc3 [\#1258](https://github.com/argoproj/argo-cd/issues/1258)

## [v0.12.0-rc5](https://github.com/argoproj/argo-cd/tree/v0.12.0-rc5) (2019-03-19)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.12.0-rc4...v0.12.0-rc5)

**Implemented enhancements:**

- Application warning when a manifest is defined twice [\#1070](https://github.com/argoproj/argo-cd/issues/1070)

**Closed issues:**

- controller concurrent map read and map write [\#1290](https://github.com/argoproj/argo-cd/issues/1290)
- Local diff seems to get confused by non namespaced resources [\#1287](https://github.com/argoproj/argo-cd/issues/1287)
- application view \(sometimes\) shows only blank canvas [\#1282](https://github.com/argoproj/argo-cd/issues/1282)
- Istio Based Canary [\#1273](https://github.com/argoproj/argo-cd/issues/1273)
- Network View [\#1271](https://github.com/argoproj/argo-cd/issues/1271)
- Webhook secrets [\#1142](https://github.com/argoproj/argo-cd/issues/1142)
- Failed to sync cluster [\#1006](https://github.com/argoproj/argo-cd/issues/1006)
- Question: Helm repository with multible values files  [\#938](https://github.com/argoproj/argo-cd/issues/938)
- provide helm chart for installing argo-cd [\#935](https://github.com/argoproj/argo-cd/issues/935)
- Application controller HA [\#894](https://github.com/argoproj/argo-cd/issues/894)
- Add style guide for contributors [\#579](https://github.com/argoproj/argo-cd/issues/579)

## [v0.12.0-rc4](https://github.com/argoproj/argo-cd/tree/v0.12.0-rc4) (2019-03-12)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.12.0-rc3...v0.12.0-rc4)

**Fixed bugs:**

- Application controller incorrectly build application objects tree [\#1252](https://github.com/argoproj/argo-cd/issues/1252)
- Excluded resources setting is ignored by controller [\#1248](https://github.com/argoproj/argo-cd/issues/1248)
- Application controller does not handle CRD creation/deletion well [\#1247](https://github.com/argoproj/argo-cd/issues/1247)
- CompareAppState\(\): runtime error: invalid memory address or nil pointer dereference [\#1233](https://github.com/argoproj/argo-cd/issues/1233)
- Application controller does not handle cluster credentials update [\#1220](https://github.com/argoproj/argo-cd/issues/1220)
- Allow using any name for secrets which store cluster credentials [\#1218](https://github.com/argoproj/argo-cd/issues/1218)

**Closed issues:**

- \[UI\] UI loads helm parameters without taking into account selected values files [\#1261](https://github.com/argoproj/argo-cd/issues/1261)
- Getting started guide - accessing admin passwd incorrect [\#1246](https://github.com/argoproj/argo-cd/issues/1246)
- go-git is unable to properly handle AWS CodeCommit [\#1243](https://github.com/argoproj/argo-cd/issues/1243)
- Daemonset fluentd-elasticsearch keep being deleted after upgrade to 0.11.2 [\#1242](https://github.com/argoproj/argo-cd/issues/1242)
- project field in 'create application' dialog is confusing [\#1236](https://github.com/argoproj/argo-cd/issues/1236)
- SSO using dex and Google as the provider, user unable to access/add resources [\#1235](https://github.com/argoproj/argo-cd/issues/1235)
- Migrate off gometalinter onto something else [\#1225](https://github.com/argoproj/argo-cd/issues/1225)
- Document basic Argo CD settings [\#1150](https://github.com/argoproj/argo-cd/issues/1150)
- \[UI\] Some componentParameterOverrides don't get rendered [\#1116](https://github.com/argoproj/argo-cd/issues/1116)
- Handle OutOfSync condition when K8s drops annotations & labels for CRDs [\#1089](https://github.com/argoproj/argo-cd/issues/1089)
- Statefullset cannot be updated  [\#1068](https://github.com/argoproj/argo-cd/issues/1068)
- Azure Git Repos Fail: unexpected client error: status code: 400 [\#1067](https://github.com/argoproj/argo-cd/issues/1067)

## [v0.12.0-rc3](https://github.com/argoproj/argo-cd/tree/v0.12.0-rc3) (2019-03-07)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.12.0-rc2...v0.12.0-rc3)

## [v0.12.0-rc2](https://github.com/argoproj/argo-cd/tree/v0.12.0-rc2) (2019-03-06)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.12.0-rc1...v0.12.0-rc2)

**Closed issues:**

- Deprecated resource kinds from 'extensions' groups are not reconciled correctly [\#1231](https://github.com/argoproj/argo-cd/issues/1231)
- App creation failed for public repository [\#1229](https://github.com/argoproj/argo-cd/issues/1229)

## [v0.12.0-rc1](https://github.com/argoproj/argo-cd/tree/v0.12.0-rc1) (2019-03-05)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.11.2...v0.12.0-rc1)

**Implemented enhancements:**

- Jsonnet Applications External Variables and/or TLA [\#1107](https://github.com/argoproj/argo-cd/issues/1107)
- Deal with 4KB cookie limit for JWT [\#1103](https://github.com/argoproj/argo-cd/issues/1103)
- Support Jsonnet external variables and top level arguments [\#1082](https://github.com/argoproj/argo-cd/issues/1082)

**Fixed bugs:**

- Deletion of an auto-synced app causes infighting [\#1208](https://github.com/argoproj/argo-cd/issues/1208)
- Partial sync with non-existant resources syncs entire app [\#1205](https://github.com/argoproj/argo-cd/issues/1205)
- Sync from CLI without get permissions panics [\#1204](https://github.com/argoproj/argo-cd/issues/1204)

**Closed issues:**

- gRPC interceptor logger is not honoring the global logrus loglevel [\#1222](https://github.com/argoproj/argo-cd/issues/1222)
- `argocd app create APPNAME --file` should fail when name conflicts [\#1201](https://github.com/argoproj/argo-cd/issues/1201)
- Set-up opinionated code imports and formatting  [\#1197](https://github.com/argoproj/argo-cd/issues/1197)
- Application history should include copy of application source.  [\#1196](https://github.com/argoproj/argo-cd/issues/1196)
- `argocd app diff --local` is broken if "application.instanceLabelKey" setting is not configured [\#1190](https://github.com/argoproj/argo-cd/issues/1190)
- omitted clusters field in excludedResources should match all clusters [\#1183](https://github.com/argoproj/argo-cd/issues/1183)
- kustomize images field is not recognized [\#1180](https://github.com/argoproj/argo-cd/issues/1180)
- UI should support raw YAML editor when creating/updating an app  [\#1176](https://github.com/argoproj/argo-cd/issues/1176)
- adding a private repo declaratively shows error SSH agent requested but SSH\_AUTH\_SOCK not-specified [\#1172](https://github.com/argoproj/argo-cd/issues/1172)
- Refactor manifests for Kustomize 2 [\#1164](https://github.com/argoproj/argo-cd/issues/1164)
- Add "argocd app patch" command [\#1162](https://github.com/argoproj/argo-cd/issues/1162)
- Improve cluster state cache collection [\#1161](https://github.com/argoproj/argo-cd/issues/1161)
- \[UI\] Deleting an application child resource from a parent application deletes the parent [\#1160](https://github.com/argoproj/argo-cd/issues/1160)
- Support wildcard globbing for RBAC groups [\#1154](https://github.com/argoproj/argo-cd/issues/1154)
- \[UI\] Cluster name in Application creation wizard [\#1152](https://github.com/argoproj/argo-cd/issues/1152)
- Pruning heredity [\#1149](https://github.com/argoproj/argo-cd/issues/1149)
- Deprecate ComponentParameterOverrides in favor of source specific config [\#1141](https://github.com/argoproj/argo-cd/issues/1141)
- PermissionDenied errors should be more specific [\#1126](https://github.com/argoproj/argo-cd/issues/1126)
- v0.12 UI enhancements feedback [\#1122](https://github.com/argoproj/argo-cd/issues/1122)
- \[UI\] Switch to text based YAML diff instead of json diff [\#1086](https://github.com/argoproj/argo-cd/issues/1086)
- Kustomize 2.0 support [\#1085](https://github.com/argoproj/argo-cd/issues/1085)
- Ability to selectively ignore differences to support fuzzy diff comparisons [\#1075](https://github.com/argoproj/argo-cd/issues/1075)
- SSO with non-public argocd url doesn't work [\#1062](https://github.com/argoproj/argo-cd/issues/1062)
- \[UI\] Enhance app creation page with auto-sync and auto-prune options [\#1058](https://github.com/argoproj/argo-cd/issues/1058)
- Support resource filtering in Application controller [\#1010](https://github.com/argoproj/argo-cd/issues/1010)
- \[UI\] YAML syntax highlighting in manifest view [\#942](https://github.com/argoproj/argo-cd/issues/942)
- Prometheus counters for app sync [\#936](https://github.com/argoproj/argo-cd/issues/936)
- Custom health assessments for CRDs [\#911](https://github.com/argoproj/argo-cd/issues/911)
- SSO should support readonly for all authorized github users [\#811](https://github.com/argoproj/argo-cd/issues/811)
- Configuration management plugins \(custom templating\) [\#701](https://github.com/argoproj/argo-cd/issues/701)

**Merged pull requests:**

- Move parameters listing from GenerateManifests to GetAppDetails [\#1221](https://github.com/argoproj/argo-cd/pull/1221) ([jessesuen](https://github.com/jessesuen))
- Issue 1065 - The '--grpc-web' flag is ignored by login command [\#1215](https://github.com/argoproj/argo-cd/pull/1215) ([alexmt](https://github.com/alexmt))
- Issue \#1065 - Correctly read chunked http response in GRPC proxy [\#1214](https://github.com/argoproj/argo-cd/pull/1214) ([alexmt](https://github.com/alexmt))
- Fix broken test for rollout health lua script [\#1213](https://github.com/argoproj/argo-cd/pull/1213) ([dthomson25](https://github.com/dthomson25))
- Fix CLI panic with sync after timeout when cli did not have get permissions [\#1209](https://github.com/argoproj/argo-cd/pull/1209) ([jessesuen](https://github.com/jessesuen))
- Deprecate componentParameterOverrides in favor of source specific config [\#1207](https://github.com/argoproj/argo-cd/pull/1207) ([jessesuen](https://github.com/jessesuen))
- Add Suspended status to Rollout health script [\#1203](https://github.com/argoproj/argo-cd/pull/1203) ([dthomson25](https://github.com/dthomson25))
- Add cli command to patch resources [\#1200](https://github.com/argoproj/argo-cd/pull/1200) ([dthomson25](https://github.com/dthomson25))
- Lints local imports. Closes \#1197 [\#1198](https://github.com/argoproj/argo-cd/pull/1198) ([alexec](https://github.com/alexec))
- Issue  \#1161 - no need to maintain map of existing CRDs since it is handled by resourceVersion usage [\#1194](https://github.com/argoproj/argo-cd/pull/1194) ([alexmt](https://github.com/alexmt))
- Issue \#1161 - Update resource version on every watch event [\#1192](https://github.com/argoproj/argo-cd/pull/1192) ([alexmt](https://github.com/alexmt))
- 'argocd app diff --local' is broken if application.instanceLabelKey setting is not configured [\#1191](https://github.com/argoproj/argo-cd/pull/1191) ([alexmt](https://github.com/alexmt))
- Rename excludedResources config key to resource.exclusions. Support hot reload [\#1189](https://github.com/argoproj/argo-cd/pull/1189) ([jessesuen](https://github.com/jessesuen))
- Include resource, action, object in permission denied errors [\#1188](https://github.com/argoproj/argo-cd/pull/1188) ([jessesuen](https://github.com/jessesuen))
- Add suspended status [\#1187](https://github.com/argoproj/argo-cd/pull/1187) ([dthomson25](https://github.com/dthomson25))
- Support patching resource using REST API [\#1186](https://github.com/argoproj/argo-cd/pull/1186) ([alexmt](https://github.com/alexmt))
- Makes the fields of excluded resources optional. Closes \#1183 [\#1185](https://github.com/argoproj/argo-cd/pull/1185) ([alexec](https://github.com/alexec))
- Introduce prometheus histogram for app reconcile performance [\#1184](https://github.com/argoproj/argo-cd/pull/1184) ([jessesuen](https://github.com/jessesuen))
- invalidate repo cache on delete [\#1182](https://github.com/argoproj/argo-cd/pull/1182) ([narg95](https://github.com/narg95))
- Fix typo [\#1179](https://github.com/argoproj/argo-cd/pull/1179) ([dulltz](https://github.com/dulltz))
- Switch to kustomize v2.0.2 \(from v2.0.1\) [\#1178](https://github.com/argoproj/argo-cd/pull/1178) ([jessesuen](https://github.com/jessesuen))
- Make test more tolerant [\#1177](https://github.com/argoproj/argo-cd/pull/1177) ([alexec](https://github.com/alexec))
- Issue \#701 - Add config management plugins documentation [\#1175](https://github.com/argoproj/argo-cd/pull/1175) ([alexmt](https://github.com/alexmt))
- Issue \#701 - Rename config management plugin command 'template' to 'generate' [\#1174](https://github.com/argoproj/argo-cd/pull/1174) ([alexmt](https://github.com/alexmt))
- Issue \#1161 - Use correct resource version in K8S watch API calls to avoid lost update events [\#1173](https://github.com/argoproj/argo-cd/pull/1173) ([alexmt](https://github.com/alexmt))
- Fix reconcile hotloop due to incorrect app source equality check [\#1170](https://github.com/argoproj/argo-cd/pull/1170) ([jessesuen](https://github.com/jessesuen))
- Switch to kustomize2 as default. Add argocd-ha install manifests [\#1169](https://github.com/argoproj/argo-cd/pull/1169) ([jessesuen](https://github.com/jessesuen))
- Use kubernetes recommended labels [\#1168](https://github.com/argoproj/argo-cd/pull/1168) ([lcostea](https://github.com/lcostea))
- Adds support for patching applications. Closes \#1162 [\#1166](https://github.com/argoproj/argo-cd/pull/1166) ([alexec](https://github.com/alexec))
- Adds support for Jsonnet External Variables and Top-Level Arguments. … [\#1165](https://github.com/argoproj/argo-cd/pull/1165) ([alexec](https://github.com/alexec))
- Let config management plugin inherit system env variables [\#1163](https://github.com/argoproj/argo-cd/pull/1163) ([alexmt](https://github.com/alexmt))
- Fix issue where `argocd app diff` reversed the left/right comparison [\#1158](https://github.com/argoproj/argo-cd/pull/1158) ([jessesuen](https://github.com/jessesuen))
- Add application sync counters as new prometheus metric. Add API-server metrics [\#1156](https://github.com/argoproj/argo-cd/pull/1156) ([jessesuen](https://github.com/jessesuen))
- Fix issue where YAML file did not split correctly when encoded in UTF-16 [\#1155](https://github.com/argoproj/argo-cd/pull/1155) ([jessesuen](https://github.com/jessesuen))
- Issue \#701 - Support for custom templaters [\#1151](https://github.com/argoproj/argo-cd/pull/1151) ([alexmt](https://github.com/alexmt))
- Allows you to exclude resources based on API group, kind, and cluster… [\#1147](https://github.com/argoproj/argo-cd/pull/1147) ([alexec](https://github.com/alexec))
- Display a warning if the JWT cookie is too large. Fixes \#1103 [\#1146](https://github.com/argoproj/argo-cd/pull/1146) ([alexec](https://github.com/alexec))
- Update CHANGELOG.md for v0.11.2 [\#1144](https://github.com/argoproj/argo-cd/pull/1144) ([jessesuen](https://github.com/jessesuen))
- Correct Redis port [\#1143](https://github.com/argoproj/argo-cd/pull/1143) ([leominov](https://github.com/leominov))
- Document custom health checks and diffing customization [\#1140](https://github.com/argoproj/argo-cd/pull/1140) ([alexmt](https://github.com/alexmt))
- Issue \#911 - Implement cert-manager CRD health checks [\#1139](https://github.com/argoproj/argo-cd/pull/1139) ([alexmt](https://github.com/alexmt))
- Adds support for Kustomize 2.0.1. Fixes \#1085 [\#1138](https://github.com/argoproj/argo-cd/pull/1138) ([alexec](https://github.com/alexec))
- Fix EncodeX509KeyPair function so it takes in account chained certificates [\#1137](https://github.com/argoproj/argo-cd/pull/1137) ([amarruedo](https://github.com/amarruedo))
- Disable authentication in dev setup [\#1136](https://github.com/argoproj/argo-cd/pull/1136) ([alexmt](https://github.com/alexmt))
- Add service manifest for redis [\#1134](https://github.com/argoproj/argo-cd/pull/1134) ([StuartHarris](https://github.com/StuartHarris))
- Issue \#1132 - Interactive application/project edit [\#1133](https://github.com/argoproj/argo-cd/pull/1133) ([alexmt](https://github.com/alexmt))
- Adds support for ARGOCD\_OPTS envvar for global variables. Fixes \#1081 [\#1131](https://github.com/argoproj/argo-cd/pull/1131) ([alexec](https://github.com/alexec))
- Issue \#1075 - Ability to selectively ignore differences to support fuzzy diff comparisons [\#1130](https://github.com/argoproj/argo-cd/pull/1130) ([alexmt](https://github.com/alexmt))
- add community blog to readme [\#1129](https://github.com/argoproj/argo-cd/pull/1129) ([edlee2121](https://github.com/edlee2121))
- Exclude metrics.k8s.io from watch [\#1128](https://github.com/argoproj/argo-cd/pull/1128) ([jessesuen](https://github.com/jessesuen))
- Issue \#1087 - Exclude hooks from local diff [\#1123](https://github.com/argoproj/argo-cd/pull/1123) ([alexmt](https://github.com/alexmt))
- Adds some instructions on how to run images locally. [\#1121](https://github.com/argoproj/argo-cd/pull/1121) ([alexec](https://github.com/alexec))
- Issue \#937 - Use redis as a shared throwaway cache [\#1120](https://github.com/argoproj/argo-cd/pull/1120) ([alexmt](https://github.com/alexmt))
- Adds client retry. Fixes \#959 [\#1119](https://github.com/argoproj/argo-cd/pull/1119) ([alexec](https://github.com/alexec))
- Prevent deletion hotloop. Improve reconciliation log for easier log queries [\#1115](https://github.com/argoproj/argo-cd/pull/1115) ([jessesuen](https://github.com/jessesuen))
- Fix issue where dex restart could cause login failures [\#1114](https://github.com/argoproj/argo-cd/pull/1114) ([jessesuen](https://github.com/jessesuen))
- Nil out application sources if source spec is equal to their zero value [\#1109](https://github.com/argoproj/argo-cd/pull/1109) ([jessesuen](https://github.com/jessesuen))
- Revert broken fix for azure repos which broke private repositories [\#1108](https://github.com/argoproj/argo-cd/pull/1108) ([jessesuen](https://github.com/jessesuen))
- Issue \#1076 - support wildcard globs for project sources & destinations [\#1106](https://github.com/argoproj/argo-cd/pull/1106) ([alexmt](https://github.com/alexmt))
- Fix rollback command help [\#1104](https://github.com/argoproj/argo-cd/pull/1104) ([leominov](https://github.com/leominov))
- fixed minor typo in docs [\#1102](https://github.com/argoproj/argo-cd/pull/1102) ([jecho](https://github.com/jecho))
- Update CONTRIBUTING.md [\#1098](https://github.com/argoproj/argo-cd/pull/1098) ([alexec](https://github.com/alexec))
- Mention `brew tap argoproj/tap` in getting started [\#1097](https://github.com/argoproj/argo-cd/pull/1097) ([alexmt](https://github.com/alexmt))
- Added a recurse option for directories. Fixes 1083 [\#1096](https://github.com/argoproj/argo-cd/pull/1096) ([alexec](https://github.com/alexec))
- Add security.md and how to build custom repo-server from Dockerfile [\#1078](https://github.com/argoproj/argo-cd/pull/1078) ([jessesuen](https://github.com/jessesuen))
- Issue \#1065 - Support using grpc-web in argocd cli [\#1077](https://github.com/argoproj/argo-cd/pull/1077) ([alexmt](https://github.com/alexmt))
- Refactor packr box usage into new assets library. Add faster DEV\_IMAGE build [\#1073](https://github.com/argoproj/argo-cd/pull/1073) ([jessesuen](https://github.com/jessesuen))
- Switch to CLI git fetch from go-git to support fetching Azure DevOps repos [\#1071](https://github.com/argoproj/argo-cd/pull/1071) ([jessesuen](https://github.com/jessesuen))
- Stop logging /cluster.ClusterService/Create [\#1069](https://github.com/argoproj/argo-cd/pull/1069) ([dthomson25](https://github.com/dthomson25))
- Add custom resource health through lua [\#1064](https://github.com/argoproj/argo-cd/pull/1064) ([dthomson25](https://github.com/dthomson25))
- Enable docker buildkit in ci builds [\#1060](https://github.com/argoproj/argo-cd/pull/1060) ([jessesuen](https://github.com/jessesuen))
- Relax ingress/service health check to accept non-empty ingress list [\#1053](https://github.com/argoproj/argo-cd/pull/1053) ([jessesuen](https://github.com/jessesuen))
- Fix test compile error [\#1052](https://github.com/argoproj/argo-cd/pull/1052) ([alexmt](https://github.com/alexmt))
- Handle case where manifests contain a null items list [\#1051](https://github.com/argoproj/argo-cd/pull/1051) ([jessesuen](https://github.com/jessesuen))
- Document v0.11.1 changes [\#1049](https://github.com/argoproj/argo-cd/pull/1049) ([alexmt](https://github.com/alexmt))
- Fix controller deadlock when checking for stale cache [\#1046](https://github.com/argoproj/argo-cd/pull/1046) ([jessesuen](https://github.com/jessesuen))
- Controller cache was susceptible to clock skew in managed cluster [\#1043](https://github.com/argoproj/argo-cd/pull/1043) ([jessesuen](https://github.com/jessesuen))
- Fix sync operation sorting [\#1042](https://github.com/argoproj/argo-cd/pull/1042) ([alexmt](https://github.com/alexmt))
- Fix ability to unset ApplicationSource specific parameters [\#1041](https://github.com/argoproj/argo-cd/pull/1041) ([jessesuen](https://github.com/jessesuen))
- Issue \#1039 - Correct redirect to login page if dex authentication is not successful [\#1040](https://github.com/argoproj/argo-cd/pull/1040) ([alexmt](https://github.com/alexmt))
- Hooks result should have Running phase by default \(given we don't have Pending state\) [\#1037](https://github.com/argoproj/argo-cd/pull/1037) ([alexmt](https://github.com/alexmt))
- Split manifests into components [\#1035](https://github.com/argoproj/argo-cd/pull/1035) ([mgoodness](https://github.com/mgoodness))
- Issue \#1033 - Fix force resource delete API [\#1034](https://github.com/argoproj/argo-cd/pull/1034) ([alexmt](https://github.com/alexmt))
- Fix PermissionDenied issue during app creation with project roles [\#1030](https://github.com/argoproj/argo-cd/pull/1030) ([jessesuen](https://github.com/jessesuen))
- Replace grpc repo-server parallelism limit interceptor with semaphore [\#1029](https://github.com/argoproj/argo-cd/pull/1029) ([alexmt](https://github.com/alexmt))
- Issue \#1025 - Fix /v1/applications/\<appName\>/manifests for app with helm dependencies [\#1026](https://github.com/argoproj/argo-cd/pull/1026) ([alexmt](https://github.com/alexmt))
- Correct "basehref " in the sample UI base path [\#1024](https://github.com/argoproj/argo-cd/pull/1024) ([infinitydon](https://github.com/infinitydon))
- Downgrade kubectl to v1.12 to regain `kubectl convert` functionality [\#1023](https://github.com/argoproj/argo-cd/pull/1023) ([jessesuen](https://github.com/jessesuen))
- Switch to a custom casbin adapter for rbac enforcment [\#1022](https://github.com/argoproj/argo-cd/pull/1022) ([jessesuen](https://github.com/jessesuen))
- Do not allow metadata.creationTimestamp to affect sync status [\#1021](https://github.com/argoproj/argo-cd/pull/1021) ([jessesuen](https://github.com/jessesuen))
- Issue \#937 - Allow using redis as a cache in repo-server [\#1020](https://github.com/argoproj/argo-cd/pull/1020) ([alexmt](https://github.com/alexmt))
- Graceful handling of clusters where API resource discovery is partially successful [\#1018](https://github.com/argoproj/argo-cd/pull/1018) ([jessesuen](https://github.com/jessesuen))
- Issue \#1013 - handle k8s resources circular dependency [\#1016](https://github.com/argoproj/argo-cd/pull/1016) ([alexmt](https://github.com/alexmt))
- Update README [\#1014](https://github.com/argoproj/argo-cd/pull/1014) ([edlee2121](https://github.com/edlee2121))
- Fix `app diff --local` command [\#1008](https://github.com/argoproj/argo-cd/pull/1008) ([jessesuen](https://github.com/jessesuen))
- Update parameters.md [\#1007](https://github.com/argoproj/argo-cd/pull/1007) ([saradhis](https://github.com/saradhis))
- Update CHANGELOG, docs to use stable tag, and tweak getting started guide [\#1005](https://github.com/argoproj/argo-cd/pull/1005) ([jessesuen](https://github.com/jessesuen))
- Moving apps between projects requires create/update in new project [\#1002](https://github.com/argoproj/argo-cd/pull/1002) ([jessesuen](https://github.com/jessesuen))
- Update docs to use v0.11.0-rc6 [\#1001](https://github.com/argoproj/argo-cd/pull/1001) ([jessesuen](https://github.com/jessesuen))
- Settings were getting re-initialized when incomplete. Session manager now uses settings manager [\#1000](https://github.com/argoproj/argo-cd/pull/1000) ([jessesuen](https://github.com/jessesuen))
- Log manifest with debug log level [\#999](https://github.com/argoproj/argo-cd/pull/999) ([alexmt](https://github.com/alexmt))
- Update docs to use v0.11.0-rc5 [\#994](https://github.com/argoproj/argo-cd/pull/994) ([jessesuen](https://github.com/jessesuen))
- Add better project policy rule validation [\#990](https://github.com/argoproj/argo-cd/pull/990) ([jessesuen](https://github.com/jessesuen))
- Use informers to load ArgoCD settings  [\#989](https://github.com/argoproj/argo-cd/pull/989) ([alexmt](https://github.com/alexmt))
- Eliminate reconcile hotloop by prevent Endpoint updates from requeuing apps [\#986](https://github.com/argoproj/argo-cd/pull/986) ([jessesuen](https://github.com/jessesuen))
- Increase QPS and Burst used in K8s client configs to 25/50 [\#984](https://github.com/argoproj/argo-cd/pull/984) ([jessesuen](https://github.com/jessesuen))
- Fix issue where custom resource objects might get synced to incorrect namespace during initial sync [\#982](https://github.com/argoproj/argo-cd/pull/982) ([jessesuen](https://github.com/jessesuen))
- Fix loading cluster connection status [\#980](https://github.com/argoproj/argo-cd/pull/980) ([alexmt](https://github.com/alexmt))
- Issue \#978 - Fix application rollback to deployment without overrides [\#979](https://github.com/argoproj/argo-cd/pull/979) ([alexmt](https://github.com/alexmt))
- Update golang to v1.11.4 [\#977](https://github.com/argoproj/argo-cd/pull/977) ([jessesuen](https://github.com/jessesuen))
- Improving documentation regarding params \(\#974\) [\#975](https://github.com/argoproj/argo-cd/pull/975) ([PaulVanStaden](https://github.com/PaulVanStaden))
- Update versions for kubectl \(v1.13.1\), helm \(v2.12.1\), ksonnet \(v0.13.1\) [\#973](https://github.com/argoproj/argo-cd/pull/973) ([jessesuen](https://github.com/jessesuen))
- Reduce timeout for checking cluster health [\#972](https://github.com/argoproj/argo-cd/pull/972) ([alexmt](https://github.com/alexmt))
- Update sample commands in project management doc [\#971](https://github.com/argoproj/argo-cd/pull/971) ([alexmt](https://github.com/alexmt))
- Add rollout progress and patch resource [\#965](https://github.com/argoproj/argo-cd/pull/965) ([dthomson25](https://github.com/dthomson25))
- Update docs to use v0.11.0-rc2 version [\#964](https://github.com/argoproj/argo-cd/pull/964) ([jessesuen](https://github.com/jessesuen))
- Use --refresh --hard-refresh flags in 'app get' 'app diff' commands [\#963](https://github.com/argoproj/argo-cd/pull/963) ([alexmt](https://github.com/alexmt))
- Issue \#916 - Use 'diff' to render actual vs target state difference [\#962](https://github.com/argoproj/argo-cd/pull/962) ([alexmt](https://github.com/alexmt))
- Show sync policy in `app list` view [\#961](https://github.com/argoproj/argo-cd/pull/961) ([jessesuen](https://github.com/jessesuen))
- Handle diff corner case where Role/ClusterRole rules are null [\#960](https://github.com/argoproj/argo-cd/pull/960) ([jessesuen](https://github.com/jessesuen))
- Load repo/cluster status in parallel to improve /repos /clusters API performance [\#958](https://github.com/argoproj/argo-cd/pull/958) ([alexmt](https://github.com/alexmt))
- Issue \#956 - Slow comparison if cluster is down [\#957](https://github.com/argoproj/argo-cd/pull/957) ([alexmt](https://github.com/alexmt))
- Issue \#950 - Application controller don't refresh app after destination update [\#951](https://github.com/argoproj/argo-cd/pull/951) ([alexmt](https://github.com/alexmt))
- Update aws-iam-authenticator to new version [\#948](https://github.com/argoproj/argo-cd/pull/948) ([lbrictson](https://github.com/lbrictson))
- Correctly drop cluster cache after CRD creation/deletion [\#947](https://github.com/argoproj/argo-cd/pull/947) ([alexmt](https://github.com/alexmt))
- Diff library handles case where live object has null secret data [\#945](https://github.com/argoproj/argo-cd/pull/945) ([jessesuen](https://github.com/jessesuen))
- Make injected application instance label configurable from default [\#944](https://github.com/argoproj/argo-cd/pull/944) ([jessesuen](https://github.com/jessesuen))
- Issue \#939 - Fix nil dereference error in Diff function [\#940](https://github.com/argoproj/argo-cd/pull/940) ([alexmt](https://github.com/alexmt))
- Issue 914 - Allow invalidating application related cache [\#931](https://github.com/argoproj/argo-cd/pull/931) ([alexmt](https://github.com/alexmt))
- Issue 906 - Support setting different base href in UI [\#930](https://github.com/argoproj/argo-cd/pull/930) ([alexmt](https://github.com/alexmt))
- Issue \#927 - Add missing handlings for deprecated extensions group kinds [\#928](https://github.com/argoproj/argo-cd/pull/928) ([alexmt](https://github.com/alexmt))
- Issue \#912 - Make ResourceNode 'tags' into a more generic 'info' struct [\#926](https://github.com/argoproj/argo-cd/pull/926) ([alexmt](https://github.com/alexmt))
- Issue \#922 - Fix nil derefrence error in 'argocd app diff' command [\#925](https://github.com/argoproj/argo-cd/pull/925) ([alexmt](https://github.com/alexmt))
- Issue \#910 - Reconstruct tree structure on the flight to avoid inconsistent state [\#921](https://github.com/argoproj/argo-cd/pull/921) ([alexmt](https://github.com/alexmt))
- Issue \#915 - Local 'argocd app diff' fails [\#920](https://github.com/argoproj/argo-cd/pull/920) ([alexmt](https://github.com/alexmt))
- Add v0.11.0-rc1 to getting\_started.md [\#919](https://github.com/argoproj/argo-cd/pull/919) ([alexmt](https://github.com/alexmt))
-  Prefix controller resource names with 'argocd-' [\#917](https://github.com/argoproj/argo-cd/pull/917) ([zcahana](https://github.com/zcahana))
- Fix issue preventing kustomize apps being multi-namespaced [\#913](https://github.com/argoproj/argo-cd/pull/913) ([jessesuen](https://github.com/jessesuen))

## [v0.11.2](https://github.com/argoproj/argo-cd/tree/v0.11.2) (2019-02-19)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.11.1...v0.11.2)

**Implemented enhancements:**

- Customize argocd cli using ARGOCD\_OPTS env variable [\#1081](https://github.com/argoproj/argo-cd/issues/1081)

**Fixed bugs:**

- ArgoCD server modifies custom tls.crt cert located in `argocd-secret` [\#1118](https://github.com/argoproj/argo-cd/issues/1118)
- Dex restart/upgrade causes login issues  [\#1113](https://github.com/argoproj/argo-cd/issues/1113)
- Resource hooks are included in CLI diff [\#1087](https://github.com/argoproj/argo-cd/issues/1087)

**Closed issues:**

- CLI should support application/project interactive edit [\#1132](https://github.com/argoproj/argo-cd/issues/1132)
- metrics.k8s.io should be an excluded group from watch [\#1127](https://github.com/argoproj/argo-cd/issues/1127)
- Reformat log output to make it easier to parse for reporting [\#1112](https://github.com/argoproj/argo-cd/issues/1112)
- \[UI\] Resources list view on Application details page [\#1101](https://github.com/argoproj/argo-cd/issues/1101)
- Update contributor guidelines  [\#1095](https://github.com/argoproj/argo-cd/issues/1095)
- Projects dropdown list stays empty [\#1093](https://github.com/argoproj/argo-cd/issues/1093)
- Resource poddisruptionbudget stays out of sync [\#1092](https://github.com/argoproj/argo-cd/issues/1092)
- When adding hostport to daemonset, argo diff does not see drift [\#1090](https://github.com/argoproj/argo-cd/issues/1090)
- Recurse app source directories [\#1083](https://github.com/argoproj/argo-cd/issues/1083)
- Question: Can ArgoCD function without access to a Git repo? [\#1080](https://github.com/argoproj/argo-cd/issues/1080)
- HPA objects can always be OutOfSync [\#1079](https://github.com/argoproj/argo-cd/issues/1079)
- Support wildcard globs for project sources & destinations [\#1076](https://github.com/argoproj/argo-cd/issues/1076)
- Consider using swagger client in Argo CD cli [\#1065](https://github.com/argoproj/argo-cd/issues/1065)
- Cannot add application via UI, errors in console log [\#1061](https://github.com/argoproj/argo-cd/issues/1061)
- helm ComparisonError [\#1057](https://github.com/argoproj/argo-cd/issues/1057)
- \[UI\] Rework application view [\#1055](https://github.com/argoproj/argo-cd/issues/1055)
- multiple application sources defined: Helm,Ksonnet [\#1054](https://github.com/argoproj/argo-cd/issues/1054)
- Force options does not work in /api/v1/applications/{name}/resource API [\#1033](https://github.com/argoproj/argo-cd/issues/1033)
- Status progressing when status.loadbalancer.ingress is non empty and does not have Hostname OR IP [\#997](https://github.com/argoproj/argo-cd/issues/997)
- Redis as a shared, distributed throwaway cache [\#937](https://github.com/argoproj/argo-cd/issues/937)
- \[UI\] It is unclear if application tree has a filter [\#929](https://github.com/argoproj/argo-cd/issues/929)
- Condensed top level application resource view [\#698](https://github.com/argoproj/argo-cd/issues/698)
- ArgoCD UI review [\#291](https://github.com/argoproj/argo-cd/issues/291)
- UI cleanup: improve empty state design [\#279](https://github.com/argoproj/argo-cd/issues/279)

## [v0.11.1](https://github.com/argoproj/argo-cd/tree/v0.11.1) (2019-01-18)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.11.0...v0.11.1)

**Closed issues:**

- Unable to deploy prometheus-operator v1.8 chart due to null ConfigMapList [\#1050](https://github.com/argoproj/argo-cd/issues/1050)
- Controller deadlock during CRD install + app refresh [\#1044](https://github.com/argoproj/argo-cd/issues/1044)
- Empty response page if user unauthorized to login [\#1039](https://github.com/argoproj/argo-cd/issues/1039)
- Namespaces are not being sorted during apply [\#1038](https://github.com/argoproj/argo-cd/issues/1038)
- Application details page rendering crash  [\#1036](https://github.com/argoproj/argo-cd/issues/1036)
- Edit button does not work for application via UI [\#1032](https://github.com/argoproj/argo-cd/issues/1032)
- Out of sync when using dependency chart postgresql [\#1031](https://github.com/argoproj/argo-cd/issues/1031)
- \[UI\] Resource details 'blink' when resource changes [\#1028](https://github.com/argoproj/argo-cd/issues/1028)
- UI should render page title to simplify navigation [\#1027](https://github.com/argoproj/argo-cd/issues/1027)
- API /v1/applications/\<appName\>/manifests issue [\#1025](https://github.com/argoproj/argo-cd/issues/1025)
- Incorrect PermissionDenied error during app creation when using project roles + user-defined RBAC [\#1019](https://github.com/argoproj/argo-cd/issues/1019)
- Controller should gracefully handle downed API aggregation services [\#1015](https://github.com/argoproj/argo-cd/issues/1015)
- Stack overflow in application-controller [\#1013](https://github.com/argoproj/argo-cd/issues/1013)
- Status out of sync sealed secret after syncing argocd [\#1009](https://github.com/argoproj/argo-cd/issues/1009)
- `argocd app diff --local` broken [\#1003](https://github.com/argoproj/argo-cd/issues/1003)
- leases.coordination.k8s.io is forbidden during sync [\#996](https://github.com/argoproj/argo-cd/issues/996)
- Sync errors with Kubernetes 1.13.1 [\#954](https://github.com/argoproj/argo-cd/issues/954)
- ArgoCD fails to compare app if ServiceCatalog CRD is installed in target cluster [\#650](https://github.com/argoproj/argo-cd/issues/650)

## [v0.11.0](https://github.com/argoproj/argo-cd/tree/v0.11.0) (2019-01-10)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.11.0-rc6...v0.11.0)

## [v0.11.0-rc6](https://github.com/argoproj/argo-cd/tree/v0.11.0-rc6) (2019-01-10)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.11.0-rc5...v0.11.0-rc6)

**Closed issues:**

- argocd-server Username or Password not working on fresh install [\#998](https://github.com/argoproj/argo-cd/issues/998)
- status sync OK but keeps synchronizing [\#992](https://github.com/argoproj/argo-cd/issues/992)
- Not able to sync helm chart, stays out of sync [\#991](https://github.com/argoproj/argo-cd/issues/991)

## [v0.11.0-rc5](https://github.com/argoproj/argo-cd/tree/v0.11.0-rc5) (2019-01-08)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.11.0-rc4...v0.11.0-rc5)

**Implemented enhancements:**

- Intelligent re-queuing of applications upon webhooks  [\#824](https://github.com/argoproj/argo-cd/issues/824)

**Closed issues:**

- console log is too large to show in front ui [\#988](https://github.com/argoproj/argo-cd/issues/988)
- cluster status failed [\#987](https://github.com/argoproj/argo-cd/issues/987)
- Leader election pattern causing reconciliation hot loop [\#985](https://github.com/argoproj/argo-cd/issues/985)
- can't login with message `unsatisfied client version constraint: \>= 0.11.0 ` [\#983](https://github.com/argoproj/argo-cd/issues/983)
- Initial sync of CRD + CRD instance installs resources into argocd namespace [\#981](https://github.com/argoproj/argo-cd/issues/981)
- Deploying pods from local code [\#846](https://github.com/argoproj/argo-cd/issues/846)

## [v0.11.0-rc4](https://github.com/argoproj/argo-cd/tree/v0.11.0-rc4) (2019-01-04)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.11.0-rc3...v0.11.0-rc4)

**Closed issues:**

- Rollback to deployment without parameter overrides works incorrectly [\#978](https://github.com/argoproj/argo-cd/issues/978)
- SIGSEGV: unexpected return pc for fmt.\(\*ss\).errorString [\#976](https://github.com/argoproj/argo-cd/issues/976)
- Unclear specification for parameters [\#974](https://github.com/argoproj/argo-cd/issues/974)

## [v0.11.0-rc3](https://github.com/argoproj/argo-cd/tree/v0.11.0-rc3) (2019-01-03)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.11.0-rc2...v0.11.0-rc3)

**Closed issues:**

- ClusterRole not permitted in project iloveu [\#970](https://github.com/argoproj/argo-cd/issues/970)
- Regression: no indication of application using overrides [\#969](https://github.com/argoproj/argo-cd/issues/969)
- problem: ingress health check always be processing [\#968](https://github.com/argoproj/argo-cd/issues/968)
- how to deal with params ? [\#967](https://github.com/argoproj/argo-cd/issues/967)
- UI error with helm charts parameters [\#966](https://github.com/argoproj/argo-cd/issues/966)
- argocd app get \<app-name\> --show-command does not print sync status [\#949](https://github.com/argoproj/argo-cd/issues/949)

## [v0.11.0-rc2](https://github.com/argoproj/argo-cd/tree/v0.11.0-rc2) (2018-12-28)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.11.0-rc1...v0.11.0-rc2)

**Fixed bugs:**

- \[UI\] JS error while syncing app [\#726](https://github.com/argoproj/argo-cd/issues/726)

**Closed issues:**

- Slow comparison if cluster is down [\#956](https://github.com/argoproj/argo-cd/issues/956)
- Unable to add a private bitbucket repository from the UI/CLI [\#955](https://github.com/argoproj/argo-cd/issues/955)
- Error listing all apps via argocd client [\#953](https://github.com/argoproj/argo-cd/issues/953)
- When deploying a Helm chart via UI the values files is not persisted [\#952](https://github.com/argoproj/argo-cd/issues/952)
- Application controller don't refresh app after destination update [\#950](https://github.com/argoproj/argo-cd/issues/950)
- timeout on external values.yaml [\#946](https://github.com/argoproj/argo-cd/issues/946)
- Handle equivalence of secrets with nils vs. empty string [\#943](https://github.com/argoproj/argo-cd/issues/943)
- Comparing resources which have fields with "invalid" values [\#941](https://github.com/argoproj/argo-cd/issues/941)
- sync error when deploying stable grafana chart with default values.yaml [\#939](https://github.com/argoproj/argo-cd/issues/939)
- Injected label `app.kubernetes.io/instance` should be configurable [\#934](https://github.com/argoproj/argo-cd/issues/934)
- Stop removing the ksonnet.io/component label [\#933](https://github.com/argoproj/argo-cd/issues/933)
- support stable helm charts with seperate values.yaml file [\#932](https://github.com/argoproj/argo-cd/issues/932)
- outofsync healthy status when there is no diff to find.  [\#927](https://github.com/argoproj/argo-cd/issues/927)
- Unauthorized error when refreshing state [\#924](https://github.com/argoproj/argo-cd/issues/924)
- Deploy application on selection of clusters [\#923](https://github.com/argoproj/argo-cd/issues/923)
- nil pointer dereference in `argocd app diff --local` [\#922](https://github.com/argoproj/argo-cd/issues/922)
- Application detail view not correct when deploying application with multiple resources [\#918](https://github.com/argoproj/argo-cd/issues/918)
- \[0.11.0\] `argocd app diff` shows invalid differences [\#916](https://github.com/argoproj/argo-cd/issues/916)
- \[0.11.0\] Local `argocd app diff` fails with "differing number of resources vs. live" error [\#915](https://github.com/argoproj/argo-cd/issues/915)
- Need a way to "force" refresh an app \(invalidate its cache\) [\#914](https://github.com/argoproj/argo-cd/issues/914)
- Make ResourceNode `tags` into a more generic `info` struct [\#912](https://github.com/argoproj/argo-cd/issues/912)
- Recover from inconsistent resource tree in live state cache  [\#910](https://github.com/argoproj/argo-cd/issues/910)
- \[UI\] filter by sync and health status [\#909](https://github.com/argoproj/argo-cd/issues/909)
- Support setting different base href in UI [\#906](https://github.com/argoproj/argo-cd/issues/906)
- GitLab Deploy Token support [\#901](https://github.com/argoproj/argo-cd/issues/901)
- argo for openshift [\#898](https://github.com/argoproj/argo-cd/issues/898)
- argocd delete CLI:  code = PermissionDenied desc = permission denied even though it works [\#830](https://github.com/argoproj/argo-cd/issues/830)
- Improve `argocd wait` output [\#759](https://github.com/argoproj/argo-cd/issues/759)

## [v0.11.0-rc1](https://github.com/argoproj/argo-cd/tree/v0.11.0-rc1) (2018-12-10)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.10.6...v0.11.0-rc1)

**Implemented enhancements:**

- Resolve ambiguous revision to commitSHA at sync time [\#818](https://github.com/argoproj/argo-cd/issues/818)
- Improve Application state reconciliation performance [\#806](https://github.com/argoproj/argo-cd/issues/806)
- Out of sync reported if Secrets with stringData are used [\#763](https://github.com/argoproj/argo-cd/issues/763)

**Fixed bugs:**

- Ingress resources always being marked as “Missing” [\#868](https://github.com/argoproj/argo-cd/issues/868)
- \[UI\] Modifying project sources erases previous ones [\#823](https://github.com/argoproj/argo-cd/issues/823)
- \[UI\] No error indication when insufficient permissions to create tokens [\#822](https://github.com/argoproj/argo-cd/issues/822)
- `kubectl auth reconcile` auto-creates namespaces [\#803](https://github.com/argoproj/argo-cd/issues/803)
- Hooks do not cascade delete [\#797](https://github.com/argoproj/argo-cd/issues/797)
- Hooks are not checked for permission in a project [\#794](https://github.com/argoproj/argo-cd/issues/794)
- Panic in application controller [\#790](https://github.com/argoproj/argo-cd/issues/790)
- Application type is incorrectly inferred as 'directory' if app source path starts with '.' [\#782](https://github.com/argoproj/argo-cd/issues/782)
- `argocd proj list` is not readable [\#769](https://github.com/argoproj/argo-cd/issues/769)
- `argocd sync` failure [\#760](https://github.com/argoproj/argo-cd/issues/760)
- API Server fails to return apps due to grpc max message size limit [\#690](https://github.com/argoproj/argo-cd/issues/690)
- API client watch helper to retry disconnections from API server [\#896](https://github.com/argoproj/argo-cd/pull/896) ([jessesuen](https://github.com/jessesuen))

**Closed issues:**

- `argocd wait` is not tolerant to disconnects [\#895](https://github.com/argoproj/argo-cd/issues/895)
- error creating SSH agent: "SSH agent requested but SSH\_AUTH\_SOCK not-specified" [\#884](https://github.com/argoproj/argo-cd/issues/884)
- Allow tracking of individual files/subdirectories [\#880](https://github.com/argoproj/argo-cd/issues/880)
- Support notifications like Argo [\#879](https://github.com/argoproj/argo-cd/issues/879)
- Support `\[skip ci\]` / `\[ci skip\]` in commit messages [\#878](https://github.com/argoproj/argo-cd/issues/878)
- OIDC config needs to be able to reference $secret.keys [\#877](https://github.com/argoproj/argo-cd/issues/877)
- Helm repositories config missing username/password [\#874](https://github.com/argoproj/argo-cd/issues/874)
- CLI support for multi-namespaced apps [\#873](https://github.com/argoproj/argo-cd/issues/873)
- Simplify, clarify, and consolidate resource datastructures [\#862](https://github.com/argoproj/argo-cd/issues/862)
- \[UI\] Using wrong endpoint when getting resource events [\#858](https://github.com/argoproj/argo-cd/issues/858)
- Hooks are being included as managed resources [\#856](https://github.com/argoproj/argo-cd/issues/856)
- pod logs does not work in multi namespaced apps [\#853](https://github.com/argoproj/argo-cd/issues/853)
- Rename 'controlled resources' to 'managed resources' [\#847](https://github.com/argoproj/argo-cd/issues/847)
- argocd-util needs static compilation [\#844](https://github.com/argoproj/argo-cd/issues/844)
- Need validation on project role OIDC group bindings [\#843](https://github.com/argoproj/argo-cd/issues/843)
- Remove git URL normalization in favor of fuzzy comparisons [\#838](https://github.com/argoproj/argo-cd/issues/838)
- Duplicate legacy helm/ksonnet fields into new fields [\#835](https://github.com/argoproj/argo-cd/issues/835)
- Set admin password on install [\#829](https://github.com/argoproj/argo-cd/issues/829)
- Self signed certificates prevents access with chrome using kubectl port-forward; identity and access control for argo cd [\#828](https://github.com/argoproj/argo-cd/issues/828)
- ks param doesnt like - \(dash\) in component names [\#825](https://github.com/argoproj/argo-cd/issues/825)
- \[UI\] Login button when external OIDC provider is configured [\#821](https://github.com/argoproj/argo-cd/issues/821)
- Purge application-controller cache when app deleted [\#802](https://github.com/argoproj/argo-cd/issues/802)
- Problems installing the istio chart [\#788](https://github.com/argoproj/argo-cd/issues/788)
- Failed to deploy helm chart into different namespaces in v0.10.6, kube v1.12.2 [\#787](https://github.com/argoproj/argo-cd/issues/787)
- Failed to deploy helm chart with local dependencies and no internet access [\#786](https://github.com/argoproj/argo-cd/issues/786)
- Ignore empty string differences in metadata.namespace [\#784](https://github.com/argoproj/argo-cd/issues/784)
- Support editing white/blacklisted resources and groups in UI [\#783](https://github.com/argoproj/argo-cd/issues/783)
- Rework application source with more structure [\#776](https://github.com/argoproj/argo-cd/issues/776)
- Failed to deploy app in v0.10.5 with v1.12.1 kubernetes [\#774](https://github.com/argoproj/argo-cd/issues/774)
- ArgoCD permissions [\#773](https://github.com/argoproj/argo-cd/issues/773)
- Guidance on custom binaries, Helm plugins, kustomize generators [\#772](https://github.com/argoproj/argo-cd/issues/772)
- \[UI\] Helm value files on App details page [\#770](https://github.com/argoproj/argo-cd/issues/770)
- Support diffing a local Helm app to the live application state [\#764](https://github.com/argoproj/argo-cd/issues/764)
- Publish hyperkube-style image [\#762](https://github.com/argoproj/argo-cd/issues/762)
- Helm: argocd app set appname -p variable=value -\> fails due to ksonnet validation [\#756](https://github.com/argoproj/argo-cd/issues/756)
- UI should render pending deletion application state [\#752](https://github.com/argoproj/argo-cd/issues/752)
- Missing health check for bare Pods and Jobs [\#751](https://github.com/argoproj/argo-cd/issues/751)
- Declaratively add helm repositories [\#747](https://github.com/argoproj/argo-cd/issues/747)
- Self service group access to project applications [\#742](https://github.com/argoproj/argo-cd/issues/742)
- \[UI\] Trim repo URL in app creation wizard [\#741](https://github.com/argoproj/argo-cd/issues/741)
- \[UI\] Cmd+Click should open app in new tab [\#732](https://github.com/argoproj/argo-cd/issues/732)
- Diff view shows incorrect base/value comparison [\#725](https://github.com/argoproj/argo-cd/issues/725)
- API server & UI should serve argocd binaries instead of linking to GitHub [\#716](https://github.com/argoproj/argo-cd/issues/716)
- Rename resources using Helm Templates does NOT delete previous objects [\#713](https://github.com/argoproj/argo-cd/issues/713)
- panic: send on closed channel in kube.copyEventsChannel [\#699](https://github.com/argoproj/argo-cd/issues/699)
- Ability for a single application to deploy into multiple namespaces [\#696](https://github.com/argoproj/argo-cd/issues/696)
- App delete should not remove finalizer until resources are deleted [\#636](https://github.com/argoproj/argo-cd/issues/636)
- ArgoUI should render three-way diffs \(like CLI\) [\#633](https://github.com/argoproj/argo-cd/issues/633)
- Error updating application \<app\>: etcdserver: request is too large [\#621](https://github.com/argoproj/argo-cd/issues/621)
- Make use of K8s recommended application labels [\#534](https://github.com/argoproj/argo-cd/issues/534)
- API for RBAC resources/verbs and move audit actions and RBAC actions to constants [\#453](https://github.com/argoproj/argo-cd/issues/453)
- Force delete option for deleting pods [\#417](https://github.com/argoproj/argo-cd/issues/417)
- Better error message for private repos [\#221](https://github.com/argoproj/argo-cd/issues/221)

**Merged pull requests:**

- Enforces looses user claims if default role is set [\#907](https://github.com/argoproj/argo-cd/pull/907) ([alexmt](https://github.com/alexmt))
- Server should accept clients with pre-release version [\#905](https://github.com/argoproj/argo-cd/pull/905) ([alexmt](https://github.com/alexmt))
- Fix discovering cluster wide resources with namespace [\#904](https://github.com/argoproj/argo-cd/pull/904) ([alexmt](https://github.com/alexmt))
- Give 'get' access to the argocd-server cluster role [\#903](https://github.com/argoproj/argo-cd/pull/903) ([jessesuen](https://github.com/jessesuen))
- Issue \#897 - Secret data not redacted in last-applied-configuration [\#902](https://github.com/argoproj/argo-cd/pull/902) ([alexmt](https://github.com/alexmt))
- Remove gracePeriod seconds option from API [\#900](https://github.com/argoproj/argo-cd/pull/900) ([dthomson25](https://github.com/dthomson25))
- Issue \#770 - Support loading app details by directory [\#893](https://github.com/argoproj/argo-cd/pull/893) ([alexmt](https://github.com/alexmt))
- Add initContainer volumeMount to custom tooling docs [\#892](https://github.com/argoproj/argo-cd/pull/892) ([mgoodness](https://github.com/mgoodness))
- Add force delete option to API [\#891](https://github.com/argoproj/argo-cd/pull/891) ([dthomson25](https://github.com/dthomson25))
- Issue \#760 - Properly read watch events to avoid nil pointer errors [\#890](https://github.com/argoproj/argo-cd/pull/890) ([alexmt](https://github.com/alexmt))
- Update to kustomize 1.0.11 [\#889](https://github.com/argoproj/argo-cd/pull/889) ([twz123](https://github.com/twz123))
- Add protection from malformed project policies being sent to casbin [\#888](https://github.com/argoproj/argo-cd/pull/888) ([jessesuen](https://github.com/jessesuen))
- Enable --auto-prune for app create if --sync-policy is automated [\#887](https://github.com/argoproj/argo-cd/pull/887) ([essh](https://github.com/essh))
- CLI support for multi-namespaced applications [\#886](https://github.com/argoproj/argo-cd/pull/886) ([jessesuen](https://github.com/jessesuen))
- Issue \#887 - OIDC config needs to be able to reference .keys [\#885](https://github.com/argoproj/argo-cd/pull/885) ([alexmt](https://github.com/alexmt))
- Update release notes for v0.11 and add more documentation [\#883](https://github.com/argoproj/argo-cd/pull/883) ([jessesuen](https://github.com/jessesuen))
- Issue \#874 - Helm repositories config missing username/password [\#882](https://github.com/argoproj/argo-cd/pull/882) ([alexmt](https://github.com/alexmt))
- build cli with packr [\#875](https://github.com/argoproj/argo-cd/pull/875) ([essh](https://github.com/essh))
- Remove parameters field from ApplicationStatus [\#872](https://github.com/argoproj/argo-cd/pull/872) ([alexmt](https://github.com/alexmt))
- Application controller does not save application parameters in app crd [\#871](https://github.com/argoproj/argo-cd/pull/871) ([alexmt](https://github.com/alexmt))
- Fix flaky e2e test [\#870](https://github.com/argoproj/argo-cd/pull/870) ([alexmt](https://github.com/alexmt))
- Issue \#868 - Filter out extensions group resources which are mirrored in apps group [\#869](https://github.com/argoproj/argo-cd/pull/869) ([alexmt](https://github.com/alexmt))
- Refactor, consolidate and rename resource type datastructures [\#867](https://github.com/argoproj/argo-cd/pull/867) ([jessesuen](https://github.com/jessesuen))
- build application-controller with packr [\#866](https://github.com/argoproj/argo-cd/pull/866) ([essh](https://github.com/essh))
- Issue \#858 - Support loading resource events for multi-network apps [\#865](https://github.com/argoproj/argo-cd/pull/865) ([alexmt](https://github.com/alexmt))
- Issue \#747 - Declaratively add helm repositories [\#864](https://github.com/argoproj/argo-cd/pull/864) ([alexmt](https://github.com/alexmt))
- Add local diff back [\#863](https://github.com/argoproj/argo-cd/pull/863) ([dthomson25](https://github.com/dthomson25))
- Use standard Scheme Convert function instead of the kubectl based converter [\#860](https://github.com/argoproj/argo-cd/pull/860) ([jessesuen](https://github.com/jessesuen))
- Proper treatment of resource lifecycle hooks [\#859](https://github.com/argoproj/argo-cd/pull/859) ([jessesuen](https://github.com/jessesuen))
- Switch to k8s recommended app.kubernetes.io/instance label [\#857](https://github.com/argoproj/argo-cd/pull/857) ([jessesuen](https://github.com/jessesuen))
- Issue \#853 - pod logs does not work in multi namespaced apps [\#855](https://github.com/argoproj/argo-cd/pull/855) ([alexmt](https://github.com/alexmt))
- Fix app diff command [\#854](https://github.com/argoproj/argo-cd/pull/854) ([alexmt](https://github.com/alexmt))
- Resources events streaming bug fixes: panic \(\#699\), stale cache detection, restaring bad watchers [\#852](https://github.com/argoproj/argo-cd/pull/852) ([alexmt](https://github.com/alexmt))
- Only run `helm dependency build` when necessary \(issue \#786\) [\#851](https://github.com/argoproj/argo-cd/pull/851) ([jessesuen](https://github.com/jessesuen))
- Rename 'controlled resources' to 'managed resources' [\#850](https://github.com/argoproj/argo-cd/pull/850) ([jessesuen](https://github.com/jessesuen))
- Remove git URL normalization in favor of fuzzy equivalence \(issue \#838\) [\#849](https://github.com/argoproj/argo-cd/pull/849) ([jessesuen](https://github.com/jessesuen))
- Normalize app spec during controller reconciliation and API server create/update [\#848](https://github.com/argoproj/argo-cd/pull/848) ([jessesuen](https://github.com/jessesuen))
- Build argocd-util as a statically linked binary [\#845](https://github.com/argoproj/argo-cd/pull/845) ([zcahana](https://github.com/zcahana))
- Issue \#696 - Support apps with static namespaces in resources [\#842](https://github.com/argoproj/argo-cd/pull/842) ([alexmt](https://github.com/alexmt))
- gRPC API client and gateway now supply user-agent. Require client min version as v0.11 [\#841](https://github.com/argoproj/argo-cd/pull/841) ([jessesuen](https://github.com/jessesuen))
- Refactor application controller [\#840](https://github.com/argoproj/argo-cd/pull/840) ([alexmt](https://github.com/alexmt))
- Serve CLI binaries directly from API server [\#837](https://github.com/argoproj/argo-cd/pull/837) ([jessesuen](https://github.com/jessesuen))
- Resolve ambiguous revisions in API server when initiating syncs \(issue \#818\) [\#834](https://github.com/argoproj/argo-cd/pull/834) ([jessesuen](https://github.com/jessesuen))
- Update kustomize base when setting image tags [\#833](https://github.com/argoproj/argo-cd/pull/833) ([jessesuen](https://github.com/jessesuen))
- Relax validation to permit app with no manifests [\#832](https://github.com/argoproj/argo-cd/pull/832) ([zcahana](https://github.com/zcahana))
- Explicitly check for namespace before running auth reconcile [\#826](https://github.com/argoproj/argo-cd/pull/826) ([jessesuen](https://github.com/jessesuen))
- Support the ability to map OIDC groups to project roles \(issue \#742\) [\#817](https://github.com/argoproj/argo-cd/pull/817) ([jessesuen](https://github.com/jessesuen))
- Add special case for secrets to base64 encode stringData before performing diff [\#814](https://github.com/argoproj/argo-cd/pull/814) ([jessesuen](https://github.com/jessesuen))
- Add declarative argocd setup docs [\#813](https://github.com/argoproj/argo-cd/pull/813) ([alexmt](https://github.com/alexmt))
- Fix repository settings deserialization [\#812](https://github.com/argoproj/argo-cd/pull/812) ([alexmt](https://github.com/alexmt))
- Ignore metadata.namespace in config when performing two-way diff \(issue \#784\) [\#810](https://github.com/argoproj/argo-cd/pull/810) ([jessesuen](https://github.com/jessesuen))
- Diff view shows incorrect base/value comparison \(issue \#725\) [\#809](https://github.com/argoproj/argo-cd/pull/809) ([jessesuen](https://github.com/jessesuen))
- Reorder `auth reconcile` after `apply` to prevent namespace creation [\#808](https://github.com/argoproj/argo-cd/pull/808) ([jessesuen](https://github.com/jessesuen))
- Defer deletion of app object until all resources have been deleted [\#807](https://github.com/argoproj/argo-cd/pull/807) ([jessesuen](https://github.com/jessesuen))
- Support for Pods as a sync hook [\#801](https://github.com/argoproj/argo-cd/pull/801) ([jessesuen](https://github.com/jessesuen))
- Issue \#621 - Fix child resource deletion [\#800](https://github.com/argoproj/argo-cd/pull/800) ([alexmt](https://github.com/alexmt))
- Restructure application sources to separate types [\#799](https://github.com/argoproj/argo-cd/pull/799) ([jessesuen](https://github.com/jessesuen))
- Use default server addresses. Use an imagePullPolicy of Always for manifests [\#796](https://github.com/argoproj/argo-cd/pull/796) ([jessesuen](https://github.com/jessesuen))
- Fix `make all` target and use archiveLogs workflow feature [\#795](https://github.com/argoproj/argo-cd/pull/795) ([jessesuen](https://github.com/jessesuen))
- Fix typo in documentation for hook delete policy [\#793](https://github.com/argoproj/argo-cd/pull/793) ([wmedlar](https://github.com/wmedlar))
- Issue \#355 - Treat 'crd-install' hooks as normal k8s resource [\#792](https://github.com/argoproj/argo-cd/pull/792) ([alexmt](https://github.com/alexmt))
- Issue \#790 - Fix application controller panic [\#791](https://github.com/argoproj/argo-cd/pull/791) ([alexmt](https://github.com/alexmt))
- Issue \#782 - Application type is incorrectly inferred as 'directory' if app source path starts with '.' [\#789](https://github.com/argoproj/argo-cd/pull/789) ([alexmt](https://github.com/alexmt))
- Move to single master image for all argocd services \(issue \#762\) [\#780](https://github.com/argoproj/argo-cd/pull/780) ([jessesuen](https://github.com/jessesuen))
- Update docs to describe how to customize repo-server \(issue \#772\) [\#778](https://github.com/argoproj/argo-cd/pull/778) ([jessesuen](https://github.com/jessesuen))
- Update CHANGELOG and docs to point to v0.10.6 [\#777](https://github.com/argoproj/argo-cd/pull/777) ([jessesuen](https://github.com/jessesuen))
- Fix issue preventing in-cluster app sync due to go-client changes \(issue \#774\) [\#775](https://github.com/argoproj/argo-cd/pull/775) ([jessesuen](https://github.com/jessesuen))
- Update CHANGELOG and docs to point to v0.10.5 install manifests [\#771](https://github.com/argoproj/argo-cd/pull/771) ([jessesuen](https://github.com/jessesuen))
- add project label to all metrics [\#767](https://github.com/argoproj/argo-cd/pull/767) ([conorfennell](https://github.com/conorfennell))
- add argo cluster permission to view logs [\#766](https://github.com/argoproj/argo-cd/pull/766) ([conorfennell](https://github.com/conorfennell))
- add metrics label for service monitor discovery [\#765](https://github.com/argoproj/argo-cd/pull/765) ([conorfennell](https://github.com/conorfennell))
- Update getting\_started to latest argo-cd version [\#761](https://github.com/argoproj/argo-cd/pull/761) ([solidnerd](https://github.com/solidnerd))
- Issue \#621 - Remove resources state from application CRD [\#758](https://github.com/argoproj/argo-cd/pull/758) ([alexmt](https://github.com/alexmt))
- Avoid overwriting an existing configmap with empty data if the argocd secret is not found [\#755](https://github.com/argoproj/argo-cd/pull/755) ([wmedlar](https://github.com/wmedlar))
- Health check is not discerning apiVersion when assessing CRDs \(issue \#753\) [\#754](https://github.com/argoproj/argo-cd/pull/754) ([jessesuen](https://github.com/jessesuen))
- Updated helm to latest version \(2.11.0\) [\#749](https://github.com/argoproj/argo-cd/pull/749) ([amarrella](https://github.com/amarrella))
- Issue \#536 - Declarative setup and configuration of ArgoCD [\#745](https://github.com/argoproj/argo-cd/pull/745) ([alexmt](https://github.com/alexmt))
- Update version to v0.10.3 in getting started guide. [\#739](https://github.com/argoproj/argo-cd/pull/739) ([tsuna](https://github.com/tsuna))
- Issue \#697 - Ability to perform field selection in API [\#736](https://github.com/argoproj/argo-cd/pull/736) ([alexmt](https://github.com/alexmt))
- Support adding name prefix in helm and kustomize [\#735](https://github.com/argoproj/argo-cd/pull/735) ([dthomson25](https://github.com/dthomson25))
- Fix ksonnet validation, take 2 [\#734](https://github.com/argoproj/argo-cd/pull/734) ([merenbach](https://github.com/merenbach))
- Fix applying TLS version settings [\#731](https://github.com/argoproj/argo-cd/pull/731) ([alexmt](https://github.com/alexmt))
- Revert "Validate Ksonnet apps through component dir presence \(\#708\)" [\#730](https://github.com/argoproj/argo-cd/pull/730) ([merenbach](https://github.com/merenbach))
- Update k8s dependencies. Fixes/refactoring to kube libraries [\#729](https://github.com/argoproj/argo-cd/pull/729) ([jessesuen](https://github.com/jessesuen))
- Update to kustomize 1.0.10 [\#728](https://github.com/argoproj/argo-cd/pull/728) ([twz123](https://github.com/twz123))
- Support for external OIDC providers and implicit login flows [\#727](https://github.com/argoproj/argo-cd/pull/727) ([jessesuen](https://github.com/jessesuen))
- Fix app refresh err when k8s patch is too slow [\#724](https://github.com/argoproj/argo-cd/pull/724) ([dthomson25](https://github.com/dthomson25))
- Fix nil pointer dereference in util/health [\#723](https://github.com/argoproj/argo-cd/pull/723) ([mduarte](https://github.com/mduarte))
- Update to kustomize 1.0.9 [\#722](https://github.com/argoproj/argo-cd/pull/722) ([twz123](https://github.com/twz123))
- Issue \#670 - Allow using Sets the value of different fields in kustomization file. [\#720](https://github.com/argoproj/argo-cd/pull/720) ([alexmt](https://github.com/alexmt))
- Changelog for v0.10.1 release [\#719](https://github.com/argoproj/argo-cd/pull/719) ([alexmt](https://github.com/alexmt))
- Issue \#657 - Use codecov to collect test coverage [\#717](https://github.com/argoproj/argo-cd/pull/717) ([alexmt](https://github.com/alexmt))
- Handle case where OIDC settings become invalid after dex server restart \(issue \#710\) [\#715](https://github.com/argoproj/argo-cd/pull/715) ([jessesuen](https://github.com/jessesuen))
- Update getting\_started to use v0.10.0 [\#714](https://github.com/argoproj/argo-cd/pull/714) ([jessesuen](https://github.com/jessesuen))
- git clean also needs to clean files under gitignore \(issue \#711\) [\#712](https://github.com/argoproj/argo-cd/pull/712) ([jessesuen](https://github.com/jessesuen))
- Validate Ksonnet apps through component dir presence [\#708](https://github.com/argoproj/argo-cd/pull/708) ([merenbach](https://github.com/merenbach))
- Make Argo CD naming consistent [\#694](https://github.com/argoproj/argo-cd/pull/694) ([tedmiston](https://github.com/tedmiston))
- Split up CRD manifests into their own folder [\#674](https://github.com/argoproj/argo-cd/pull/674) ([twz123](https://github.com/twz123))

## [v0.10.6](https://github.com/argoproj/argo-cd/tree/v0.10.6) (2018-11-15)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.10.5...v0.10.6)

## [v0.10.5](https://github.com/argoproj/argo-cd/tree/v0.10.5) (2018-11-14)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.10.4...v0.10.5)

**Fixed bugs:**

- `app create --upsert` does not update --dest-server and --dest-namespace from git [\#750](https://github.com/argoproj/argo-cd/issues/750)

**Closed issues:**

- \[UI\] Application create wizard crashes  [\#768](https://github.com/argoproj/argo-cd/issues/768)
- Declarative setup and configuration of ArgoCD [\#536](https://github.com/argoproj/argo-cd/issues/536)

## [v0.10.4](https://github.com/argoproj/argo-cd/tree/v0.10.4) (2018-11-08)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.10.3...v0.10.4)

**Implemented enhancements:**

- Support customizing name prefix in helm and kustomize [\#721](https://github.com/argoproj/argo-cd/issues/721)
- Support using external OIDC providers [\#671](https://github.com/argoproj/argo-cd/issues/671)

**Fixed bugs:**

- CI job does not update code coverage results [\#657](https://github.com/argoproj/argo-cd/issues/657)
- Better detection of ksonnet apps [\#594](https://github.com/argoproj/argo-cd/issues/594)
- \[UI\] won't show operation unless status.operationState exists [\#410](https://github.com/argoproj/argo-cd/issues/410)

**Closed issues:**

- Health check is not discerning apiVersion for CRDs [\#753](https://github.com/argoproj/argo-cd/issues/753)
- Installing crds via argocd [\#748](https://github.com/argoproj/argo-cd/issues/748)
- Ability to perform field selection in API [\#697](https://github.com/argoproj/argo-cd/issues/697)
- Rename ComponentParameter to just Parameter [\#557](https://github.com/argoproj/argo-cd/issues/557)

## [v0.10.3](https://github.com/argoproj/argo-cd/tree/v0.10.3) (2018-10-29)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.10.2...v0.10.3)

**Fixed bugs:**

- Unable to delete application in K8s v1.12 [\#718](https://github.com/argoproj/argo-cd/issues/718)

**Closed issues:**

- Update to client-go 8.0 to use new dynamic client [\#353](https://github.com/argoproj/argo-cd/issues/353)

## [v0.10.2](https://github.com/argoproj/argo-cd/tree/v0.10.2) (2018-10-25)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.10.1...v0.10.2)

**Implemented enhancements:**

- Allow using `kustomize edit set` during kustomize app resource generation [\#670](https://github.com/argoproj/argo-cd/issues/670)

**Fixed bugs:**

- ODICProvider should no longer be memoized [\#710](https://github.com/argoproj/argo-cd/issues/710)

## [v0.10.1](https://github.com/argoproj/argo-cd/tree/v0.10.1) (2018-10-24)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.10.0...v0.10.1)

**Closed issues:**

- git clean also needs to clean files under gitignore [\#711](https://github.com/argoproj/argo-cd/issues/711)

## [v0.10.0](https://github.com/argoproj/argo-cd/tree/v0.10.0) (2018-10-19)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.9.2...v0.10.0)

**Fixed bugs:**

- Application details page don't allow editing parameter if parameter name has '.' [\#707](https://github.com/argoproj/argo-cd/issues/707)
- \[UI\] Input type text instead of password on Connect repo panel [\#693](https://github.com/argoproj/argo-cd/issues/693)
- Resource is always out of sync if it has only 'ksonnet.io/component' label [\#686](https://github.com/argoproj/argo-cd/issues/686)
- \[UI\] Better update conflict error handing during app editing in UI [\#685](https://github.com/argoproj/argo-cd/issues/685)
- \[UI\] Resource nodes are 'jumping' on app details page [\#683](https://github.com/argoproj/argo-cd/issues/683)
- Operation stuck in 'in progress' state if application has no resources [\#682](https://github.com/argoproj/argo-cd/issues/682)
- \[UI\] Sync always suggest using latest revision instead of target [\#669](https://github.com/argoproj/argo-cd/issues/669)
- App creation UI should allow specifying values files outside of helm app directory [\#658](https://github.com/argoproj/argo-cd/issues/658)
- \[UI\] Create role UI does not allow specify policies [\#656](https://github.com/argoproj/argo-cd/issues/656)
- \[UI\] Generate role token click resets policy changes [\#655](https://github.com/argoproj/argo-cd/issues/655)

**Closed issues:**

- Install manifests RBAC are missing event list permissions [\#704](https://github.com/argoproj/argo-cd/issues/704)
- x509: certificate signed by unknown authority on private git repo [\#703](https://github.com/argoproj/argo-cd/issues/703)
- \[UI\] Display init container logs [\#681](https://github.com/argoproj/argo-cd/issues/681)
- Default project is created without permission to deploy cluster level resources [\#679](https://github.com/argoproj/argo-cd/issues/679)
- Adding kustomize app via UI fails in v0.9.2 [\#673](https://github.com/argoproj/argo-cd/issues/673)
- Metrics endpoint not reachable through the metrics kubernetes service [\#672](https://github.com/argoproj/argo-cd/issues/672)
- stateful set is showing progressing  [\#668](https://github.com/argoproj/argo-cd/issues/668)
- We should be able to select the order of the `yaml` files while creating a Helm App [\#664](https://github.com/argoproj/argo-cd/issues/664)
- App Creation UI for Helm Apps shows only files prefixed with `values-`   [\#663](https://github.com/argoproj/argo-cd/issues/663)
- Remove RollbackOperation in favor of Sync with ParameterOverrides [\#628](https://github.com/argoproj/argo-cd/issues/628)
- Cluster watch needs to be restarted when CRDs get created [\#627](https://github.com/argoproj/argo-cd/issues/627)
- \[UI\] Support ability to use a helm values files from a URL [\#624](https://github.com/argoproj/argo-cd/issues/624)
- \[UI\] Project should influence options for cluster and namespace during app creation [\#592](https://github.com/argoproj/argo-cd/issues/592)
- Extract status.history.params into separate API call [\#556](https://github.com/argoproj/argo-cd/issues/556)
- Repo-server performance improvements [\#554](https://github.com/argoproj/argo-cd/issues/554)
- Need a way to sync a single resource  [\#508](https://github.com/argoproj/argo-cd/issues/508)
- Support public not-connected repo in app creation UI [\#426](https://github.com/argoproj/argo-cd/issues/426)
- \[UI\] Redirect to /auth/login instead of /login when SSO token expires [\#348](https://github.com/argoproj/argo-cd/issues/348)

**Merged pull requests:**

- Issue \#628 - Remove RollbackOperation in favor of Sync with ParameterOverrides [\#706](https://github.com/argoproj/argo-cd/pull/706) ([alexmt](https://github.com/alexmt))
- RBAC for cluster wide install was missing permissions to list events across namespaces \(resolves \#704\) [\#705](https://github.com/argoproj/argo-cd/pull/705) ([jessesuen](https://github.com/jessesuen))
- Add 0.10 changelog [\#700](https://github.com/argoproj/argo-cd/pull/700) ([alexmt](https://github.com/alexmt))
- Issue \#672 - Metrics endpoint not reachable through the metrics kubenetes service [\#692](https://github.com/argoproj/argo-cd/pull/692) ([alexmt](https://github.com/alexmt))
- Issue \#690 - Increase GRPC message limit [\#691](https://github.com/argoproj/argo-cd/pull/691) ([alexmt](https://github.com/alexmt))
- Add argocd-util cluster-kubeconfig command [\#689](https://github.com/argoproj/argo-cd/pull/689) ([alexmt](https://github.com/alexmt))
- Issue \#686 - Resource is always out of sync if it has only 'ksonnet.io/component' label [\#688](https://github.com/argoproj/argo-cd/pull/688) ([alexmt](https://github.com/alexmt))
- Issue \#682 - Operation stuck in 'in progress' state if application has no resources [\#684](https://github.com/argoproj/argo-cd/pull/684) ([alexmt](https://github.com/alexmt))
- Issue \#679 - Default project is created without permission to deploy cluster level resources [\#680](https://github.com/argoproj/argo-cd/pull/680) ([alexmt](https://github.com/alexmt))
- Issue \#627 - Cluster watch needs to be restarted when CRDs get created [\#678](https://github.com/argoproj/argo-cd/pull/678) ([alexmt](https://github.com/alexmt))
- Issue \#426 - Support public not-connected repo in app creation UI [\#675](https://github.com/argoproj/argo-cd/pull/675) ([alexmt](https://github.com/alexmt))
- Upgrade testify [\#667](https://github.com/argoproj/argo-cd/pull/667) ([merenbach](https://github.com/merenbach))
- Allow more fine-grained sync \(\#508\) [\#666](https://github.com/argoproj/argo-cd/pull/666) ([merenbach](https://github.com/merenbach))
- Add project get permission automatically to roles [\#665](https://github.com/argoproj/argo-cd/pull/665) ([dthomson25](https://github.com/dthomson25))
- Add v0.9.2 changelog [\#662](https://github.com/argoproj/argo-cd/pull/662) ([alexmt](https://github.com/alexmt))
- Issue \#650 - Temporary ignore service catalog resources [\#661](https://github.com/argoproj/argo-cd/pull/661) ([alexmt](https://github.com/alexmt))
- Update generated files [\#660](https://github.com/argoproj/argo-cd/pull/660) ([merenbach](https://github.com/merenbach))
- Normalize policies by always adding space after comma [\#659](https://github.com/argoproj/argo-cd/pull/659) ([dthomson25](https://github.com/dthomson25))
- Switch to go-git for all remote git interactions including auth \(issue \#651\) [\#652](https://github.com/argoproj/argo-cd/pull/652) ([jessesuen](https://github.com/jessesuen))
- Remove default params from app history [\#649](https://github.com/argoproj/argo-cd/pull/649) ([dthomson25](https://github.com/dthomson25))
- Add errgroup dependency for Packr [\#648](https://github.com/argoproj/argo-cd/pull/648) ([merenbach](https://github.com/merenbach))
- Add version check during release to ensure compiled version is accurate [\#646](https://github.com/argoproj/argo-cd/pull/646) ([jessesuen](https://github.com/jessesuen))
- Do not append .git extension during normalization for Azure hosted git \(issue \#643\) [\#645](https://github.com/argoproj/argo-cd/pull/645) ([jessesuen](https://github.com/jessesuen))
- update to kustomize 1.0.8 [\#644](https://github.com/argoproj/argo-cd/pull/644) ([essh](https://github.com/essh))
- Documentation clarifications and fixes [\#642](https://github.com/argoproj/argo-cd/pull/642) ([jessesuen](https://github.com/jessesuen))
- Update getting\_started.md with new version; update releasing steps [\#641](https://github.com/argoproj/argo-cd/pull/641) ([alexmt](https://github.com/alexmt))
- Issue \#639 - Repo server unable to execute ls-remote for private repos [\#640](https://github.com/argoproj/argo-cd/pull/640) ([alexmt](https://github.com/alexmt))
- Use ksonnet CLI instead of ksonnet libs \(\#590\) [\#626](https://github.com/argoproj/argo-cd/pull/626) ([merenbach](https://github.com/merenbach))
- Update documentation with auto-sync and projects \(issue \#521\) [\#616](https://github.com/argoproj/argo-cd/pull/616) ([jessesuen](https://github.com/jessesuen))

## [v0.9.2](https://github.com/argoproj/argo-cd/tree/v0.9.2) (2018-09-28)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.9.1...v0.9.2)

**Closed issues:**

- argocd-server logs credentials in plain text when adding git repositories [\#653](https://github.com/argoproj/argo-cd/issues/653)
- Credentials not being accepted for Google Source Repositories [\#651](https://github.com/argoproj/argo-cd/issues/651)
- Missing errgroup package when installing via dep ensure [\#647](https://github.com/argoproj/argo-cd/issues/647)
- Azure Repos do not work as a repository [\#643](https://github.com/argoproj/argo-cd/issues/643)
- Repo server unable to execute ls-remote for private repos  [\#639](https://github.com/argoproj/argo-cd/issues/639)
- Use ksonnet CLI instead of ksonnet lib [\#590](https://github.com/argoproj/argo-cd/issues/590)
- Documentation for auto-sync [\#521](https://github.com/argoproj/argo-cd/issues/521)

## [v0.9.1](https://github.com/argoproj/argo-cd/tree/v0.9.1) (2018-09-24)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.9.0...v0.9.1)

## [v0.9.0](https://github.com/argoproj/argo-cd/tree/v0.9.0) (2018-09-24)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.8.2...v0.9.0)

**Implemented enhancements:**

- Use 'kubectl auth reconcile' before 'kubectl apply' [\#523](https://github.com/argoproj/argo-cd/issues/523)

**Fixed bugs:**

- Helm hooks are being deployed as resources [\#605](https://github.com/argoproj/argo-cd/issues/605)
- Disagreement in three way diff calculation [\#597](https://github.com/argoproj/argo-cd/issues/597)
- SIGSEGV in kube.GetResourcesWithLabel [\#587](https://github.com/argoproj/argo-cd/issues/587)
- ArgoCD fails to deploy resources list [\#584](https://github.com/argoproj/argo-cd/issues/584)
- App comparison state fails if target resource type is not supported [\#583](https://github.com/argoproj/argo-cd/issues/583)

**Closed issues:**

- Improve event messages [\#635](https://github.com/argoproj/argo-cd/issues/635)
- Simplify the RBAC resources to remove unnecessary sub-resources [\#629](https://github.com/argoproj/argo-cd/issues/629)
- ClusterRoleBinding defaults not detected [\#620](https://github.com/argoproj/argo-cd/issues/620)
- Rename server.crt and server.key to tls.crt and tls.key [\#617](https://github.com/argoproj/argo-cd/issues/617)
- \[UI\] Ability to modify application from UI [\#615](https://github.com/argoproj/argo-cd/issues/615)
- Trim url being pasted into Connect Repository [\#614](https://github.com/argoproj/argo-cd/issues/614)
- Special handling of CustomResourceDefinitions [\#613](https://github.com/argoproj/argo-cd/issues/613)
- Support restricting TLS version [\#609](https://github.com/argoproj/argo-cd/issues/609)
- \[UI\] Error after sync [\#601](https://github.com/argoproj/argo-cd/issues/601)
- Lazy enforcement of unknown cluster/namespace restricted resources [\#599](https://github.com/argoproj/argo-cd/issues/599)
- Improve remarshalling to use reflection/schema builders to handle all k8s core types [\#595](https://github.com/argoproj/argo-cd/issues/595)
- Helm example fails with no such file or directory when choosing values file [\#591](https://github.com/argoproj/argo-cd/issues/591)
- \[SPIKE\] check compatibility with ArgoCD and nginx-ingress [\#585](https://github.com/argoproj/argo-cd/issues/585)
- ArgoCD should download helm chart dependencies [\#582](https://github.com/argoproj/argo-cd/issues/582)
- Expand RBAC for application-controller/api-server [\#580](https://github.com/argoproj/argo-cd/issues/580)
- \[UI\] indicate when operation is in progress or has failed [\#566](https://github.com/argoproj/argo-cd/issues/566)
- \[UI\] Creating apps from directories is not obvious [\#565](https://github.com/argoproj/argo-cd/issues/565)
- Split out dex into it's own deployment \(instead of sidecar\) [\#555](https://github.com/argoproj/argo-cd/issues/555)
- \[UI\] Project Role/Token management from UI [\#548](https://github.com/argoproj/argo-cd/issues/548)
- Add readiness probe to repo server [\#529](https://github.com/argoproj/argo-cd/issues/529)
- Investigate running API server with replicas [\#515](https://github.com/argoproj/argo-cd/issues/515)
- Export ArgoCD stats as prometheus style metrics [\#513](https://github.com/argoproj/argo-cd/issues/513)
- Support IAM Authentication for managing external K8s clusters [\#482](https://github.com/argoproj/argo-cd/issues/482)
- Projects need controls on cluster-scoped resources [\#330](https://github.com/argoproj/argo-cd/issues/330)
- Utilize a rolling update strategy in api-server and repo-server [\#224](https://github.com/argoproj/argo-cd/issues/224)
- ksonnet show should order resources by k8s proper creation order [\#102](https://github.com/argoproj/argo-cd/issues/102)

**Merged pull requests:**

- Update changelog and fix release command dependency [\#638](https://github.com/argoproj/argo-cd/pull/638) ([alexmt](https://github.com/alexmt))
- Add more event sources and provide better detail in event messages \(issue \#635\) [\#637](https://github.com/argoproj/argo-cd/pull/637) ([jessesuen](https://github.com/jessesuen))
- Update manifests to support in-cluster installations [\#634](https://github.com/argoproj/argo-cd/pull/634) ([jessesuen](https://github.com/jessesuen))
- Issue \#613 - Don't delete CRD [\#630](https://github.com/argoproj/argo-cd/pull/630) ([alexmt](https://github.com/alexmt))
- Support ability to use a helm values files from a URL \(issue \#624\) [\#625](https://github.com/argoproj/argo-cd/pull/625) ([jessesuen](https://github.com/jessesuen))
- Trim git url during normalization \(issue \#614\) [\#623](https://github.com/argoproj/argo-cd/pull/623) ([jessesuen](https://github.com/jessesuen))
- Fix false OutOfSync condition when an explicit namespace is set in the config [\#622](https://github.com/argoproj/argo-cd/pull/622) ([jessesuen](https://github.com/jessesuen))
- Simplify the RBAC resources to remove unnecessary sub-resources \(issue \#629\) [\#619](https://github.com/argoproj/argo-cd/pull/619) ([jessesuen](https://github.com/jessesuen))
- Rename server.crt/server.key to tls.crt/tls.key to integrate with Ingress \(issue \#617\) [\#618](https://github.com/argoproj/argo-cd/pull/618) ([jessesuen](https://github.com/jessesuen))
- Issue \#599 - Lazy enforcement of unknown cluster/namespace restricted resources [\#612](https://github.com/argoproj/argo-cd/pull/612) ([alexmt](https://github.com/alexmt))
- 609 - Support restricting TLS version [\#610](https://github.com/argoproj/argo-cd/pull/610) ([alexmt](https://github.com/alexmt))
- Support for exporting prometheus metrics about ArgoCD applications \(issue \#513\) [\#608](https://github.com/argoproj/argo-cd/pull/608) ([jessesuen](https://github.com/jessesuen))
- Update getting started to point to v0.8.2 [\#607](https://github.com/argoproj/argo-cd/pull/607) ([kuy](https://github.com/kuy))
- Support helm charts with dependencies and namespace sensitivity \(issue \#582\) [\#606](https://github.com/argoproj/argo-cd/pull/606) ([jessesuen](https://github.com/jessesuen))
- Improve three-way diff to provide more accurate Sync status and diff result \(issue \#597\) [\#604](https://github.com/argoproj/argo-cd/pull/604) ([jessesuen](https://github.com/jessesuen))
- Improve remarshalling to use reflection/schema builders to handle all k8s core types [\#603](https://github.com/argoproj/argo-cd/pull/603) ([jessesuen](https://github.com/jessesuen))
- Issue \#515 - handle concurrent settings initialization by api servers [\#602](https://github.com/argoproj/argo-cd/pull/602) ([alexmt](https://github.com/alexmt))
- Issue \#523 - Use 'kubectl auth reconcile' for RBAC resources [\#600](https://github.com/argoproj/argo-cd/pull/600) ([alexmt](https://github.com/alexmt))
- Issue \#584 - ArgoCD fails to deploy resources list [\#598](https://github.com/argoproj/argo-cd/pull/598) ([alexmt](https://github.com/alexmt))
- Fix comparison failure when app contains unregistered custom resource \(issue \#583\) [\#596](https://github.com/argoproj/argo-cd/pull/596) ([jessesuen](https://github.com/jessesuen))
- Fix race in kube.GetResourcesWithLabel with concurrent appends to a list \(issue \#587\) [\#593](https://github.com/argoproj/argo-cd/pull/593) ([jessesuen](https://github.com/jessesuen))
- Cherry pick v0.8.2 fixes [\#589](https://github.com/argoproj/argo-cd/pull/589) ([jessesuen](https://github.com/jessesuen))
- Issue \#482 - Support IAM Authentication for managing external K8s clusters [\#588](https://github.com/argoproj/argo-cd/pull/588) ([alexmt](https://github.com/alexmt))
- Add iat as path param for delete token http call [\#586](https://github.com/argoproj/argo-cd/pull/586) ([dthomson25](https://github.com/dthomson25))
- Issue \#577 - Add rbac metrics non resourcce url policy for argocd-manager-role [\#578](https://github.com/argoproj/argo-cd/pull/578) ([conorfennell](https://github.com/conorfennell))
- Update getting started and docs to point to v0.8.1 [\#575](https://github.com/argoproj/argo-cd/pull/575) ([jessesuen](https://github.com/jessesuen))
- In-memory implementation of ls-remote using go-git to reduce repo lock contention [\#574](https://github.com/argoproj/argo-cd/pull/574) ([jessesuen](https://github.com/jessesuen))
- Add link to SigApps video and update CHANGELOG for v0.8.1 [\#572](https://github.com/argoproj/argo-cd/pull/572) ([jessesuen](https://github.com/jessesuen))
- Introduce ability to automatically sync upon detected git changes [\#571](https://github.com/argoproj/argo-cd/pull/571) ([jessesuen](https://github.com/jessesuen))
- Fix controller hot loop when app source contains bad manifests \(issue \#568\) [\#570](https://github.com/argoproj/argo-cd/pull/570) ([jessesuen](https://github.com/jessesuen))
- Fix issue where branch checkout did not have accurate git tree state \(issue \#567\) [\#569](https://github.com/argoproj/argo-cd/pull/569) ([jessesuen](https://github.com/jessesuen))
- Derive dedicated Dex deployment [\#564](https://github.com/argoproj/argo-cd/pull/564) ([merenbach](https://github.com/merenbach))
- Issue \#553 - Turn on TLS for repo server [\#563](https://github.com/argoproj/argo-cd/pull/563) ([alexmt](https://github.com/alexmt))
- Issue \#540 - Support raw jsonnet as an application source [\#561](https://github.com/argoproj/argo-cd/pull/561) ([alexmt](https://github.com/alexmt))
- Issue \#470 - K8s secrets need to be redacted in API server [\#560](https://github.com/argoproj/argo-cd/pull/560) ([alexmt](https://github.com/alexmt))
- Issue \#527 - Support --in-cluster authentication without providing a kubeconfig [\#559](https://github.com/argoproj/argo-cd/pull/559) ([alexmt](https://github.com/alexmt))
- Issue \#330 - Projects need controls on cluster-scoped resources [\#558](https://github.com/argoproj/argo-cd/pull/558) ([alexmt](https://github.com/alexmt))
- Issue 499 - fileFiles path should be relative to app directory [\#552](https://github.com/argoproj/argo-cd/pull/552) ([alexmt](https://github.com/alexmt))
- Reorder K8s resources to correct creation order [\#551](https://github.com/argoproj/argo-cd/pull/551) ([dthomson25](https://github.com/dthomson25))
- Update documentation for v0.8.0 [\#550](https://github.com/argoproj/argo-cd/pull/550) ([jessesuen](https://github.com/jessesuen))

## [v0.8.2](https://github.com/argoproj/argo-cd/tree/v0.8.2) (2018-09-12)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.8.1...v0.8.2)

**Implemented enhancements:**

- App creation wizard should allow specifying source revision [\#562](https://github.com/argoproj/argo-cd/issues/562)

**Closed issues:**

- Support prometheus operator rbac defaults [\#577](https://github.com/argoproj/argo-cd/issues/577)
- \[UI\] Projects filter does not work when application got changed [\#573](https://github.com/argoproj/argo-cd/issues/573)
- \[UI\] provide a YAML view of resources [\#396](https://github.com/argoproj/argo-cd/issues/396)
- Auto-sync option in application CRD instance [\#79](https://github.com/argoproj/argo-cd/issues/79)

**Merged pull requests:**

- Downgrade ksonnet from v0.12.0 to v0.11.0 due to quote unescape regression [\#581](https://github.com/argoproj/argo-cd/pull/581) ([jessesuen](https://github.com/jessesuen))

## [v0.8.1](https://github.com/argoproj/argo-cd/tree/v0.8.1) (2018-09-10)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.8.0...v0.8.1)

**Implemented enhancements:**

- \[UI\]Support helm values files in App creation wizard [\#499](https://github.com/argoproj/argo-cd/issues/499)

**Closed issues:**

- Controller hot loop when application source has bad manifests [\#568](https://github.com/argoproj/argo-cd/issues/568)
- Branch tracking not working properly [\#567](https://github.com/argoproj/argo-cd/issues/567)
- Turn on TLS for repo server [\#553](https://github.com/argoproj/argo-cd/issues/553)
- Support raw jsonnet as an application source [\#540](https://github.com/argoproj/argo-cd/issues/540)
- \[UI\] Clear indication of pods with containers not ready [\#539](https://github.com/argoproj/argo-cd/issues/539)
- Support --in-cluster authentication without providing a kubeconfig [\#527](https://github.com/argoproj/argo-cd/issues/527)
- Add documentation on project tokens [\#519](https://github.com/argoproj/argo-cd/issues/519)
- \[UI\] Top level indicator when app is overriding parameters [\#503](https://github.com/argoproj/argo-cd/issues/503)
- K8s secrets need to be redacted in API server [\#470](https://github.com/argoproj/argo-cd/issues/470)
- Improve resource diff rendering [\#457](https://github.com/argoproj/argo-cd/issues/457)
- Example repository using ArgoCD from a CI pipeline [\#436](https://github.com/argoproj/argo-cd/issues/436)

## [v0.8.0](https://github.com/argoproj/argo-cd/tree/v0.8.0) (2018-09-05)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.7.2...v0.8.0)

**Fixed bugs:**

- \[UI\] Project deletion does not wait for confirmation [\#545](https://github.com/argoproj/argo-cd/issues/545)
- argocd app wait is showing incorrect information [\#542](https://github.com/argoproj/argo-cd/issues/542)
- unable to retrieve the complete list of server APIs: servicecatalog.k8s.io/v1beta1: the server is currently unable to handle the request [\#524](https://github.com/argoproj/argo-cd/issues/524)

**Closed issues:**

- \[UI\] Project management UI does not allow `\*` additions [\#547](https://github.com/argoproj/argo-cd/issues/547)
- Rebuild images to run containers as non-root users [\#528](https://github.com/argoproj/argo-cd/issues/528)
- Proper health checks and liveness/readiness probes [\#520](https://github.com/argoproj/argo-cd/issues/520)
- Default project needs to be auto created on startup [\#514](https://github.com/argoproj/argo-cd/issues/514)
- Translate Unknown errors into proper gRPC codes [\#512](https://github.com/argoproj/argo-cd/issues/512)
- Health check for PVCs [\#501](https://github.com/argoproj/argo-cd/issues/501)
- Argo CD may benefit from a code of conduct [\#93](https://github.com/argoproj/argo-cd/issues/93)

**Merged pull requests:**

- Minor improvements to token CLI [\#549](https://github.com/argoproj/argo-cd/pull/549) ([jessesuen](https://github.com/jessesuen))
- Run all containers as an unprivileged user \(resolves \#528\) [\#546](https://github.com/argoproj/argo-cd/pull/546) ([jessesuen](https://github.com/jessesuen))
- Make manifests friendly to `kubectl apply` semantics by omitting `data:` field [\#544](https://github.com/argoproj/argo-cd/pull/544) ([jessesuen](https://github.com/jessesuen))
- Fix `argocd app wait` printing incorrect Sync output \(resolves \#542\) [\#543](https://github.com/argoproj/argo-cd/pull/543) ([jessesuen](https://github.com/jessesuen))
- Fix issue where argocd could not sync to a tag [\#541](https://github.com/argoproj/argo-cd/pull/541) ([jessesuen](https://github.com/jessesuen))
- Refactor Makefile/build to use a single Dockerfile. Update kustomize to v1.0.7 [\#538](https://github.com/argoproj/argo-cd/pull/538) ([jessesuen](https://github.com/jessesuen))
- Add PVC healthcheck to controller \(\#501\) [\#537](https://github.com/argoproj/argo-cd/pull/537) ([dthomson25](https://github.com/dthomson25))
- Create default project on startup [\#535](https://github.com/argoproj/argo-cd/pull/535) ([dthomson25](https://github.com/dthomson25))
- Add documentation on project roles and JWT tokens [\#533](https://github.com/argoproj/argo-cd/pull/533) ([dthomson25](https://github.com/dthomson25))
- Use gRPC error codes instead of fmt.Errorf [\#532](https://github.com/argoproj/argo-cd/pull/532) ([merenbach](https://github.com/merenbach))
- API discovery becomes best effort when partial resource list is returned \(resolves \#524\) [\#526](https://github.com/argoproj/argo-cd/pull/526) ([jessesuen](https://github.com/jessesuen))
- Add health check on API server [\#522](https://github.com/argoproj/argo-cd/pull/522) ([merenbach](https://github.com/merenbach))
- Fix typo in sso.md [\#518](https://github.com/argoproj/argo-cd/pull/518) ([dthomson25](https://github.com/dthomson25))
- Fix build breakage [\#517](https://github.com/argoproj/argo-cd/pull/517) ([jessesuen](https://github.com/jessesuen))
- Use named FIFO so we can exit with non-zero status [\#516](https://github.com/argoproj/argo-cd/pull/516) ([merenbach](https://github.com/merenbach))
- Add ability to delete a single application resource \(issue \#262\) [\#511](https://github.com/argoproj/argo-cd/pull/511) ([jessesuen](https://github.com/jessesuen))
- Support for kustomize app directories [\#510](https://github.com/argoproj/argo-cd/pull/510) ([jessesuen](https://github.com/jessesuen))
- Issue 435 - pump ci logs to s3 [\#509](https://github.com/argoproj/argo-cd/pull/509) ([alexmt](https://github.com/alexmt))
- Upgrade Ksonnet [\#506](https://github.com/argoproj/argo-cd/pull/506) ([merenbach](https://github.com/merenbach))
- Issue \#458 - Add api to load project events [\#504](https://github.com/argoproj/argo-cd/pull/504) ([alexmt](https://github.com/alexmt))
- Issue \#489 - Static assets are being browser cached between upgrades [\#502](https://github.com/argoproj/argo-cd/pull/502) ([alexmt](https://github.com/alexmt))
- Enable code coverage [\#500](https://github.com/argoproj/argo-cd/pull/500) ([merenbach](https://github.com/merenbach))
- Add ability to create tokens [\#498](https://github.com/argoproj/argo-cd/pull/498) ([dthomson25](https://github.com/dthomson25))
- Support explicit deny [\#497](https://github.com/argoproj/argo-cd/pull/497) ([merenbach](https://github.com/merenbach))
- Support UI cluster creation [\#469](https://github.com/argoproj/argo-cd/pull/469) ([merenbach](https://github.com/merenbach))

## [v0.7.2](https://github.com/argoproj/argo-cd/tree/v0.7.2) (2018-08-21)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.7.1...v0.7.2)

**Fixed bugs:**

- \[SPIKE\] Controller fails to update deployment history [\#449](https://github.com/argoproj/argo-cd/issues/449)

**Closed issues:**

- Feature Request - Support argocd-cli installation on archlinux [\#488](https://github.com/argoproj/argo-cd/issues/488)
- Add project tokens for automated access [\#472](https://github.com/argoproj/argo-cd/issues/472)
- Upgrade to ksonnet v0.12 [\#465](https://github.com/argoproj/argo-cd/issues/465)
- \[SPIKE\]Project management UI [\#458](https://github.com/argoproj/argo-cd/issues/458)
- Support kustomize as an application source [\#451](https://github.com/argoproj/argo-cd/issues/451)
- \[SPIKE\] Enable log artifacts for CI jobs to store test logs in a S3 bucket [\#435](https://github.com/argoproj/argo-cd/issues/435)
- support explicit deny in RBAC policy [\#322](https://github.com/argoproj/argo-cd/issues/322)
- Support cluster creation from UI [\#303](https://github.com/argoproj/argo-cd/issues/303)
- Enable code coverage [\#302](https://github.com/argoproj/argo-cd/issues/302)
- Need a way to delete a single resource [\#262](https://github.com/argoproj/argo-cd/issues/262)

**Merged pull requests:**

- API discovery becomes best effort when partial resource list is returned \(resolves \#524\) [\#525](https://github.com/argoproj/argo-cd/pull/525) ([jessesuen](https://github.com/jessesuen))

## [v0.7.1](https://github.com/argoproj/argo-cd/tree/v0.7.1) (2018-08-03)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.7.0...v0.7.1)

**Fixed bugs:**

- Failed e2e test does not fail CI workflow [\#491](https://github.com/argoproj/argo-cd/issues/491)
- \[UI\] application card is showing green, when health is degraded [\#480](https://github.com/argoproj/argo-cd/issues/480)
- AppProjectSpec SourceRepos mislabeled [\#476](https://github.com/argoproj/argo-cd/issues/476)
- Delete application confirmation incorrect redirect [\#447](https://github.com/argoproj/argo-cd/issues/447)
- Improve data loading error notification [\#446](https://github.com/argoproj/argo-cd/issues/446)

**Closed issues:**

- \[UI\] indicator when refresh is still in progress [\#493](https://github.com/argoproj/argo-cd/issues/493)
- Application events does not work end-to-end [\#478](https://github.com/argoproj/argo-cd/issues/478)
- argocd relogin should not ask for username [\#475](https://github.com/argoproj/argo-cd/issues/475)
- ListApps API does not scale [\#474](https://github.com/argoproj/argo-cd/issues/474)
- Surface helm parameters to the application level [\#463](https://github.com/argoproj/argo-cd/issues/463)
- Improve application creation wizard [\#459](https://github.com/argoproj/argo-cd/issues/459)
- Documentation for Helm and Manifest directory application sources [\#444](https://github.com/argoproj/argo-cd/issues/444)
- \[SPIKE\] Design API to create system users [\#228](https://github.com/argoproj/argo-cd/issues/228)

**Merged pull requests:**

- Update manifests and install instructions for v0.7.1 [\#496](https://github.com/argoproj/argo-cd/pull/496) ([jessesuen](https://github.com/jessesuen))
- Fix 404 error in repo API [\#495](https://github.com/argoproj/argo-cd/pull/495) ([alexmt](https://github.com/alexmt))
- Issue \#474 - ListApps API does not scale [\#494](https://github.com/argoproj/argo-cd/pull/494) ([alexmt](https://github.com/alexmt))
- Issue \#491 - Failed e2e test does not fail CI workflow [\#492](https://github.com/argoproj/argo-cd/pull/492) ([alexmt](https://github.com/alexmt))
- Issue \#476 - AppProjectSpec SourceRepos mislabeled [\#490](https://github.com/argoproj/argo-cd/pull/490) ([alexmt](https://github.com/alexmt))
- Update getting\_started.md [\#487](https://github.com/argoproj/argo-cd/pull/487) ([chocopowwwa](https://github.com/chocopowwwa))
- Fix failure in identifying app source type when path was '.' [\#486](https://github.com/argoproj/argo-cd/pull/486) ([jessesuen](https://github.com/jessesuen))
- Issue \#463 - Surface helm parameters to the application level [\#485](https://github.com/argoproj/argo-cd/pull/485) ([alexmt](https://github.com/alexmt))
- Fix issue where application server was retrieving events from incorrect cluster \(resolves \#478\) [\#484](https://github.com/argoproj/argo-cd/pull/484) ([jessesuen](https://github.com/jessesuen))
- Infer username from claims during an `argocd relogin` \(resolves \#475\) [\#483](https://github.com/argoproj/argo-cd/pull/483) ([jessesuen](https://github.com/jessesuen))
- Expand RBAC role to be able to create application events. Fix username claims extraction. [\#479](https://github.com/argoproj/argo-cd/pull/479) ([jessesuen](https://github.com/jessesuen))
- Create update-manifests.sh script to support manifest generation for personal images [\#477](https://github.com/argoproj/argo-cd/pull/477) ([jessesuen](https://github.com/jessesuen))
- Update getting\_started.md with relogin command during password change [\#473](https://github.com/argoproj/argo-cd/pull/473) ([jessesuen](https://github.com/jessesuen))

## [v0.7.0](https://github.com/argoproj/argo-cd/tree/v0.7.0) (2018-07-28)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.6.2...v0.7.0)

**Fixed bugs:**

- UI does not allow to select project [\#442](https://github.com/argoproj/argo-cd/issues/442)

**Closed issues:**

- e2e tests are broken [\#461](https://github.com/argoproj/argo-cd/issues/461)
- Make use of dex refresh tokens [\#455](https://github.com/argoproj/argo-cd/issues/455)
- Reposerver is logging git username/password/ssh key in the clear [\#445](https://github.com/argoproj/argo-cd/issues/445)
- UI changes for selecting Helm and manifest app directories [\#443](https://github.com/argoproj/argo-cd/issues/443)
- New audit logging interceptor is logging passwords in the clear [\#438](https://github.com/argoproj/argo-cd/issues/438)
- Support directory of YAML/JSON as an application source [\#431](https://github.com/argoproj/argo-cd/issues/431)
- Rework local authentication to expire tokens when password is updated [\#424](https://github.com/argoproj/argo-cd/issues/424)
- \[SPIKE\] Controller is using too much memory [\#419](https://github.com/argoproj/argo-cd/issues/419)
- One bad health check can affect all resource health checks [\#398](https://github.com/argoproj/argo-cd/issues/398)
- CLI needs progress to be shown during sync \(similar to wait\) [\#389](https://github.com/argoproj/argo-cd/issues/389)
- Expose Ksonnet version in the argocd version command [\#346](https://github.com/argoproj/argo-cd/issues/346)
- Need complete audit trail [\#340](https://github.com/argoproj/argo-cd/issues/340)

**Merged pull requests:**

- Bump VERSION to v0.7.0 [\#471](https://github.com/argoproj/argo-cd/pull/471) ([jessesuen](https://github.com/jessesuen))
- Add `argocd relogin` command as a convenience around login to current context [\#468](https://github.com/argoproj/argo-cd/pull/468) ([jessesuen](https://github.com/jessesuen))
- Issue \#376 - Fix saving default connection status for repos and clusters [\#467](https://github.com/argoproj/argo-cd/pull/467) ([alexmt](https://github.com/alexmt))
- Add documentation for helm, application sources, and parameter overrides [\#466](https://github.com/argoproj/argo-cd/pull/466) ([jessesuen](https://github.com/jessesuen))
- Issue \#461 - Fix broken e2e tests [\#464](https://github.com/argoproj/argo-cd/pull/464) ([alexmt](https://github.com/alexmt))
- Add verbose flag to tests [\#462](https://github.com/argoproj/argo-cd/pull/462) ([merenbach](https://github.com/merenbach))
- Issue \#443 - API to list helm apps [\#460](https://github.com/argoproj/argo-cd/pull/460) ([alexmt](https://github.com/alexmt))
- Make use of dex refresh tokens and store them into local config [\#456](https://github.com/argoproj/argo-cd/pull/456) ([jessesuen](https://github.com/jessesuen))
- Update install manifests to v0.6.2 [\#452](https://github.com/argoproj/argo-cd/pull/452) ([jessesuen](https://github.com/jessesuen))
- Expire local superuser tokens when their password changes [\#450](https://github.com/argoproj/argo-cd/pull/450) ([merenbach](https://github.com/merenbach))
- Issue 419 - Clean watch resources [\#448](https://github.com/argoproj/argo-cd/pull/448) ([alexmt](https://github.com/alexmt))
- Issue \#438 - audit logging interceptor is logging passwords in the clear [\#441](https://github.com/argoproj/argo-cd/pull/441) ([alexmt](https://github.com/alexmt))
- Issue \#340 - create application/project events for audit [\#440](https://github.com/argoproj/argo-cd/pull/440) ([alexmt](https://github.com/alexmt))
- Health check was using wrong converter for statefulsets, daemonset, replicasets [\#439](https://github.com/argoproj/argo-cd/pull/439) ([jessesuen](https://github.com/jessesuen))
- Resolves 398 -\> errors set status to known, status details set to error message [\#437](https://github.com/argoproj/argo-cd/pull/437) ([JazminGonzalez-Rivero](https://github.com/JazminGonzalez-Rivero))
- Issue \#340 - add gRPC payload logging interceptor [\#434](https://github.com/argoproj/argo-cd/pull/434) ([alexmt](https://github.com/alexmt))
- Add ksonnet version to version endpoint [\#433](https://github.com/argoproj/argo-cd/pull/433) ([dthomson25](https://github.com/dthomson25))
- Support helm charts and yaml directories as an application source [\#432](https://github.com/argoproj/argo-cd/pull/432) ([jessesuen](https://github.com/jessesuen))
- Show CLI progress for sync and rollback [\#393](https://github.com/argoproj/argo-cd/pull/393) ([merenbach](https://github.com/merenbach))

## [v0.6.2](https://github.com/argoproj/argo-cd/tree/v0.6.2) (2018-07-21)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.6.1...v0.6.2)

## [v0.6.1](https://github.com/argoproj/argo-cd/tree/v0.6.1) (2018-07-18)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.6.0...v0.6.1)

**Closed issues:**

- Error from server \(Forbidden\) [\#428](https://github.com/argoproj/argo-cd/issues/428)
- dex [\#427](https://github.com/argoproj/argo-cd/issues/427)
- SSO login failure due to restricted access gives obtuse dex error page [\#351](https://github.com/argoproj/argo-cd/issues/351)

**Merged pull requests:**

- Issue \#428 - Add GKE specific installation instructions [\#430](https://github.com/argoproj/argo-cd/pull/430) ([alexmt](https://github.com/alexmt))
- Add UI GIF, docs for application health, resource hooks, tweaks to README.md [\#429](https://github.com/argoproj/argo-cd/pull/429) ([jessesuen](https://github.com/jessesuen))
- Issue \#351 - forward dex error message to login page [\#425](https://github.com/argoproj/argo-cd/pull/425) ([alexmt](https://github.com/alexmt))
- Add RBAC documentation [\#423](https://github.com/argoproj/argo-cd/pull/423) ([alexmt](https://github.com/alexmt))
- Fix app creation command in getting\_started.md [\#422](https://github.com/argoproj/argo-cd/pull/422) ([alexmt](https://github.com/alexmt))

## [v0.6.0](https://github.com/argoproj/argo-cd/tree/v0.6.0) (2018-07-17)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.5.4...v0.6.0)

**Closed issues:**

- nil pointer in `argocd cluster add`  [\#414](https://github.com/argoproj/argo-cd/issues/414)
- API server nil pointer dereference in GetSpecErrors\(\) [\#407](https://github.com/argoproj/argo-cd/issues/407)
- \[UI\] button to terminate a operation [\#406](https://github.com/argoproj/argo-cd/issues/406)
- `argocd app unset` is broken [\#404](https://github.com/argoproj/argo-cd/issues/404)
- App deployment history don't display parameter overrides [\#402](https://github.com/argoproj/argo-cd/issues/402)
- \[UI\] Provide a link to swagger UI [\#400](https://github.com/argoproj/argo-cd/issues/400)
- swagger-ui does not load due to swagger.json missing [\#399](https://github.com/argoproj/argo-cd/issues/399)
- \[UI\] UI does not show the tracking revision of an app [\#392](https://github.com/argoproj/argo-cd/issues/392)
- \[UI\] should show sync hooks in operation pane [\#390](https://github.com/argoproj/argo-cd/issues/390)
- Sync and health status should be unknown if controller unable to load target/live state [\#385](https://github.com/argoproj/argo-cd/issues/385)
- ArgoCD controller is too chatty [\#381](https://github.com/argoproj/argo-cd/issues/381)
- Optimize repo server to lookup cache when query is for a git SHA1 [\#380](https://github.com/argoproj/argo-cd/issues/380)
- Need a way to cancel a running operation [\#379](https://github.com/argoproj/argo-cd/issues/379)
- \[UI\] Operation details should show message [\#378](https://github.com/argoproj/argo-cd/issues/378)
- API and CLI to change password [\#377](https://github.com/argoproj/argo-cd/issues/377)
- Repos and clusters which have just been added should start out in Connected state [\#376](https://github.com/argoproj/argo-cd/issues/376)
- Add private key & password reset into argocd-util [\#375](https://github.com/argoproj/argo-cd/issues/375)
- Hook cleanup policies [\#374](https://github.com/argoproj/argo-cd/issues/374)
- When operations complete, force refresh comparison state [\#365](https://github.com/argoproj/argo-cd/issues/365)
- StatefulSets, DaemonSets, ReplicaSets need same treatment as Deployments [\#364](https://github.com/argoproj/argo-cd/issues/364)
- PostSync hook should wait until application is Healthy [\#363](https://github.com/argoproj/argo-cd/issues/363)
- \[UI\] resource names are almost always truncated [\#352](https://github.com/argoproj/argo-cd/issues/352)
- argocd wait passes when an ingress object failed [\#349](https://github.com/argoproj/argo-cd/issues/349)
- Getting permission Denied application destination is not permitted in project [\#343](https://github.com/argoproj/argo-cd/issues/343)
- Resources should have a "Missing" health state when they aren't deployed [\#342](https://github.com/argoproj/argo-cd/issues/342)
- \[UI\] add refresh button in app view [\#341](https://github.com/argoproj/argo-cd/issues/341)
- Git authentication using ssh key is broken [\#339](https://github.com/argoproj/argo-cd/issues/339)
- Add sock shop in argocd example apps [\#338](https://github.com/argoproj/argo-cd/issues/338)
- \[UI\] remember my resource filtering preferences [\#337](https://github.com/argoproj/argo-cd/issues/337)
- \[UI\] Terminating pods are shown as Running [\#336](https://github.com/argoproj/argo-cd/issues/336)
- ArgoCD injected app-name does not play nice with `apps/v1beta1` deployment specs [\#335](https://github.com/argoproj/argo-cd/issues/335)
- \[UI\] events message is cut off [\#333](https://github.com/argoproj/argo-cd/issues/333)
- API server --diable-auth flag does not properly disable RBAC [\#332](https://github.com/argoproj/argo-cd/issues/332)
- \[UI\] Remove namespace from UI URL [\#324](https://github.com/argoproj/argo-cd/issues/324)
- app sync should support retries. [\#323](https://github.com/argoproj/argo-cd/issues/323)
- argocd-server-role is too privileged [\#319](https://github.com/argoproj/argo-cd/issues/319)
- UI should allow redeploying most recent successful deployment from history [\#306](https://github.com/argoproj/argo-cd/issues/306)
- Add spec validation condition to application CRD [\#304](https://github.com/argoproj/argo-cd/issues/304)
- cascade deletion should be a decision made at deletion time, not as a automatic finalizer [\#301](https://github.com/argoproj/argo-cd/issues/301)
- ArgoCD Project CRD [\#295](https://github.com/argoproj/argo-cd/issues/295)
- \[UI\] Cluster list page [\#290](https://github.com/argoproj/argo-cd/issues/290)
- \[UI\] Support app force delete in UI and CLI [\#288](https://github.com/argoproj/argo-cd/issues/288)
- \[UI\] Application conditions on application details page [\#287](https://github.com/argoproj/argo-cd/issues/287)
- Warning message if controller detect resources which belongs to multiple application [\#277](https://github.com/argoproj/argo-cd/issues/277)
- Remove `argocd \(un\)install` into argocd-util [\#254](https://github.com/argoproj/argo-cd/issues/254)
- Application controller should not erase previously collected info about resources if repo or cluster is not avaiable [\#210](https://github.com/argoproj/argo-cd/issues/210)
- App controller should use kubernetes converter instead of refetching service/depoyment [\#192](https://github.com/argoproj/argo-cd/issues/192)
- Support for event hooks as Argo workflows [\#87](https://github.com/argoproj/argo-cd/issues/87)
- Argo workflow integration [\#85](https://github.com/argoproj/argo-cd/issues/85)

**Merged pull requests:**

- Rework installation process to apply from install.yaml [\#421](https://github.com/argoproj/argo-cd/pull/421) ([jessesuen](https://github.com/jessesuen))
- Issue \#419 - Add ability to dump heap profile by sending SIGUSR2 [\#420](https://github.com/argoproj/argo-cd/pull/420) ([alexmt](https://github.com/alexmt))
- Clean up RBAC policy rule format for non-project based resources [\#418](https://github.com/argoproj/argo-cd/pull/418) ([jessesuen](https://github.com/jessesuen))
- Issue \#414 - fix nil pointer in 'argocd cluster add' [\#416](https://github.com/argoproj/argo-cd/pull/416) ([alexmt](https://github.com/alexmt))
- Rename users service into account service [\#415](https://github.com/argoproj/argo-cd/pull/415) ([jessesuen](https://github.com/jessesuen))
- Switch repo-server to use in-memory cache in lieu of redis. Periodically dump stats [\#413](https://github.com/argoproj/argo-cd/pull/413) ([jessesuen](https://github.com/jessesuen))
- Add support for hook deletion policies \(OnSuccess, OnFailure\) [\#412](https://github.com/argoproj/argo-cd/pull/412) ([jessesuen](https://github.com/jessesuen))
- create User Api [\#411](https://github.com/argoproj/argo-cd/pull/411) ([JazminGonzalez-Rivero](https://github.com/JazminGonzalez-Rivero))
- Add ability to terminate a running operation \(resolves \#379\) [\#409](https://github.com/argoproj/argo-cd/pull/409) ([jessesuen](https://github.com/jessesuen))
- Issue \#407 - fix nil pointer dereference in GetSpecErrors [\#408](https://github.com/argoproj/argo-cd/pull/408) ([alexmt](https://github.com/alexmt))
- Issue \#404 - Unset application parameters [\#405](https://github.com/argoproj/argo-cd/pull/405) ([alexmt](https://github.com/alexmt))
- Fix production swagger [\#403](https://github.com/argoproj/argo-cd/pull/403) ([merenbach](https://github.com/merenbach))
- Fix issue where ingress incorrectly was converting to v1 instead of extensions/v1beta1 [\#397](https://github.com/argoproj/argo-cd/pull/397) ([jessesuen](https://github.com/jessesuen))
- Apply logic was ignoring  `kubectl apply` failures [\#395](https://github.com/argoproj/argo-cd/pull/395) ([jessesuen](https://github.com/jessesuen))
- \( resolves 375 \) add admin password util. [\#394](https://github.com/argoproj/argo-cd/pull/394) ([JazminGonzalez-Rivero](https://github.com/JazminGonzalez-Rivero))
- Issue \#364 - Assess health of StatefulSets, DaemonSets, ReplicaSets [\#391](https://github.com/argoproj/argo-cd/pull/391) ([alexmt](https://github.com/alexmt))
- Change default connection state [\#388](https://github.com/argoproj/argo-cd/pull/388) ([merenbach](https://github.com/merenbach))
- Label hooks so that the cluster resource watch will be notified about completions [\#387](https://github.com/argoproj/argo-cd/pull/387) ([jessesuen](https://github.com/jessesuen))
- Issue \#385 - Sync and health status should be unknown if controller unable to load target/live state [\#386](https://github.com/argoproj/argo-cd/pull/386) ([alexmt](https://github.com/alexmt))
- Issue \#349 - argocd wait passes when an ingress object failed [\#384](https://github.com/argoproj/argo-cd/pull/384) ([alexmt](https://github.com/alexmt))
- Improve logging. Prevent some unecessary patches to app [\#383](https://github.com/argoproj/argo-cd/pull/383) ([jessesuen](https://github.com/jessesuen))
- If manifest query is a commit sha, check cache first to prevent locking git repo [\#382](https://github.com/argoproj/argo-cd/pull/382) ([jessesuen](https://github.com/jessesuen))
- Force app refresh after sync [\#373](https://github.com/argoproj/argo-cd/pull/373) ([merenbach](https://github.com/merenbach))
- Issue \#277 - Warning message if controller detect resources which belongs to multiple application [\#372](https://github.com/argoproj/argo-cd/pull/372) ([alexmt](https://github.com/alexmt))
- Issue \#304 - Print information about app conditions [\#371](https://github.com/argoproj/argo-cd/pull/371) ([alexmt](https://github.com/alexmt))
- Various fixes to sync logic: [\#370](https://github.com/argoproj/argo-cd/pull/370) ([jessesuen](https://github.com/jessesuen))
- Use hook strategy as the default when performing a sync [\#368](https://github.com/argoproj/argo-cd/pull/368) ([jessesuen](https://github.com/jessesuen))
- Issue \#210 - Application controller should not erase previously collected info about resources if repo or cluster is not available [\#367](https://github.com/argoproj/argo-cd/pull/367) ([alexmt](https://github.com/alexmt))
- app-name label was inadvertently injected into spec.selector [\#366](https://github.com/argoproj/argo-cd/pull/366) ([jessesuen](https://github.com/jessesuen))
- Fix git authentication implementation when using using SSH key \(resolves \#339\) [\#362](https://github.com/argoproj/argo-cd/pull/362) ([jessesuen](https://github.com/jessesuen))
- Cascade deletion is decided during app deletion, instead of app creation \(resolves \#301\) [\#361](https://github.com/argoproj/argo-cd/pull/361) ([jessesuen](https://github.com/jessesuen))
- Issue \#304 - Add spec validation condition to application CRD [\#360](https://github.com/argoproj/argo-cd/pull/360) ([alexmt](https://github.com/alexmt))
- Remove unnecessary role privileges from api server [\#359](https://github.com/argoproj/argo-cd/pull/359) ([jessesuen](https://github.com/jessesuen))
- Consolidate printing of hook resources with application resources [\#358](https://github.com/argoproj/argo-cd/pull/358) ([jessesuen](https://github.com/jessesuen))
- Move k8s health assessment into a stand-alone library. Use kubectl convert to statically assess health [\#356](https://github.com/argoproj/argo-cd/pull/356) ([jessesuen](https://github.com/jessesuen))
- Add missing health status [\#354](https://github.com/argoproj/argo-cd/pull/354) ([merenbach](https://github.com/merenbach))
- Support for PreSync, Sync, PostSync resource hooks [\#350](https://github.com/argoproj/argo-cd/pull/350) ([jessesuen](https://github.com/jessesuen))
- Retry sync and rollback [\#347](https://github.com/argoproj/argo-cd/pull/347) ([merenbach](https://github.com/merenbach))
- Issue \#343 - Getting permission Denied application destination is not… [\#345](https://github.com/argoproj/argo-cd/pull/345) ([alexmt](https://github.com/alexmt))
- Fix issue where --disable-auth did not disable RBAC properly \(resolves \#332\) [\#344](https://github.com/argoproj/argo-cd/pull/344) ([jessesuen](https://github.com/jessesuen))
- Remove namespace from app URL [\#334](https://github.com/argoproj/argo-cd/pull/334) ([merenbach](https://github.com/merenbach))
- Issue \#295 - Add repositories to project spec [\#331](https://github.com/argoproj/argo-cd/pull/331) ([alexmt](https://github.com/alexmt))
- Idempotent cluster create [\#328](https://github.com/argoproj/argo-cd/pull/328) ([merenbach](https://github.com/merenbach))
- move install functionality from custom to kubectl apply [\#327](https://github.com/argoproj/argo-cd/pull/327) ([JazminGonzalez-Rivero](https://github.com/JazminGonzalez-Rivero))
- RBAC import/export [\#325](https://github.com/argoproj/argo-cd/pull/325) ([merenbach](https://github.com/merenbach))
- Idempotent repo add [\#321](https://github.com/argoproj/argo-cd/pull/321) ([merenbach](https://github.com/merenbach))
- Enable CGO on Linux builds [\#320](https://github.com/argoproj/argo-cd/pull/320) ([merenbach](https://github.com/merenbach))
- Rm swagger.json from reposerver [\#318](https://github.com/argoproj/argo-cd/pull/318) ([merenbach](https://github.com/merenbach))
- Issue \#295 - Allow editing project destinations using CLI [\#317](https://github.com/argoproj/argo-cd/pull/317) ([alexmt](https://github.com/alexmt))
- Projects feature bug fixes [\#313](https://github.com/argoproj/argo-cd/pull/313) ([alexmt](https://github.com/alexmt))
- Issue \#295 - implement app destination permissions validation [\#310](https://github.com/argoproj/argo-cd/pull/310) ([alexmt](https://github.com/alexmt))
- add validation to argocd app set -p [\#309](https://github.com/argoproj/argo-cd/pull/309) ([JazminGonzalez-Rivero](https://github.com/JazminGonzalez-Rivero))
- Support cluster management using the internal k8s API address https://kubernetes.default.svc [\#307](https://github.com/argoproj/argo-cd/pull/307) ([jessesuen](https://github.com/jessesuen))
- Issue \#295 - add project CRD, basic API and CLI implementation [\#299](https://github.com/argoproj/argo-cd/pull/299) ([alexmt](https://github.com/alexmt))
- Generate swagger files [\#278](https://github.com/argoproj/argo-cd/pull/278) ([merenbach](https://github.com/merenbach))

## [v0.5.4](https://github.com/argoproj/argo-cd/tree/v0.5.4) (2018-06-27)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.5.3...v0.5.4)

**Closed issues:**

- Error messages in idempotent application/repo create use wrong variable [\#329](https://github.com/argoproj/argo-cd/issues/329)
- Cluster creation should be idempotent [\#326](https://github.com/argoproj/argo-cd/issues/326)
- `argocd repo add` using https messes up git global username/password  [\#315](https://github.com/argoproj/argo-cd/issues/315)
- argocd linux binary fails unless user is root [\#311](https://github.com/argoproj/argo-cd/issues/311)
- `argocd repo add` should be idempotent [\#305](https://github.com/argoproj/argo-cd/issues/305)
- Support cluster addition using internal hostname: https://kubernetes.default.svc.cluster.local [\#300](https://github.com/argoproj/argo-cd/issues/300)
- import/export should now support RBAC configmap [\#294](https://github.com/argoproj/argo-cd/issues/294)
- Enable Swagger UI for argo api server [\#267](https://github.com/argoproj/argo-cd/issues/267)
- Update Cluster privileges of ArgoCD- read only for whole cluster, read write for only namespace [\#212](https://github.com/argoproj/argo-cd/issues/212)
- argocd app set -p to have validation [\#195](https://github.com/argoproj/argo-cd/issues/195)
- Restrict the cluster namespace access [\#183](https://github.com/argoproj/argo-cd/issues/183)
- E2E tests for main features [\#159](https://github.com/argoproj/argo-cd/issues/159)

## [v0.5.3](https://github.com/argoproj/argo-cd/tree/v0.5.3) (2018-06-20)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.5.2...v0.5.3)

**Closed issues:**

- \[UI\] Support `prune` option for app sync operation on app details page [\#289](https://github.com/argoproj/argo-cd/issues/289)
- argocd get app --refresh [\#269](https://github.com/argoproj/argo-cd/issues/269)
- `argocd app diff` should support local repository comparison [\#239](https://github.com/argoproj/argo-cd/issues/239)

**Merged pull requests:**

- Support diffing a local ksonnet app to the live application state \(resolves \#239\) [\#298](https://github.com/argoproj/argo-cd/pull/298) ([jessesuen](https://github.com/jessesuen))
- Add ability to show last operation result in `app get`. Show path in `app list -o wide` [\#297](https://github.com/argoproj/argo-cd/pull/297) ([jessesuen](https://github.com/jessesuen))
- Update dependencies: ksonnet v0.11, golang v1.10, debian v9.4 [\#296](https://github.com/argoproj/argo-cd/pull/296) ([jessesuen](https://github.com/jessesuen))
- Add ability to force a refresh of an app during get \(resolves \#269\) [\#293](https://github.com/argoproj/argo-cd/pull/293) ([jessesuen](https://github.com/jessesuen))
- Automatically restart API server upon certificate changes [\#292](https://github.com/argoproj/argo-cd/pull/292) ([jessesuen](https://github.com/jessesuen))

## [v0.5.2](https://github.com/argoproj/argo-cd/tree/v0.5.2) (2018-06-14)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.5.1...v0.5.2)

**Closed issues:**

- \[UI\] Resource events tab on application details page [\#286](https://github.com/argoproj/argo-cd/issues/286)
- Application controller fails to get app state if app has resource without name [\#275](https://github.com/argoproj/argo-cd/issues/275)
- \[UI\] Display pod status on application details page [\#231](https://github.com/argoproj/argo-cd/issues/231)

**Merged pull requests:**

- Issue \#275 - Application controller fails to get app state if app has resource without name [\#285](https://github.com/argoproj/argo-cd/pull/285) ([alexmt](https://github.com/alexmt))

## [v0.5.1](https://github.com/argoproj/argo-cd/tree/v0.5.1) (2018-06-13)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.5.0...v0.5.1)

**Closed issues:**

- API server incorrectly compose application fully qualified name for RBAC check [\#283](https://github.com/argoproj/argo-cd/issues/283)
- It is impossible to restrict application access by repository URL [\#280](https://github.com/argoproj/argo-cd/issues/280)
- Rate limiter is preventing force refreshes \(e.g. webhook\) from functioning [\#260](https://github.com/argoproj/argo-cd/issues/260)

**Merged pull requests:**

- Issue \#260 - Rate limiter is preventing force refreshes \(e.g. webhook\) from functioning [\#282](https://github.com/argoproj/argo-cd/pull/282) ([alexmt](https://github.com/alexmt))

## [v0.5.0](https://github.com/argoproj/argo-cd/tree/v0.5.0) (2018-06-12)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.4.7...v0.5.0)

**Closed issues:**

- `argocd app sync` should fail if `prune` flag is required [\#273](https://github.com/argoproj/argo-cd/issues/273)
- \[UI\] Redesign application list page [\#243](https://github.com/argoproj/argo-cd/issues/243)
- \[UI\] Application creation wizard [\#242](https://github.com/argoproj/argo-cd/issues/242)
- \[UI\] Repositories list page [\#241](https://github.com/argoproj/argo-cd/issues/241)
- Repo browsing to improve app creation process [\#71](https://github.com/argoproj/argo-cd/issues/71)

**Merged pull requests:**

- Issue \#283 - API server incorrectly compose application fully qualified name for RBAC check [\#284](https://github.com/argoproj/argo-cd/pull/284) ([alexmt](https://github.com/alexmt))
- Issue \#280 - It is impossible to restrict application access by repository URL [\#281](https://github.com/argoproj/argo-cd/pull/281) ([alexmt](https://github.com/alexmt))
- Fail app sync if `prune` flag is required [\#276](https://github.com/argoproj/argo-cd/pull/276) ([merenbach](https://github.com/merenbach))
- Issue \#271 - perform three way diff only if resource has expected state and live state with last-applied-configuration annotation [\#274](https://github.com/argoproj/argo-cd/pull/274) ([alexmt](https://github.com/alexmt))
- Take into account number of unavailable replicas to decided if deployment is healthy or not [\#270](https://github.com/argoproj/argo-cd/pull/270) ([alexmt](https://github.com/alexmt))
- Tweak flags for import/export, thanks @jessesuen [\#268](https://github.com/argoproj/argo-cd/pull/268) ([merenbach](https://github.com/merenbach))
- fix \#120 refactor the rbac code to support customizable claims enforc… [\#265](https://github.com/argoproj/argo-cd/pull/265) ([wanghong230](https://github.com/wanghong230))
- Add path to API /application/{repo}/ksonnet response [\#264](https://github.com/argoproj/argo-cd/pull/264) ([alexmt](https://github.com/alexmt))
- Implement RBAC support \(issue \#120\) [\#263](https://github.com/argoproj/argo-cd/pull/263) ([jessesuen](https://github.com/jessesuen))
- Introduce `argocd app manifests` to print application manifests from git or live [\#261](https://github.com/argoproj/argo-cd/pull/261) ([jessesuen](https://github.com/jessesuen))
- Implement workaround for https://github.com/golang/go/issues/21955 [\#256](https://github.com/argoproj/argo-cd/pull/256) ([alexmt](https://github.com/alexmt))
- Support resource import/export [\#255](https://github.com/argoproj/argo-cd/pull/255) ([merenbach](https://github.com/merenbach))
- Clean up .proto definitions for consistency and reduction of pointer usage [\#253](https://github.com/argoproj/argo-cd/pull/253) ([jessesuen](https://github.com/jessesuen))
- ListDir should not fail if Redis is down [\#252](https://github.com/argoproj/argo-cd/pull/252) ([alexmt](https://github.com/alexmt))
- Fix bug secret controller which is causing update loop in secret controller [\#251](https://github.com/argoproj/argo-cd/pull/251) ([alexmt](https://github.com/alexmt))
- GET /cluster/\<clustername\> API should not panic if invalid cluster url is provided [\#250](https://github.com/argoproj/argo-cd/pull/250) ([alexmt](https://github.com/alexmt))
- Wrap method signatures [\#249](https://github.com/argoproj/argo-cd/pull/249) ([merenbach](https://github.com/merenbach))
- Issue \#244 - Cluster/Repository connection status [\#248](https://github.com/argoproj/argo-cd/pull/248) ([alexmt](https://github.com/alexmt))

## [v0.4.7](https://github.com/argoproj/argo-cd/tree/v0.4.7) (2018-06-07)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.4.6...v0.4.7)

**Closed issues:**

- argocd-server gets into a catch-22 failed startup situation with unhealthy ELB [\#272](https://github.com/argoproj/argo-cd/issues/272)
- nil pointer dereference in application controller [\#271](https://github.com/argoproj/argo-cd/issues/271)
- Linux argocd-server cannot start due to builtin-policy.csv: no such file or directory [\#266](https://github.com/argoproj/argo-cd/issues/266)
- `argocd app wait` should retry connection errors [\#259](https://github.com/argoproj/argo-cd/issues/259)
- Repo names do not accept underscore [\#258](https://github.com/argoproj/argo-cd/issues/258)
- Import & export of app instances  [\#213](https://github.com/argoproj/argo-cd/issues/213)
- RBAC controls [\#120](https://github.com/argoproj/argo-cd/issues/120)

## [v0.4.6](https://github.com/argoproj/argo-cd/tree/v0.4.6) (2018-06-06)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.4.5...v0.4.6)

**Closed issues:**

- \[UI\] Argo CD UI overwrites history [\#257](https://github.com/argoproj/argo-cd/issues/257)
- Create/Update endpoints should wrap resource objects into request objects [\#246](https://github.com/argoproj/argo-cd/issues/246)
- Cluster/Repository connection status [\#244](https://github.com/argoproj/argo-cd/issues/244)
- argocd-util enhancements [\#179](https://github.com/argoproj/argo-cd/issues/179)

## [v0.4.5](https://github.com/argoproj/argo-cd/tree/v0.4.5) (2018-05-31)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.4.4...v0.4.5)

## [v0.4.4](https://github.com/argoproj/argo-cd/tree/v0.4.4) (2018-05-30)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.4.3...v0.4.4)

**Implemented enhancements:**

- Repo and cluster need an error field [\#169](https://github.com/argoproj/argo-cd/issues/169)
- Add error fields to cluster/repo, shell output [\#200](https://github.com/argoproj/argo-cd/pull/200) ([merenbach](https://github.com/merenbach))

**Fixed bugs:**

- Controller does not persist rollback operation result [\#233](https://github.com/argoproj/argo-cd/issues/233)

**Closed issues:**

- `argo cd app get` should show parameters and overrides [\#240](https://github.com/argoproj/argo-cd/issues/240)
- `argocd app create --upsert` [\#238](https://github.com/argoproj/argo-cd/issues/238)
- \[UI\] Allow viewing pod side car container logs [\#235](https://github.com/argoproj/argo-cd/issues/235)
- \[UI\] Resource filtering on Application Details page [\#232](https://github.com/argoproj/argo-cd/issues/232)
- \[UI\] Display operation state on application details page [\#230](https://github.com/argoproj/argo-cd/issues/230)
- A lot of test namespaces in ArgoCD ci cluster [\#223](https://github.com/argoproj/argo-cd/issues/223)
- Prevent ELB timeouts with keep alives [\#217](https://github.com/argoproj/argo-cd/issues/217)
- Changes to argocd configmap/secret should get applied automatically [\#174](https://github.com/argoproj/argo-cd/issues/174)
- App sync frequently fails due to concurrent app modification [\#147](https://github.com/argoproj/argo-cd/issues/147)
- Implement API to load application resource events [\#135](https://github.com/argoproj/argo-cd/issues/135)

**Merged pull requests:**

- Add ability to show parameters and overrides in CLI \(resolves \#240\) [\#247](https://github.com/argoproj/argo-cd/pull/247) ([jessesuen](https://github.com/jessesuen))
- Issue \#238 - add upsert flag to 'argocd app create' command [\#245](https://github.com/argoproj/argo-cd/pull/245) ([alexmt](https://github.com/alexmt))
- Add Events API endpoint [\#237](https://github.com/argoproj/argo-cd/pull/237) ([merenbach](https://github.com/merenbach))
- Issue \#233 - Controller does not persist rollback operation result [\#234](https://github.com/argoproj/argo-cd/pull/234) ([alexmt](https://github.com/alexmt))
- Add repo browsing endpoint [\#229](https://github.com/argoproj/argo-cd/pull/229) ([merenbach](https://github.com/merenbach))
- Support subscribing to settings updates and auto-restart of dex and API server \(resolves \#174\) [\#227](https://github.com/argoproj/argo-cd/pull/227) ([jessesuen](https://github.com/jessesuen))
- Issue \#147 - App sync frequently fails due to concurrent app modification [\#226](https://github.com/argoproj/argo-cd/pull/226) ([alexmt](https://github.com/alexmt))
- Issue \# 223 - Remove app finalizers during e2e fixture teardown [\#225](https://github.com/argoproj/argo-cd/pull/225) ([alexmt](https://github.com/alexmt))

## [v0.4.3](https://github.com/argoproj/argo-cd/tree/v0.4.3) (2018-05-21)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.4.2...v0.4.3)

**Closed issues:**

- ArgoCD fails to fetch latest git repository updates [\#185](https://github.com/argoproj/argo-cd/issues/185)

**Merged pull requests:**

- Move local branch deletion as part of git Reset\(\) \(resolves \#185\) [\#222](https://github.com/argoproj/argo-cd/pull/222) ([jessesuen](https://github.com/jessesuen))
- Fix exit code for app wait [\#219](https://github.com/argoproj/argo-cd/pull/219) ([merenbach](https://github.com/merenbach))

## [v0.4.2](https://github.com/argoproj/argo-cd/tree/v0.4.2) (2018-05-21)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.4.1...v0.4.2)

**Closed issues:**

- Add `argo app ensure-healthy` cli [\#211](https://github.com/argoproj/argo-cd/issues/211)

**Merged pull requests:**

- CLI improvements [\#218](https://github.com/argoproj/argo-cd/pull/218) ([jessesuen](https://github.com/jessesuen))

## [v0.4.1](https://github.com/argoproj/argo-cd/tree/v0.4.1) (2018-05-18)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.4.0...v0.4.1)

**Implemented enhancements:**

- Add `argocd app wait` command [\#216](https://github.com/argoproj/argo-cd/pull/216) ([merenbach](https://github.com/merenbach))
- Manifest endpoint [\#207](https://github.com/argoproj/argo-cd/pull/207) ([merenbach](https://github.com/merenbach))

**Closed issues:**

- UI doesn't show app history [\#215](https://github.com/argoproj/argo-cd/issues/215)
- /api/v1/application/\<appname\>/manifests to generate manifests [\#197](https://github.com/argoproj/argo-cd/issues/197)
- Allow downloading of argocd binaries directly from API server [\#184](https://github.com/argoproj/argo-cd/issues/184)

## [v0.4.0](https://github.com/argoproj/argo-cd/tree/v0.4.0) (2018-05-17)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.4.0-alpha1...v0.4.0)

**Closed issues:**

- Update to ksonnet 0.10.2 [\#208](https://github.com/argoproj/argo-cd/issues/208)
- App "Synced" status does not detect removed fields [\#199](https://github.com/argoproj/argo-cd/issues/199)
- Always run kubectl apply during application sync [\#194](https://github.com/argoproj/argo-cd/issues/194)
- ArgoCD UI s/rollback/history/ and s/deploy/sync/ [\#191](https://github.com/argoproj/argo-cd/issues/191)
- Cleanup types [\#190](https://github.com/argoproj/argo-cd/issues/190)
- ArgoCD UI should switch to Spec.Destination.Server/Namespace fields [\#189](https://github.com/argoproj/argo-cd/issues/189)
- ArgoCD UI should supply return\_url param to /auth/login [\#188](https://github.com/argoproj/argo-cd/issues/188)
- Migration utility for 0.3 to 0.4 [\#178](https://github.com/argoproj/argo-cd/issues/178)
- ArgoCD applications should have a rolled up health status [\#146](https://github.com/argoproj/argo-cd/issues/146)
- Redact cluster and repo credentials from get/list cluster repo responses [\#134](https://github.com/argoproj/argo-cd/issues/134)
- gRPC error not being shown in CLI [\#131](https://github.com/argoproj/argo-cd/issues/131)
- Before applying resources, run kubectl apply --dry-run to detect issues [\#122](https://github.com/argoproj/argo-cd/issues/122)
- Move application sync logic to the controller [\#119](https://github.com/argoproj/argo-cd/issues/119)
- app delete should be done through controller using finalizers [\#118](https://github.com/argoproj/argo-cd/issues/118)

**Merged pull requests:**

- Add intelligence in diff libray to perform three-way diff from last-applied-configuration annotation \(resolves \#199\) [\#214](https://github.com/argoproj/argo-cd/pull/214) ([jessesuen](https://github.com/jessesuen))
- Make sure api server started during fixture setup [\#209](https://github.com/argoproj/argo-cd/pull/209) ([alexmt](https://github.com/alexmt))
- Issue \#118 - app delete should be done through controller using finalizers [\#206](https://github.com/argoproj/argo-cd/pull/206) ([alexmt](https://github.com/alexmt))
- Implement App management and repo management e2e tests [\#205](https://github.com/argoproj/argo-cd/pull/205) ([alexmt](https://github.com/alexmt))
- Add last update time to operation status, fix operation status patching [\#204](https://github.com/argoproj/argo-cd/pull/204) ([alexmt](https://github.com/alexmt))
- Add connect timeouts when interacting with SSH git repos \(resolves \#131\) [\#203](https://github.com/argoproj/argo-cd/pull/203) ([jessesuen](https://github.com/jessesuen))
- Refactor application controller sync/apply loop [\#202](https://github.com/argoproj/argo-cd/pull/202) ([jessesuen](https://github.com/jessesuen))
- Rename recent deployments to history [\#201](https://github.com/argoproj/argo-cd/pull/201) ([merenbach](https://github.com/merenbach))
- Issue \#146 - Render health status information in 'app list' and 'app get' commands [\#198](https://github.com/argoproj/argo-cd/pull/198) ([alexmt](https://github.com/alexmt))
- Add 'database' library for CRUD operations against repos and clusters [\#196](https://github.com/argoproj/argo-cd/pull/196) ([jessesuen](https://github.com/jessesuen))
- Add 0.3.0 to 0.4.0 migration utility [\#186](https://github.com/argoproj/argo-cd/pull/186) ([merenbach](https://github.com/merenbach))

## [v0.4.0-alpha1](https://github.com/argoproj/argo-cd/tree/v0.4.0-alpha1) (2018-05-11)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.3.3...v0.4.0-alpha1)

**Fixed bugs:**

- argocd uninstall should not remove CRD by default [\#167](https://github.com/argoproj/argo-cd/issues/167)

**Closed issues:**

- App server and namespace should be required and possibly immutable [\#175](https://github.com/argoproj/argo-cd/issues/175)
- CLI support for SSO login [\#172](https://github.com/argoproj/argo-cd/issues/172)
- Kubernetes API resource discovery needs to be cached [\#170](https://github.com/argoproj/argo-cd/issues/170)
- Improve CI job stability [\#166](https://github.com/argoproj/argo-cd/issues/166)
- API  POST /v1/repo should trim whitespaces before saving username [\#161](https://github.com/argoproj/argo-cd/issues/161)
- If argocd token is expired server should return 401 instead of 500 [\#157](https://github.com/argoproj/argo-cd/issues/157)
- SSO Login Button in UI [\#154](https://github.com/argoproj/argo-cd/issues/154)
- Application create should be idempotent [\#151](https://github.com/argoproj/argo-cd/issues/151)
- Repo server should memoize results [\#143](https://github.com/argoproj/argo-cd/issues/143)
- Normalize git repo URLs [\#138](https://github.com/argoproj/argo-cd/issues/138)
- GitHub webhook handling [\#121](https://github.com/argoproj/argo-cd/issues/121)

**Merged pull requests:**

- Issue \#187 - implement `argo settings install` command [\#193](https://github.com/argoproj/argo-cd/pull/193) ([alexmt](https://github.com/alexmt))
- Support OAuth2 login flow from CLI \(resolves \#172\) [\#181](https://github.com/argoproj/argo-cd/pull/181) ([jessesuen](https://github.com/jessesuen))
- Move sync logic to contoller [\#180](https://github.com/argoproj/argo-cd/pull/180) ([alexmt](https://github.com/alexmt))
- Make ApplicationSpec.Destination non-optional, non-pointer [\#177](https://github.com/argoproj/argo-cd/pull/177) ([merenbach](https://github.com/merenbach))
- Cache kubernetes API resource discovery \(resolves \#170\) [\#176](https://github.com/argoproj/argo-cd/pull/176) ([jessesuen](https://github.com/jessesuen))
- Do not delete namespace or CRD during uninstall unless explicitly stated \(resolves \#167\) [\#173](https://github.com/argoproj/argo-cd/pull/173) ([jessesuen](https://github.com/jessesuen))
- Trim spaces server-side in GitHub usernames [\#171](https://github.com/argoproj/argo-cd/pull/171) ([merenbach](https://github.com/merenbach))
- Don't fail when new app has same spec as old [\#168](https://github.com/argoproj/argo-cd/pull/168) ([merenbach](https://github.com/merenbach))
- Introduce caching layer to repo server to improve query response times [\#165](https://github.com/argoproj/argo-cd/pull/165) ([jessesuen](https://github.com/jessesuen))
- Issue \#146 - ArgoCD applications should have a rolled up health status [\#164](https://github.com/argoproj/argo-cd/pull/164) ([alexmt](https://github.com/alexmt))
- Refactor repo server and git client [\#163](https://github.com/argoproj/argo-cd/pull/163) ([jessesuen](https://github.com/jessesuen))
- Expand Git repo URL normalization [\#162](https://github.com/argoproj/argo-cd/pull/162) ([merenbach](https://github.com/merenbach))
- Add GitHub webhook handling to fast-track controller application reprocessing [\#160](https://github.com/argoproj/argo-cd/pull/160) ([jessesuen](https://github.com/jessesuen))
- Issue \#157 - If argocd token is expired server should return 401 instead of 500 [\#158](https://github.com/argoproj/argo-cd/pull/158) ([alexmt](https://github.com/alexmt))

## [v0.3.3](https://github.com/argoproj/argo-cd/tree/v0.3.3) (2018-05-03)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.3.2...v0.3.3)

**Implemented enhancements:**

- Add workflow for blue-green deployments [\#148](https://github.com/argoproj/argo-cd/pull/148) ([merenbach](https://github.com/merenbach))

**Closed issues:**

- Application update failes due to concurrent access [\#155](https://github.com/argoproj/argo-cd/issues/155)
- Workflow for blue-green deployment while preserving original deployment name [\#77](https://github.com/argoproj/argo-cd/issues/77)
- SSO integration [\#70](https://github.com/argoproj/argo-cd/issues/70)

**Merged pull requests:**

- Issue \#155 - Application update failes due to concurrent access [\#156](https://github.com/argoproj/argo-cd/pull/156) ([alexmt](https://github.com/alexmt))
- Add settings endpoint so frontend can show/hide SSO login button. Ren… [\#153](https://github.com/argoproj/argo-cd/pull/153) ([jessesuen](https://github.com/jessesuen))
- Initial SSO support using dex [\#152](https://github.com/argoproj/argo-cd/pull/152) ([jessesuen](https://github.com/jessesuen))
- Redact sensitive repo/cluster information upon retrieval [\#150](https://github.com/argoproj/argo-cd/pull/150) ([merenbach](https://github.com/merenbach))

## [v0.3.2](https://github.com/argoproj/argo-cd/tree/v0.3.2) (2018-05-01)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.3.1...v0.3.2)

**Fixed bugs:**

- Service is always out of sync if spec contains empty loadBalancerSourceRanges [\#136](https://github.com/argoproj/argo-cd/issues/136)
- Deployment is always out of sync if resource limit is specified [\#132](https://github.com/argoproj/argo-cd/issues/132)

**Closed issues:**

- App comparator should lookup resources by name [\#141](https://github.com/argoproj/argo-cd/issues/141)
- Application sync should delete 'unexpected' resources [\#139](https://github.com/argoproj/argo-cd/issues/139)
- repo-server is not tracking HEAD properly [\#133](https://github.com/argoproj/argo-cd/issues/133)
- An error when configuring a repo will expose git token in CLI output [\#130](https://github.com/argoproj/argo-cd/issues/130)
- Better error message when running commands without ~/.argocd/config [\#128](https://github.com/argoproj/argo-cd/issues/128)
- `argocd app sync --dry-run` appends items to history [\#127](https://github.com/argoproj/argo-cd/issues/127)
- Rollback ignores parameter overrides [\#110](https://github.com/argoproj/argo-cd/issues/110)
- argocd context management [\#103](https://github.com/argoproj/argo-cd/issues/103)
- DeletionPropagation policy when deleting an application [\#99](https://github.com/argoproj/argo-cd/issues/99)

**Merged pull requests:**

- Issue \#147 - Use patch to update recentDeployments field [\#149](https://github.com/argoproj/argo-cd/pull/149) ([alexmt](https://github.com/alexmt))
- Issue \#136 - Use custom formatter to get desired state of deployment and service [\#145](https://github.com/argoproj/argo-cd/pull/145) ([alexmt](https://github.com/alexmt))
-  Issue \#139 - Application sync should delete 'unexpected' resources [\#144](https://github.com/argoproj/argo-cd/pull/144) ([alexmt](https://github.com/alexmt))
- Improve comparator to fall back to looking up a resource by name [\#142](https://github.com/argoproj/argo-cd/pull/142) ([jessesuen](https://github.com/jessesuen))
- Refactor git library [\#140](https://github.com/argoproj/argo-cd/pull/140) ([jessesuen](https://github.com/jessesuen))
- Populated 'unexpected' resources while comparing target and live states [\#137](https://github.com/argoproj/argo-cd/pull/137) ([alexmt](https://github.com/alexmt))
- Don't ask for user credentials if username and password are specified as arguments [\#129](https://github.com/argoproj/argo-cd/pull/129) ([alexmt](https://github.com/alexmt))

## [v0.3.1](https://github.com/argoproj/argo-cd/tree/v0.3.1) (2018-04-24)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.3.0...v0.3.1)

**Closed issues:**

- CLI should have `argo app history` `argo app rollback` [\#125](https://github.com/argoproj/argo-cd/issues/125)
- argocd login fails if .argocd does not exist [\#123](https://github.com/argoproj/argo-cd/issues/123)

**Merged pull requests:**

- Allow overriding server or namespace separately [\#126](https://github.com/argoproj/argo-cd/pull/126) ([alexmt](https://github.com/alexmt))
- Issue \#123 - Create .argocd directory before saving config file [\#124](https://github.com/argoproj/argo-cd/pull/124) ([alexmt](https://github.com/alexmt))
- Issue \#110 - Rollback ignores parameter overrides [\#117](https://github.com/argoproj/argo-cd/pull/117) ([alexmt](https://github.com/alexmt))

## [v0.3.0](https://github.com/argoproj/argo-cd/tree/v0.3.0) (2018-04-23)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.2.0...v0.3.0)

**Implemented enhancements:**

- gRPC interceptor to transform kubernetes errors to gRPC errors [\#30](https://github.com/argoproj/argo-cd/issues/30)
- Add JWT support [\#60](https://github.com/argoproj/argo-cd/pull/60) ([merenbach](https://github.com/merenbach))

**Fixed bugs:**

- Argo CD server and repo server fails to start if installed into custom namespace [\#80](https://github.com/argoproj/argo-cd/issues/80)
- Upgrading installs fails on final step [\#54](https://github.com/argoproj/argo-cd/issues/54)
- Set session cookies, errors appropriately [\#100](https://github.com/argoproj/argo-cd/pull/100) ([merenbach](https://github.com/merenbach))

**Closed issues:**

- Update app via `argocd app set` [\#116](https://github.com/argoproj/argo-cd/issues/116)
- application-controller stuck in loop \(Failed to watch clusters: unknown \(get secrets\), retrying in 10s\) [\#115](https://github.com/argoproj/argo-cd/issues/115)
- `app add/rm` should be `app create/delete` [\#114](https://github.com/argoproj/argo-cd/issues/114)
- argocd-server service should expose 80 and 443 during install [\#112](https://github.com/argoproj/argo-cd/issues/112)
- App controller incorrectly report that app is out of sync [\#108](https://github.com/argoproj/argo-cd/issues/108)
- Deployment of resources needs to use `kubectl apply` semantics [\#106](https://github.com/argoproj/argo-cd/issues/106)
- Add ability to display unexpected resources [\#104](https://github.com/argoproj/argo-cd/issues/104)
- Application create/update should validate cluster/repo/app/environment [\#97](https://github.com/argoproj/argo-cd/issues/97)
- Rework local config to support multiple servers/credentials [\#95](https://github.com/argoproj/argo-cd/issues/95)
- UI needs to support auth tokens [\#91](https://github.com/argoproj/argo-cd/issues/91)
- Allow CLI tokens to be specified manually [\#88](https://github.com/argoproj/argo-cd/issues/88)
- Unable to add repo if repo URL contains upper case characters [\#81](https://github.com/argoproj/argo-cd/issues/81)
- Implement live application CRD update [\#76](https://github.com/argoproj/argo-cd/issues/76)
- Delete pods using ArgoCD UI [\#75](https://github.com/argoproj/argo-cd/issues/75)
- Pods logs on application details page [\#74](https://github.com/argoproj/argo-cd/issues/74)
- `argocd login` and local ~/.argocd/config for CLI [\#72](https://github.com/argoproj/argo-cd/issues/72)
- Resources tree view on application details page [\#67](https://github.com/argoproj/argo-cd/issues/67)
- Migrate to ksonnet 0.10 [\#66](https://github.com/argoproj/argo-cd/issues/66)
- Delete all the kube object once app is removed [\#65](https://github.com/argoproj/argo-cd/issues/65)
- Argocd install failure [\#63](https://github.com/argoproj/argo-cd/issues/63)
- Rollback feature enhancement [\#33](https://github.com/argoproj/argo-cd/issues/33)
- TLS and authentication support on API Server [\#29](https://github.com/argoproj/argo-cd/issues/29)
- Persist environment parameters during deployment [\#27](https://github.com/argoproj/argo-cd/issues/27)
- Deployed application resources should be annotated [\#25](https://github.com/argoproj/argo-cd/issues/25)

**Merged pull requests:**

- Expose port 443 to proxy to port 8080 [\#113](https://github.com/argoproj/argo-cd/pull/113) ([merenbach](https://github.com/merenbach))
- Add application validation [\#111](https://github.com/argoproj/argo-cd/pull/111) ([jessesuen](https://github.com/jessesuen))
- Issue \#108 - App controller incorrectly report that app is out of sync [\#109](https://github.com/argoproj/argo-cd/pull/109) ([alexmt](https://github.com/alexmt))
- Add syncPolicy field to application CRD [\#107](https://github.com/argoproj/argo-cd/pull/107) ([merenbach](https://github.com/merenbach))
- Fix null pointer error in controller [\#105](https://github.com/argoproj/argo-cd/pull/105) ([alexmt](https://github.com/alexmt))
- Rework local config to support multiple servers/credentials [\#101](https://github.com/argoproj/argo-cd/pull/101) ([jessesuen](https://github.com/jessesuen))
- Allow ignoring recource deletion related errors while deleting application [\#98](https://github.com/argoproj/argo-cd/pull/98) ([alexmt](https://github.com/alexmt))
- Add server auth cookies [\#96](https://github.com/argoproj/argo-cd/pull/96) ([merenbach](https://github.com/merenbach))
- Add tests for lowercase repo secret names [\#94](https://github.com/argoproj/argo-cd/pull/94) ([merenbach](https://github.com/merenbach))
- Support manual token use for CLI commands [\#90](https://github.com/argoproj/argo-cd/pull/90) ([merenbach](https://github.com/merenbach))
- Convert Kubernetes errors to gRPC errors [\#89](https://github.com/argoproj/argo-cd/pull/89) ([merenbach](https://github.com/merenbach))
- Add session gateway [\#84](https://github.com/argoproj/argo-cd/pull/84) ([merenbach](https://github.com/merenbach))
- Issue \#69 - Auto-sync option in application CRD instance [\#83](https://github.com/argoproj/argo-cd/pull/83) ([alexmt](https://github.com/alexmt))
- Add `argocd login` command [\#82](https://github.com/argoproj/argo-cd/pull/82) ([merenbach](https://github.com/merenbach))
- TLS support. HTTP/HTTPS/gRPC all serving on single port [\#78](https://github.com/argoproj/argo-cd/pull/78) ([jessesuen](https://github.com/jessesuen))
- Issue \#67 - Application controller should persist ksonnet app parameters in app CRD [\#73](https://github.com/argoproj/argo-cd/pull/73) ([alexmt](https://github.com/alexmt))
- Update ksonnet and begin using it as library instead of parsing stdout [\#69](https://github.com/argoproj/argo-cd/pull/69) ([jessesuen](https://github.com/jessesuen))
- Issue \#67 - Persist resources tree in application CRD [\#68](https://github.com/argoproj/argo-cd/pull/68) ([alexmt](https://github.com/alexmt))
- Installer and settings management refactoring [\#64](https://github.com/argoproj/argo-cd/pull/64) ([jessesuen](https://github.com/jessesuen))
- Update grpc-gateway version [\#62](https://github.com/argoproj/argo-cd/pull/62) ([alexmt](https://github.com/alexmt))
- Add authentication endpoints [\#61](https://github.com/argoproj/argo-cd/pull/61) ([merenbach](https://github.com/merenbach))

## [v0.2.0](https://github.com/argoproj/argo-cd/tree/v0.2.0) (2018-03-29)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/v0.1.0...v0.2.0)

**Implemented enhancements:**

- Allow use of public repos without prior registration [\#36](https://github.com/argoproj/argo-cd/pull/36) ([merenbach](https://github.com/merenbach))
- Support -f/--file flag in `argocd app add` [\#35](https://github.com/argoproj/argo-cd/pull/35) ([merenbach](https://github.com/merenbach))

**Fixed bugs:**

- ArgoCD failed to add cluster [\#43](https://github.com/argoproj/argo-cd/issues/43)
- Don't overwrite application status in tryRefreshAppStatus, thanks @alexmt [\#42](https://github.com/argoproj/argo-cd/pull/42) ([merenbach](https://github.com/merenbach))

**Closed issues:**

- Server fails to read config map from namespace on cluster with RBAC enabled [\#57](https://github.com/argoproj/argo-cd/issues/57)
- Project should have a license [\#49](https://github.com/argoproj/argo-cd/issues/49)
- Support 'ephemeral' applications  [\#41](https://github.com/argoproj/argo-cd/issues/41)
- Support ssh git URLs and ssh key authentication [\#34](https://github.com/argoproj/argo-cd/issues/34)
- `argocd app add` should support an -f option [\#31](https://github.com/argoproj/argo-cd/issues/31)
- Argo CD should try to clone app repo without password if repo is not registered [\#28](https://github.com/argoproj/argo-cd/issues/28)

**Merged pull requests:**

- Maintain list of recent deployments in app CRD [\#59](https://github.com/argoproj/argo-cd/pull/59) ([alexmt](https://github.com/alexmt))
- Issue \#57 - Add configmaps into argocd server role [\#58](https://github.com/argoproj/argo-cd/pull/58) ([alexmt](https://github.com/alexmt))
- Fix deleting resources which do not support 'deletecollection' method but support 'delete' method [\#56](https://github.com/argoproj/argo-cd/pull/56) ([alexmt](https://github.com/alexmt))
- Argo server should not fail if configmap name is not provided or config map does not exist [\#55](https://github.com/argoproj/argo-cd/pull/55) ([alexmt](https://github.com/alexmt))
- Add application source and component parameters into recentDeployment field of application CRD [\#53](https://github.com/argoproj/argo-cd/pull/53) ([alexmt](https://github.com/alexmt))
- Replace ephemeral environments with override parameters [\#52](https://github.com/argoproj/argo-cd/pull/52) ([alexmt](https://github.com/alexmt))
- Add password hashing [\#51](https://github.com/argoproj/argo-cd/pull/51) ([merenbach](https://github.com/merenbach))
- Delete child dependents while deleting app resources [\#48](https://github.com/argoproj/argo-cd/pull/48) ([alexmt](https://github.com/alexmt))
- Add install configmap override flag [\#47](https://github.com/argoproj/argo-cd/pull/47) ([merenbach](https://github.com/merenbach))
- Use ksonnet release version and fix app copy command [\#46](https://github.com/argoproj/argo-cd/pull/46) ([alexmt](https://github.com/alexmt))
- Disable strict host key checking while cloning repo in repo-server [\#45](https://github.com/argoproj/argo-cd/pull/45) ([alexmt](https://github.com/alexmt))
- Issue \#43 - Don't setup RBAC resources for clusters with basic authentication [\#44](https://github.com/argoproj/argo-cd/pull/44) ([alexmt](https://github.com/alexmt))
- Support deploying/destroying ephemeral environments [\#40](https://github.com/argoproj/argo-cd/pull/40) ([alexmt](https://github.com/alexmt))
- Persist parameters during deployment \(Sync\) [\#39](https://github.com/argoproj/argo-cd/pull/39) ([merenbach](https://github.com/merenbach))
- Add new dependency to CONTRIBUTING.md [\#38](https://github.com/argoproj/argo-cd/pull/38) ([merenbach](https://github.com/merenbach))
- Issue \#34 - Support ssh git URLs and ssh key authentication [\#37](https://github.com/argoproj/argo-cd/pull/37) ([alexmt](https://github.com/alexmt))
- Update CONTRIBUTING.md [\#32](https://github.com/argoproj/argo-cd/pull/32) ([merenbach](https://github.com/merenbach))

## [v0.1.0](https://github.com/argoproj/argo-cd/tree/v0.1.0) (2018-03-12)

[Full Changelog](https://github.com/argoproj/argo-cd/compare/a67038ae2e9cb9b9b16423702f98b41e36601001...v0.1.0)

**Closed issues:**

- Example app [\#26](https://github.com/argoproj/argo-cd/issues/26)
- Support rollback without updating spec in git repo [\#21](https://github.com/argoproj/argo-cd/issues/21)
- Move Kubernetes manifest generation into separate service [\#19](https://github.com/argoproj/argo-cd/issues/19)

**Merged pull requests:**

- Add README.md [\#24](https://github.com/argoproj/argo-cd/pull/24) ([jessesuen](https://github.com/jessesuen))
- Issue \#21 - Support rollback without updating spec in git repo [\#23](https://github.com/argoproj/argo-cd/pull/23) ([alexmt](https://github.com/alexmt))
- Issue \#19 - Add repo-server into installer [\#22](https://github.com/argoproj/argo-cd/pull/22) ([alexmt](https://github.com/alexmt))
- Issue \#19 - Move Kubernetes manifest generation into separate service [\#20](https://github.com/argoproj/argo-cd/pull/20) ([alexmt](https://github.com/alexmt))
- Correctly finish app streaming call when app watch requests stopped [\#18](https://github.com/argoproj/argo-cd/pull/18) ([alexmt](https://github.com/alexmt))
- Fix git client CloneOrFetch: remove local branches after fetching rem… [\#17](https://github.com/argoproj/argo-cd/pull/17) ([alexmt](https://github.com/alexmt))
- Fix in-cluster kube config creation [\#16](https://github.com/argoproj/argo-cd/pull/16) ([alexmt](https://github.com/alexmt))
- Implement api server installer [\#15](https://github.com/argoproj/argo-cd/pull/15) ([alexmt](https://github.com/alexmt))
- Support server sent events formatting in streaming api [\#14](https://github.com/argoproj/argo-cd/pull/14) ([alexmt](https://github.com/alexmt))
- Application target state [\#13](https://github.com/argoproj/argo-cd/pull/13) ([alexmt](https://github.com/alexmt))
- Implement application watch API [\#12](https://github.com/argoproj/argo-cd/pull/12) ([alexmt](https://github.com/alexmt))
- Add argo ci config, temporary skip failing tests [\#11](https://github.com/argoproj/argo-cd/pull/11) ([alexmt](https://github.com/alexmt))
- Implement controller e2e tests [\#10](https://github.com/argoproj/argo-cd/pull/10) ([alexmt](https://github.com/alexmt))
- Allow segregate applications among controllers using label [\#9](https://github.com/argoproj/argo-cd/pull/9) ([alexmt](https://github.com/alexmt))
- Implement application controller installer [\#8](https://github.com/argoproj/argo-cd/pull/8) ([alexmt](https://github.com/alexmt))
- Add controller-image make target [\#7](https://github.com/argoproj/argo-cd/pull/7) ([alexmt](https://github.com/alexmt))
- Synchronize repo processing in app manager [\#6](https://github.com/argoproj/argo-cd/pull/6) ([alexmt](https://github.com/alexmt))
- More application manager tests [\#5](https://github.com/argoproj/argo-cd/pull/5) ([alexmt](https://github.com/alexmt))
- Implement application manager unit tests [\#4](https://github.com/argoproj/argo-cd/pull/4) ([alexmt](https://github.com/alexmt))
- Update application CRD types, implement application manager draft [\#3](https://github.com/argoproj/argo-cd/pull/3) ([alexmt](https://github.com/alexmt))
- Implement Application controller boilerplate [\#2](https://github.com/argoproj/argo-cd/pull/2) ([alexmt](https://github.com/alexmt))
- Add application CRD definition. Implement CRD install CLI command [\#1](https://github.com/argoproj/argo-cd/pull/1) ([alexmt](https://github.com/alexmt))



\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*
