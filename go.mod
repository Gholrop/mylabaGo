module hierarchical-classes

go 1.21

replace (
    tehnica => ./tehnica
    transport => ./transport
    naztransport => ./naztransport
)

require (
    naztransport v0.0.0
    tehnica v0.0.0
    transport v0.0.0
)
