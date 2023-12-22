"""Go Workspace Configuration for Repository"""

load("@bazel_gazelle//:deps.bzl", "go_repository")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

def workspace_go():
    go_rules_dependencies()
    go_register_toolchains(version = "1.21.3")

def go_dependencies():
    """Third Party Go Dependencies"""
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_cncf_udpa_go",
        importpath = "github.com/cncf/udpa/go",
        sum = "h1:QQ3GSy+MqSHxm/d8nCtnAiZdYFd45cYZPs8vOOIYKfk=",
        version = "v0.0.0-20220112060539-c52dc94e7fbe",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:/inchEIKaYC1Akx+H+gqO04wryn5h75LSazbRlnya1k=",
        version = "v0.0.0-20230607035331-e9ce68804cb4",
    )
    go_repository(
        name = "com_github_creack_pty",
        importpath = "github.com/creack/pty",
        sum = "h1:uDmaGzcdjhF4i/plgjmEsriH11Y0o7RKapEf/LDaM3w=",
        version = "v1.1.9",
    )

    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:wSUXTlLfiAQRWs2F+p+EKOY9rUyis1MyGqJ2DIk5HpM=",
        version = "v0.11.1",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:QkIBuU5k+x7/QXPvPPnWXWlCdaBFApVqftFV6k087DA=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:DVjP2PbBOzHyzA+dn3WhHIq4NdVu3Q+pvivFICf/7fo=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:KhyjKVUg7Usr/dYsdSqoFveMYd5ko72D+zANwlG1mmg=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=",
        version = "v0.5.9",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:KjJaJ9iWZ3jOFZIf1Lqf4laDRCasjl0BCmnEGxkdLb4=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_jackc_pgpassfile",
        importpath = "github.com/jackc/pgpassfile",
        sum = "h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgservicefile",
        importpath = "github.com/jackc/pgservicefile",
        sum = "h1:bbPeKD0xmW/Y25WS6cokEszi5g+S0QxI/d45PkRi7Nk=",
        version = "v0.0.0-20221227161230-091c0ba34f0a",
    )
    go_repository(
        name = "com_github_jackc_pgx_v5",
        importpath = "github.com/jackc/pgx/v5",
        sum = "h1:cxFyXhxlvAifxnkKKdlxv8XqUf59tDlYjnV5YYfsJJY=",
        version = "v5.4.3",
    )
    go_repository(
        name = "com_github_jackc_puddle_v2",
        importpath = "github.com/jackc/puddle/v2",
        sum = "h1:RhxXJtFG022u4ibrCSMSiu5aOq1i77R3OHKNJj77OAk=",
        version = "v2.2.1",
    )
    go_repository(
        name = "com_github_jinzhu_inflection",
        importpath = "github.com/jinzhu/inflection",
        sum = "h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jinzhu_now",
        importpath = "github.com/jinzhu/now",
        sum = "h1:/o9tlHleP7gOFmsnYNz3RGnqzefHA47wQpKrrdTIwXQ=",
        version = "v1.1.5",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:WgNl7dwNpEZ6jJ9k1snq4pZsg7DOEN8hP9Xw0Tsjwk0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )

    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:exVL4IDcn6na9z1rAb56Vxr+CgyK3nn3O+epU5NdKM8=",
        version = "v1.12.0",
    )

    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=",
        version = "v1.8.4",
    )
    go_repository(
        name = "com_google_cloud_go_compute",
        importpath = "cloud.google.com/go/compute",
        sum = "h1:tP41Zoavr8ptEqaW6j+LQOnyBBhO7OkOMAGrgLopTwY=",
        version = "v1.23.0",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:mg4jlk7mCAj6xXp9UJ4fjI9VUI5rubuGBW5aJ7UnBMY=",
        version = "v0.2.3",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
        version = "v1.0.0-20201130134442-10cb98267c6c",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "io_gorm_driver_postgres",
        importpath = "gorm.io/driver/postgres",
        sum = "h1:Iyrp9Meh3GmbSuyIAGyjkN+n9K+GHX9b9MqsTL4EJCo=",
        version = "v1.5.4",
    )
    go_repository(
        name = "io_gorm_gorm",
        importpath = "gorm.io/gorm",
        sum = "h1:zR9lOiiYf09VNh5Q1gphfyia1JpiClIWG9hQaxB/mls=",
        version = "v1.25.5",
    )

    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        sum = "h1:IhEN5q69dyKagZPYMSdIjS2HqprW324FRQZJcGqPAsM=",
        version = "v1.6.8",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:SeZZZx0cP0fqUyA+oRzP9k7cSwJlvDFiROO72uwD6i0=",
        version = "v0.0.0-20231002182017-d307bd883b97",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:W18sezcAYs+3tDZX4F80yctqa12jcP1PUS2gQu1zTPU=",
        version = "v0.0.0-20231002182017-d307bd883b97",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:6GQBEOdGkX6MMTLT9V+TjtIRZCw9VPD5Z+yHY9wMgS0=",
        version = "v0.0.0-20231002182017-d307bd883b97",
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        sum = "h1:6FQAR0kM31P6MRdeluor2w2gPaS4SVNrD/DNTxrQ15k=",
        version = "v1.60.0",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        importpath = "google.golang.org/protobuf",
        sum = "h1:g0LDEJHgrBl9N9r17Ru3sqWhkIx2NB67okBHPwC7hs8=",
        version = "v1.31.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:wBqGXzWJW6m1XrIKlAH0Hs1JJ7+9KBwnIO8v66Q9cHc=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_mod",
        importpath = "golang.org/x/mod",
        sum = "h1:KENHtAZL2y3NLMYZeHY9DW8HW8V+kQyJsY/V9JlKvCs=",
        version = "v0.9.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:7eBu7KsSvFDtSXUIDbh3aqlK4DPsZ1rByC8PFfBThos=",
        version = "v0.16.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:jDDenyj+WgFtmV3zYVoi8aE2BwtXFLWOA67ZfNWftiY=",
        version = "v0.13.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        importpath = "golang.org/x/sync",
        sum = "h1:zxkM55ReGkDlKSM+Fu41A+zmbZuaPVbGMzvvdUPznYQ=",
        version = "v0.4.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:Af8nKPmuFypiUBjVoU9V20FiaFXOcuZI21p0ycVYYGE=",
        version = "v0.13.0",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:bb+I9cTfFazGW51MZqBVmZy7+JEJMouUHTUSKVQLBek=",
        version = "v0.13.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:ablQoSUd0tRdKxZewP80B+BaqeKJuVhuRxj/dkrun3k=",
        version = "v0.13.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        sum = "h1:BOw41kyTf3PuCW1pVQf8+Cyg8pMlkYB1oo9iJ6D/lKM=",
        version = "v0.6.0",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:E7g+9GITq07hpfrRu66IVDexMakfv52eLZ2CXBWiKr4=",
        version = "v0.0.0-20191204190536-9bdfabe68543",
    )
