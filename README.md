[![Coverage Status](https://coveralls.io/repos/github/xplorfin/ozzo-validators/badge.svg?branch=master)](https://coveralls.io/github/xplorfin/ozzo-validators?branch=master)
[![Renovate enabled](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://app.renovatebot.com/dashboard#github/xplorfin/ozzo-validators)
[![Build status](https://github.com/xplorfin/ozzo-validators/workflows/test/badge.svg)](https://github.com/xplorfin/ozzo-validators/actions?query=workflow%3Atest)
[![Build status](https://github.com/xplorfin/ozzo-validators/workflows/goreleaser/badge.svg)](https://github.com/xplorfin/ozzo-validators/actions?query=workflow%3Agoreleaser)
[![](https://godoc.org/github.com/xplorfin/ozzo-validators?status.svg)](https://pkg.go.dev/github.com/xplorfin/ozzo-validators)

# Custom Validators for Ozzo

This is a series of custom validators for ozzo used by [entropy](http://entropy.rocks/). In particular this adds a new rule type [`IntRule`](rules/int_rule.go) based on [StringRule](https://pkg.go.dev/github.com/go-ozzo/ozzo-validation#StringRule).

This also adds several new rules available in the godoc.