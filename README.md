# Monorepo ADO – Front + Back

## Requisitos del agente Self-Hosted

- Node 20.19.x (o NodeTool@0 en pipeline)
- Go 1.22+ (o el que uses)
- Acceso a Internet para bajar deps

## Estructura

- /front → React/Vite (build en `front/dist`)
- /back → Go (binario en `back/bin/server`)
- azure-pipelines.yml → CI en ADO (Self-Hosted)

## Correr local

cd front && npm ci && npm run dev
cd back && go run main.go # http://localhost:8080

## Pipeline

- Trigger: push a `main`
- Jobs: Build_Front, Build_Back
- Artefactos: `front-dist`, `back-bin`
