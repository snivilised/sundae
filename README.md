# 🍦sundae: ___assistant for cobra based cli applications___

[![A B](https://img.shields.io/badge/branching-commonflow-informational?style=flat)](https://commonflow.org)
[![A B](https://img.shields.io/badge/merge-rebase-informational?style=flat)](https://git-scm.com/book/en/v2/Git-Branching-Rebasing)
[![A B](https://img.shields.io/badge/branch%20history-linear-blue?style=flat)](https://docs.github.com/en/repositories/configuring-branches-and-merges-in-your-repository/defining-the-mergeability-of-pull-requests/managing-a-branch-protection-rule)
[![Go Reference](https://pkg.go.dev/badge/github.com/snivilised/sundae.svg)](https://pkg.go.dev/github.com/snivilised/sundae)
[![Go report](https://goreportcard.com/badge/github.com/snivilised/sundae)](https://goreportcard.com/report/github.com/snivilised/sundae)
[![Coverage Status](https://coveralls.io/repos/github/snivilised/sundae/badge.svg?branch=main)](https://coveralls.io/github/snivilised/sundae?branch=main&kill_cache=1)
[![Sundae Continuous Integration](https://github.com/snivilised/sundae/actions/workflows/ci-workflow.yml/badge.svg)](https://github.com/snivilised/sundae/actions/workflows/ci-workflow.yml)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
[![A B](https://img.shields.io/badge/commit-conventional-commits?style=flat)](https://www.conventionalcommits.org/)

<!-- MD013/Line Length -->
<!-- MarkDownLint-disable MD013 -->

<!-- MD014/commands-show-output: Dollar signs used before commands without showing output mark down lint -->
<!-- MarkDownLint-disable MD014 -->

<!-- MD033/no-inline-html: Inline HTML -->
<!-- MarkDownLint-disable MD033 -->

<!-- MD040/fenced-code-language: Fenced code blocks should have a language specified -->
<!-- MarkDownLint-disable MD040 -->

<!-- MD028/no-blanks-blockquote: Blank line inside blockquote -->
<!-- MarkDownLint-disable MD028 -->

<p align="left">
  <a href="https://go.dev"><img src="resources/images/go-logo-light-blue.png" width="50" alt="Go" /></a>
</p>

## 🔰 Introduction

_[Cobra](https://cobra.dev/) is an excellent framework for the development of command line applications, but is missing a few features that would make it a bit easier to work with. This package aims to fulfil this purpose, especially in regards to creation of commands, encapsulating commands into a container and providing an export mechanism to re-create cli data in a form that is free from cobra (and indeed sundae) abstractions. The aim of this last aspect to to be able to inject data into the core of an application in a way that removes tight coupling to the `Cobra` framework, which is achieved by representing data only in terms of client defined (native) abstractions.

## 📚 Usage

... tbd

## 🎀 Features

<p align="left">
  <a href="https://onsi.github.io/ginkgo/"><img src="https://onsi.github.io/ginkgo/images/ginkgo.png" width="100" alt="ginkgo" /></a>
  <a href="https://onsi.github.io/gomega/"><img src="https://onsi.github.io/gomega/images/gomega.png" width="100" alt="ginkgo" /></a>
</p>

+ Cobra container; collection of cobra commands that can be independently referenced by name as opposed to via child/parent relationship. The container also takes care of adding commands to the root or any other as required.
+ A `parameter set` groups together all the flag option values, so that they don't have to be handled separately. A single entity (the ___ParamSet___) can be created and passed into the core of the client application.
+ unit testing with [Ginkgo](https://onsi.github.io/ginkgo/)/[Gomega](https://onsi.github.io/gomega/)
+ implemented with [🐍 Cobra](https://cobra.dev/) cli framework
+ i18n with [go-i18n](https://github.com/nicksnyder/go-i18n)
+ linting configuration and pre-commit hooks, (see: [linting-golang](https://freshman.tech/linting-golang/)).

### 🎁 Cobra Container

The container serves as a repository for `Cobra` commands and `Cobrass` parameter sets. Commands in `Cobra` are related to each other via parent child relationships. The container, flattens this hierarchy so that a command can be queried for, simply by its name, as opposed to getting the commands by parent command, ie ___parentCommand.Commands()___.

The methods on the container, should not fail. Any failures that occur are due to programming errors. For this reason, when an error scenario occurs, a panic is raised.

Registering commands/parameter sets with the container, obviates the need to use specific `Cobra` api calls as they are handled on the clients behalf by the container. For parameter sets, the type specific methods on the various ___FlagSet___ definitions, such as ___Float32Var___, do not have to be called by the client. For commands, ___AddCommand___ does not have to be called explicitly either.

### 💎 Param Set

The rationale behind the concept of a parameter set came from initial discovery of how the `Cobra` api worked. Capturing user defined command line input requires binding option values into disparate variables. Having to manage independently defined variables usually at a package level could lead to a scattering of these variables on an adhoc basis. Having to then pass all these items independently into the core of a client application could easily become disorganised.

To manage this, the concept of a `parameter set` was introduced to bring about a consistency of design to the implementation of multiple cli applications. The aim of this is to reduce the number package level global variables that have to be managed. Instead of handling multiple option variables independently, the client can group them together into a parameter set.

Each `Cobra` command can define multiple parameter sets which reflects the different ways that a particular command can be invoked by the user. However, to reduce complexity, it's probably best to stick with a single parameter set per command. Option values not defined by the user can already be defaulted by the `Cobra` api itself, but it may be, that distinguishing the way that a command is invoked (ie what combination of flags/options appear on the command line) may be significant to the application, in which case the client can define multiple parameter sets.

The ___ParamSet___ also handles flag definition on each command. The client defines the flag info and passes this into the appropriate `binder` method depending on the option value type. The binder methods are of the form:

+ ___Bind\<Type>___ : where ___\<Type>___ represents the type, (eg ___BindString___), the client passes in '___info___', a ___FlagInfo___ object and '___to___' a pointer to a variable to which `Cobra` will bind the option value to.

### 🎭 Alternative Flag Set

By default, binding a flag is performed on the default flag set. This flag set is the one you get from calling ___command.Flags()___ (this is performed automatically by ___NewFlagInfo___). However, there are a few more options for defining flags in `Cobra`. There are multiple flag set methods on the `Cobra` command, eg ___command.PersistentFlags()___. To utilise an alternative flag set, the client should use ___NewFlagInfoOnFlagSet___ instead of ___NewFlagInfo___. ___NewFlagInfoOnFlagSet___ requires that an extra parameter be provided and that is the alternative flag set, which can be derived from calling the appropriate method on the `command`, eg:

```go
  paramSet.BindString(
    assistant.NewFlagInfoOnFlagSet("pattern", "p", "default-pattern",
      widgetCommand.PersistentFlags()), &paramSet.Native.Pattern,
  )
```

The flag set defined for the flag (in the above case 'pattern'), will always override the default one defined on the parameter set.

### ✨ Code Generation

Please see [Powershell Code Generation](resources/doc/CODE-GEN.md)
