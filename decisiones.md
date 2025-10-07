# Decisiones

## Stack y estructura

- Front: React + Vite por velocidad y simplicidad de build.
- Back: Go por binarios livianos y test fáciles.
- Monorepo: `/front`, `/back` para versionar todo junto.

## Diseño del pipeline

- YAML + Self-Hosted para evitar límites de Microsoft-hosted y controlar toolchains.
- Stage único `CI` con dos jobs independientes (front/back) para aislar fallas.
- Publicación de artefactos (`front-dist`, `back-bin`) para trazabilidad.

## Evidencias

- Capturas en `/evidencias`

## Lecciones y problemas

- Ajusté el `pool: SelfHosted` para evitar “No hosted parallelism…”.
- Aseguré Node 20 con NodeTool y Go local instalado.
- Separé jobs para identificar cuál rompe si falla.
