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

### Add configuration into package.json For For Next.Js

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

### Add configuration into package.json For For Svelte

```
"scripts": {
    "prepare": "cd .. && husky svelte_project_ecommerce/.husky",
}
```

# **\*** Backend Project Setup **\***

## 1. Golang

```
go mod init project_name
```

## 2. Node (ts)

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

### Add configuration into package.json For For Node.ts

```
"scripts": {
    "prepare": "cd .. && husky node_project_ecommerce/.husky",
    "lint": "next lint",
    "format": "npx prettier . --write",
    "format:check": "npx prettier . --check"
}
```

## 3. Django

```

```
