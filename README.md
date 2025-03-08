# **\*** MicroFrontend and MicroBackend **\***

## **\*** Mapping

```
Angular -> Golang -> Postgres
```

```
Next.Js -> Node -> MongoDB
```

```
Svelte -> Django -> MySQL
```

## **\*** Backend

```
Golang -> Postgres -> ProdcutService
```

```
Node -> MongoDB -> AuthService
```

```
Django -> MySQL -> OrderService
```

## **\*** Frontend

```
Angular -> Product and Order
```

```
Next.Js -> User
```

```
Svelte -> Dashbard
```

# **\*** Frontend Project Setup **\***

## 1. Angular

```
ng new project_name
```

### Setup Husky Prettier & Eslint For Angular

```
npm install --save-dev husky
npx husky init
npm install --save-dev --save-exact prettier
ng lint
```

### Add configuration into package.json For Anuglar

```
"scripts": {
    "prepare": "cd .. && husky angular_project_ecommerce/.husky",
    "lint": "ng lint",
    "format": "npx prettier . --write",
    "format:check": "npx prettier . --check"
}
```

## 2. Next.Js

```
npx create-next-app@latest project_name
```

### Setup Husky Prettier & Eslint For Next.Js

```
npm install --save-dev husky
npx husky init
npm install --save-dev --save-exact prettier
```

### Add configuration into package.json For Next.Js

```
"scripts": {
    "prepare": "cd .. && husky next_project_ecommerce/.husky",
    "lint": "next lint",
    "format": "npx prettier . --write",
    "format:check": "npx prettier . --check"
}
```

## 3. Svelte

```
npx sv create project_name
```

### Setup Husky Prettier & Eslint For Svelte

```
npm install --save-dev husky
npx husky init
```

### Add configuration into package.json For Svelte

```
"scripts": {
    "prepare": "cd .. && husky svelte_project_ecommerce/.husky",
}
```

# **\*** Backend Project Setup **\***

## 1. Node (ts)

```
npm init or npm init -y
npm install --save-dev typescript
npx tsc --init
```

### Setup Husky Prettier & Eslint For Node.ts

```
npm install --save-dev husky
npx husky init
npm install --save-dev eslint
npm init @eslint/config@latest
npm install --save-dev --save-exact prettier
```

### Add configuration into package.json For Node.ts

```
"scripts": {
    "prepare": "cd .. && husky node_project_ecommerce/.husky",
    "lint": "next lint",
    "format": "npx prettier . --write",
    "format:check": "npx prettier . --check"
}
```

## 2. Golang

```
go mod init project_name
```

### Setup Pre-commit For Django

```
go install github.com/golangci/golangci-lint/cmd/golangci-lint@<version>
pre-commit install
create .pre-commit-config.yaml file
```

### Add configuration into .pre-commit-config.yaml For golang

```
- repo: https://github.com/pre-commit/pre-commit-hooks
  rev: <lastest-version>
  hooks:
    - id: check-merge-conflict
    - id: check-yaml
    - id: end-of-file-fixer
    - id: trailing-whitespace
- repo: https://github.com/dnephin/pre-commit-golang
  rev: master
  hooks:
    - id: go-fmt
    - id: go-vet
    - id: go-lint
    - id: go-imports
    - id: go-cyclo
      args: [-over=15]
    - id: validate-toml
    - id: no-go-testing
    - id: golangci-lint
    - id: go-critic
    - id: go-unit-tests
    - id: go-build
    - id: go-mod-tidy
```

## 3. Django

```
django-admin startproject project_name
```

### Setup Pre-commit For Django

```
pip install pre-commit
add .pre-commit-config.yaml
create .pre-commit-config.yaml file
```

### Add configuration into .pre-commit-config.yaml For golang

```
repos:
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v4.1.0
    hooks:
      - id: check-yaml
      - id: trailing-whitespace
  - repo: https://github.com/psf/black
    rev: 22.10.0
    hooks:
      - id: black
```

# **\*** Backstage **\***

```
npx @backstage/create-app@latest backstage_name
```

## Run BackStage

```
yarn dev
```

### Setup Husky For BackStage

```
yarn add --dev husky
yarn husky init
```

### Add configuration into package.json For BackStage

```
"scripts": {
  "prepare": "cd .. && husky backstage-micro-fe-be/.husky",
}
```
