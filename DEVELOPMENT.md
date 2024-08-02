# Stargazer Development Guide

Here we will describe the steps required to setup a development environment with Stargazer.  

- [Stargazer Development Guide](#stargazer-development-guide)
  - [Setting up a development environment](#setting-up-a-development-environment)
    - [Pre-requisites](#pre-requisites)
    - [1. Fork the Repo](#1-fork-the-repo)
    - [2. Clone the repo down](#2-clone-the-repo-down)
    - [3. Install dependencies](#3-install-dependencies)

## Setting up a development environment

In this section we will walk you through the general process of setting up your development environment to get started.

### Pre-requisites

The following must be installed in order to get started. The details of how to install them is outside the scope of this doc, but generally they should be able to be installed with your systems package manager (apt, yum, brew, choco, etc).

- Go
- Bun
- Node.js

### 1. Fork the Repo

[Fork the repo](https://docs.github.com/en/get-started/quickstart/fork-a-repo) onto your own GitHub account for developing.  

### 2. Clone the repo down

### 3. Install dependencies

```bash
# Install Backent dependencies
go mod tidy

# Install Frontend dependencies
bun install
```
